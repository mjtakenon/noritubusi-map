package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"noritubusi-map/backend/app/db"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
	"github.com/srinathgs/mysqlstore"
)

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// TODO: 返し方を見直す(JSONやText,もしくはNoContentか…)
func getBuildingInfoInRange(c echo.Context) error {
	beginLat := c.QueryParam("begin_latitude")
	beginLong := c.QueryParam("begin_longitude")
	endLat := c.QueryParam("end_latitude")
	endLong := c.QueryParam("end_longitude")

	if beginLat == "" || beginLong == "" || endLat == "" || endLong == "" {
		return c.String(http.StatusBadRequest, "invalid paramater")
	}

	stationInfo, err := DB.GetBuildingInfoInRange(beginLat, beginLong, endLat, endLong)
	if err != nil {
		log.Println("/stations get info error:", err)
		return c.String(http.StatusInternalServerError, "server error")
	}

	return c.JSON(http.StatusOK, stationInfo)
}

func getStationInfoByID(c echo.Context) error {
	stationID, err := strconv.Atoi(c.Param("stationid"))
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid parameter")
	}

	stationInfo, err := DB.GetStationInfoByID(stationID)
	if err != nil {
		return c.String(http.StatusNotFound, "not found")
	}

	return c.JSON(http.StatusOK, []db.StationInfo{stationInfo})
}

func getStationInfoByBuildingID(c echo.Context) error {
	buildingID, err := strconv.Atoi(c.Param("buildingid"))
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid parameter")
	}

	stationInfos, err := DB.GetStationInfoByBuildingID(buildingID)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid parameter")
	}

	if len(stationInfos) == 0 {
		return c.String(http.StatusNotFound, "not found")
	}

	return c.JSON(http.StatusOK, stationInfos)
}

func getStationNameSuggestion(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	if keyword == "" {
		return c.String(http.StatusBadRequest, "invalid parameter")
	}

	stationInfos, err := DB.GetStationsInfoByKeyword(keyword)
	if err != nil {
		return c.String(http.StatusInternalServerError, "server error")
	}

	return c.JSON(http.StatusOK, stationInfos)
}

func getRailwaysInfoAll(c echo.Context) error {
	railwayInfos, err := DB.GetRailwaysInfoAll()

	if err != nil {
		return c.String(http.StatusInternalServerError, "server error")
	}

	return c.JSON(http.StatusOK, railwayInfos)
}

func getRailwaysInfoByQuery(c echo.Context) error {
	query := c.Param("query")
	if query == "" {
		return c.String(http.StatusBadRequest, "invalid parameter")
	}

	id, err := strconv.Atoi(query)

	railwayInfos := []db.RailwayInfo{}
	if err != nil {
		railwayInfos, err = DB.GetRailwaysInfoByName(query)
	} else {
		railwayInfos, err = DB.GetRailwaysInfoByID(id)
	}

	if err != nil {
		return c.String(http.StatusNotFound, "not found")
	}

	return c.JSON(http.StatusOK, railwayInfos)
}

func createUser(c echo.Context) error {
	//パラメータ検査
	userID := c.FormValue("userid")
	password := c.FormValue("password")
	if userID == "" || password == "" {
		return c.String(http.StatusBadRequest, "invalid parameter")
	}

	//ハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return c.String(http.StatusInternalServerError, "server error")
	}

	err = DB.InsertUser(userID, string(hash))
	if err != nil {
		return c.String(http.StatusBadRequest, "exist")
	}

	//仮に有効期限1週間に設定
	if err = saveSession(userID, 60*60*24*7, c); err != nil {
		return c.String(http.StatusInternalServerError, "cookie store failed")
	}

	return c.String(http.StatusCreated, "ok")
}

// remember meオプション等によって有効期限変える？
func saveSession(userID string, maxAge int, c echo.Context) error {
	sess, err := session.Get("session", c)

	if err != nil {
		return err
	}

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   maxAge,
		HttpOnly: true,
	}

	sess.Values["userID"] = userID

	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return err
	}

	return nil
}

var (
	DB db.DB
)

func main() {

	stationInfoDBUserName := "user"
	stationInfoDBPassword := "password"
	stationInfoDBAddress := os.Getenv("STATION_DB_ADDRESS")
	if stationInfoDBAddress == "" {
		stationInfoDBAddress = "localhost:3314"
	}
	stationInfoDBName := "noritubusi_map"

	// Echo instance
	e := echo.New()

	// stationInfo DB connect
	err := DB.New(stationInfoDBUserName, stationInfoDBPassword, stationInfoDBAddress, stationInfoDBName)
	if err != nil {
		e.Logger.Fatal("station DB Connection Error:", err)
	}

	// session store connect
	sqlstore, err := mysqlstore.NewMySQLStoreFromConnection(DB.DB.DB, "session", "/", 60*60*24*7, []byte("sessionid"))

	if err != nil {
		e.Logger.Fatal("Session Store Error:", err)
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sqlstore))

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"},
	}))

	// Routes
	e.GET("/", hello)

	e.GET("/railways", getRailwaysInfoAll)
	e.GET("/railways/:query", getRailwaysInfoByQuery)

	e.GET("/buildings", getBuildingInfoInRange)
	e.GET("/buildings/:buildingid", getStationInfoByBuildingID)

	e.GET("/stations/:stationid", getStationInfoByID)
	e.GET("/stations/suggest", getStationNameSuggestion)

	e.POST("/signup", createUser)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

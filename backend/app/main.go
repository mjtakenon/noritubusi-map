package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"noritubusi-map/backend/app/db"

	"github.com/asaskevich/govalidator"
	"github.com/boj/redistore"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
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

	stationInfo, err := DB.GetStationInfoInRange(beginLat, beginLong, endLat, endLong)
	if err != nil {
		log.Println("/stations get info error:", err)
		return c.String(http.StatusInternalServerError, "server error")
	}

	return c.JSON(http.StatusOK, convertStationInfo2BuildingInfo(stationInfo))
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

	return c.JSON(http.StatusOK, stationInfo)
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

func getBuildingNameSuggestion(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	if keyword == "" {
		return c.String(http.StatusBadRequest, "invalid parameter")
	}

	stationInfos, err := DB.GetStationsInfoByKeyword(keyword)
	if err != nil {
		return c.String(http.StatusInternalServerError, "server error")
	}

	return c.JSON(http.StatusOK, convertStationInfo2BuildingInfo(stationInfos))
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

	stationInfos := []db.StationInfo{}
	if err != nil {
		// パーセントエンコーディングをデコード
		query, err = url.QueryUnescape(query)
		if err != nil {
			return c.String(http.StatusInternalServerError, "server error")
		}

		stationInfos, err = DB.GetStationsInfoByRailwayName(query)
	} else {
		stationInfos, err = DB.GetStationsInfoByRailwayID(id)
	}

	if err != nil {
		return c.String(http.StatusInternalServerError, "server error")
	}

	if len(stationInfos) == 0 {
		return c.String(http.StatusNotFound, "not found")
	} else {
		return c.JSON(http.StatusOK, stationInfos)
	}
}

func getUserInfo(c echo.Context) error {
	userid := c.Param("userid")

	userInfo, err := DB.GetUserInfoByUserID(userid)
	if err != nil {
		return c.String(http.StatusNotFound, "not found")
	}

	return c.JSON(http.StatusOK, userInfo)
}

func putUserInfo(c echo.Context) error {
	currentPass := c.FormValue("current_password")
	newPass := c.FormValue("new_password")

	// セッション取得
	sess, err := session.Get("session", c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "session not found")
	}
	userID, ok := sess["userID"]
	if ok == false {
		return c.String(http.StatusUnauthorized, "you should login before changing password")
	}

	//ユーザ情報取得
	userInfo, err := DB.GetUserInfoByUserID(userID)
	if err != nil { // DBとsessionの整合性が取れてないとき？
		return c.String(http.StatusInternalServerError, "internal error")
	}

	//パスワード比較
	if err := bcrypt.CompareHashAndPassword([]byte(userInfo.HashedPassword), []byte(currentPass)); err != nil {
		return c.String(http.StatusUnauthorized, "login failed")
	}

	//ハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(newPass), bcrypt.DefaultCost)
	if err != nil {
		return c.String(http.StatusInternalServerError, "server error")
	}

	err = DB.UpdateUser(userid, string(hash))
	if err != nil {
		return c.String(http.StatusBadRequest, "failed")
	}

	return c.JSON(http.StatusOK, "ok")
}

func deleteUserInfo(c echo.Context) error {
	userid := c.FormValue("userid")
	password := c.FormValue("password")

	if userid == "" || password == "" {
		return c.String(http.StatusBadRequest, "invalid parameter")
	}

	//ユーザ情報取得
	userInfo, err := DB.GetUserInfoByUserID(userid)
	if err != nil {
		return c.String(http.StatusUnauthorized, "login failed")
	}

	//パスワード比較
	if err := bcrypt.CompareHashAndPassword([]byte(userInfo.HashedPassword), []byte(password)); err != nil {
		return c.String(http.StatusUnauthorized, "login failed")
	}

	err = DB.DeleteUser(userid)
	if err != nil {
		return c.String(http.StatusBadRequest, "failed")
	}

	return c.JSON(http.StatusOK, "ok")
}

func createUser(c echo.Context) error {
	//パラメータ検査
	userID := c.FormValue("userid")
	password := c.FormValue("password")
	if userID == "" || password == "" ||
		!govalidator.IsAlphanumeric(userID) || !govalidator.IsByteLength(userID, 1, 128) {
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

func signin(c echo.Context) error {
	userID := c.FormValue("userid")
	password := c.FormValue("password")
	if userID == "" || password == "" {
		return c.String(http.StatusBadRequest, "invalid parameter")
	}

	//ユーザ情報取得
	userInfo, err := DB.GetUserInfoByUserID(userID)
	if err != nil {
		return c.String(http.StatusUnauthorized, "login failed")
	}

	//パスワード比較
	if err := bcrypt.CompareHashAndPassword([]byte(userInfo.HashedPassword), []byte(password)); err != nil {
		return c.String(http.StatusUnauthorized, "login failed")
	}

	//ログイン成功 , セッション保存
	if err := saveSession(userID, 60*60*24*7, c); err != nil {
		return c.String(http.StatusInternalServerError, "login failed")
	}

	return c.String(http.StatusOK, "success login")
}

func signout(c echo.Context) error {
	// Cookieの削除
	if err := saveSession("", -1, c); err != nil {
		return c.String(http.StatusInternalServerError, "failed logout")
	}

	return c.String(http.StatusOK, "success logout")
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

// 構造体
type BuildingInfo struct {
	BuildingID int64  `json:"building_id"`
	Name       string `json:"name"`
	Latitude   string `json:"latitude"`
	Longitude  string `json:"longitude"`
	Lines      []Line `json:"lines"`
}

type Line struct {
	RailwayName    string `json:"railway_name"`
	StationID      int64  `json:"station_id"`
	OrderInRailWay int64  `json:"order_in_railway"`
}

func convertStationInfo2BuildingInfo(stationInfos []db.StationInfo) []BuildingInfo {
	//現在の建物番号
	prevID := int64(0)

	var buildingInfo BuildingInfo
	ret := []BuildingInfo{}

	// BuildingInfo構造体に詰め替え
	for _, info := range stationInfos {
		// 前と異なる建物番号
		if prevID != info.BuildingId {
			if prevID != 0 {
				ret = append(ret, buildingInfo)
			}
			prevID = info.BuildingId
			buildingInfo = BuildingInfo{
				BuildingID: info.BuildingId,
				Name:       info.Name,
				Latitude:   info.Latitude,
				Longitude:  info.Longitude,
			}
		}

		// 路線情報追加
		buildingInfo.Lines = append(buildingInfo.Lines, Line{RailwayName: info.RailwayName,
			StationID:      info.StationId,
			OrderInRailWay: info.OrderInRailway,
		})
	}

	// 最後に触ったBuildingInfoの追加
	if prevID != 0 {
		ret = append(ret, buildingInfo)
	}

	return ret
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
	sqlstore, err := redistore.NewRediStore(10, "tcp", "redis:6379", "", []byte("secret-key"))
	if err != nil {
		e.Logger.Fatal("Session Store Error:", err)
	}
	defer sqlstore.Close()

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
	e.GET("/buildings/suggest", getBuildingNameSuggestion)

	e.GET("/stations/:stationid", getStationInfoByID)
	e.GET("/stations/suggest", getStationNameSuggestion)

	e.GET("/users/:userid", getUserInfo)
	e.PUT("/users", putUserInfo)
	e.POST("/users/delete", deleteUserInfo)

	e.POST("/signup", createUser)

	e.POST("/signin", signin)
	e.DELETE("/signin", signout)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

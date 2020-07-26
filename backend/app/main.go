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
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getBuildingInfoInRange(c echo.Context) error {
	beginLat := c.QueryParam("begin_latitude")
	beginLong := c.QueryParam("begin_longitude")
	endLat := c.QueryParam("end_latitude")
	endLong := c.QueryParam("end_longitude")

	if beginLat == "" || beginLong == "" || endLat == "" || endLong == "" {
		return c.JSON(http.StatusBadRequest, Response{Status: 200, Message: "invalid paramater"})
	}

	stationInfo, err := DB.GetStationInfoInRange(beginLat, beginLong, endLat, endLong)
	if err != nil {
		log.Println("/stations get info error:", err)
		return c.JSON(http.StatusInternalServerError, Response{Status: 900, Message: "server error"})
	}

	return c.JSON(http.StatusOK, convertStationInfo2BuildingInfo(stationInfo))
}

func getStationInfoByID(c echo.Context) error {
	stationID, err := strconv.Atoi(c.Param("stationid"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{Status: 200, Message: "invalid paramater"})
	}

	stationInfo, err := DB.GetStationInfoByID(stationID)
	if err != nil {
		return c.JSON(http.StatusNotFound, Response{Status: 103, Message: "not found"})
	}

	return c.JSON(http.StatusOK, stationInfo)
}

func getStationInfoByBuildingID(c echo.Context) error {
	buildingID, err := strconv.Atoi(c.Param("buildingid"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{Status: 200, Message: "invalid paramater"})
	}

	stationInfos, err := DB.GetStationInfoByBuildingID(buildingID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Status: 901, Message: "db error"})
	}

	if len(stationInfos) == 0 {
		return c.JSON(http.StatusNotFound, Response{Status: 103, Message: "not found"})
	}

	return c.JSON(http.StatusOK, stationInfos)
}

func getStationNameSuggestion(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	if keyword == "" {
		return c.JSON(http.StatusBadRequest, Response{Status: 200, Message: "invalid paramater"})
	}

	stationInfos, err := DB.GetStationsInfoByKeyword(keyword)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Status: 900, Message: "server error"})
	}

	return c.JSON(http.StatusOK, stationInfos)
}

func getBuildingNameSuggestion(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	if keyword == "" {
		return c.JSON(http.StatusBadRequest, Response{Status: 200, Message: "invalid paramater"})
	}

	stationInfos, err := DB.GetStationsInfoByKeyword(keyword)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Status: 900, Message: "server error"})
	}

	return c.JSON(http.StatusOK, convertStationInfo2BuildingInfo(stationInfos))
}

func getRailwaysInfoAll(c echo.Context) error {
	railwayInfos, err := DB.GetRailwaysInfoAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Status: 900, Message: "server error"})
	}

	return c.JSON(http.StatusOK, railwayInfos)
}

func getRailwaysInfoByQuery(c echo.Context) error {
	query := c.Param("query")
	if query == "" {
		return c.JSON(http.StatusBadRequest, Response{Status: 200, Message: "invalid paramater"})
	}

	id, err := strconv.Atoi(query)

	stationInfos := []db.StationInfo{}
	if err != nil {
		// パーセントエンコーディングをデコード
		query, err = url.QueryUnescape(query)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Response{Status: 900, Message: "server error"})
		}

		stationInfos, err = DB.GetStationsInfoByRailwayName(query)
	} else {
		stationInfos, err = DB.GetStationsInfoByRailwayID(id)
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Status: 900, Message: "server error"})
	}

	if len(stationInfos) == 0 {
		return c.JSON(http.StatusNotFound, Response{Status: 103, Message: "not found"})
	} else {
		return c.JSON(http.StatusOK, stationInfos)
	}
}

func getUserInfo(c echo.Context) error {
	userid := c.Param("userid")

	userInfo, err := DB.GetUserInfoByUserID(userid)
	if err != nil {
		return c.JSON(http.StatusNotFound, Response{Status: 103, Message: "not found"})
	}

	return c.JSON(http.StatusOK, userInfo)
}

func putUserInfo(c echo.Context) error {
	currentPass := c.FormValue("current_password")
	newPass := c.FormValue("new_password")
	userID := c.Param("userid")

	// セッション取得
	sess, err := session.Get("session", c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Status: 400, Message: "session not found"})
	}

	// セッションからユーザID取得
	sessionUserID, ok := sess.Values["userID"].(string)
	if ok == false {
		return c.JSON(http.StatusUnauthorized, Response{Status: 501, Message: "you should login"})
	}

	// 変更するユーザとログインしているユーザが異なる
	if userID != sessionUserID {
		return c.JSON(http.StatusForbidden, Response{Status: 300, Message: "you don't have permission"})
	}

	//ユーザ情報取得
	userInfo, err := DB.GetUserInfoByUserID(sessionUserID)
	if err != nil { // DBとsessionの整合性が取れてないとき？
		return c.JSON(http.StatusInternalServerError, Response{Status: 901, Message: "db error"})
	}

	//パスワード比較
	if err := bcrypt.CompareHashAndPassword([]byte(userInfo.HashedPassword), []byte(currentPass)); err != nil {
		return c.JSON(http.StatusUnauthorized, Response{Status: 500, Message: "login failed"})
	}

	//ハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(newPass), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Status: 900, Message: "server error"})
	}

	err = DB.UpdateUser(sessionUserID, string(hash))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Status: 901, Message: "db error"})
	}

	return c.JSON(http.StatusOK, Response{Status: 0, Message: "ok"})
}

func deleteUserInfo(c echo.Context) error {
	m := echo.Map{}
	if err := c.Bind(&m); err != nil {
		return err
	}
	password := m["password"].(string)

	if password == "" {
		return c.JSON(http.StatusBadRequest, Response{Status: 200, Message: "invalid paramater"})
	}

	// セッション取得
	sess, err := session.Get("session", c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Status: 400, Message: "session not found"})
	}

	// セッションからユーザIDを取得
	sessionUserID, ok := sess.Values["userID"].(string)
	if ok == false {
		return c.JSON(http.StatusUnauthorized, Response{Status: 501, Message: "you should login"})
	}

	// パスパラメータ取得
	userID := c.Param("userid")

	// 変更するユーザとログインしているユーザが異なる
	if userID != sessionUserID {
		return c.JSON(http.StatusForbidden, Response{Status: 300, Message: "you don't have permission"})
	}

	//ユーザ情報取得
	userInfo, err := DB.GetUserInfoByUserID(sessionUserID)
	if err != nil { // DBとsessionの整合性が取れてないとき？
		return c.JSON(http.StatusInternalServerError, Response{Status: 901, Message: "db error"})
	}

	//パスワード比較
	// TODO: 権限チェック，管理者なら一般ユーザを削除可能にするかも？
	if err := bcrypt.CompareHashAndPassword([]byte(userInfo.HashedPassword), []byte(password)); err != nil {
		return c.JSON(http.StatusUnauthorized, Response{Status: 502, Message: "wrong password"})
	}

	err = DB.DeleteUser(sessionUserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Status: 901, Message: "db error"})
	}

	// cookie削除
	if err := saveSession("", -1, c); err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Status: 401, Message: "cookie delete failed"})
	}

	return c.JSON(http.StatusOK, Response{Status: 0, Message: "ok"})
}

func createUser(c echo.Context) error {
	//パラメータ検査
	userID := c.FormValue("userid")
	password := c.FormValue("password")
	if userID == "" || password == "" ||
		!govalidator.IsAlphanumeric(userID) || !govalidator.IsByteLength(userID, 1, 128) {
		return c.JSON(http.StatusBadRequest, Response{Status: 200, Message: "invalid paramater"})
	}

	//ハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Status: 900, Message: "server error"})
	}

	err = DB.InsertUser(userID, string(hash))
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{Status: 102, Message: "user exist"})
	}

	//仮に有効期限1週間に設定
	if err = saveSession(userID, 60*60*24*7, c); err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Status: 402, Message: "cookie store failed"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"userId": userID})
}

func signin(c echo.Context) error {
	userID := c.FormValue("userid")
	password := c.FormValue("password")
	if userID == "" || password == "" {
		return c.JSON(http.StatusBadRequest, Response{Status: 200, Message: "invalid paramater"})
	}

	//ユーザ情報取得
	userInfo, err := DB.GetUserInfoByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, Response{Status: 500, Message: "login failed"})
	}

	//パスワード比較
	if err := bcrypt.CompareHashAndPassword([]byte(userInfo.HashedPassword), []byte(password)); err != nil {
		return c.JSON(http.StatusUnauthorized, Response{Status: 500, Message: "login failed"})
	}

	//ログイン成功 , セッション保存
	if err := saveSession(userID, 60*60*24*7, c); err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Status: 402, Message: "cookie store failed"})
	}

	return c.JSON(http.StatusOK, map[string]string{"userId": userID})
}

func signout(c echo.Context) error {
	// Cookieの削除
	if err := saveSession("", -1, c); err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Status: 401, Message: "cookie delete failed"})
	}

	return c.JSON(http.StatusOK, Response{Status: 0, Message: "ok"})
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
	BuildingID int64  `json:"buildingId"`
	Name       string `json:"name"`
	Latitude   string `json:"latitude"`
	Longitude  string `json:"longitude"`
	Lines      []Line `json:"lines"`
}

type Line struct {
	RailwayName    string `json:"railwayName"`
	StationID      int64  `json:"stationId"`
	OrderInRailWay int64  `json:"orderInRailway"`
}

type Response struct {
	Status  int64  `json:"status"`
	Message string `json:"message"`
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
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowCredentials: true,
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
	e.PUT("/users/:userid", putUserInfo)
	e.DELETE("/users/:userid", deleteUserInfo)

	e.POST("/signup", createUser)

	e.POST("/signin", signin)
	e.DELETE("/signin", signout)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

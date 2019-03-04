package main

import (
	"log"
	"net/http"
	"noritubusi-map/backend/app/station"
	"strconv"

	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getStationInfomationInRange(c echo.Context) error {
	beginLat := c.QueryParam("begin_latitude")
	beginLong := c.QueryParam("begin_longitude")
	endLat := c.QueryParam("end_latitude")
	endLong := c.QueryParam("end_longitude")

	if beginLat == "" || beginLong == "" || endLat == "" || endLong == "" {
		return c.String(http.StatusBadRequest, "invalid paramater")
	}

	stationInfo, err := stationInfoDB.GetStationInfoInRange(beginLat, beginLong, endLat, endLong)
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
	stationInfo, err := stationInfoDB.GetStationInfoByID(stationID)
	if err != nil {
		return c.String(http.StatusNotFound, "not found")
	}
	return c.JSON(http.StatusOK, stationInfo)
}

var (
	stationInfoDB station.StationDB
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

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// stationInfo DB connect
	err := stationInfoDB.New(stationInfoDBUserName, stationInfoDBPassword, stationInfoDBAddress, stationInfoDBName)
	if err != nil {
		e.Logger.Fatal("station DB Connection Error:", err)
	}

	// Routes
	e.GET("/", hello)
	e.GET("/stations", getStationInfomationInRange)
	e.GET("/stations/:stationid", getStationInfoByID)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
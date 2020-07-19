package db

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type BuildingInfo struct {
	Id                   int64  `db:"id" json:"id"`
	Name                 string `db:"name" json:"name"`
	Latitude             string `db:"lat" json:"latitude"`
	Longitude            string `db:"long" json:"longitude"`
	ConnectedRailwaysNum int64  `db:"connected_railways_num" json:"connectedRailwaysNum"`
}

type RailwayInfo struct {
	Id   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type StationInfo struct {
	BuildingId     int64  `db:"building_id" json:"buildingId"`
	StationId      int64  `db:"station_id" json:"stationId"`
	Name           string `db:"station_name" json:"name"`
	Latitude       string `db:"lat" json:"latitude"`
	Longitude      string `db:"long" json:"longitude"`
	RailwayName    string `db:"railway_line_name" json:"railwayLineName"`
	OrderInRailway int64  `db:"num_in_railway" json:"orderInRailway"`
}

func (s StationInfo) String() string {
	return fmt.Sprintf("[%4d, %4d]%s, (%s,%s),  %s ,%d", s.BuildingId, s.StationId, s.Name, s.Latitude, s.Longitude, s.RailwayName, s.OrderInRailway)
}

type DB struct {
	DB *sqlx.DB
}

func (d *DB) New(userName, password, address, dbName string) error {
	//userName:password@protocol(adress)/dbName
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", userName, password, address, dbName))

	if err != nil {
		log.Println("DB Connection Error:", err)
		return err
	}

	d.DB = db
	return nil
}

func (d *DB) GetBuildingInfoInRange(beginLat, beginLong, endLat, endLong string) ([]BuildingInfo, error) {
	query := `SELECT id,name,ST_X(latlong) AS 'lat',ST_Y(latlong) AS 'long' ,connected_railways_num FROM buildings WHERE MBRContains(ST_GeomFromText(CONCAT("LINESTRING(",?," ",?,",",?," ", ?,")")),latlong) ORDER BY id`

	infos := []BuildingInfo{}
	err := d.DB.Select(&infos, query, beginLat, beginLong, endLat, endLong)

	return infos, err
}

func (d *DB) GetStationInfoInRange(beginLat, beginLong, endLat, endLong string) ([]StationInfo, error) {
	query := `SELECT buildings.id AS building_id,stations.id AS station_id,stations.name AS station_name,ST_X(latlong) AS 'lat',ST_Y(latlong) AS 'long',railways.name AS railway_line_name ,num_in_railway FROM stations INNER JOIN railways on stations.railway_id=railways.id INNER JOIN buildings on stations.building_id=buildings.id WHERE MBRContains(ST_GeomFromText(CONCAT("LINESTRING(",?," ",?,",",?," ", ?,")")),latlong) ORDER BY building_id`

	infos := []StationInfo{}
	err := d.DB.Select(&infos, query, beginLat, beginLong, endLat, endLong)

	return infos, err
}

func (d *DB) GetStationInfoByID(id int) (StationInfo, error) {
	query := `SELECT buildings.id AS building_id,stations.id AS station_id,stations.name AS station_name,ST_X(latlong) AS 'lat',ST_Y(latlong) AS 'long',railways.name AS railway_line_name ,num_in_railway FROM stations INNER JOIN railways on stations.railway_id=railways.id INNER JOIN buildings on stations.building_id=buildings.id WHERE stations.id = ? ORDER BY building_id`

	var info StationInfo
	err := d.DB.Get(&info, query, id)

	// Getは何も存在しないとerrorを返すのでerrorチェックの必要がない
	return info, err
}

func (d *DB) GetStationsInfoByKeyword(keyword string) ([]StationInfo, error) {
	query := `SELECT buildings.id AS building_id,stations.id AS station_id,stations.name AS station_name,ST_X(latlong) AS 'lat',ST_Y(latlong) AS 'long',railways.name AS railway_line_name ,num_in_railway FROM stations INNER JOIN railways on stations.railway_id=railways.id INNER JOIN buildings on stations.building_id=buildings.id WHERE stations.name like concat("%",?,"%") ORDER BY building_id`

	suggestedStations := []StationInfo{}
	err := d.DB.Select(&suggestedStations, query, keyword)

	return suggestedStations, err
}

func (d *DB) GetStationInfoByBuildingID(buildingID int) ([]StationInfo, error) {
	query := `SELECT buildings.id AS building_id,stations.id AS station_id,stations.name AS station_name,ST_X(latlong) AS 'lat',ST_Y(latlong) AS 'long',railways.name AS railway_line_name ,num_in_railway FROM stations INNER JOIN railways on stations.railway_id=railways.id INNER JOIN buildings on stations.building_id=buildings.id WHERE buildings.id = ?  ORDER BY station_id`

	stationInfos := []StationInfo{}
	err := d.DB.Select(&stationInfos, query, buildingID)

	return stationInfos, err
}

func (d *DB) GetRailwaysInfoAll() ([]RailwayInfo, error) {
	query := `SELECT * FROM railways ORDER BY id`

	railways := []RailwayInfo{}
	err := d.DB.Select(&railways, query)

	return railways, err
}

func (d *DB) GetRailwaysInfoByID(id int) ([]RailwayInfo, error) {
	query := `SELECT * FROM railways WHERE id = ? ORDER BY id`

	railways := []RailwayInfo{}
	err := d.DB.Select(&railways, query, id)

	return railways, err
}

func (d *DB) GetRailwaysInfoByName(name string) ([]RailwayInfo, error) {
	query := `SELECT * FROM railways WHERE name = ? ORDER BY id`

	railways := []RailwayInfo{}
	err := d.DB.Select(&railways, query, name)

	return railways, err
}

func (d *DB) GetStationsInfoByRailwayID(id int) ([]StationInfo, error) {
	query := `SELECT buildings.id AS building_id,stations.id AS station_id,stations.name AS station_name,ST_X(latlong) AS 'lat',ST_Y(latlong) AS 'long',railways.name AS railway_line_name ,num_in_railway FROM stations INNER JOIN railways on stations.railway_id=railways.id INNER JOIN buildings on stations.building_id=buildings.id WHERE railways.id = ?  ORDER BY station_id`

	stations := []StationInfo{}
	err := d.DB.Select(&stations, query, id)

	return stations, err
}

func (d *DB) GetStationsInfoByRailwayName(name string) ([]StationInfo, error) {
	query := `SELECT buildings.id AS building_id,stations.id AS station_id,stations.name AS station_name,ST_X(latlong) AS 'lat',ST_Y(latlong) AS 'long',railways.name AS railway_line_name ,num_in_railway FROM stations INNER JOIN railways on stations.railway_id=railways.id INNER JOIN buildings on stations.building_id=buildings.id WHERE railways.name = ?  ORDER BY station_id`

	stations := []StationInfo{}
	err := d.DB.Select(&stations, query, name)

	return stations, err
}

package db

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type BuildingInfo struct {
	Id        int64  `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	Latitude  string `db:"lat" json:"latitude"`
	Longitude string `db:"long" json:"longitude"`
}

type RailwayInfo struct {
	Id                  int64  `db:"id" json:"id"`
	Name                string `db:"name" json:"name"`
	Type                int64  `db:"type" json:"type"`
	Company             string `db:"company_name" json:"company"`
	ServiceProviderType int64  `db:"service_provider_type" json:"serviceProviderType"`
}

type StationInfo struct {
	BuildingId          int64  `db:"building_id" json:"building_id"`
	StationId           int64  `db:"station_id" json:"station_id"`
	Name                string `db:"station_name" json:"name"`
	Latitude            string `db:"lat" json:"latitude"`
	Longitude           string `db:"long" json:"longitude"`
	Company             string `db:"company_name" json:"company"`
	ServiceProviderType int64  `db:"service_provider_type" json:"serviceProviderType"`
	RailwayName         string `db:"railway_line_name" json:"railwayName"`
}

func (s StationInfo) String() string {
	return fmt.Sprintf("[%4d, %4d]%s, (%s,%s), %s, %d, %s", s.BuildingId, s.StationId, s.Name, s.Latitude, s.Longitude, s.Company, s.ServiceProviderType, s.RailwayName)
}

type DB struct {
	DB *sqlx.DB
}

func (d *DB) New(userName, password, address, dbName string) error {
	//userName:password@protocol(adress)/dbName
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", userName, password, address, dbName))

	if err != nil {
		log.Println("DB Connection Error:", err)
		return err
	}

	d.DB = db
	return nil
}

func (d *DB) GetBuildingInfoInRange(beginLat, beginLong, endLat, endLong string) ([]BuildingInfo, error) {
	query := `SELECT id,name,ST_X(latlong) AS 'lat',ST_Y(latlong) AS 'long' FROM buildings WHERE MBRContains(ST_GeomFromText(CONCAT("LINESTRING(",?," ",?,",",?," ", ?,")")),latlong) ORDER BY id`

	infos := []BuildingInfo{}
	err := d.DB.Select(&infos, query, beginLat, beginLong, endLat, endLong)

	return infos, err
}

func (d *DB) GetStationInfoByID(id int) (StationInfo, error) {
	query := `SELECT buildings.id AS building_id,stations.id AS station_id,stations.name AS station_name,ST_X(latlong) AS 'lat',ST_Y(latlong) AS 'long',company_name,service_provider_type,railways.name AS railway_line_name FROM stations INNER JOIN railways on stations.railway_id=railways.id INNER JOIN buildings on stations.building_id=buildings.id WHERE stations.id = ? ORDER BY building_id`

	var info StationInfo
	err := d.DB.Get(&info, query, id)

	// Getは何も存在しないとerrorを返すのでerrorチェックの必要がない
	return info, err
}

func (d *DB) GetStationsInfoByKeyword(keyword string) ([]StationInfo, error) {
	query := `SELECT buildings.id AS building_id,stations.id AS station_id,stations.name AS station_name,ST_X(latlong) AS 'lat',ST_Y(latlong) AS 'long',company_name,service_provider_type,railways.name AS railway_line_name FROM stations INNER JOIN railways on stations.railway_id=railways.id INNER JOIN buildings on stations.building_id=buildings.id WHERE stations.name like concat("%",?,"%") ORDER BY building_id`

	suggestedStations := []StationInfo{}
	err := d.DB.Select(&suggestedStations, query, keyword)

	return suggestedStations, err
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

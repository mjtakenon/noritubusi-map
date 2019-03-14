package station

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type stationInfo struct {
	Id                  int64  `db:"id" json:"id"`
	Name                string `db:"station_name" json:"name"`
	Latitude            string `db:"lat" json:"latitude"`
	Longitude           string `db:"long" json:"longitude"`
	Company             string `db:"operation_company" json:"company"`
	ServiceProviderType int64  `db:"service_provider_type" json:"serviceProviderType"`
	RailwayName         string `db:"railway_line_name" json:"railwayName"`
	RailwayType         int64  `db:"railway_type" json:"railwayType"`
}

func (s stationInfo) String() string {
	return fmt.Sprintf("[%4d]%s, (%s,%s), %s, %d, %s, %d", s.Id, s.Name, s.Latitude, s.Longitude, s.Company, s.ServiceProviderType, s.RailwayName, s.RailwayType)
}

type StationDB struct {
	DB *sqlx.DB
}

func (s *StationDB) New(userName, password, address, dbName string) error {
	//userName:password@protocol(adress)/dbName
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", userName, password, address, dbName))

	if err != nil {
		log.Println("DB Connection Error:", err)
		return err
	}

	s.DB = db
	return nil
}

func (s *StationDB) GetStationInfoInRange(beginLat, beginLong, endLat, endLong string) ([]stationInfo, error) {
	query := `select id,station_name,ST_X(center_latlong) AS 'lat',ST_Y(center_latlong) AS 'long',operation_company,service_provider_type,railway_line_name,railway_type from stations where MBRContains(GeomFromText(CONCAT("LINESTRING(",?," ",?,",",?," ", ?,")")),center_latlong) order by id`

	infos := []stationInfo{}
	err := s.DB.Select(&infos, query, beginLat, beginLong, endLat, endLong)
	if err != nil {
		return nil, err
	}

	return infos, nil
}

func (s *StationDB) GetStationInfoByID(id int) (stationInfo, error) {
	query := `select id,station_name,ST_X(center_latlong) AS 'lat',ST_Y(center_latlong) AS 'long',operation_company,service_provider_type,railway_line_name,railway_type from stations where id = ? order by id`

	var info stationInfo
	err := s.DB.Get(&info, query, id)

	// Getは何も存在しないとerrorを返すのでerrorチェックの必要がない
	return info, err
}

func (s *StationDB) GetStationsInfoByKeyword(keyword string) ([]stationInfo, error) {
	query := `select id,station_name,ST_X(center_latlong) AS 'lat',ST_Y(center_latlong) AS 'long',operation_company,service_provider_type,railway_line_name,railway_type from stations where station_name like concat("%",?,"%") order by id;`

	suggestedStations := []stationInfo{}
	err := s.DB.Select(&suggestedStations, query, keyword)
	if err != nil {
		return nil, err
	}
	return suggestedStations, nil
}

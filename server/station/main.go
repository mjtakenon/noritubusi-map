package station

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type stationInfo struct {
	Id                  int64  `db:"id"`
	Name                string `db:"station_name"`
	Latitude            string `db:"lat"`
	Longitude           string `db:"long"`
	Company             string `db:"operation_company"`
	ServiceProviderType int64  `db:"service_provider_type"`
	RailwayName         string `db:"railway_line_name"`
	RailwayType         int64  `db:"railway_type"`
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

func (s *StationDB) getStationInfoInRange(beginLat, beginLong, endLat, endLong string) ([]stationInfo, error) {
	query := `select id,station_name,X(center_latlong) AS 'lat',Y(center_latlong) AS 'long',operation_company,service_provider_type,railway_line_name,railway_type from stations where MBRContains(GeomFromText(CONCAT("LINESTRING(",?," ",?,",",?," ", ?,")")),center_latlong)`

	infos := []stationInfo{}
	err := s.DB.Select(&infos, query, beginLat, beginLong, endLat, endLong)
	if err != nil {
		return nil, err
	}

	return infos, nil
}

package station

import (
	"fmt"
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func TestStationData_New(t *testing.T) {
	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		userName string
		password string
		address  string
		dbName   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "success Case",
			fields: fields{DB: nil},
			args: args{
				userName: "user",
				password: "password",
				address:  "localhost:3314",
				dbName:   "noritubusi_map",
			},
			wantErr: false,
		},
		{
			name:   "Error Case",
			fields: fields{DB: nil},
			args: args{
				userName: "go",
				password: "go",
				address:  "localhost:765",
				dbName:   "go",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StationDB{
				DB: tt.fields.DB,
			}
			if err := s.New(tt.args.userName, tt.args.password, tt.args.address, tt.args.dbName); (err != nil) != tt.wantErr {
				t.Errorf("StationData.New() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStationDB_GetStationInfoInRange(t *testing.T) {
	type fields struct {
		DB *sqlx.DB
	}

	//DBセットアップ
	userName := "user"
	password := "password"
	address := "localhost:3314"
	dbName := "noritubusi_map"
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", userName, password, address, dbName))

	if err != nil {
		t.Errorf("DB Connection Error:%v", err)
		return
	}
	type args struct {
		beginLat  string
		beginLong string
		endLat    string
		endLong   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []stationInfo
		wantErr bool
	}{
		{
			name:   "near Hamamatsu Station",
			fields: fields{DB: db},
			args: args{
				beginLat:  "34.705549",
				beginLong: "137.729265",
				endLat:    "34.702267",
				endLong:   "137.741667",
			},
			want: []stationInfo{
				{
					Id:                  437,
					Name:                "新浜松",
					Latitude:            "34.70341",
					Longitude:           "137.73246",
					Company:             "遠州鉄道",
					ServiceProviderType: 4,
					RailwayName:         "鉄道線",
					RailwayType:         12,
				},
				{
					Id:                  5796,
					Name:                "浜松",
					Latitude:            "34.70376",
					Longitude:           "137.7353775",
					Company:             "東海旅客鉄道",
					ServiceProviderType: 1,
					RailwayName:         "東海道新幹線",
					RailwayType:         11,
				},
				{
					Id:                  5871,
					Name:                "浜松",
					Latitude:            "34.70406",
					Longitude:           "137.7351175",
					Company:             "東海旅客鉄道",
					ServiceProviderType: 2,
					RailwayName:         "東海道線",
					RailwayType:         11,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StationDB{
				DB: tt.fields.DB,
			}
			got, err := s.GetStationInfoInRange(tt.args.beginLat, tt.args.beginLong, tt.args.endLat, tt.args.endLong)
			if (err != nil) != tt.wantErr {
				t.Errorf("StationDB.getStationInfoInRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StationDB.getStationInfoInRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

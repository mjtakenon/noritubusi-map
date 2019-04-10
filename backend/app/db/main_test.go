package db

import (
	"fmt"
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func SetUpDB(t *testing.T) *sqlx.DB {
	//DBセットアップ
	userName := "user"
	password := "password"
	address := "localhost:3314"
	dbName := "noritubusi_map"
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", userName, password, address, dbName))

	if err != nil {
		t.Errorf("DB Connection Error:%v", err)
		return nil
	}
	return db
}

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
			s := DB{
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
		want    []StationInfo
		wantErr bool
	}{
		{
			name:   "near Hamamatsu Station",
			fields: fields{DB: SetUpDB(t)},
			args: args{
				beginLat:  "34.705549",
				beginLong: "137.729265",
				endLat:    "34.702267",
				endLong:   "137.741667",
			},
			want: []StationInfo{
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
			s := &DB{
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

func TestStationDB_GetStationInfoByID(t *testing.T) {
	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    StationInfo
		wantErr bool
	}{
		{
			name:   "Shin-Hamamatsu(No.437)",
			fields: fields{DB: SetUpDB(t)},
			args:   args{id: 437},
			want: StationInfo{
				Id:                  437,
				Name:                "新浜松",
				Latitude:            "34.70341",
				Longitude:           "137.73246",
				Company:             "遠州鉄道",
				ServiceProviderType: 4,
				RailwayName:         "鉄道線",
				RailwayType:         12,
			},
			wantErr: false,
		},
		{
			name:    "not found id number",
			fields:  fields{DB: SetUpDB(t)},
			args:    args{id: 12345},
			want:    StationInfo{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &DB{
				DB: tt.fields.DB,
			}
			got, err := s.GetStationInfoByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("StationDB.GetStationInfoByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StationDB.GetStationInfoByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStationDB_GetStationsInfoByKeyword(t *testing.T) {
	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		keyword string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []StationInfo
		wantErr bool
	}{
		{
			name:   "%浜松%",
			fields: fields{DB: SetUpDB(t)},
			args: args{
				keyword: "浜松",
			},
			want: []StationInfo{
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
				{
					Id:                  6005,
					Name:                "モノレール浜松町",
					Latitude:            "35.655745",
					Longitude:           "139.75667",
					Company:             "東京モノレール",
					ServiceProviderType: 4,
					RailwayName:         "東京モノレール羽田線",
					RailwayType:         15,
				},
				{
					Id:                  7708,
					Name:                "浜松町",
					Latitude:            "35.65541",
					Longitude:           "139.757125",
					Company:             "東日本旅客鉄道",
					ServiceProviderType: 2,
					RailwayName:         "東海道線",
					RailwayType:         11,
				},
			},
			wantErr: false,
		},
		{
			name:   "error case",
			fields: fields{DB: SetUpDB(t)},
			args: args{
				keyword: "ああああ",
			},
			want:    []StationInfo{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &DB{
				DB: tt.fields.DB,
			}
			got, err := s.GetStationsInfoByKeyword(tt.args.keyword)
			if (err != nil) != tt.wantErr {
				t.Errorf("StationDB.GetStationsInfoByKeyword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StationDB.GetStationsInfoByKeyword() = %v, want %v", got, tt.want)
			}
		})
	}
}

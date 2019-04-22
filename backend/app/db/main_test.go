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
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", userName, password, address, dbName))

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

func TestStationDB_GetBuildingInfoInRange(t *testing.T) {
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
		want    []BuildingInfo
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
			want: []BuildingInfo{
				{
					Id:        423,
					Name:      "新浜松",
					Latitude:  "34.70341",
					Longitude: "137.73246",
				},
				{
					Id:        5297,
					Name:      "浜松",
					Latitude:  "34.70376",
					Longitude: "137.7353775",
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
			got, err := s.GetBuildingInfoInRange(tt.args.beginLat, tt.args.beginLong, tt.args.endLat, tt.args.endLong)
			if (err != nil) != tt.wantErr {
				t.Errorf("StationDB.getBuildingInfoInRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StationDB.getBuildingInfoInRange() = %v, want %v", got, tt.want)
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
				BuildingId:          423,
				StationId:           437,
				Name:                "新浜松",
				Latitude:            "34.70341",
				Longitude:           "137.73246",
				Company:             "遠州鉄道",
				ServiceProviderType: 4,
				RailwayName:         "鉄道線",
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
					BuildingId:          423,
					StationId:           437,
					Name:                "新浜松",
					Latitude:            "34.70341",
					Longitude:           "137.73246",
					Company:             "遠州鉄道",
					ServiceProviderType: 4,
					RailwayName:         "鉄道線",
				},
				{
					BuildingId:          5297,
					StationId:           5796,
					Name:                "浜松",
					Latitude:            "34.70376",
					Longitude:           "137.7353775",
					Company:             "東海旅客鉄道",
					ServiceProviderType: 1,
					RailwayName:         "東海道新幹線",
				},
				{
					BuildingId:          5297,
					StationId:           5871,
					Name:                "浜松",
					Latitude:            "34.70376",
					Longitude:           "137.7353775",
					Company:             "東海旅客鉄道",
					ServiceProviderType: 2,
					RailwayName:         "東海道線",
				},
				{
					BuildingId:          5479,
					StationId:           6005,
					Name:                "モノレール浜松町",
					Latitude:            "35.655745",
					Longitude:           "139.75667",
					Company:             "東京モノレール",
					ServiceProviderType: 4,
					RailwayName:         "東京モノレール羽田線",
				},
				{
					BuildingId:          6908,
					StationId:           7708,
					Name:                "浜松町",
					Latitude:            "35.65541",
					Longitude:           "139.757125",
					Company:             "東日本旅客鉄道",
					ServiceProviderType: 2,
					RailwayName:         "東海道線",
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

func TestDB_GetStationInfoByBuildingID(t *testing.T) {
	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		buildingID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []StationInfo
		wantErr bool
	}{
		{
			name:   "浜松駅",
			fields: fields{DB: SetUpDB(t)},
			args:   args{buildingID: 5297},
			want: []StationInfo{
				{
					BuildingId:          5297,
					StationId:           5796,
					Name:                "浜松",
					Latitude:            "34.70376",
					Longitude:           "137.7353775",
					Company:             "東海旅客鉄道",
					ServiceProviderType: 1,
					RailwayName:         "東海道新幹線",
				},
				{
					BuildingId:          5297,
					StationId:           5871,
					Name:                "浜松",
					Latitude:            "34.70376",
					Longitude:           "137.7353775",
					Company:             "東海旅客鉄道",
					ServiceProviderType: 2,
					RailwayName:         "東海道線",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DB{
				DB: tt.fields.DB,
			}
			got, err := d.GetStationInfoByBuildingID(tt.args.buildingID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DB.GetStationInfoByBuildingID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.GetStationInfoByBuildingID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_GetRailwaysInfoAll(t *testing.T) {
	type fields struct {
		DB *sqlx.DB
	}

	tt := struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		name:    "Number of railways info",
		fields:  fields{DB: SetUpDB(t)},
		want:    605,
		wantErr: false,
	}

	t.Run(tt.name, func(t *testing.T) {
		d := &DB{
			DB: tt.fields.DB,
		}
		got, err := d.GetRailwaysInfoAll()
		if (err != nil) != tt.wantErr {
			t.Errorf("DB.GetRailwaysInfoAll() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(len(got), tt.want) {
			t.Errorf("len(DB.GetRailwaysInfoAll()) = %d, want %d", len(got), tt.want)
		}
	})
}

func TestDB_GetRailwaysInfoByID(t *testing.T) {
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
		want    []RailwayInfo
		wantErr bool
	}{
		{
			name:   "Odakyu Odawara Line",
			fields: fields{DB: SetUpDB(t)},
			args: args{
				id: 214,
			},
			want: []RailwayInfo{
				{
					Id:                  214,
					Name:                "小田原線",
					Type:                12,
					Company:             "小田急電鉄",
					ServiceProviderType: 4,
				},
			},
			wantErr: false,
		},
		{
			name:   "error case",
			fields: fields{DB: SetUpDB(t)},
			args: args{
				id: 999,
			},
			want:    []RailwayInfo{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DB{
				DB: tt.fields.DB,
			}
			got, err := d.GetRailwaysInfoByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DB.GetRailwaysInfoByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.GetRailwaysInfoByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_GetRailwaysInfoByName(t *testing.T) {
	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []RailwayInfo
		wantErr bool
	}{
		{
			name:   "Odakyu Odawara Line",
			fields: fields{DB: SetUpDB(t)},
			args: args{
				name: "小田原線",
			},
			want: []RailwayInfo{
				{
					Id:                  214,
					Name:                "小田原線",
					Type:                12,
					Company:             "小田急電鉄",
					ServiceProviderType: 4,
				},
			},
			wantErr: false,
		},
		{
			name:   "error case",
			fields: fields{DB: SetUpDB(t)},
			args: args{
				name: "ああああ",
			},
			want:    []RailwayInfo{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DB{
				DB: tt.fields.DB,
			}
			got, err := d.GetRailwaysInfoByName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("DB.GetRailwaysInfoByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.GetRailwaysInfoByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

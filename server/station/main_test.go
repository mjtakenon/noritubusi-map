package station

import (
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
			s := StationData{
				DB: tt.fields.DB,
			}
			if err := s.New(tt.args.userName, tt.args.password, tt.args.address, tt.args.dbName); (err != nil) != tt.wantErr {
				t.Errorf("StationData.New() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

package db

import (
	"reflect"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
)

func TestDB_insertUser(t *testing.T) {
	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		userID         string
		hashedPassword string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:   "success",
			fields: fields{DB: SetUpDB(t)},
			args: args{
				userID:         "john",
				hashedPassword: "$2a$10$gjfXvNjweHaIl9G0Kyc7w.Dg.tRCJKexpDVWlKijxnoFiAVSGd15S", //password:john
			},
			wantErr: false,
		},
		{
			name:   "failed",
			fields: fields{DB: SetUpDB(t)},
			args: args{
				userID:         "mjtakenon",
				hashedPassword: "$2a$10$TY73Cx/B978COKuT9X0Kvubq7HupRFX7NoFtp0O3pKKgUJ9YAP.WC",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DB{
				DB: tt.fields.DB,
			}
			err := d.InsertUser(tt.args.userID, tt.args.hashedPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("DB.InsertUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
	//テストデータ削除
	SetUpDB(t).Exec(`DELETE FROM users WHERE id = ?`, tests[0].args.userID)
}

func TestDB_GetUserInfoByUserID(t *testing.T) {
	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		userID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    UserInfo
		wantErr bool
	}{
		{
			name:   "mjtakenon",
			fields: fields{SetUpDB(t)},
			args:   args{userID: "mjtakenon"},
			want: UserInfo{
				Id:             "mjtakenon",
				HashedPassword: "$2a$10$uy.XzaOpSaPVPTCo6PW6k.C3x9mB9ZIrpiotuRwflR3JYzXIEeovy",
				CreateTime:     time.Date(2019, 4, 1, 0, 0, 0, 0, time.UTC),
				ChangeTime:     time.Date(2019, 4, 1, 0, 0, 0, 0, time.UTC),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DB{
				DB: tt.fields.DB,
			}
			got, err := d.GetUserInfoByUserID(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DB.GetUserInfoByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.GetUserInfoByUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_UpdateUser(t *testing.T) {
	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		userID            string
		newHashedPassword string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "success",
			fields: fields{DB: SetUpDB(t)},
			args: args{
				userID:            "john",
				newHashedPassword: "$2a$10$gjfXvNjweHaIl9G0Kyc7w.Dg.tRCJKexpDVWlKijxnoFiAVSGd15S", //password:john
			},
			wantErr: false,
		},
	}
	//テストデータ生成
	SetUpDB(t).Exec(`INSERT INTO users (id,hashed_password) VALUES (?,?)`, tests[0].args.userID, "$2a$10$uy.XzaOpSaPVPTCo6PW6k.C3x9mB9ZIrpiotuRwflR3JYzXIEeovy")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DB{
				DB: tt.fields.DB,
			}
			if err := d.UpdateUser(tt.args.userID, tt.args.newHashedPassword); (err != nil) != tt.wantErr {
				t.Errorf("DB.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	//テストデータ削除
	SetUpDB(t).Exec(`DELETE FROM users WHERE id = ?`, tests[0].args.userID)
}

func TestDB_DeleteUser(t *testing.T) {
	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		userID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "success",
			fields: fields{DB: SetUpDB(t)},
			args: args{
				userID: "john",
			},
			wantErr: false,
		},
	}
	//テストデータ生成
	SetUpDB(t).Exec(`INSERT INTO users (id,hashed_password) VALUES (?,?)`, tests[0].args.userID, "$2a$10$uy.XzaOpSaPVPTCo6PW6k.C3x9mB9ZIrpiotuRwflR3JYzXIEeovy")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DB{
				DB: tt.fields.DB,
			}
			if err := d.DeleteUser(tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("DB.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

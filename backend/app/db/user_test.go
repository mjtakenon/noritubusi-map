package db

import (
	"testing"

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

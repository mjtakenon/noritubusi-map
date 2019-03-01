package station

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type StationData struct {
	DB *sqlx.DB
}

func (s StationData) New(userName, password, address, dbName string) error {
	//userName:password@protocol(adress)/dbName
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", userName, password, address, dbName))

	if err != nil {
		log.Println("DB Connection Error:", err)
		return err
	}

	s.DB = db
	return nil
}

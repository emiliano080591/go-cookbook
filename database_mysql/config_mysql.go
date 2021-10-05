package database_mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

type Example struct {
	Name string
	Created *time.Time
}

func Setup() (*sql.DB,error) {
	dns:=fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",os.Getenv("USER"),os.Getenv("PASSWORD"),
		os.Getenv("HOST"),os.Getenv("DB"))
	db,err:=sql.Open("mysql", dns)
	if err!=nil {
		return nil, err
	}

	db.SetMaxOpenConns(15)
	db.SetMaxIdleConns(15)

	return db,nil
}

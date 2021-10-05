package database

import (
	"database/sql"
	"fmt"
)

func Query(db *sql.DB) error {
	name:="Emiliano"
	rows,err:=db.Query("SELECT name,created FROM example WHERE name=?",name)
	if err != nil {
		return err
	}
	defer db.Close()
	for rows.Next() {
		var e Example
		if err:=rows.Scan(&e.Name,&e.Created);err != nil {
			return err
		}
		fmt.Printf("Results:\n\tName:%s\n\tCreated: %v\n",e.Name,e.Created)
	}
	return rows.Err()
}

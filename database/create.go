package database

import "database/sql"

func Create(db *sql.DB) error {
	if _,err:=db.Exec("CREATE TABLE example(name VARCHAR(20),created DATETIME)");err!=nil{
		return err
	}
	if _,err:=db.Exec(`INSERT INTO example(name,created) VALUES("Emiliano",NOW())`);err!=nil{
		return err
	}
	return nil
}

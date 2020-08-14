package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/lotus703/SJob/db"
)

func main() {
	sql := &db.Sql{
		//dbDDb       *sqlx.DB
		Driver: "mysql",
		User:   "root",
		Pass:   "A212796619",
		DbName: "mydb",
	}
	sql.Connect()
	defer sql.Close()
}

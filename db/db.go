package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

type Sql struct {
	Db     *sqlx.DB
	Driver string
	User   string
	Pass   string
	DbName string
}

func (s *Sql) Connect() {
	connectString := fmt.Sprintf("%s:%s@/%s", s.User, s.Pass, s.DbName)
	sqlx.Open(s.Driver, connectString)
	if err := s.Db.Ping(); err != nil {
		log.Error(err.Error())
		return
	}
	fmt.Println("Connect Database OK")
}
func (s *Sql) Close() {
	s.Db.Close()
}

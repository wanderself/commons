package sqls

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

type Conns struct {
	DB *sql.DB
}

func (c *Conns) NewDataSource(user, pwd, host, database string, MaxOpenConns, MaxIdleConns int) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", user, pwd, host, database)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	db.SetMaxOpenConns(MaxOpenConns)
	db.SetMaxIdleConns(MaxIdleConns)

	if ping := db.Ping(); ping != nil {
		log.Println(ping.Error())
		return ping
	}

	c.DB = db

	return nil
}

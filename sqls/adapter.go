package sqls

import (
	"fmt"
	"database/sql"
	"github.com/wanderself/commons/errs"
	_ "github.com/go-sql-driver/mysql"
)

type Conns struct {
	DB *sql.DB
}

func (c *Conns) NewDataSource(user, pwd, host, database string, MaxOpenConns, MaxIdleConns int) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", user, pwd, host, database)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		errs.ErrCheck(err, "NewDataSource sql.Open: ")
		return err
	}

	db.SetMaxOpenConns(MaxOpenConns)
	db.SetMaxIdleConns(MaxIdleConns)

	c.DB = db

	return nil
}

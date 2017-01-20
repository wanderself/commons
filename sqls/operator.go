package sqls

import (
	"github.com/wanderself/commons/conf"
	"github.com/wanderself/commons/errs"
	"log"
)

var conn *Conns

func Connect(user, pwd, host, database string, MaxOpenConns, MaxIdleConns int) error {
	conn = new(Conns)
	if err := conn.NewDataSource(user, pwd, host, database, MaxOpenConns, MaxIdleConns); err != nil {
		errs.ErrCheck(err, "Connect() conn.NewDataSource: ")
		return err
	}

	if err := conn.DB.Ping(); err != nil {
		errs.ErrCheck(err, "Connect() conn.DB.Ping: ")
		return err
	} else {
		log.Println("mysql connection established!	Database:", conf.GetConf().Database)
		return nil
	}
}

func Ins(sql string, args ...interface{}) error {
	stmtIns, err := conn.DB.Prepare(sql)
	if err != nil {
		errs.ErrCheck(err, "Ins() conn.DB.Prepare: ")
		return err
	}
	defer stmtIns.Close()

	if _, err := stmtIns.Exec(args ...); err != nil {
		errs.ErrCheck(err, "Ins() stmtIns.Exec: ")
		return err
	}

	return nil
}

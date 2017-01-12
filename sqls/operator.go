package sqls

import (
	"github.com/wanderself/commons/conf"
	"log"
)

var conn *Conns

func Connect(user, pwd, host, database string, MaxOpenConns, MaxIdleConns int) error {
	conn = new(Conns)
	if err := conn.NewDataSource(user, pwd, host, database, MaxOpenConns, MaxIdleConns); err != nil {
		log.Println(err.Error())
		return err
	}

	if err := conn.DB.Ping(); err != nil {
		log.Println(err.Error())
		return err
	} else {
		log.Println("mysql connection established!	Database:", conf.GetConf().Database)
		return nil
	}
}

package conf

import "wander/test/sqlConner/common/errs"



package conf

import (
"github.com/wanderself/commons/errs"
)

type Config struct {
	Host         string
	Database     string
	Username     string
	Password     string
	MaxOpenConns int
	MaxIdleConns int
}

var (
	conf *Config
	scaner *config.Config
)

func init() {
	var err error
	scaner, err = config.ReadDefault("config.conf")
	errs.ErrCheck(err, "configReaderUnexpected exception be thrown.")

	conf = new(Config)
	conf.Host = getStringDefault("base", "host", "localhost:3306")
	conf.Database = getStringDefault("base", "database", "push")
	conf.Username = getStringDefault("base", "user", "root")
	conf.Password = getStringDefault("base", "passwd", "")
	conf.MaxOpenConns = getIntDefault("base", "MaxOpenConns", 10)
	conf.MaxIdleConns = getIntDefault("base", "MaxIdleConns", 1)
}

func GetConf() *Config {
	if conf == nil {
		conf = new(Config)
	}
	return conf
}

func getStringDefault(section, option, defaltVal string) string {
	val, err := scaner.String(section, option)
	if err != nil {
		return defaltVal
	} else {
		return val
	}
}

func getIntDefault(section, option string, defaultVal int) int {
	val, err := scaner.Int(section, option)
	if err != nil {
		return defaultVal
	} else {
		return val
	}

}



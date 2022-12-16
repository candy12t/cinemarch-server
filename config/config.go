package config

import (
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

func DSN() string {
	conf := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		DBName:               os.Getenv("DB_DATABASE"),
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	return conf.FormatDSN()
}

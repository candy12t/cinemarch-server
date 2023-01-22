package config

import (
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

func DSN() string {
	conf := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		DBName:               os.Getenv("DB_DATABASE"),
		Collation:            "utf8mb4_general_ci",
		Loc:                  time.UTC,
		MaxAllowedPacket:     4 << 20,
		AllowNativePasswords: true,
		CheckConnLiveness:    true,
		ParseTime:            true,
	}
	return conf.FormatDSN()
}

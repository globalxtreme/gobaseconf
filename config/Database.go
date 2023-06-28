package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConf struct {
	Host      string
	Port      any
	Username  string
	Password  string
	Database  string
	Charset   string
	ParseTime bool
	Loc       string
	Collation string
}

func Connect(conn DBConf) *gorm.DB {
	option := "?"

	if len(conn.Charset) > 0 {
		option += "charset=" + conn.Charset
	} else {
		option += "charset=utf8mb4"
	}

	if conn.ParseTime {
		option += "&parseTime=True"
	} else {
		option += "&parseTime=False"
	}

	if len(conn.Loc) > 0 {
		option += "&loc=" + conn.Loc
	} else {
		option += "&loc=Local"
	}

	if len(conn.Collation) > 0 {
		option += "&collation=" + conn.Collation
	} else {
		option += "&collation=utf8mb4_unicode_ci"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%s",
		conn.Username, conn.Password, conn.Host, conn.Port, conn.Database, option)
	driver, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return driver
}

func SetMigration(conn *gorm.DB) *gorm.DB {
	return conn.Set("gorm:table_options", "COLLATE=utf8mb4_unicode_ci")
}

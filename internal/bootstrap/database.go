package bootstrap

import (
	"database/sql"
	"fmt"

	"github.com/go-gorp/gorp/v3"
	_ "github.com/go-sql-driver/mysql" // defines that I use mysql driver
)

type DBOption struct {
	Host                 string `env:"DB_HOST"`
	Port                 int    `env:"DB_PORT"`
	Username             string `env:"DB_USERNAME"`
	Password             string `env:"DB_PASSWORD"`
	DBName               string `env:"DB_NAME"`
	AdditionalParameters string `env:"DB_ADDITIONAL_PARAMS"`
}

// NewMysqlDatabase return gorp dbmap object with MySQL options param
func NewMysqlDatabase(option DBOption) (*gorp.DbMap, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", option.Username, option.Password, option.Host, option.Port, option.DBName, option.AdditionalParameters))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	gorp := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}
	return gorp, nil
}

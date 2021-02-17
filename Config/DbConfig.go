package Config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type DBConfig struct {
	user string
	pass string
	host string
	port string
	db   string

}

/**
	private constructor to keep singleton this struct
 */
func newDBConfig(user string, pass string, host string, port string, db string) *DBConfig {
	return &DBConfig{user: user, pass: pass, host: host, port: port, db: db}
}



var config *DBConfig = nil


func CreateDbConfig(params map[string]string) *DBConfig {
	if config == nil {
		config = newDBConfig(
			params["db_config.user"],
			params["db_config.password"],
			params["db_config.host"],
			params["db_config.port"],
			params["db_config.db"])
	}
	return config
}


func GetConnection() (*sql.DB, error) {
	var sb strings.Builder
	sb.WriteString(config.user)
	sb.WriteString(":")
	sb.WriteString(config.pass)
	sb.WriteString("@tcp(")
	sb.WriteString(config.host)
	sb.WriteString(":")
	sb.WriteString(config.port)
	sb.WriteString(")/")
	sb.WriteString(config.db)
	var connectionUrl = sb.String()
	db, err := sql.Open("mysql", connectionUrl)
	return db,err

}

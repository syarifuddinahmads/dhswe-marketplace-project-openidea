package db

import (
	"github.com/jmoiron/sqlx"
	"sync"
	// "gorm.io/gorm"
)

var (
	// dbConn *gorm.DB
	dbConn *sqlx.DB
	// we use sync.Once for make sure we create connection only once
	once sync.Once
)

// CreateConnection is a function for creating new connection with database
// you can choose you want use mysql or postgresql
func CreateConnection() {

	cnf, err := NewParsedConfig()
	if err != nil {
		return
	}

	conf := dbConfig{
		Host: cnf.Database.Host,
		Port: cnf.Database.Port,
		User: cnf.Database.User,
		Pass: cnf.Database.Password,
		Name: cnf.Database.Name,
	}

	// mysql := mysqlConfig{dbConfig: conf}
	// if you use postgres, you can uncomment code bellow

	postgres := postgresqlConfig{dbConfig: conf}

	once.Do(func() {
		// mysql.Connect()
		postgres.Connect()

	})
}

// GetConnection is a faction for return connection or return value dbConn
// because we set var dbConn is private
func GetConnection() *sqlx.DB {
	if dbConn == nil {
		CreateConnection()
	}
	return dbConn
}

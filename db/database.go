package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func (conf dbConfig) ConnectDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", conf.Host, conf.User, conf.Pass, conf.Name, conf.Port)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CreateConnectionDB() {
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

	once.Do(func() {
		var err error
		dbConn, err = conf.ConnectDB()
		if err != nil {
			panic(err)
		}
	})
}

func GetConnectionDB() *sql.DB {
	if dbConn == nil {
		CreateConnectionDB()
	}
	return dbConn
}

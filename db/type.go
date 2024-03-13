package db

import (
	"database/sql"
	"sync"

	_ "github.com/lib/pq"
)

var (
	dbConn *sql.DB
	once   sync.Once
)

type dbConfig struct {
	Host string
	User string
	Pass string
	Port int
	Name string
}

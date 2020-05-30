package db

import (
	"time"
	"github.com/go-pg/pg"
	"github.com/jeroldleslie/golang-web-scraper/api/lib"
)

var db *pg.DB

func GetDB() *pg.DB {
	if db == nil {
		db = pg.Connect(&pg.Options{
			Database:              lib.GetEnv("DB_DATABASE","cityfalcon"),
			User:                  lib.GetEnv("DB_USER","cityfalcon"),
			Password:              lib.GetEnv("DB_PASSWORD","cityfalcon"),
			Addr:                  lib.GetEnv("POSTGRESS_ADDRESS","localhost:5432"),
			RetryStatementTimeout: true,
			MaxRetries:            4,
			MinRetryBackoff:       250 * time.Millisecond,
		})
	}
	//defer db.Close()
	return db
}

package initialize

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/LeMinh0706/fakeapi/config"
	_ "github.com/lib/pq"
)

func InitDb() *sql.DB {

	connStr := config.Cfg.Databases.Source.GetPostgresDSN()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database!")
	return db
}

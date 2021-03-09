package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB
	err    error
	Db     *sql.DB
)

func InitDb() {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		"root", "Krish@knight8", "127.0.0.1:3306", "Graphql_db",
	)
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")

	Db = Client
}

/*
func Migrate() {

	if err := Db.Ping(); err != nil {
		log.Fatal(err)
	}
	//fmt.Println(Db)
	driver, _ := mysql.WithInstance(Db, &mysql.Config{})
	//fmt.Println(driver)
	m, _ := migrate.NewWithDatabaseInstance(
		"file://internal/pkg/db/database",
		"mysql",
		driver,
	)
	//fmt.Println(m)

	err := m.Up()

	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

}
*/

package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// https://github.com/golang-migrate/migrate

func main() {
	fmt.Println("main")
	POSTGRES_USER := "aaaa"
	POSTGRES_PASSWORD := "aaaa"
	POSTGRES_DB := "aaaa"
	db, err1 := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@localhost:25432/%s?sslmode=disable", POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB))
	if err1 != nil {
		fmt.Println("Error opening database: ", err1)
		panic(err1)
	}
	driver, err2 := postgres.WithInstance(db, &postgres.Config{})
	if err2 != nil {
		fmt.Println("Error WithInstance: err2", err2)
		panic(err2)
	}

	currentDir, err4 := os.Getwd()
	if err4 != nil {
		fmt.Println("Error getting current directory:", err4)
	} else {
		fmt.Println("Current directory:", currentDir)
	}

	m, err3 := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s/ddl", currentDir), "postgres", driver)
	if err3 != nil {
		fmt.Println("Error NewWithDatabaseInstance: err3  ", err3)
		panic(err3)
	}
	m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run
	// m.Down()
	fmt.Println("finish successfuly")
	// m.Step(1)
}

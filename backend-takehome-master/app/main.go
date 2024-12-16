package main

import (
	"app/handler"
	"app/provider"
	"app/service"
	"app/usecase"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Server is running on http://localhost:8080")
	db := initDB()
	defer db.Close()
	provider := &provider.Provider{
		DB: db,
	}
	usecase := &usecase.Usecase{
		Provider: provider,
	}
	service := &service.Service{
		Usecase: usecase,
	}
	handler := &handler.Handler{
		Service: service,
	}

	provider.InitProvider()
	handler.HandleRequest()
	http.ListenAndServe(":8080", nil)
}

func initDB() *sql.DB {
	dsn := "root:abc123@tcp(db:3306)/appdb"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS maindb")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("USE maindb")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

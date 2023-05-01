package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type employee struct {
	ID      int
	Ename   string
	Esalary int
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "WalkingDreamzdb"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

// func dbConn() (db *sql.DB) {
// 	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/forestdb")
// 	if err != nil {
// 		panic(err)
// 	}
// 	return db
// }

func CreateDatabase(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	_, err := db.Exec("CREATE DATABASE WalkingDreamzdb")
	if err != nil {
		panic(err)
	}
}

func CreateTable(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	_, err := db.Exec("CREATE TABLE Employee(ID INT PRIMARY KEY NOT NULL AUTO_INCREMENT, Ename VARCHAR(50), Esalary INT))")
	if err != nil {
		panic(err)
	}
	fmt.Println("table created")
	fmt.Fprintln(w, "Table created ")
}

func InsertTable(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	_, err := db.Exec("INSERT INTO employee(Ename , Esalary) VALUES('Swapnil' , 45)")
	if err != nil {
		panic(err)
	}
	fmt.Println("Values inserted into table")
	fmt.Fprintf(w, "Data Inserted")
}

func main() {
	dbConn()

	//log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", CreateTable)
	http.HandleFunc("/inserttable", InsertTable)
	http.ListenAndServe(":8080", nil)
}

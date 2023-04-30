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

func dbConnect() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/WalkingDreamzdb")
	if err != nil {
		panic(err)
	}
	return db
}
func CreateDatabase(w http.ResponseWriter, r *http.Request) {
	db := dbConnect()
	_, err := db.Exec("CREATE DATABASE WalkingDreamzdb")
	if err != nil {
		panic(err)
	}
	fmt.Println("Database created successfully")
}

func CreateTable(w http.ResponseWriter, r *http.Request) {
	db := dbConnect()
	_, err := db.Exec("CREATE TABLE employee(ID PRIMARY KEY NOT NULL AUTO_INCREMENT, Ename VARCHAR(50) , Esalary INT)")
	if err != nil {
		panic(err)
	}
	fmt.Println("table created")
	fmt.Fprintln(w, "Table created ")
}

func InsertTable(w http.ResponseWriter, r *http.Request) {
	db := dbConnect()
	_, err := db.Exec("INSERT INTO employee(Ename , Esalary) VALUES('Swapnil' , 45)")
	if err != nil {
		panic(err)
	}
	fmt.Println("Values inserted into table")
	fmt.Fprintf(w, "Data Inserted")
}

func main() {
	http.HandleFunc("/createdb", CreateDatabase)
	http.HandleFunc("/createtable", CreateTable)
	fmt.Println(http.ListenAndServe(":8763", nil))
}

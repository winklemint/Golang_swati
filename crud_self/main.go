package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Jungle struct {
	ID    int
	Aname string
	Claws int
}

func main() {
	// 	// 	// Open a connection to the MySQL database
	// 	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/forestdb")
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}

	// 	CreateTableQuery := `CREATE TABLE IF NOT EXISTS jungle(Id INT AUTO_INCREMENT PRIMARY KEY,
	// 	Aname VARCHAR(50),
	// 	Claws Int,
	// )`

	// 	_, err := db.Exec(CreateTableQuery)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	fmt.Println("Table created successfully")
	// }

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/forestdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// CreateTableQuery := `CREATE TABLE IF NOT EXISTS jungle (
	// 	Id INT AUTO_INCREMENT PRIMARY KEY,
	// 	Aname VARCHAR(50),
	// 	claws INT
	// )`

	// _, err = db.Exec(CreateTableQuery)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// fmt.Println("Table created successfully")

	sql := `INSERT INTO jungle(Aname , Claws) VALUES("Lion" , 4)`
	res, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {

		log.Fatal(err)
	}

	fmt.Printf("the last inserted row id %d\n", lastId)

	// Create a new database named "forestdb"
	// 	_, err = db.Exec("CREATE DATABASE forestdb")
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	fmt.Println("Database created successfully")
}

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type students struct {
	Id         int
	Email      string
	First_name string
	Last_name  string
}

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/classroom")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	sql := "INSERT INTO students(Email, First_name, Last_name) VALUES ('admin@gmail.com', 'admin','admin')"

	res, err := db.Exec(sql)

	if err != nil {
		panic(err.Error())
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The last inserted row id: %d\n", lastId)

	result, err := db.Query("SELECT * FROM students")

	if err != nil {
		log.Fatal(err)
	}

	for result.Next() {

		var student students
		err := result.Scan(&student.Id, &student.Email, &student.First_name, &student.Last_name)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", student)
	}

}

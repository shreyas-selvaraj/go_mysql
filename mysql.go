package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //importing a package solely for its side-effects, registering sql as driver without importing any other functions
)

type User struct {
	Name string `json: "name"`
}

func main() {
	fmt.Println("Go MySQL")

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/testdb")

	if err != nil {
		panic(err.Error()) //exit out of function if unexpected error
	}

	defer db.Close() // defers the execution of a function until the surrounding function returns

	/* // INSERT INTO DATABASE
	insert, err := db.Query("INSERT INTO users VALUES('SHREYAS')")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	fmt.Println("Successfully inserted into user table")
	*/

	results, err := db.Query("SELECT name FROM users")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var user User
		err = results.Scan(&user.Name)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(user.Name)
	}

}

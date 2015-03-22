package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {

	// 1レコードを SELECT してみる
	cnn, err := sql.Open("mysql", "root:@/aims_100")
	if err != nil {
		panic(err.Error())
	}
	defer cnn.Close()

	id := 100
	var name string

	if err := cnn.QueryRow("SELECT user_name FROM users WHERE user_id = ? LIMIT 1", id).Scan(&name); err != nil {
		log.Fatal(err)
	}

	fmt.Println(id, name)

	// 複数レコードを SELECT してみる
	rows, err := cnn.Query("SELECT user_id, user_name FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

}

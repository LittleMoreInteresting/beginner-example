package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 连接到 MySQL 数据库
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/sys")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	// 查询数据
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// 处理结果
	for rows.Next() {
		var id int
		var name string
		var email string
		err = rows.Scan(&id, &name, &email)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", id, name, email)
	}

	row := db.QueryRow("SELECT * FROM users where id = ?", 1)
	var id int
	var name string
	var email string
	if err = row.Scan(&id, &name, &email); err != nil {
		if err == sql.ErrNoRows {

		}
	}

	result, err := db.Exec("INSERT INTO `users` (`name`,`email`) VALUES (?, ?)", "Lilei", "789@golang.com")
	if err != nil {
		fmt.Printf("add err: %v", err)
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("addAlbum: %v", err)
	}
	fmt.Println(lastId)
}

package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id    uint
	Name  string
	Email string
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/sys?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	var user User
	err = db.Find(&user).Error
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(user.Id, user.Name, user.Email)

	new_user := User{Name: "Tommy", Email: "121@mail.com"}

	result := db.Create(&new_user)
	fmt.Println(
		new_user.Id,
		result.Error,
		result.RowsAffected)
}

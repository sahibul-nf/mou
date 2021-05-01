package main

import (
	"fmt"
	"log"
	"mou/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// membuat koneksi ke db mysql
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "uSNF:uSNF@tcp(127.0.0.1:3306)/my_dictionary?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connection to db is good :)")

	var tb_user []user.User
	length := len(tb_user)

	fmt.Println(length)

	db.Table("tb_user").Select("password", "email").Scan(&tb_user)

	length = len(tb_user)
	fmt.Println(length)

	for _, user := range tb_user {
		fmt.Println(user.Email)
		fmt.Println(user.Password)
		fmt.Println()
	}

	// result := map[string]interface{}{}
	// a := db.Table("tb_user").Take(&result)
	// fmt.Println(a)
}

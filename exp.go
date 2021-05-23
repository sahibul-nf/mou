// contoh cara koneksi database mysql
// dan membuat entpoint + decode to json

// package main

// import (
// 	"fmt"
// 	"log"
// 	"mou/user"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// func main() {
// 	// membuat koneksi ke db mysql
// 	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
// 	// dsn := "uSNF:uSNF@tcp(127.0.0.1:3306)/moyu?charset=utf8mb4&parseTime=True&loc=Local"
// 	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	// if err != nil {
// 	// 	log.Fatal(err.Error())
// 	// }

// 	router := gin.Default()
// 	router.GET("/register", handler)
// 	router.Run()
// }

// func handler(c *gin.Context) {
// 	// membuat koneksi ke db mysql
// 	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
// 	dsn := "uSNF:uSNF@tcp(127.0.0.1:3306)/moyu?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	fmt.Println("Connection to db is good :)")

// 	var users []user.User

// 	db.Find(&users)
// 	// db.Table("users").Select("password_hash", "email", "created_at").Scan(&users)

// 	c.JSON(http.StatusOK, users)
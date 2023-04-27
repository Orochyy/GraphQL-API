package database

import (
	"GraphQL-API/dbmodel"
	"fmt"
	"github.com/joho/godotenv"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DBInstance *gorm.DB

var err error

//var CONNECTION_STRING string = "root:root@tcp(localhost:3306)/?charset=utf8&parseTime=True&loc=Local"

func ConnectDB() {

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("MYSQL_ROOT_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)

	ConnectionURI := dsn

	DBInstance, err = gorm.Open("mysql", ConnectionURI)
	if err != nil {
		fmt.Println(err)
		panic("Database connection attempt was unsuccessful.....")
	} else {
		fmt.Println("Database Connected successfully.....")
	}

	DBInstance.LogMode(true)
}

func CreateDB() {
	DBInstance.Exec("CREATE DATABASE IF NOT EXISTS Blog_Posts")
	DBInstance.Exec("USE Blog_Posts")
}

func MigrateDB() {
	DBInstance.AutoMigrate(&dbmodel.Post{})
	fmt.Println("Database migration completed....")
}

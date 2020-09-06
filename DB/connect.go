package db

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Connect ...DBとの接続
func Connect() *gorm.DB {
	/*
		.envが取得できない
		ssl := "?sslmode=disable"
		err := godotenv.Load("../.env")
		if err != nil {
			ssl = ""
			log.Print("Error loading .env file")
		}
	*/

	//connstr := "自分のDATABASE_URL"
	connstr := os.Getenv("DATABASE_URL")
	db, err := gorm.Open("postgres", connstr)
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}

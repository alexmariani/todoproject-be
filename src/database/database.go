package database

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDatabase() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Errore nel file .env")
	}

	host := os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	user := os.Getenv("USERNAME")
	dbname := os.Getenv("DB_NAME")
	pass := os.Getenv("PASSWORD")

	fmt.Printf("host=%s port=%d user=%s dbname=%s pass=%s \n", host, port, user, dbname, pass)

	psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, pass)
	db, errSql := gorm.Open(postgres.Open(psqlSetup), &gorm.Config{TranslateError: true})

	if errSql != nil {
		fmt.Println("Errore nella connessione al db ", errSql)
		panic(errSql)
	}

	Db = db
	fmt.Println("Connessione al db avvenuta con successo!")

}

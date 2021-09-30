package repository

import (
	"database/sql"
	"fmt"
	"github.com/sinamna/PaymentGateway/internal/models"
	"os"
)

var DB Database

func init() {
	initPG()
}

type Database interface {
	SaveTransaction(req *models.CreateTransactionReq) error
}

func initPG() {
	host := os.Getenv("PG_GATEWAY_HOST")
	port := os.Getenv("PG_GATEWAY_PORT")
	user := os.Getenv("PG_GATEWAY_USER")
	password := os.Getenv("PG_GATEWAY_PASSWORD")
	dbname := os.Getenv("PG_GATEWAY_DBNAME")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("error in connecting to db: ", err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("error in pinging to db: ", err)
	}
	DB = &PgDatabase{client: db}
}

func initInMemmory() {

}
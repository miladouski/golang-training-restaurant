package main

import (
	"log"
	"os"
	"time"

	"github.com/miladouski/golang-training-restaurant/restaurant/pkg/data"
	"github.com/miladouski/golang-training-restaurant/restaurant/pkg/db"
)

var (
	host     = os.Getenv("DB_USERS_HOST")
	port     = os.Getenv("DB_USERS_PORT")
	user     = os.Getenv("DB_USERS_USER")
	dbname   = os.Getenv("DB_USERS_DBNAME")
	password = os.Getenv("DB_USERS_PASSWORD")
	sslmode  = os.Getenv("DB_USERS_SSL")
)

func init() {
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5432"
	}
	if user == "" {
		user = "postgres"
	}
	if dbname == "" {
		dbname = "restaurant"
	}
	if password == "" {
		password = "1111"
	}
	if sslmode == "" {
		sslmode = "disable"
	}
}

func main() {
	conn, err := db.GetConnection(host, port, user, dbname, password, sslmode)
	if err != nil {
		log.Fatalf("can't connect to database, error: %v", err)
	}
	orderData := data.NewOrderDate(conn)
	orders, err := orderData.ReadAll()
	if err != nil {
		log.Println(err)
	}
	order, err := orderData.Read(2)
	if err != nil {
		log.Println(err)
	}
	err = orderData.Create(data.Order{Id: 6, Date: time.Now(), Table: 2, WaiterId: 3, Price: 356, Payment: true})
	if err != nil {
		log.Println(err)
	}
	err = orderData.Update(2, 500, false)
	if err != nil {
		log.Println(err)
	}
	err = orderData.Delete(6)
	if err != nil {
		log.Println(err)
	}
	log.Println(orders)
	log.Println(order)
}

package main

import (
	"github.com/gorilla/mux"
	"github.com/miladouski/golang-training-restaurant/restaurant/pkg/api"
	"github.com/miladouski/golang-training-restaurant/restaurant/pkg/data"
	"github.com/miladouski/golang-training-restaurant/restaurant/pkg/db"
	"log"
	"net"
	"net/http"
	"os"
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
	r := mux.NewRouter()
	orderData := data.NewOrderData(conn)
	api.ServeOrderResource(r, *orderData)
	r.Use(mux.CORSMethodMiddleware(r))
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Server Listen port...")
	}
	if err := http.Serve(listener, r); err != nil {
		log.Fatal("Server has been crashed...")
	}
}

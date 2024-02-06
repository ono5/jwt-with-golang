package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ono5/jwt-with-golang/controllers"
	"github.com/ono5/jwt-with-golang/driver"
	"github.com/subosito/gotenv"
)

var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {

	db = driver.ConnectDB()

	router := mux.NewRouter()

	controller := controllers.Controller{}

	router.HandleFunc("/signup", controller.Signup(db)).Methods("POST")
	router.HandleFunc("/login", controller.Login(db)).Methods("POST")
	router.HandleFunc("/protected", controller.TokenVerifyMiddleware(controller.ProtectedEndpoint())).Methods("GET")

	log.Println("Liten on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}

package main

import (
	"log"
	"net/http"
	"separation/auth/db"
	"separation/auth/handlers"
	"separation/auth/utils"
)

func main() {
	err := utils.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	err = db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.CloseDB()

	http.HandleFunc("/signup", handlers.SignUp)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/delete/", handlers.Delete)
	http.HandleFunc("/refreshtoken", handlers.RefreshToken)
	http.HandleFunc("/validate", handlers.Validate)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

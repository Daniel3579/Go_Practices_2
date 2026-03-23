package main

import (
	"log"
	"net/http"
	"separation/task/db"
	h "separation/task/handlers"
	mid "separation/task/middleware"
	"separation/task/utils"
)

func main() {
	err := utils.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	err = db.ConnectDB("DATABASE_URL")
	if err != nil {
		log.Fatal(err)
	}
	defer db.CloseDB()

	http.Handle("/insert", mid.ValidateMiddleware(http.HandlerFunc(h.InsertHandler)))
	http.Handle("/selectall", mid.ValidateMiddleware(http.HandlerFunc(h.SelectAllHandler)))
	http.Handle("/select/", mid.ValidateMiddleware(http.HandlerFunc(h.SelectCurrentHandler)))
	http.Handle("/update/", mid.ValidateMiddleware(http.HandlerFunc(h.UpdateHandler)))
	http.Handle("/delete/", mid.ValidateMiddleware(http.HandlerFunc(h.DeleteHandler)))
	log.Fatal(http.ListenAndServe(":8081", nil))
}

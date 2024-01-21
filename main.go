package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var tpl = template.Must(template.ParseFiles("templates/index.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)

	http.Handle("/", r)

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

	log.Println("Server started on 8080")
	http.ListenAndServe(":8080", handlers.CORS(allowedMethods, allowedOrigins)(r))
}

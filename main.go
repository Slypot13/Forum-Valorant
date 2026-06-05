package main

import (
	"fmt"
	"html/template"
	"net/http"
	"forum-valorant/config"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func main() {

	config.LoadEnv()

	db := config.InitDB()
	defer db.Close()

	http.HandleFunc("/", homeHandler)

	http.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static")),
		),
	)

	fmt.Println("Serveur lancé sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
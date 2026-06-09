package main

import (
	"fmt"
	"html/template"
	"net/http"

	"forum-valorant/config"
	"forum-valorant/controllers"
	"forum-valorant/repositories"
	"forum-valorant/services"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func main() {
	config.LoadEnv()

	db := config.InitDB()
	defer db.Close()

	userRepository := repositories.InitUserRepository(db)
	userService := services.InitUserService(userRepository)
	authController := controllers.InitAuthController(userService)

	threadRepository := repositories.InitThreadRepository(db)
	threadService := services.InitThreadService(threadRepository)
	threadController := controllers.InitThreadController(threadService)

	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			authController.ShowRegister(w, r)
		}

		if r.Method == http.MethodPost {
			authController.Register(w, r)
		}
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			authController.ShowLogin(w, r)
		}

		if r.Method == http.MethodPost {
			authController.Login(w, r)
		}
	})

	http.HandleFunc("/threads/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			threadController.ShowCreateThread(w, r)
		}

		if r.Method == http.MethodPost {
			threadController.CreateThread(w, r)
		}
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Serveur lancé sur http://localhost:8080")
	fmt.Println("FT1 inscription : http://localhost:8080/register")
	fmt.Println("FT2 connexion : http://localhost:8080/login")
	fmt.Println("FT3 création fil : http://localhost:8080/threads/create")

	http.ListenAndServe(":8080", nil)
}
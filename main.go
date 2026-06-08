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

	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			authController.ShowRegister(w, r)
		}

		if r.Method == http.MethodPost {
			authController.Register(w, r)
		}
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Serveur lancé sur http://localhost:8080")
	fmt.Println("FT1 sur http://localhost:8080/register")
	http.ListenAndServe(":8080", nil)
}

http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		authController.ShowLogin(w, r)
	}

	if r.Method == http.MethodPost {
		authController.Login(w, r)
	}
})
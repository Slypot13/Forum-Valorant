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

func main() {
	config.LoadEnv()

	db := config.InitDB()
	defer db.Close()

	userRepository := repositories.InitUserRepository(db)
	userService := services.InitUserService(userRepository)
	authController := controllers.InitAuthController(userService)

	threadRepository := repositories.InitThreadRepository(db)
	threadService := services.InitThreadService(threadRepository)

	messageRepository := repositories.InitMessageRepository(db)
	messageService := services.InitMessageService(messageRepository, threadRepository)

	reactionRepository := repositories.InitReactionRepository(db)
	reactionService := services.InitReactionService(reactionRepository)
	reactionController := controllers.InitReactionController(reactionService)

	threadController := controllers.InitThreadController(threadService, messageService)
	messageController := controllers.InitMessageController(messageService)

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			authController.ShowRegister(w, r)
			return
		}

		if r.Method == http.MethodPost {
			authController.Register(w, r)
			return
		}
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			authController.ShowLogin(w, r)
			return
		}

		if r.Method == http.MethodPost {
			authController.Login(w, r)
			return
		}
	})

	http.HandleFunc("/threads/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			threadController.ShowCreateThread(w, r)
			return
		}

		if r.Method == http.MethodPost {
			threadController.CreateThread(w, r)
			return
		}
	})

	http.HandleFunc("/threads/view", func(w http.ResponseWriter, r *http.Request) {
		threadController.ShowThreadDetail(w, r)
	})

	http.HandleFunc("/messages/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			messageController.CreateMessage(w, r)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	http.HandleFunc("/reactions/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			reactionController.React(w, r)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		threads, err := threadService.GetVisibleThreads()

		if err != nil {
			http.Error(w, "Erreur lors du chargement des sujets", http.StatusInternalServerError)
			return
		}

		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, threads)
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Serveur lancé sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
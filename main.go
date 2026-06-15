package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"forum-valorant/config"
	"forum-valorant/controllers"
	"forum-valorant/models"
	"forum-valorant/repositories"
	"forum-valorant/services"
)

type HomePage struct {
	Threads []models.Thread
	Limit   string
	Page    int
	Tag     string
	Search  string
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

	http.HandleFunc("/threads/edit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			threadController.ShowEditThread(w, r)
			return
		}

		if r.Method == http.MethodPost {
			threadController.EditThread(w, r)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	http.HandleFunc("/threads/delete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			threadController.DeleteThread(w, r)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
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

	http.HandleFunc("/messages/delete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			messageController.DeleteMessage(w, r)
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

		page := 1
		limit := 10

		limitString := r.URL.Query().Get("limit")
		pageString := r.URL.Query().Get("page")
		tagString := r.URL.Query().Get("tag")
		search := r.URL.Query().Get("search")

		if search != "" {
			_, err := services.GetUserIdFromRequest(r)

			if err != nil {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
		}

		if pageString != "" {
			value, err := strconv.Atoi(pageString)

			if err == nil && value > 0 {
				page = value
			}
		}

		if limitString == "20" {
			limit = 20
		} else if limitString == "30" {
			limit = 30
		} else if limitString == "all" {
			limit = 100000
		} else {
			limitString = "10"
		}

		offset := (page - 1) * limit

		var threads []models.Thread
		var err error

		if search != "" {
			threads, err = threadService.SearchVisibleThreads(search, limit, offset)
		} else if tagString != "" {
			tagId, convertErr := strconv.Atoi(tagString)

			if convertErr == nil {
				threads, err = threadService.GetVisibleThreadsByTagPaginated(tagId, limit, offset)
			} else {
				threads, err = threadService.GetVisibleThreadsPaginated(limit, offset)
			}
		} else {
			threads, err = threadService.GetVisibleThreadsPaginated(limit, offset)
		}

		if err != nil {
			http.Error(w, "Erreur lors du chargement des sujets", http.StatusInternalServerError)
			return
		}

		data := HomePage{
			Threads: threads,
			Limit:   limitString,
			Page:    page,
			Tag:     tagString,
			Search:  search,
		}

		funcMap := template.FuncMap{
			"plus": func(a int, b int) int {
				return a + b
			},
			"minus": func(a int, b int) int {
				return a - b
			},
		}

		tmpl := template.Must(
			template.New("index.html").
				Funcs(funcMap).
				ParseFiles("templates/index.html"),
		)

		tmpl.Execute(w, data)
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Serveur lancé sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
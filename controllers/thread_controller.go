package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	"forum-valorant/models"
	"forum-valorant/services"
)

type ThreadController struct {
	threadService  *services.ThreadService
	messageService *services.MessageService
}

func InitThreadController(threadService *services.ThreadService, messageService *services.MessageService) *ThreadController {
	return &ThreadController{
		threadService:  threadService,
		messageService: messageService,
	}
}

type ThreadDetailPage struct {
	Thread   models.Thread
	Messages []models.Message
}

func (c *ThreadController) ShowCreateThread(w http.ResponseWriter, r *http.Request) {
	_, err := services.GetUserIdFromRequest(r)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/create_thread.html"))
	tmpl.Execute(w, nil)
}

func (c *ThreadController) CreateThread(w http.ResponseWriter, r *http.Request) {
	userId, err := services.GetUserIdFromRequest(r)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")

	err = c.threadService.CreateThread(title, content, userId)

	if err != nil {
		tmpl := template.Must(template.ParseFiles("templates/create_thread.html"))
		tmpl.Execute(w, err.Error())
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (c *ThreadController) ShowThreadDetail(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, "Sujet introuvable", http.StatusBadRequest)
		return
	}

	thread, err := c.threadService.GetThreadById(id)

	if err != nil {
		http.Error(w, "Sujet introuvable ou archivé", http.StatusNotFound)
		return
	}

	messages, err := c.messageService.GetMessagesByThreadId(id)

	if err != nil {
		http.Error(w, "Erreur lors du chargement des messages", http.StatusInternalServerError)
		return
	}

	page := ThreadDetailPage{
		Thread:   thread,
		Messages: messages,
	}

	tmpl, err := template.ParseFiles("templates/thread_detail.html")

	if err != nil {
		http.Error(w, "Erreur chargement template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, page)

	if err != nil {
		http.Error(w, "Erreur affichage template", http.StatusInternalServerError)
		return
	}
}
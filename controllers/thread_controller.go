package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	"forum-valorant/models"
	"forum-valorant/services"
)

// gère les sujets.
type ThreadController struct {
	threadService  *services.ThreadService
	messageService *services.MessageService
}

// initialise le contrôleur.
func InitThreadController(threadService *services.ThreadService, messageService *services.MessageService) *ThreadController {
	return &ThreadController{
		threadService:  threadService,
		messageService: messageService,
	}
}

// données envoyées à la page détail.
type ThreadDetailPage struct {
	Thread      models.Thread
	Messages    []models.Message
	Sort        string
	Limit       string
	MessagePage int
}

// affiche la page de création.
func (c *ThreadController) ShowCreateThread(w http.ResponseWriter, r *http.Request) {
	_, err := services.GetUserIdFromRequest(r)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/create_thread.html"))
	tmpl.Execute(w, nil)
}

// traite la création.
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

// affiche un sujet et ses messages.
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

	sort := r.URL.Query().Get("sort")

	messagePage := 1
	limit := 10
	limitString := r.URL.Query().Get("limit")
	pageString := r.URL.Query().Get("messagePage")

	if pageString != "" {
		value, err := strconv.Atoi(pageString)

		if err == nil && value > 0 {
			messagePage = value
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

	offset := (messagePage - 1) * limit

	messages, err := c.messageService.GetMessagesByThreadId(id, sort, limit, offset)

	if err != nil {
		http.Error(w, "Erreur lors du chargement des messages", http.StatusInternalServerError)
		return
	}

	page := ThreadDetailPage{
		Thread:      thread,
		Messages:    messages,
		Sort:        sort,
		Limit:       limitString,
		MessagePage: messagePage,
	}

	funcMap := template.FuncMap{
		"plus": func(a int, b int) int {
			return a + b
		},
		"minus": func(a int, b int) int {
			return a - b
		},
	}

	tmpl, err := template.New("thread_detail.html").Funcs(funcMap).ParseFiles("templates/thread_detail.html")

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

// affiche la page de modification.
func (c *ThreadController) ShowEditThread(w http.ResponseWriter, r *http.Request) {
	userId, err := services.GetUserIdFromRequest(r)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	role, err := services.GetUserRoleFromRequest(r)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, "Sujet introuvable", http.StatusBadRequest)
		return
	}

	thread, err := c.threadService.GetThreadById(id)

	if err != nil {
		http.Error(w, "Sujet introuvable", http.StatusNotFound)
		return
	}

	if thread.UserId != userId && role != "admin" {
		http.Error(w, "Accès refusé", http.StatusForbidden)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/edit_thread.html"))
	tmpl.Execute(w, thread)
}

// traite la modification.
func (c *ThreadController) EditThread(w http.ResponseWriter, r *http.Request) {
	userId, err := services.GetUserIdFromRequest(r)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	role, err := services.GetUserRoleFromRequest(r)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	idString := r.FormValue("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, "Sujet introuvable", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")

	err = c.threadService.UpdateThread(id, title, content, userId, role)

	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	http.Redirect(w, r, "/threads/view?id="+idString, http.StatusSeeOther)
}

// supprime un sujet.
func (c *ThreadController) DeleteThread(w http.ResponseWriter, r *http.Request) {
	userId, err := services.GetUserIdFromRequest(r)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	role, err := services.GetUserRoleFromRequest(r)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	idString := r.FormValue("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, "Sujet introuvable", http.StatusBadRequest)
		return
	}

	err = c.threadService.DeleteThread(id, userId, role)

	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
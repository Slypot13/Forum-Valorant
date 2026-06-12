package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	"forum-valorant/models"
	"forum-valorant/services"
)

// gère les sujets (sujets, messages, etc.).
type ThreadController struct {
	threadService  *services.ThreadService
	messageService *services.MessageService
}

// InitThreadController initialise le contrôleur.
func InitThreadController(threadService *services.ThreadService, messageService *services.MessageService) *ThreadController {
	return &ThreadController{
		threadService:  threadService,
		messageService: messageService,
	}
}

// ThreadDetailPage structure de données pour la page de détail.
type ThreadDetailPage struct {
	Thread   models.Thread
	Messages []models.Message
}

// ShowCreateThread affiche la page pour créer un sujet.
func (c *ThreadController) ShowCreateThread(w http.ResponseWriter, r *http.Request) {
	_, err := services.GetUserIdFromRequest(r)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/create_thread.html"))
	tmpl.Execute(w, nil)
}

// CreateThread traite le formulaire de création de sujet.
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

// ShowThreadDetail affiche un sujet et ses messages.
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

// ShowEditThread affiche la page de modification.
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

// EditThread traite la modification.
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

// DeleteThread supprime un sujet.
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
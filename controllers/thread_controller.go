package controllers

import (
	"html/template"
	"net/http"

	"forum-valorant/services"
)

type ThreadController struct {
	threadService *services.ThreadService
}

func InitThreadController(threadService *services.ThreadService) *ThreadController {
	return &ThreadController{
		threadService: threadService,
	}
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

	err = c.threadService.CreateThread(
		title,
		content,
		userId,
	)

	if err != nil {
		tmpl := template.Must(template.ParseFiles("templates/create_thread.html"))
		tmpl.Execute(w, err.Error())
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
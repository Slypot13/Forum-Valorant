package controllers

import (
	"net/http"
	"strconv"

	"forum-valorant/services"
)

// gère la publication de messages.
type MessageController struct {
	messageService *services.MessageService
}

// initialise le contrôleur.
func InitMessageController(messageService *services.MessageService) *MessageController {
	return &MessageController{
		messageService: messageService,
	}
}

// traite le formulaire de nouveau message.
func (c *MessageController) CreateMessage(w http.ResponseWriter, r *http.Request) {

	userId, err := services.GetUserIdFromRequest(r)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	threadIdString := r.FormValue("thread_id")

	threadId, err := strconv.Atoi(threadIdString)

	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	content := r.FormValue("content")

	err = c.messageService.CreateMessage(
		content,
		threadId,
		userId,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(
		w,
		r,
		"/threads/view?id="+threadIdString,
		http.StatusSeeOther,
	)
}
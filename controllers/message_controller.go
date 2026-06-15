package controllers

import (
	"net/http"
	"strconv"

	"forum-valorant/services"
)

type MessageController struct {
	messageService *services.MessageService
}

func InitMessageController(messageService *services.MessageService) *MessageController {
	return &MessageController{
		messageService: messageService,
	}
}

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

	err = c.messageService.CreateMessage(content, threadId, userId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/threads/view?id="+threadIdString, http.StatusSeeOther)
}

func (c *MessageController) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	userId, err := services.GetUserIdFromRequest(r)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	messageIdString := r.FormValue("message_id")
	threadId := r.FormValue("thread_id")

	messageId, err := strconv.Atoi(messageIdString)

	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err = c.messageService.DeleteMessage(messageId, userId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	http.Redirect(w, r, "/threads/view?id="+threadId, http.StatusSeeOther)
}
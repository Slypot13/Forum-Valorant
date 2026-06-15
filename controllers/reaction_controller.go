package controllers

import (
	"net/http"
	"strconv"

	"forum-valorant/services"
)

// gère l'ajout des likes/dislikes.
type ReactionController struct {
	reactionService *services.ReactionService
}

//  initialise le contrôleur.
func InitReactionController(reactionService *services.ReactionService) *ReactionController {
	return &ReactionController{
		reactionService: reactionService,
	}
}

// React traite le clic sur un like/dislike.
func (c *ReactionController) React(w http.ResponseWriter, r *http.Request) {
	userId, err := services.GetUserIdFromRequest(r)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	messageIdString := r.FormValue("message_id")
	threadId := r.FormValue("thread_id")
	reactionType := r.FormValue("type")

	messageId, err := strconv.Atoi(messageIdString)

	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err = c.reactionService.ReactToMessage(messageId, userId, reactionType)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/threads/view?id="+threadId, http.StatusSeeOther)
}
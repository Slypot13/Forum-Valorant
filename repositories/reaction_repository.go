package repositories

import (
	"database/sql"

	"forum-valorant/models"
)

//gère les requêtes SQL des réactions.
type ReactionRepository struct {
	db *sql.DB
}

//initialise le dépôt.
func InitReactionRepository(db *sql.DB) *ReactionRepository {
	return &ReactionRepository{db}
}

// enregistre ou modifie une réaction.
func (r *ReactionRepository) SaveReaction(reaction models.Reaction) error {
	query := `
	INSERT INTO reactions (message_id, user_id, type)
	VALUES (?, ?, ?)
	ON DUPLICATE KEY UPDATE type = ?
	`

	_, err := r.db.Exec(
		query,
		reaction.MessageId,
		reaction.UserId,
		reaction.Type,
		reaction.Type,
	)

	return err
}
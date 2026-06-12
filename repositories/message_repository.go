package repositories

import (
	"database/sql"

	"forum-valorant/models"
)

//  gère les requêtes SQL des messages.
type MessageRepository struct {
	db *sql.DB
}

// initialise le dépôt.
func InitMessageRepository(db *sql.DB) *MessageRepository {
	return &MessageRepository{db}
}

// ajoute un message.
func (r *MessageRepository) CreateMessage(message models.Message) error {
	query := `
	INSERT INTO messages
	(content, thread_id, user_id)
	VALUES (?, ?, ?)
	`

	_, err := r.db.Exec(
		query,
		message.Content,
		message.ThreadId,
		message.UserId,
	)

	return err
}

// trouve les messages d'un sujet (triés par score).
func (r *MessageRepository) ReadByThreadId(threadId int) ([]models.Message, error) {
	var messages []models.Message

	query := `
	SELECT
		m.id,
		m.content,
		m.thread_id,
		m.user_id,
		m.created_at,
		COALESCE(
			SUM(
				CASE
					WHEN r.type = 'like' THEN 1
					WHEN r.type = 'dislike' THEN -1
					ELSE 0
				END
			),
		0) AS score
	FROM messages m
	LEFT JOIN reactions r ON m.id = r.message_id
	WHERE m.thread_id = ?
	GROUP BY m.id, m.content, m.thread_id, m.user_id, m.created_at
	ORDER BY score DESC
	`

	rows, err := r.db.Query(query, threadId)

	if err != nil {
		return messages, err
	}

	defer rows.Close()

	for rows.Next() {
		var message models.Message

		err := rows.Scan(
			&message.Id,
			&message.Content,
			&message.ThreadId,
			&message.UserId,
			&message.CreatedAt,
			&message.Score,
		)

		if err != nil {
			continue
		}

		messages = append(messages, message)
	}

	return messages, nil
}
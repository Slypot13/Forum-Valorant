package repositories

import (
	"database/sql"

	"forum-valorant/models"
)

// gère les requêtes SQL pour les messages.
type MessageRepository struct {
	db *sql.DB
}

// initialise le dépôt de messages.
func InitMessageRepository(db *sql.DB) *MessageRepository {
	return &MessageRepository{db}
}

// insère un message en base de données.
func (r *MessageRepository) CreateMessage(message models.Message) error {
	query := `
	INSERT INTO messages
	(content, thread_id, user_id)
	VALUES (?, ?, ?)
	`

	_, err := r.db.Exec(query, message.Content, message.ThreadId, message.UserId)

	return err
}

// récupère les messages d'un fil avec tri et pagination.
func (r *MessageRepository) ReadByThreadId(threadId int, sort string, limit int, offset int) ([]models.Message, error) {
	var messages []models.Message

	orderBy := "m.created_at DESC"

	if sort == "oldest" {
		orderBy = "m.created_at ASC"
	}

	if sort == "popular" {
		orderBy = "score DESC"
	}

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
	ORDER BY ` + orderBy + `
	LIMIT ? OFFSET ?
	`

	rows, err := r.db.Query(query, threadId, limit, offset)

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

// récupère un message par son ID.
func (r *MessageRepository) ReadById(id int) (models.Message, error) {
	var message models.Message

	query := `
	SELECT id, content, thread_id, user_id, created_at
	FROM messages
	WHERE id = ?
	`

	err := r.db.QueryRow(query, id).Scan(
		&message.Id,
		&message.Content,
		&message.ThreadId,
		&message.UserId,
		&message.CreatedAt,
	)

	return message, err
}

// supprime un message.
func (r *MessageRepository) DeleteMessage(id int) error {
	query := `
	DELETE FROM messages
	WHERE id = ?
	`

	_, err := r.db.Exec(query, id)

	return err
}
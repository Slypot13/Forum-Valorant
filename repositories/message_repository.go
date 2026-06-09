package repositories

import (
	"database/sql"

	"forum-valorant/models"
)

type MessageRepository struct {
	db *sql.DB
}

func InitMessageRepository(db *sql.DB) *MessageRepository {
	return &MessageRepository{db}
}

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

func (r *MessageRepository) ReadByThreadId(threadId int) ([]models.Message, error) {
	var messages []models.Message

	query := `
	SELECT id, content, thread_id, user_id, created_at
	FROM messages
	WHERE thread_id = ?
	ORDER BY created_at DESC
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
		)

		if err != nil {
			continue
		}

		messages = append(messages, message)
	}

	return messages, nil
}
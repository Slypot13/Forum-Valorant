package repositories

import (
	"database/sql"
	"forum-valorant/models"
)

type ThreadRepository struct {
	db *sql.DB
}

func InitThreadRepository(db *sql.DB) *ThreadRepository {
	return &ThreadRepository{db}
}

func (r *ThreadRepository) CreateThread(thread models.Thread) (int, error) {
	query := `
	INSERT INTO threads
	(title, content, status, user_id)
	VALUES (?, ?, ?, ?)
	`

	result, err := r.db.Exec(
		query,
		thread.Title,
		thread.Content,
		thread.Status,
		thread.UserId,
	)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}
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

func (r *ThreadRepository) ReadVisibleThreads() ([]models.Thread, error) {
	var threads []models.Thread

	query := `
	SELECT id, title, content, status, user_id, created_at
	FROM threads
	WHERE status != 'archivé'
	ORDER BY created_at DESC
	`

	rows, err := r.db.Query(query)

	if err != nil {
		return threads, err
	}

	defer rows.Close()

	for rows.Next() {
		var thread models.Thread

		err := rows.Scan(
			&thread.Id,
			&thread.Title,
			&thread.Content,
			&thread.Status,
			&thread.UserId,
			&thread.CreatedAt,
		)

		if err != nil {
			continue
		}

		threads = append(threads, thread)
	}

	return threads, nil
}

func (r *ThreadRepository) ReadById(id int) (models.Thread, error) {
	var thread models.Thread

	query := `
	SELECT id, title, content, status, user_id, created_at
	FROM threads
	WHERE id = ? AND status != 'archivé'
	`

	err := r.db.QueryRow(query, id).Scan(
		&thread.Id,
		&thread.Title,
		&thread.Content,
		&thread.Status,
		&thread.UserId,
		&thread.CreatedAt,
	)

	return thread, err
}

func (r *ThreadRepository) UpdateThread(thread models.Thread) error {
	query := `
	UPDATE threads
	SET title = ?, content = ?
	WHERE id = ?
	`

	_, err := r.db.Exec(
		query,
		thread.Title,
		thread.Content,
		thread.Id,
	)

	return err
}

func (r *ThreadRepository) DeleteThread(id int) error {
	query := `
	DELETE FROM threads
	WHERE id = ?
	`

	_, err := r.db.Exec(query, id)

	return err
}
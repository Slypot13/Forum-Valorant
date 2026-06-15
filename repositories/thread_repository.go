package repositories

import (
	"database/sql"

	"forum-valorant/models"
)

// gère les requêtes SQL des sujets.
type ThreadRepository struct {
	db *sql.DB
}

// initialise le dépôt.
func InitThreadRepository(db *sql.DB) *ThreadRepository {
	return &ThreadRepository{db}
}

// insère un sujet et retourne son ID.
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

// ajoute une catégorie à un sujet.
func (r *ThreadRepository) AddTagToThread(threadId int, tagId int) error {
	query := `
	INSERT INTO thread_tags
	(thread_id, tag_id)
	VALUES (?, ?)
	`

	_, err := r.db.Exec(query, threadId, tagId)

	return err
}

// récupère les sujets non archivés.
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

// récupère les sujets avec pagination.
func (r *ThreadRepository) ReadVisibleThreadsPaginated(limit int, offset int) ([]models.Thread, error) {
	var threads []models.Thread

	query := `
	SELECT id, title, content, status, user_id, created_at
	FROM threads
	WHERE status != 'archivé'
	ORDER BY created_at DESC
	LIMIT ? OFFSET ?
	`

	rows, err := r.db.Query(query, limit, offset)

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

// récupère les sujets d'une catégorie avec pagination.
func (r *ThreadRepository) ReadVisibleThreadsByTagPaginated(tagId int, limit int, offset int) ([]models.Thread, error) {
	var threads []models.Thread

	query := `
	SELECT t.id, t.title, t.content, t.status, t.user_id, t.created_at
	FROM threads t
	INNER JOIN thread_tags tt ON t.id = tt.thread_id
	WHERE t.status != 'archivé'
	AND tt.tag_id = ?
	ORDER BY t.created_at DESC
	LIMIT ? OFFSET ?
	`

	rows, err := r.db.Query(query, tagId, limit, offset)

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

// trouve un sujet par son ID.
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

// modifie un sujet.
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

// supprime un sujet.
func (r *ThreadRepository) DeleteThread(id int) error {
	query := `
	DELETE FROM threads
	WHERE id = ?
	`

	_, err := r.db.Exec(query, id)

	return err
}
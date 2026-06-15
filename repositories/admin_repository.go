package repositories

import (
	"database/sql"

	"forum-valorant/models"
)

type AdminRepository struct {
	db *sql.DB
}

func InitAdminRepository(db *sql.DB) *AdminRepository {
	return &AdminRepository{db}
}

func (r *AdminRepository) ReadAllThreads() ([]models.Thread, error) {
	var threads []models.Thread

	query := `
	SELECT t.id, t.title, t.content, t.status, t.user_id, u.username, t.created_at
	FROM threads t
	INNER JOIN users u ON t.user_id = u.id
	ORDER BY t.created_at DESC
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
			&thread.Username,
			&thread.CreatedAt,
		)

		if err != nil {
			continue
		}

		threads = append(threads, thread)
	}

	return threads, nil
}

func (r *AdminRepository) UpdateThreadStatus(threadId int, status string) error {
	query := `
	UPDATE threads
	SET status = ?
	WHERE id = ?
	`

	_, err := r.db.Exec(query, status, threadId)

	return err
}

func (r *AdminRepository) DeleteThread(threadId int) error {
	query := `
	DELETE FROM threads
	WHERE id = ?
	`

	_, err := r.db.Exec(query, threadId)

	return err
}

func (r *AdminRepository) DeleteMessage(messageId int) error {
	query := `
	DELETE FROM messages
	WHERE id = ?
	`

	_, err := r.db.Exec(query, messageId)

	return err
}

func (r *AdminRepository) ReadAllMessages() ([]models.Message, error) {
	var messages []models.Message

	query := `
	SELECT m.id, m.content, m.thread_id, m.user_id, u.username, m.created_at
	FROM messages m
	INNER JOIN users u ON m.user_id = u.id
	ORDER BY m.created_at DESC
	`

	rows, err := r.db.Query(query)

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
			&message.Username,
			&message.CreatedAt,
		)

		if err != nil {
			continue
		}

		messages = append(messages, message)
	}

	return messages, nil
}

func (r *AdminRepository) ReadAllUsers() ([]models.User, error) {
	var users []models.User

	query := `
	SELECT id, username, email, role, banned, created_at
	FROM users
	ORDER BY created_at DESC
	`

	rows, err := r.db.Query(query)

	if err != nil {
		return users, err
	}

	defer rows.Close()

	for rows.Next() {
		var user models.User

		err := rows.Scan(
			&user.Id,
			&user.Username,
			&user.Email,
			&user.Role,
			&user.IsBanned,
			&user.CreatedAt,
		)

		if err != nil {
			continue
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *AdminRepository) BanUser(userId int) error {
	query := `
	UPDATE users
	SET banned = TRUE
	WHERE id = ?
	`

	_, err := r.db.Exec(query, userId)

	return err
}
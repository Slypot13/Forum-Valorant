package repositories

import (
	"database/sql"
	"forum-valorant/models"
)

type UserRepository struct {
	db *sql.DB
}

func InitUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) CreateUser(user models.User) error {

	query := `
	INSERT INTO users
	(username, email, password, role, banned)
	VALUES (?, ?, ?, ?, ?)
	`

	_, err := r.db.Exec(
		query,
		user.Username,
		user.Email,
		user.Password,
		user.Role,
		user.IsBanned,
	)

	return err
}

func (r *UserRepository) FindByIdentifier(identifier string) (models.User, error) {
	var user models.User

	query := `
	SELECT id, username, email, password, role, banned, created_at
	FROM users
	WHERE username = ? OR email = ?
	`

	err := r.db.QueryRow(query, identifier, identifier).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.IsBanned,
		&user.CreatedAt,
	)

	return user, err
}
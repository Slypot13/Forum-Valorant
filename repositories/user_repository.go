package repositories

import (
	"database/sql"
	"forum-valorant/models"
)

//gère les requêtes SQL des utilisateurs.
type UserRepository struct {
	db *sql.DB
}

// initialise le dépôt.
func InitUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

// ajoute un utilisateur en base.
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

//cherche un utilisateur par pseudo/email.
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
package repositories

import (
	"github.com/NurbekDos/funk/internal/db"
	"github.com/NurbekDos/funk/internal/models"
)

func CreateUser(user *models.User) (uint, error) {
	query := `
		INSERT INTO users(email, password)
		VALUES($1, $2) RETURNING id
	`

	var id uint

	row := db.DB.QueryRow(query, user.Email, user.Password)
	err := row.Scan(&id)

	return id, err
}

func GetUser(email string) (*models.User, error) {
	query := `
		SELECT id, password, email_verified_at
		FROM users
		WHERE email = $1 AND deleted_at IS NULL
	`

	var user models.User

	row := db.DB.QueryRow(query, email)
	err := row.Scan(
		&user.ID,
		&user.Password,
		&user.EmailVerifiedAt,
	)
	if err != nil {
		return nil, err
	}

	user.Email = email
	return &user, nil
}

func IsUserExists(email string) bool {
	query := `
		SELECT EXISTS(
			SELECT 1 FROM users WHERE email = $1 AND deleted_at IS NULL
		);
	`

	var exists bool
	row := db.DB.QueryRow(query, email)
	err := row.Scan(&exists)

	if err != nil {
		return true
	}
	return exists
}

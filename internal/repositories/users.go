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

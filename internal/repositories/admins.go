package repositories

import (
	"github.com/NurbekDos/funk/internal/db"
	"github.com/NurbekDos/funk/internal/models"
)

func CreateAdmin(admin models.Admin) (uint, error) {
	query := `
		INSERT INTO admin(username, role, password)
		VALUES($1, $2, $3) RETURNING id
	`

	var id uint

	row := db.DB.QueryRow(query, admin.Username, admin.Role, admin.Username)
	err := row.Scan(&id)

	return id, err
}

func GetAdmin(username string) (*models.Admin, error) {
	query := `
		SELECT id, role, password
		FROM admin
		WHERE username = $1 AND deleted_at IS NULL
	`

	var admin models.Admin

	row := db.DB.QueryRow(query, username)
	err := row.Scan(
		&admin.ID,
		&admin.Role,
		&admin.Password,
	)
	if err != nil {
		return nil, err
	}

	admin.Username = username
	return &admin, nil
}

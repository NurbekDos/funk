package repositories

import (
	"github.com/NurbekDos/funk/internal/db"
	"github.com/NurbekDos/funk/internal/models"
)

func CreateIssuer(issuer models.Issuer, adminId uint) (uint, error) {
	query := `
		INSERT INTO issuer(email, phone_number, password, created_by)
		VALUES($1, $2, $3, $4) RETURNING id
	`

	var id uint

	row := db.DB.QueryRow(query, issuer.Email, issuer.PhoneNumber, issuer.Password, adminId)
	err := row.Scan(&id)

	return id, err
}

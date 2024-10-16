package repositories

import (
	"github.com/NurbekDos/funk/internal/db"
	"github.com/NurbekDos/funk/internal/models"
)

func GetCases() ([]models.Cases, error) {
	query := `
		SELECT id, case_name
		FROM cases
		WHERE deleted_at IS NULL
		ORDER BY id
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cases []models.Cases

	for rows.Next() {
		var c models.Cases
		err := rows.Scan(&c.ID, &c.CaseName)
		if err != nil {
			return nil, err
		}
		// Добавляем запись в срез
		cases = append(cases, c)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cases, nil
}

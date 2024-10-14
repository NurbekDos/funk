package repositories

import (
	"github.com/NurbekDos/funk/internal/db"
	"github.com/NurbekDos/funk/internal/models"
)

func CreateCase(cases models.Cases) (uint, error) {
	query := `
		INSERT INTO cases(case_name)
		VALUES($1) RETURNING id
	`

	var id uint

	row := db.DB.QueryRow(query, cases.CaseName)
	err := row.Scan(&id)

	return id, err
}

func GetCases() ([]models.Cases, error) {
	query := `
		SELECT id, case_name
		FROM cases
		WHERE deleted_at IS NULL
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

func CaseExists(caseName string) (bool, error) {
	query := `
		SELECT COUNT(*)
		FROM cases
		WHERE case_name = $1 AND deleted_at IS NULL
	`

	var count int
	err := db.DB.QueryRow(query, caseName).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

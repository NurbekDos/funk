package repositories

import (
	"fmt"

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

func ExistsInTable(tableName string, columnName string, value interface{}) (bool, error) {
	query := fmt.Sprintf(`
		SELECT COUNT(*)
		FROM %s
		WHERE %s = $1 AND deleted_at IS NULL
	`, tableName, columnName)

	var count int
	err := db.DB.QueryRow(query, value).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

/*
func ValidateExistsToInsert(c *gin.Context, tableName string, columnName string, value interface{}) error {
	exists, err := ExistsInTable(tableName, columnName, value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existence"})
		return err
	}

	if exists {
		// Return conflict error if the record does not exist
		c.JSON(http.StatusConflict, gin.H{"message": fmt.Sprintf("%s уже существует", value)})
		return errors.New("record already exists")
	}
}*/

/*
func UpdateCase(c models.Cases) error {
	// Проверяем, существует ли кейс с данным ID
	queryCheck := `
		SELECT COUNT(*)
		FROM cases
		WHERE id = $1 AND deleted_at IS NULL
	`

	fmt.Println("Case name: ", c)
	var count int
	err := db.DB.QueryRow(queryCheck, c.ID).Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("case with id %d does not exist", c.ID)
	}

	exists, err := CaseExists(c.CaseName)
	if err != nil {
		return err
	}
	if exists {
		// Возвращаем форматированную ошибку, если кейс уже существует
		return fmt.Errorf("case с названием %s уже существует", c.CaseName)
	}

	// Если кейс существует, обновляем его данные
	queryUpdate := `
		UPDATE cases
		SET case_name = $1
		WHERE id = $2
	`

	_, err = db.DB.Exec(queryUpdate, c.CaseName, c.ID)
	if err != nil {
		return err
	}

	return nil
}*/

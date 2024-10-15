package repositories

import (
	"database/sql"
	"fmt"
	"reflect"
)

// UniversalRepository - общий репозиторий
type UniversalRepository struct {
	DB *sql.DB
}

// Создание записи (INSERT)
func (r *UniversalRepository) Create(table string, columns []string, values []interface{}) (uint, error) {
	columnsString := fmt.Sprintf("(%s)", joinColumns(columns))
	placeholders := generatePlaceholders(len(values))

	fmt.Println("columnsString: ", columnsString)
	fmt.Println("placeholders: ", placeholders)
	fmt.Println("values: ", values)

	query := fmt.Sprintf("INSERT INTO %s %s VALUES (%s) RETURNING id", table, columnsString, placeholders)

	var id uint
	err := r.DB.QueryRow(query, values...).Scan(&id)
	return id, err
}

// Обновление записи (UPDATE)
func (r *UniversalRepository) Update(table string, id uint, columns []string, values []interface{}) error {
	setClause := generateSetClause(columns)

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", table, setClause, len(values)+1)
	values = append(values, id) // добавляем ID в конец для WHERE

	_, err := r.DB.Exec(query, values...)
	return err
}

// Удаление записи (DELETE)
func (r *UniversalRepository) Delete(table string, id uint) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", table)
	_, err := r.DB.Exec(query, id)
	return err
}

// Вспомогательная функция для объединения столбцов в строку
func joinColumns(columns []string) string {
	return fmt.Sprintf("%s", reflect.ValueOf(columns).Index(0).Interface())
}

// Генерация плейсхолдеров для SQL-запросов (например, $1, $2, $3)
func generatePlaceholders(count int) string {
	placeholders := make([]string, count)
	for i := 1; i <= count; i++ {
		placeholders[i-1] = fmt.Sprintf("$%d", i)
	}
	return fmt.Sprintf("%s", reflect.ValueOf(placeholders).Index(0).Interface())
}

// Генерация строки SET для UPDATE-запроса
func generateSetClause(columns []string) string {
	setClause := ""
	for i, col := range columns {
		setClause += fmt.Sprintf("%s = $%d", col, i+1)
		if i < len(columns)-1 {
			setClause += ", "
		}
	}
	return setClause
}

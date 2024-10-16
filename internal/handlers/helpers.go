package handlers

import (
	"crypto/sha256"
	"fmt"
	"reflect"

	"github.com/NurbekDos/funk/internal/db"
)

func hashPassword(password string) string {
	h := sha256.New()
	h.Write([]byte(password))

	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}

func buildUpdateQuery(req interface{}) ([]string, []interface{}) {
	columns := []string{}
	values := []interface{}{}

	// Получаем рефлективное значение и тип структуры
	v := reflect.ValueOf(req)
	t := reflect.TypeOf(req)

	// Проверка, является ли переданный объект структурой
	if v.Kind() != reflect.Struct {
		fmt.Println("Переданный объект не является структурой")
		return nil, nil
	}

	// Проходим по всем полям структуры
	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		fieldName := t.Field(i).Tag.Get("json") // Получаем имя из тега json

		// Пропускаем поле, если нет тега json
		if fieldName == "" || fieldName == "-" {
			continue
		}

		// Проверяем тип поля и добавляем только непустые значения
		switch fieldValue.Kind() {
		case reflect.String:
			if fieldValue.String() != "" {
				columns = append(columns, fieldName)
				values = append(values, fieldValue.String())
			}
		case reflect.Int, reflect.Int64:
			if fieldValue.Int() != 0 {
				columns = append(columns, fieldName)
				values = append(values, fieldValue.Int())
			}
		case reflect.Uint, reflect.Uint64:
			if fieldValue.Uint() != 0 {
				columns = append(columns, fieldName)
				values = append(values, fieldValue.Uint())
			}
		case reflect.Float64:
			if fieldValue.Float() != 0 {
				columns = append(columns, fieldName)
				values = append(values, fieldValue.Float())
			}
		case reflect.Bool:
			// Добавляем только если значение true
			if fieldValue.Bool() {
				columns = append(columns, fieldName)
				values = append(values, fieldValue.Bool())
			}
		// Добавьте другие типы данных, если необходимо
		default:
			// Можно добавить обработку других типов, если они нужны
			fmt.Printf("Тип поля %s не поддерживается\n", fieldName)
		}
	}

	return columns, values
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

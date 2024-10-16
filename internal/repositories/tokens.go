package repositories

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/NurbekDos/funk/internal/db"
	"github.com/NurbekDos/funk/internal/models"
	"github.com/gin-gonic/gin"
)

type Token struct {
	ID             uint    `json:"id"`
	CaseID         uint    `json:"case_id"`
	Type           string  `json:"type"`
	Symbol         string  `json:"symbol"`
	Name           string  `json:"name"`
	Price          float64 `json:"price"`
	IssuerNumber   uint    `json:"issuer_number"`
	CompanyArea    string  `json:"company_area"`
	CompanyCapital float64 `json:"company_capital"`
	Description    string  `json:"description"`
}

func GetTokens(c *gin.Context) ([]models.Tokens, error) {
	// Получаем параметры запроса (они могут быть пустыми)
	issuerNumber := c.Query("issuer_number")
	caseID := c.Query("case_id")

	fmt.Println("issuerNumber: ", issuerNumber, " caseID: ", caseID)
	// Базовый SQL-запрос
	query := `SELECT id, case_id, type, symbol, name, price, issuer_number, company_area, company_capital, description 
		FROM tokens 
		WHERE deleted_at IS NULL`
	conditions := []string{}
	args := []interface{}{}

	// Добавляем условия, если параметры переданы
	if issuerNumber != "" {
		conditions = append(conditions, "issuer_number = $1")
		args = append(args, issuerNumber)
	}
	if caseID != "" {
		conditions = append(conditions, "case_id = $2")
		args = append(args, caseID)
	}

	// Если есть условия, добавляем их к запросу
	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}
	query += " order by id"

	fmt.Println("query: ", query)
	// Выполняем запрос
	rows, err := db.DB.Query(query, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при выполнении запроса"})
		return nil, err
	}
	defer rows.Close()

	var tokens []models.Tokens
	for rows.Next() {
		var token models.Tokens
		if err := rows.Scan(&token.ID, &token.CaseID, &token.Type, &token.Symbol, &token.Name, &token.Price, &token.IssuerNumber, &token.CompanyArea, &token.CompanyCapital, &token.Description); err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
	}

	// Возвращаем результат
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tokens, nil
}

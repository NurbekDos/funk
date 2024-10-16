package handlers

import (
	"fmt"
	"net/http"

	"github.com/NurbekDos/funk/internal/db"
	"github.com/NurbekDos/funk/internal/middlewares"
	"github.com/NurbekDos/funk/internal/repositories"
	"github.com/gin-gonic/gin"
)

type createTokenRequest struct {
	CaseID         uint    `json:"case_id" binding:"required"`
	Type           string  `json:"type" binding:"required"`
	Symbol         string  `json:"symbol" binding:"required"`
	Name           string  `json:"name" binding:"required"`
	Price          float64 `json:"price" binding:"required"`
	IssuerNumber   int     `json:"issuer_number" binding:"required"`
	CompanyArea    string  `json:"company_area" binding:"required"`
	CompanyCapital float64 `json:"company_capital"`
	Description    string  `json:"description"`
}

type TokensResponse struct {
	Id             uint    `json:"id"`
	CaseID         uint    `json:"case_id"`
	Symbol         string  `json:"symbol"`
	Type           string  `json:"type"`
	Name           string  `json:"name"`
	Price          float64 `json:"price"`
	IssuerNumber   uint    `json:"issuer_number"`
	CompanyArea    string  `json:"company_area"`
	CompanyCapital float64 `json:"company_capital"`
	Description    string  `json:"description"`
}

func AdminGetToken(c *gin.Context) {
	// Получаем токены из репозитория
	tokens, err := repositories.GetTokens(c)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tokens"})
		return
	}

	var resp []TokensResponse

	// Проходим по всем токенам и заполняем массив ответа
	for _, tokenItem := range tokens {
		resp = append(resp, TokensResponse{
			Id:             tokenItem.ID,
			CaseID:         tokenItem.CaseID,
			Symbol:         tokenItem.Symbol,
			Type:           tokenItem.Type,
			Name:           tokenItem.Name,
			Price:          tokenItem.Price,
			IssuerNumber:   tokenItem.IssuerNumber,
			CompanyArea:    tokenItem.CompanyArea,
			CompanyCapital: tokenItem.CompanyCapital,
			Description:    tokenItem.Description,
		})
	}

	// Возвращаем JSON-ответ с массивом токенов
	c.JSON(http.StatusOK, resp)
}

func AdminCreateToken(c *gin.Context) {
	// Проверка прав администратора
	ok := middlewares.CheckAdminRole(c)
	if !ok {
		// Если проверка не пройдена, выходим
		return
	}

	// Парсим тело запроса
	var req createTokenRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body (возможно не хватает переданных данных)"})
		return
	}
	fmt.Println("req:", req)

	// Проверяем, существует ли кейс с таким ID
	caseExists, err := ExistsInTable("cases", "id", req.CaseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check case existence"})
		return
	}
	if !caseExists {
		// Если кейс не существует, возвращаем ошибку
		c.JSON(http.StatusNotFound, gin.H{"message": "Case не существует"})
		return
	}

	// Проверяем, существует ли уже токен с таким символом (symbol)
	symbolExists, err := ExistsInTable("tokens", "symbol", req.Symbol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check token existence"})
		return
	}
	if symbolExists {
		// Если токен с таким символом уже существует, возвращаем ошибку
		c.JSON(http.StatusConflict, gin.H{"message": "Token с таким символом уже существует"})
		return
	}

	// Проверяем, существует ли уже токен с таким именем (name)
	nameExists, err := ExistsInTable("tokens", "name", req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check token existence"})
		return
	}
	if nameExists {
		// Если токен с таким именем уже существует, возвращаем ошибку
		c.JSON(http.StatusConflict, gin.H{"message": "Token с таким именем уже существует"})
		return
	}

	// Создаём токен
	repo := repositories.UniversalRepository{DB: db.DB}

	columns := []string{"case_id", "type", "symbol", "name", "price", "issuer_number", "company_area", "company_capital", "description"}
	values := []interface{}{req.CaseID, req.Type, req.Symbol, req.Name, req.Price, req.IssuerNumber, req.CompanyArea, req.CompanyCapital, req.Description}

	id, err := repo.Create("tokens", columns, values)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании токена"})
		return
	}

	// Возвращаем успешный ответ с ID созданного токена
	c.JSON(http.StatusOK, gin.H{"id": id, "message": "Token создан успешно"})
}

type updateTokenRequest struct {
	ID             uint    `json:"id" binding:"required"`
	CaseID         uint    `json:"case_id"`
	Type           string  `json:"type"`
	Symbol         string  `json:"symbol"`
	Name           string  `json:"name"`
	Price          float64 `json:"price"`
	IssuerNumber   int     `json:"issuer_number"`
	CompanyArea    string  `json:"company_area"`
	CompanyCapital float64 `json:"company_capital"`
	Description    string  `json:"description"`
}

func AdminUpdateToken(c *gin.Context) {
	// Проверка прав администратора
	ok := middlewares.CheckAdminRole(c)
	if !ok {
		return
	}

	// Парсим тело запроса
	var req updateTokenRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	fmt.Println("req:", req)

	// Получаем массивы для обновления через buildUpdateQuery
	columns, values := buildUpdateQuery(req)

	fmt.Println("columns:", columns)
	fmt.Println("values:", values)

	if len(columns) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No data to update"})
		return
	}

	// Проверяем существование токена
	tokenExists, err := ExistsInTable("tokens", "id", req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check token existence"})
		return
	}
	if !tokenExists {
		c.JSON(http.StatusNotFound, gin.H{"message": "Token не существует"})
		return
	}

	// Проверяем, существует ли кейс с таким ID
	if req.CaseID != 0 {
		caseExists, err := ExistsInTable("cases", "id", req.CaseID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check case existence"})
			return
		}
		if !caseExists {
			// Если кейс не существует, возвращаем ошибку
			c.JSON(http.StatusNotFound, gin.H{"message": "Case не существует"})
			return
		}
	}

	// Проверка уникальности символа (если символ изменяется)
	if req.Symbol != "" {
		symbolExists, err := ExistsInTable("tokens", "symbol", req.Symbol)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check token existence"})
			return
		}
		if symbolExists {
			c.JSON(http.StatusConflict, gin.H{"message": "Token с таким символом уже существует"})
			return
		}
	}

	// Проверка уникальности имени (если имя изменяется)
	if req.Name != "" {
		nameExists, err := ExistsInTable("tokens", "name", req.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check token existence"})
			return
		}
		if nameExists {
			c.JSON(http.StatusConflict, gin.H{"message": "Token с таким именем уже существует"})
			return
		}
	}

	// Обновляем токен
	repo := repositories.UniversalRepository{DB: db.DB}
	err = repo.Update("tokens", req.ID, columns, values)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении токена"})
		return
	}

	// Успешный ответ
	c.JSON(http.StatusOK, gin.H{"message": "Token обновлён успешно"})
}

type deleteTokenRequest struct {
	ID uint `json:"id" binding:"required"`
}

func AdminDeleteToken(c *gin.Context) {
	// Проверка прав администратора
	ok := middlewares.CheckAdminRole(c)
	if !ok {
		// Если проверка не пройдена, выходим
		return
	}

	// Парсим тело запроса
	var req deleteTokenRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Проверяем, существует ли токен с таким ID
	tokenExists, err := ExistsInTable("tokens", "id", req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check token existence"})
		return
	}
	if !tokenExists {
		// Если токен не существует, возвращаем ошибку
		c.JSON(http.StatusNotFound, gin.H{"message": "Token не существует"})
		return
	}

	// Удаляем токен
	repo := repositories.UniversalRepository{DB: db.DB}
	err = repo.Delete("tokens", req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении токена"})
		return
	}

	// Возвращаем успешный ответ
	c.JSON(http.StatusOK, gin.H{"message": "Token успешно удалён"})
}

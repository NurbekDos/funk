package handlers

import (
	"fmt"
	"net/http"

	"github.com/NurbekDos/funk/internal/db"
	"github.com/NurbekDos/funk/internal/middlewares"
	"github.com/NurbekDos/funk/internal/repositories"
	"github.com/gin-gonic/gin"
)

type CasesResponse struct {
	Id       uint   `json:"id"`
	CaseName string `json:"case_name"`
}

func Cases(c *gin.Context) {

	cases, err := repositories.GetCases()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve cases"})
		return
	}

	var resp []CasesResponse

	// Проходим по всем кейсам и заполняем массив ответа
	for _, caseItem := range cases {
		resp = append(resp, CasesResponse{
			Id:       caseItem.ID,
			CaseName: caseItem.CaseName,
		})
	}

	c.JSON(http.StatusOK, resp)
}

type createCaseRequest struct {
	CaseName string `json:"case_name"`
}

type createCaseResponse struct {
	Id uint `json:"id"`
}

func AdminCreateCase(c *gin.Context) {
	ok := middlewares.CheckAdminRole(c)
	if !ok {
		return
	}

	req := createCaseRequest{}
	if err := c.BindJSON(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	fmt.Println("req.CaseName", req.CaseName)
	// Проверяем, существует ли уже кейс с таким же именем Если да то выводит ошибку
	exists, err := repositories.ExistsInTable("cases", "case_name", req.CaseName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check case existence"})
		return
	}
	if exists {
		// Возвращаем сообщение, если кейс с таким именем уже существует
		c.JSON(http.StatusConflict, gin.H{"message": "Case уже существует"})
		return
	}

	repo := repositories.UniversalRepository{DB: db.DB}

	columns := []string{"case_name"}
	values := []interface{}{req.CaseName}

	id, err := repo.Create("cases", columns, values)
	if err != nil {
		fmt.Println("Ошибка при создании записи:", err)
	} else {
		fmt.Printf("Создана запись с ID: %d\n", id)
	}
	c.JSON(http.StatusOK, createCaseResponse{Id: id})
}

type updateCaseRequest struct {
	Id       uint   `json:"id"`
	CaseName string `json:"case_name"`
}

func AdminUpdateCase(c *gin.Context) {
	ok := middlewares.CheckAdminRole(c)
	if !ok {
		// Если проверка не пройдена, функция уже вернула ответ и завершила запрос
		return
	}

	updateCase := updateCaseRequest{}
	if err := c.BindJSON(&updateCase); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	fmt.Println("Case name: ", updateCase)

	// Проверяем, существует ли уже кейс с таким же именем и ID
	exists, err := repositories.ExistsInTable("cases", "id", updateCase.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check case existence"})
		return
	}
	if !exists {
		// Возвращаем сообщение, если кейс с таким id не существует
		c.JSON(http.StatusConflict, gin.H{"message": "Case не существует"})
		return
	}

	exists, err = repositories.ExistsInTable("cases", "case_name", updateCase.CaseName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check case existence"})
		return
	}
	if exists {
		// Возвращаем сообщение, если кейс с таким именем уже существует
		c.JSON(http.StatusConflict, gin.H{"message": "Case уже существует"})
		return
	}

	repo := repositories.UniversalRepository{DB: db.DB}

	columns := []string{"case_name"}
	values := []interface{}{updateCase.CaseName}

	err = repo.Update("cases", updateCase.Id, columns, values)
	if err != nil {
		fmt.Println("Ошибка при обновлении записи:", err)
	} else {
		fmt.Println("Запись успешно обновлена")
	}

	c.JSON(http.StatusOK, gin.H{"message": "Case updated successfully"})
}

type deleteCaseRequest struct {
	Id uint `json:"id"`
}

func AdminDeleteCase(c *gin.Context) {
	// Проверка прав администратора
	ok := middlewares.CheckAdminRole(c)
	if !ok {
		return
	}

	// Парсим тело запроса
	req := deleteCaseRequest{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Проверяем, существует ли запись с таким ID
	exists, err := repositories.ExistsInTable("cases", "id", req.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check case existence"})
		return
	}

	if !exists {
		// Возвращаем сообщение, если кейс с таким id не существует
		c.JSON(http.StatusNotFound, gin.H{"message": "Case не существует"})
		return
	}

	// Выполняем удаление записи
	repo := repositories.UniversalRepository{DB: db.DB}
	err = repo.Delete("cases", req.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении записи"})
		return
	}

	// Возвращаем успешный ответ
	c.JSON(http.StatusOK, gin.H{"message": "Case успешно удалён"})
}

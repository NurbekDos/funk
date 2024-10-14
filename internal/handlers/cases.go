package handlers

import (
	"fmt"
	"net/http"

	"github.com/NurbekDos/funk/internal/cfg"
	"github.com/NurbekDos/funk/internal/models"
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
	adminStr, ok := c.Get("admin")
	if !ok {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	admin, ok := adminStr.(models.Admin)
	if !ok {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	if admin.Role != cfg.AdminRole_Super && admin.Role != cfg.AdminRole_Admin {
		c.JSON(http.StatusForbidden, gin.H{"error": "Нет доступа"})
		return
	}

	req := createCaseRequest{}
	if err := c.BindJSON(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	cases := models.Cases{
		CaseName: req.CaseName,
	}

	// Проверяем, существует ли уже кейс с таким же именем
	exists, err := repositories.CaseExists(cases.CaseName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check case existence"})
		return
	}

	if exists {
		// Возвращаем сообщение, если кейс с таким именем уже существует
		c.JSON(http.StatusConflict, gin.H{"message": "Case уже существует"})
		return
	}

	id, err := repositories.CreateCase(cases)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add new case"})
		return
	}

	c.JSON(http.StatusOK, createCaseResponse{Id: id})
}

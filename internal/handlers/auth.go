package handlers

import (
	"fmt"
	"net/http"

	"github.com/NurbekDos/funk/internal/models"
	"github.com/NurbekDos/funk/internal/repositories"
	"github.com/gin-gonic/gin"
)

type registerRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
}

type registerResponse struct {
	Id uint `json:"id"`
}

func Register(c *gin.Context) {
	req := &registerRequest{}

	if err := c.BindJSON(req); err != nil {
		fmt.Println("error")
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	if req.Password == "" || req.Email == "" {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	if req.Password != req.Password2 {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	// TODO check email
	// TODO check user

	id, err := repositories.CreateUser(&models.User{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	// TODO send email

	resp := registerResponse{Id: id}

	c.JSON(http.StatusOK, resp)
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	req := &loginRequest{}

	err := c.BindJSON(req)
	if err != nil {
		fmt.Println("CHE ZA HUINYA????")
		c.JSON(http.StatusBadRequest, nil)
	}

	fmt.Println(req)
}

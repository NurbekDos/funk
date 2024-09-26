package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type registerRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
}

func Register(c *gin.Context) {
	req := &registerRequest{}

	if err := c.BindJSON(req); err != nil {
		fmt.Println("error")
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	if req.Password != req.Password2 {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	fmt.Println(req)
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

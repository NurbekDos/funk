package handlers

import (
	"fmt"
	"net/http"

	"github.com/NurbekDos/funk/internal/cfg"
	"github.com/NurbekDos/funk/internal/models"
	"github.com/NurbekDos/funk/internal/repositories"
	"github.com/gin-gonic/gin"
)

type createAdminRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
}

type createAdminResponse struct {
	Id uint `json:"id"`
}

func AdminCreateAdmin(c *gin.Context) {
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

	if admin.Role != cfg.AdminRole_Super {
		c.JSON(http.StatusForbidden, nil)
		return
	}

	req := createAdminRequest{}
	if err := c.BindJSON(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	if req.Password == "" || req.Username == "" {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	if req.Password != req.Password2 {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	hPassword := hashPassword(req.Password)

	id, err := repositories.CreateAdmin(models.Admin{
		Username: req.Username,
		Role:     cfg.AdminRole_Admin,
		Password: hPassword,
	})
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, createAdminResponse{Id: id})
}

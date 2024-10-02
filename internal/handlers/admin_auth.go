package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/NurbekDos/funk/internal/cfg"
	"github.com/NurbekDos/funk/internal/repositories"
	"github.com/NurbekDos/funk/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type adminLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func AdminLogin(c *gin.Context) {
	req := &adminLoginRequest{}

	err := c.BindJSON(req)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	admin, err := repositories.GetAdmin(req.Username)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	hPassword := hashPassword(req.Password)
	if hPassword != admin.Password {
		fmt.Println("password err")
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	standardClaims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // TODO time.HOUR * 24 -> const
	}
	tokenClaims := services.TokenClaims{
		UserId:         admin.ID,
		Username:       admin.Username,
		Role:           admin.Role,
		Type:           cfg.UserType_Admin,
		StandardClaims: standardClaims,
	}

	token, err := services.GenerateToken(tokenClaims)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	resp := loginResponse{
		Id:    admin.ID,
		Token: token,
	}

	c.JSON(http.StatusOK, resp)
}

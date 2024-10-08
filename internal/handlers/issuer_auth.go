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

type issuerLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func IssuerLogin(c *gin.Context) {
	req := &issuerLoginRequest{}

	err := c.BindJSON(req)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	issuer, err := repositories.GetIssuer(req.Email)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	hPassword := hashPassword(req.Password)
	if hPassword != issuer.Password {
		fmt.Println("password err")
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	standardClaims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // TODO time.HOUR * 24 -> const
	}
	tokenClaims := services.TokenClaims{
		UserId:         issuer.ID,
		Email:          issuer.Email,
		Type:           cfg.UserType_Issuer,
		StandardClaims: standardClaims,
	}

	token, err := services.GenerateToken(tokenClaims)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	resp := loginResponse{
		Id:    issuer.ID,
		Token: token,
	}

	c.JSON(http.StatusOK, resp)
}

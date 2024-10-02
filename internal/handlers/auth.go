package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/NurbekDos/funk/internal/cfg"
	"github.com/NurbekDos/funk/internal/models"
	"github.com/NurbekDos/funk/internal/repositories"
	"github.com/NurbekDos/funk/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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
	if repositories.IsUserExists(req.Email) {
		fmt.Println("email exists")
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	hPassword := hashPassword(req.Password)
	id, err := repositories.CreateUser(&models.User{
		Email:    req.Email,
		Password: hPassword,
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

type loginResponse struct {
	Id    uint   `json:"id"`
	Token string `json:"token"`
}

func Login(c *gin.Context) {
	req := &loginRequest{}

	err := c.BindJSON(req)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	user, err := repositories.GetUser(req.Email)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	hPassword := hashPassword(req.Password)
	if hPassword != user.Password {
		fmt.Println("password err")
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	standardClaims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // TODO time.HOUR * 24 -> const
	}
	tokenClaims := services.TokenClaims{
		UserId:         user.ID,
		Email:          user.Email,
		Type:           cfg.UserType_User,
		StandardClaims: standardClaims,
	}

	token, err := services.GenerateToken(tokenClaims)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	resp := loginResponse{
		Id:    user.ID,
		Token: token,
	}

	c.JSON(http.StatusOK, resp)
}

type meResponse struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
}

func Me(c *gin.Context) {
	userStr, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	user, ok := userStr.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	resp := meResponse{
		Id:    user.ID,
		Email: user.Email,
	}

	c.JSON(http.StatusOK, resp)
}

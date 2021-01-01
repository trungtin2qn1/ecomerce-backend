package controllers

import (
	"ecommerce-backend/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	service services.UserHandler
}

func NewUser() *User {
	return &User{}
}

func NewUserWithValue(userHandler services.UserHandler) *User {
	return &User{
		service: userHandler,
	}
}

func (user *User) Register(c *gin.Context) {
	//Register feature here:
	registerdUser, err := user.service.Register("trungtin2qn1@gmail.com", "1234567")
	if err != nil {
		panic(err)
	}
	log.Println("registerdUser:", registerdUser)
	c.JSON(http.StatusOK, registerdUser)
}

package router

import (
	"ecommerce-backend/controllers"
	"ecommerce-backend/services"

	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
}

func Init() *Router {
	var userServiceHandler services.UserHandler
	userServiceHandler = services.NewUser()
	var userControllerHandler controllers.UserHandler
	userControllerHandler = controllers.NewUserWithValue(userServiceHandler)

	engine := gin.Default()
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	engine.POST("register", userControllerHandler.Register)
	return &Router{
		engine: engine,
	}
}

func NewRouter() *Router {
	return &Router{}
}

func (router *Router) Start() {
	router.engine.Run(":5000")
}

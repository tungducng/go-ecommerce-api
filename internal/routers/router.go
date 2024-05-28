package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tungducng/go-ecommerce-api/internal/controller"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("v1")
	{
		v1.GET("/user/1", controller.NewUserController().GetUserById)
	}

	return r
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

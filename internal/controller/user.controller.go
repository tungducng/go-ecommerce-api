package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tungducng/go-ecommerce-api/internal/service"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// controller->service->repo->models->db
func (u *UserController) GetUserById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": u.userService.GetInfoUser(),
		"users":   []string{"user1", "user2"},
	})
}

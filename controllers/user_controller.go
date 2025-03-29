package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/niteshKrr/gin-framework/models"
	"github.com/niteshKrr/gin-framework/services"
	"golang.org/x/crypto/bcrypt"
)

type User_controller struct {
	user_service services.User_service
}

func (n *User_controller) GetUsers() gin.HandlerFunc {
	users, find := n.user_service.Get_all_users()
	return func(c *gin.Context) {
		if !find {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Kuch to gadbad hai",
			})
		}
		c.JSON(http.StatusOK, users)
	}
}

func (n *User_controller) GetUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		user, found := n.user_service.Get_user_by_id(id)
		if !found {
			c.JSON(http.StatusNotFound, gin.H{
				"Message": "User not found",
			})
		}
		c.JSON(http.StatusOK, user)
	}
}

func (n *User_controller) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser models.User

		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		fmt.Println("Stored Password brfore Hash:", newUser.Password)
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.MinCost)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "problem in password hashing...",
			})
			return
		}
		newUser.Password = string(hashedPassword)
		createdUser := n.user_service.Create_user(newUser)
		c.JSON(http.StatusCreated, createdUser)
	}
}

func (n *User_controller) DeleteUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		delete_user := n.user_service.Delete_user(id)
		if !delete_user {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "something went wrong",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"Message": "user deleted successfully with id : " + id,
		})
	}
}

func (n *User_controller) UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var reqBody struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"-"`
		}

		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		user, err := n.user_service.Update_user(id, reqBody.Name, reqBody.Email, reqBody.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}
		c.JSON(http.StatusOK, user)
	}
}

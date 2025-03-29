package services

import (
	"fmt"

	"github.com/niteshKrr/gin-framework/config"
	"github.com/niteshKrr/gin-framework/models"
	"golang.org/x/crypto/bcrypt"
)

type User_service struct{}

func (n *User_service) Get_all_users() ([]*models.User, bool) {
	var users []*models.User
	err := config.DB.Find(&users).Error
	if err != nil {
		return nil, false
	}

	return users, true
}

func (n *User_service) Get_user_by_id(id string) (*models.User, bool) {
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, false
	}
	return &user, true
}

func (n *User_service) Create_user(new_user models.User) *models.User {
	config.DB.Create(&new_user)
	return &new_user
}

func (n *User_service) Delete_user(id string) bool {
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return false
	}
	config.DB.Delete(&user)
	return true
}

func (n *User_service) Update_user(id string, name string, email string, password string) (*models.User, error) {
	var user models.User
	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	fmt.Println("name is kdfdklf :", user.Name)
	
	if name != "" {
		user.Name = name
	}
	if email != "" {
		user.Email = email
	}
	if password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashedPassword)
	}

	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

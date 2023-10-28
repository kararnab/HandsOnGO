package db

import (
	"github.com/google/uuid"
	"github.com/kararnab/handsongo/pkg/domain/models"
)

func GetUser(user *models.User, username string) error {
	user.ID = uuid.New()
	user.Username = username
	user.Password = "password"
	user.Role = models.Admin
	user.Name = "John Doe"
	user.Age = 31
	user.Gender = models.Male
	user.Contact = "9876543210"

	return nil
}

func SaveUser(user *models.User) {
	//TODO: Save user
	return
}

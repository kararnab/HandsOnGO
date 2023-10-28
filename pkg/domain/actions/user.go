package actions

import (
	"github.com/kararnab/handsongo/pkg/auth/middleware"
	mockDb "github.com/kararnab/handsongo/pkg/db"
	"github.com/kararnab/handsongo/pkg/domain/models"
	"golang.org/x/crypto/bcrypt"
	"html"
	"strings"
)

const OneHourInNanoseconds = 1000000000 * 60 * 60

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string) (string, error) {
	var err error
	u := models.User{}

	//TODO: Login user find
	err = mockDb.GetUser(&u, username)
	//err = DB.Model(User{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}
	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	maker, err := middleware.JWTMaker()
	if err != nil {
		return "", err
	}
	token, err := maker.GenerateToken(u.Username, OneHourInNanoseconds)

	if err != nil {
		return "", err
	}

	return token, nil
}

func SaveUser(u *models.User) (*models.User, error) {
	mockDb.SaveUser(u)
	return u, nil
}

func BeforeSave(u *models.User) error {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil

}

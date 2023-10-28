package middleware

import (
	"github.com/gin-gonic/gin"
	AuthToken "github.com/kararnab/handsongo/pkg/auth/token"
	"github.com/kararnab/handsongo/pkg/initialize"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
)

var maker AuthToken.Maker
var makerErr error

func init() {
	initialize.LoadEnv()
	var jwtKey = os.Getenv("JWT_SECRET_KEY")
	maker, makerErr = AuthToken.NewJWTMaker(jwtKey)
}

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		if makerErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": makerErr.Error()})
			c.Abort()
			return
		}
		_, err := maker.VerifyToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}

func JWTMaker() (AuthToken.Maker, error) {
	if makerErr != nil {
		log.Error().Msg(makerErr.Error())
		return nil, makerErr
	}
	return maker, nil
}

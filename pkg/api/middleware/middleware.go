package middleware

import (
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
			ErrorHandler: func(w http.ResponseWriter, r *http.Request, err string) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			},
			ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
				return []byte(secretKey), nil
			},
			SigningMethod: jwt.SigningMethodHS256,
		})

		if err := jwtMiddleware.CheckJWT(c.Writer, c.Request); err != nil {
			c.Abort()
			return
		}

		c.Next()
	}
}

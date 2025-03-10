package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/MohamedOuhami/TeamGamifierWithGo/initializers"
	"github.com/MohamedOuhami/TeamGamifierWithGo/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func IsAuth(c *gin.Context) {
	fmt.Println("I'm in middlware")

	// Get the token from the cookie
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		// Check the expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {

			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Find the user with the sub id
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {

			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Attach the user to the req
		c.Set("user", user)

		// Continue
		c.Next()

	} else {

		c.AbortWithStatus(http.StatusUnauthorized)
	}

}

// Here, we're going to setup the functions needed for the users to accomplish their tasks
package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/MohamedOuhami/TeamGamifierWithGo/dao"
	"github.com/MohamedOuhami/TeamGamifierWithGo/initializers"
	"github.com/MohamedOuhami/TeamGamifierWithGo/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Signing the user up
func Signup(c *gin.Context) {

	// The body expected should be of type usersignupreq
	var body dao.UserSignUpReq

	// Get the credentials from the req
	// Binding the req to the body, but It can raise an error, so we must check if the error is nil

	if c.Bind(&body) != nil {
		// Raise an error
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error binding the request",
		})

		// Stop the request
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error hashing the password",
		})
		return
	}

	// Store the new user in the database
	user := models.User{FirstName: body.FirstName, LastName: body.LastName, Username: body.Username, Email: body.Email, Password: string(hashedPassword)}

	result := initializers.DB.Create(&user) // pass pointer of data to Create

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Error creating the new user" + result.Error.Error(),
		})
	}

	// Return a response

	c.JSON(http.StatusCreated, gin.H{
		"User": user,
	})
}

// How to login the new user
func Login(c *gin.Context) {

	// Get the email and the password off the req

	var user dao.UserLogin

	if c.Bind(&user) != nil {
		// Return a bad request
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error binding the request",
		})
		return
	}

	// Look up the user in the database
	var foundUser models.User
	initializers.DB.Where("Email = ?", user.Email).First(&foundUser)

	if foundUser.ID == 0 {
		// Return a bad request
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not find the User in the database",
		})
		return
	}

	// Check if the password match the one given in the req
	if bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password)) != nil {
		// Return a bad request
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "The password is incorrect",
		})
		return
	}

	// Generate a JWT token
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": foundUser.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {

		// Return It with the response
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "There was an error creating the token " + err.Error(),
		})
		return
	}

	// Store the token in the cookies
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

// Validate the logged in
func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Logged In cuh",
	})
}

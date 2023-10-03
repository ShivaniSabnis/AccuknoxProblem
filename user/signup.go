package user

import (
	"AccuknoxProblem/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var users = &models.UserStore{}

func Signup(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, "Invalid request body")
		return
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		c.JSON(400, "Invalid or missing request body")
		return
	}

	if existing := users.CheckExistingUser(user.Email); existing {
		c.JSON(400, "User already exists, try with another email")
		return
	}

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		c.JSON(500, "Oops! Its us, not you - please try again.")
		return
	}

	user.Password = string(hashedPassword)
	users.Users = append(users.Users, &user)

	c.JSON(200, "Successfully signed up")
}

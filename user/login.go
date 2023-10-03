package user

import (
	"AccuknoxProblem/models"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, "Invaid request body")
		return
	}

	signedup, userDetail := users.CheckIfSignedUp(user.Email)
	if !signedup {
		c.JSON(401, "User has not signed up , please signup or check your email.")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userDetail.Password), []byte(user.Password)); err != nil {
		c.JSON(401, "Invalid password")
		return
	}

	if userDetail.ExpiryTime.After(time.Now()) {
		c.JSON(200, "User is already logged in")
		return
	}
	userDetail.Sid = uuid.NewString()
	userDetail.ExpiryTime = time.Now().Add(5 * time.Minute)
	log.Println("-------------", userDetail)
	users.UpdateUserDetails(userDetail)

	user.Sid = userDetail.Sid
	user.ExpiryTime = userDetail.ExpiryTime
	user.Password = ""

	c.JSON(200, user)
}

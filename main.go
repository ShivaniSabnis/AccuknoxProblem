package main

import (
	"AccuknoxProblem/user"

	"github.com/gin-gonic/gin"
)

func main() {
	gin := gin.Default()

	gin.POST("/signup", user.Signup)
	gin.POST("/login", user.Login)
	gin.GET("/notes", user.GetNotes)
	gin.POST("/notes", user.PostNotes)
	gin.DELETE("/notes", user.DeleteNote)

	gin.Run("localhost:8080")
}

package main

import (
	"AccuknoxProblem/notes"
	"AccuknoxProblem/user"

	"github.com/gin-gonic/gin"
)

func main() {
	gin := gin.Default()

	gin.POST("/signup", user.Signup)
	gin.POST("/login", user.Login)
	gin.GET("/notes", notes.GetNotes)
	gin.POST("/notes", notes.PostNotes)
	gin.DELETE("/notes", notes.DeleteNote)

	gin.Run("localhost:8080")
}

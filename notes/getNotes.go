package notes

import (
	"AccuknoxProblem/models"
	"AccuknoxProblem/user"

	"github.com/gin-gonic/gin"
)

var notesStore = &models.NotesStore{}

func GetNotes(c *gin.Context) {
	var notes models.NotesRequest

	if err := c.BindJSON(&notes); err != nil {
		c.JSON(400, "Invaid request body")
		return
	}
	users := user.GetUserStore()
	expired := users.CheckSidExpired(notes.Sid)
	if expired {
		c.JSON(401, "Session Expired, please login again.")
		return
	}

	notesArr := notesStore.GetNotes(notes.Sid)
	c.JSON(200, notesArr)

}

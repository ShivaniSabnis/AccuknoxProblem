package notes

import (
	"AccuknoxProblem/models"
	"AccuknoxProblem/user"

	"github.com/gin-gonic/gin"
)

func DeleteNote(c *gin.Context) {
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

	for i, nd := range notesStore.NoteDetails {
		if nd.Sid == notes.Sid {
			for j, n := range nd.Notes {
				if n.Id == notes.Id {
					nd.Notes = append(nd.Notes[:j], nd.Notes[j+1:]...)
				}
			}
			if len(nd.Notes) == 0 {
				notesStore.NoteDetails = append(notesStore.NoteDetails[:i], notesStore.NoteDetails[i+1:]...)
			}
		}
	}
	c.JSON(200, "Successfully Deleted")
}

package notes

import (
	"AccuknoxProblem/models"
	"AccuknoxProblem/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

var NoteId = 0

func PostNotes(c *gin.Context) {
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

	var n models.Notes
	n.Id = NoteId + 1
	NoteId++
	n.Note = notes.Note
	present := false
	for _, nd := range notesStore.NoteDetails {
		if nd.Sid == notes.Sid {
			nd.Notes = append(nd.Notes, n)
			present = true
		}
	}
	if !present {
		notesStore.NoteDetails = append(notesStore.NoteDetails, &models.NoteDetails{Sid: notes.Sid, Notes: []models.Notes{n}})
	}

	c.JSON(200, "Successfully posted noted with "+strconv.Itoa(n.Id))
}

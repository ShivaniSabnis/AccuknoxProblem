package models

type NotesRequest struct {
	Sid  string `json:"sid" validate:"required"`
	Id   int    `json:"id"`
	Note string `json:"note"`
}

type NoteDetails struct {
	Sid   string
	Notes []Notes
}

type Notes struct {
	Id   int    `json:"id"`
	Note string `json:"note"`
}

type NotesStore struct {
	NoteDetails []*NoteDetails
}

func (n *NotesStore) GetNotes(sid string) []Notes {
	if n == nil {
		return nil
	}

	for _, note := range n.NoteDetails {
		if note.Sid == sid {
			return note.Notes
		}
	}
	return nil
}

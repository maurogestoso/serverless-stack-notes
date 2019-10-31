package note

import gouuid "github.com/satori/go.uuid"

// New constructor
func New(userID, content string) (Note, error) {
	id, err := gouuid.NewV4()
	if err != nil {
		return Note{}, err
	}
	return Note{
		NoteID:  id.String(),
		UserID:  userID,
		Content: content,
	}, nil
}

// Note struct
type Note struct {
	UserID  string `json:"userId"`
	NoteID  string `json:"noteId"`
	Content string `json:"content"`
}

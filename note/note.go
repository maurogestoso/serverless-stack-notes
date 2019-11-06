package note

import gouuid "github.com/satori/go.uuid"

// New constructor
func New(userID, content, attachment string) (Note, error) {
	id, err := gouuid.NewV4()
	if err != nil {
		return Note{}, err
	}
	return Note{
		NoteID:     id.String(),
		UserID:     userID,
		Content:    content,
		Attachment: attachment,
	}, nil
}

// Note struct
type Note struct {
	UserID     string `json:"userId"`
	NoteID     string `json:"noteId"`
	Content    string `json:"content"`
	Attachment string `json:"attachment"`
}

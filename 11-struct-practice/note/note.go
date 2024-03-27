package note

import (
	"errors"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func New(title, content string) (Note, error) {

	if title == "" || content == "" {
		return Note{}, errors.New("ERROR! invalid input. Please try again.")
	}
	return Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}

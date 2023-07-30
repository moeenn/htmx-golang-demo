package todo

import (
	"github.com/google/uuid"
)

type Todo struct {
	Id       string `json:"id"`
	Text     string `json:"text"`
	Complete bool   `json:"complete"`
}

func New(text string, complete bool) Todo {
	id := uuid.New()

	return Todo{
		Id:       id.String(),
		Text:     text,
		Complete: complete,
	}
}

package model

import "github.com/gofrs/uuid/v5"

type Player struct {
	Name string `json:"name" example:"Zezima" doc:"The player name" maxLength:"12"`
	Id   uuid.UUID `json:"id" doc:"The player ID"`
}

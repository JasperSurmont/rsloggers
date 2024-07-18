package api

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
)

type Player struct {
	Name string `json:"name" example:"Zezima" doc:"The player name" maxLength:"12"`
	Id   string `json:"id" doc:"The player ID"`
}

func (c Controller) setupPlayer() {
	huma.Get(c.api, "/player/{name}", getPlayer)
	huma.Post(c.api, "/player", createPlayer)
}

type getPlayerOutput struct {
	Body struct {
		Player
	}
}

type getPlayerInput struct {
	Name string `path:"name" maxLength:"12" example:"Zezima" doc:"Name of the RuneScape character"`
}

func getPlayer(ctx context.Context, input *getPlayerInput) (*getPlayerOutput, error) {
	res := &getPlayerOutput{}
	res.Body.Player = Player{Name: input.Name, Id: "abc123"}
	return res, nil
}

type createPlayerInput struct {
	Body struct {
		Name string `json:"name" example:"Zezima" doc:"The player name" maxLength:"12"`
	}
}

type createPlayerOutput struct {
	Body struct {
		Id string `json:"id" doc:"The newly created player ID"`
	}
}

func createPlayer(ctx context.Context, input *createPlayerInput) (*createPlayerOutput, error) {
	res := &createPlayerOutput{}
	res.Body.Id = input.Body.Name + "-random"
	return res, nil
}

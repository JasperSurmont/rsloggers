package api

import (
	"context"
	"errors"
	"log/slog"
  "net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/gofrs/uuid/v5"
	"github.com/jaspersurmont/rsloggers-api/model"
	"github.com/jaspersurmont/rsloggers-api/storage"
)

type playerController struct {
  api huma.API
  store storage.PlayerStore
}

func newPlayerController(api huma.API, store storage.PlayerStore) playerController {
  pc := playerController{api: api, store: store}

	huma.Register(api, huma.Operation{
    OperationID: "get-player-by-name",
    Method: http.MethodGet,
    Path: "/player/name/{name}",
    Summary: "Get player by name",
  }, pc.getPlayerByName)

	huma.Get(api, "/player/{id}", pc.getPlayerById)

	huma.Register(api, huma.Operation{
    OperationID: "post-player",
    Method: http.MethodPost,
    Path: "/player",
    Summary: "Create a player",
  }, pc.addPlayer)

  return pc
}

type getPlayerOutput struct {
	Body struct {
		model.Player
	}
}

type getPlayerByNameInput struct {
	Name string `path:"name" maxLength:"12" example:"Zezima" doc:"Name of the RuneScape character"`
}

func (pc playerController) getPlayerByName(ctx context.Context, input *getPlayerByNameInput) (*getPlayerOutput, error) {
  res := &getPlayerOutput{}
  p, err := pc.store.GetPlayerByName(ctx, input.Name)
  if err != nil {
    if errors.Is(err, storage.ErrNotExists) {
      return nil, huma.Error404NotFound("the player was not found")
    }
    return nil, huma.Error500InternalServerError("something went wrong when looking up the player")
  }
  res.Body.Player = p
	return res, nil
}

type getPlayerByIdInput struct {
  Id string `path:"id" format:"uuid" example:"f32f0ef6-249c-4e15-9eb8-550bfaa7175d" doc:"ID of the player"`
}

func (pc playerController) getPlayerById(ctx context.Context, input *getPlayerByIdInput) (*getPlayerOutput, error) {
  res := &getPlayerOutput{}
  id, err := uuid.FromString(input.Id)
  if err != nil {
    return nil, huma.Error400BadRequest("the given id is not a valid uuid")
  }

  p, err := pc.store.GetPlayerById(ctx, id)
  if err != nil {
    if errors.Is(err, storage.ErrNotExists) {
      return nil, huma.Error404NotFound("the player was not found")
    }
    return nil, huma.Error500InternalServerError("something went wrong when looking up the player")
  }
  res.Body.Player = p
	return res, nil
}

type addPlayerInput struct {
	Body struct {
		Name string `json:"name" example:"Zezima" doc:"The player name" maxLength:"12"`
	}
}

type addPlayerOutput struct {
	Body struct {
    model.Player
	}
}

func (pc playerController) addPlayer(ctx context.Context, input *addPlayerInput) (*addPlayerOutput, error) {
	res := &addPlayerOutput{}
  p, err := pc.store.AddPlayer(ctx, input.Body.Name)
  if errors.Is(err, storage.ErrExists) {
    return nil, huma.Error400BadRequest("a player with that name already exists")
  }

  if err != nil {
    slog.Debug("internal server error", "call", "addPlayer", "name", input.Body.Name, "err", err)
    return nil, huma.Error500InternalServerError("the request failed")
  }

	res.Body.Player = p
	return res, nil
}

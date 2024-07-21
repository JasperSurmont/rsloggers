package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/gofrs/uuid/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jaspersurmont/rsloggers-api/model"
)

type PlayerStore interface {
  AddPlayer(ctx context.Context, name string) (model.Player, error)
  GetPlayerByName(ctx context.Context, name string) (model.Player, error)
  GetPlayerById(ctx context.Context, id uuid.UUID) (model.Player, error)
  DeletePlayer(ctx context.Context, id uuid.UUID) error
  UpdatePlayer(ctx context.Context, id uuid.UUID, name string) (model.Player, error)
}

type playerStore struct {
  db *pgxpool.Pool
}

func newPlayerStore(db *pgxpool.Pool) playerStore {
  return playerStore{db: db}
}

func (ps playerStore) AddPlayer(ctx context.Context, name string) (model.Player, error) {
  var exists bool
  err := ps.db.QueryRow(ctx, "SELECT EXISTS (SELECT 1 FROM players WHERE name = $1)", name).Scan(&exists)
  if exists {
    return model.Player{}, fmt.Errorf("cannot add player (%s): %w", name, ErrExists)
  }
  if err != nil {
    return model.Player{}, fmt.Errorf("could not check existence of player: %w", err)
  }

  id, err := uuid.NewV4()
  if err != nil {
    return model.Player{}, fmt.Errorf("could not create uuid: %w", err)
  }

  _, err = ps.db.Exec(ctx, "INSERT INTO players (id, name) VALUES ($1, $2)", id, name)
  if err != nil {
    return model.Player{}, fmt.Errorf("could not add player to db: %w", err)
  }

  return model.Player{Id: id, Name: name}, nil
}

func (ps playerStore) GetPlayerByName(ctx context.Context, name string) (model.Player, error) {
  var id uuid.UUID
  err := ps.db.QueryRow(ctx, "SELECT id FROM players WHERE name = $1", name).Scan(&id)
  if errors.Is(err, pgx.ErrNoRows) {
    return model.Player{}, ErrNotExists
  }
  if err != nil {
    return model.Player{}, fmt.Errorf("could not get player by name: %w", err)
  }
  return model.Player{Id: id, Name: name}, nil
}

func (ps playerStore) GetPlayerById(ctx context.Context, id uuid.UUID) (model.Player, error) {
  var name string
  err := ps.db.QueryRow(ctx, "SELECT name FROM players WHERE id = $1", id).Scan(&name)
  if errors.Is(err, pgx.ErrNoRows) {
    return model.Player{}, ErrNotExists
  }
  if err != nil {
    return model.Player{}, fmt.Errorf("could not get player by id: %w", err)
  }
  return model.Player{Id: id, Name: name}, nil
}

func (ps playerStore) DeletePlayer(ctx context.Context, id uuid.UUID) error {
  return nil
}

func (ps playerStore) UpdatePlayer(ctx context.Context, id uuid.UUID, name string) (model.Player, error) {
  return model.Player{}, nil
}

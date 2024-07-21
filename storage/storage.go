package storage

import (
	"fmt"
	"os"
  "context"

	"github.com/jackc/pgx/v5/pgxpool"
)


type StoreProvider struct {
  pool *pgxpool.Pool

  PlayerStore PlayerStore 
}

func NewStoreProvider() (*StoreProvider, error)  {
  sp := StoreProvider{}

  err := sp.setup()
  if err != nil {
    return nil, err
  }

  // Add the different stores here
  sp.PlayerStore = newPlayerStore(sp.pool)

  return &sp, nil
}

func (sp *StoreProvider) setup()  error {
  pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
  if err != nil {
    return fmt.Errorf("unable to connect to database: %w", err) 
  }

  sp.pool = pool

  return nil
}

func (sp *StoreProvider) TearDown() {
  sp.pool.Close()
}

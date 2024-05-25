package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/dagulv/train-api/internal/env"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/webmafia/pg"
)

func Connect(ctx context.Context, env env.Env) (*pg.DB, error) {
	pool, err := pgxpool.New(ctx, env.DatabaseUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}

	return pg.NewDB(pool), err
}

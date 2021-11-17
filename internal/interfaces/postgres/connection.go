package postgres

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"main/internal/configuration"
	"sync"
)

var once sync.Once

type PGCon struct {
	Pool *pgxpool.Pool
}

var (
	instance *PGCon
)

func GetPGCon() *PGCon {
	once.Do(func() {
		cfg := configuration.GetConfig()
		var err error

		config, err := pgxpool.ParseConfig(cfg.DB.DSN())
		if err != nil {
			panic(err)
		}
		config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
			return nil
		}

		pool, err := pgxpool.ConnectConfig(context.Background(), config)

		if err != nil {
			panic(err)
		}

		instance = &PGCon{Pool: pool}
	})

	return instance
}

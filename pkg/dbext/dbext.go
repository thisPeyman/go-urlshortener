package dbext

import (
	"context"

	"github.com/jackc/pgx/v5"
	"go.uber.org/fx"
)

type DBConfig struct {
	PostgresDSN string `mapstructure:"POSTGRES_DSN"`
}

type DBConfigGetter interface {
	GetDBConfig() *DBConfig
}

func (c *DBConfig) GetDBConfig() *DBConfig {
	return c
}

func ProvideDatabase(ctx context.Context, lc fx.Lifecycle, cfgGetter DBConfigGetter) (*pgx.Conn, error) {
	cfg := cfgGetter.GetDBConfig()
	conn, err := pgx.Connect(ctx, cfg.PostgresDSN)
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return conn.Close(ctx)
		},
	})

	return conn, nil
}

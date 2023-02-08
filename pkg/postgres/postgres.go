package postgres

import (
	"context"
	"fmt"
	"net/url"

	"sales-project/cmd/config"

	"github.com/jackc/pgx/v4"
)

func NewConn(cfg config.Config) (*pgx.Conn, error) {
	dsn := url.URL{
		Scheme: "postgres",
		Host:   fmt.Sprintf("%s:%s", cfg.DbHost, cfg.DbPort),
		User:   url.UserPassword(cfg.DbUser, cfg.DbPassword),
		Path:   cfg.DbName,
	}

	q := dsn.Query()
	q.Add("sslmode", "disable")

	dsn.RawQuery = q.Encode()

	conn, err := pgx.Connect(context.Background(), dsn.String())
	if err != nil {
		return nil, fmt.Errorf("pgx.Connect%w", err)
	}

	return conn, nil
}

package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/url"
	"strconv"

	"github.com/hermeznetwork/price-updater-service/config"
	_ "github.com/lib/pq"
)

type Connection struct {
	db *sql.DB
}

func NewConnection(ctx context.Context, cfg *config.PostgresConfig) *Connection {
	dbURL := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cfg.User, cfg.Password),
		Host:   net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port)),
		Path:   cfg.Database,
	}
	q := dbURL.Query()

	if !cfg.SslModeEnabled {
		q.Add("sslmode", "disable")
	}

	q.Add("connect_timeout", "600")
	dbURL.RawQuery = q.Encode()
	dbURLStr := dbURL.String()

	db, err := sql.Open("postgres", dbURLStr)
	if err != nil {
		panic(fmt.Sprintf("can't connect to postgres error: %s", err.Error()))
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)

	return &Connection{
		db: db,
	}
}

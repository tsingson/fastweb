package pgxutils

import (
	"github.com/jackc/pgx"
)

type (
	PostgresConfig struct {
		User     string     `json:"User"`
		Password string     `json:"Password"`
		Database string     `json:"Database"`
		Port     int        `json:"Port"`
		Host     string     `json:"Host"`
		Logger   pgx.Logger `json:"-"`
	}
)

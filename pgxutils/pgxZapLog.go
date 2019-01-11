package pgxutils

import (
	"time"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/log/zapadapter"
)

const (
// PgxMaxConnect = 5
)

type (
	PostgresZapConfig struct {
		User     string             `json:"User"`
		Password string             `json:"Password"`
		Database string             `json:"Database"`
		Port     int                `json:"Port"`
		Host     string             `json:"Host"`
		Logger   *zapadapter.Logger `json:"-"`
	}
)

func NewPgxPoolAfterConnectZap(config *PostgresZapConfig, afterConnectMap func(*pgx.Conn) error, debug bool) (*pgx.ConnPool, error) {

	var pgxConfig pgx.ConnConfig

	pgxConfig = pgx.ConnConfig{
		Host:     config.Host,
		User:     config.User,
		Password: config.Password,
		Database: config.Database,
		Logger:   config.Logger,
		LogLevel: pgx.LogLevelDebug, // pgx.LogLevelInfo,pgx.LogLevelInfo, // pgx.LogLevelError,
	}
	if debug {
		pgxConfig.LogLevel = pgx.LogLevelDebug

	} else {
		pgxConfig.LogLevel = pgx.LogLevelError

	}

	connPoolConfig := pgx.ConnPoolConfig{
		ConnConfig:     pgxConfig,
		MaxConnections: PgxMaxConnect,
		AfterConnect:   afterConnectMap,
		AcquireTimeout: 5 * time.Second,
	}
	return pgx.NewConnPool(connPoolConfig)
}

func NewPgxPool(config *PostgresConfig) (*pgx.ConnPool, error) {
	// logger := zerologadapter.NewLogger(log)
	pgxConfig := pgx.ConnConfig{
		Host:     config.Host,
		User:     config.User,
		Password: config.Password,
		Database: config.Database,
		Logger:   config.Logger,
		LogLevel: pgx.LogLevelError,
	}
	connPoolConfig := pgx.ConnPoolConfig{
		ConnConfig:     pgxConfig,
		MaxConnections: PgxMaxConnect,
		// 	AfterConnect:   afterConnectMap,
	}
	return pgx.NewConnPool(connPoolConfig)
}

package pgxutils

import (
	"errors"
	"time"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/log/zerologadapter"
	"github.com/rs/zerolog"
)

const (
	PgxMaxConnect = 5
)

func (postgresConfig *PostgresConfig) SetLogger(logger zerolog.Logger) {
	//	postgresConfig := &config.CmsPostgresConfig
	//	zerolog.TimeFieldFormat = ""
	//	log = zerolog.New(os.Stdout)
	postgresConfig.Logger = zerologadapter.NewLogger(logger)
}

func AfterConnectMap(conn *pgx.Conn) error {
	var err error
	prepare := map[string]string{
		"insertTerminal": `INSERT  into terminal.valid_terminal (serial, code,filename )  VALUES ($1,$2,$3)`}

	if len(prepare) == 0 {
		err = errors.New("null map ")
		return err
	}
	for key, value := range prepare {
		_, err = conn.Prepare(key, value)
		if err != nil {
			return err
		}
	}
	return nil

}

func NewPgxPoolAfterConnect(config *PostgresConfig, afterConnectMap func(*pgx.Conn) error, debug bool) (*pgx.ConnPool, error) {
	//logger := zerologadapter.NewLogger(log)

	var pgxConfig pgx.ConnConfig

	if debug {
		pgxConfig = pgx.ConnConfig{
			Host:     config.Host,
			User:     config.User,
			Password: config.Password,
			Database: config.Database,
			Logger:   config.Logger,
			LogLevel: pgx.LogLevelDebug, //pgx.LogLevelInfo,pgx.LogLevelInfo, // pgx.LogLevelError,
		}
	} else {
		pgxConfig = pgx.ConnConfig{
			Host:     config.Host,
			User:     config.User,
			Password: config.Password,
			Database: config.Database,
			Logger:   config.Logger,
			LogLevel: pgx.LogLevelError, //pgx.LogLevelInfo,pgx.LogLevelInfo, // pgx.LogLevelError,
		}
	}

	connPoolConfig := pgx.ConnPoolConfig{
		ConnConfig:     pgxConfig,
		MaxConnections: PgxMaxConnect,
		AfterConnect:   afterConnectMap,
		AcquireTimeout: 5 * time.Second,
	}
	return pgx.NewConnPool(connPoolConfig)
}

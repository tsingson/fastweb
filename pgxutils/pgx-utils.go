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

/**
func NewPgxPool(config *PostgresConfig) (*pgx.ConnPool, error) {
	//logger := zerologadapter.NewLogger(log)
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
		//	AfterConnect:   afterConnectMap,
	}
	return pgx.NewConnPool(connPoolConfig)
}
*/
/**

// Field is ignored by this package.
Field int `json:"-"`

// Field appears in JSON as key "myName".
Field int `json:"myName"`

// Field appears in JSON as key "myName" and
// the field is omitted from the object if its value is empty,
// as defined above.
Field int `json:"myName,omitempty"`

// Field appears in JSON as key "Field" (the default), but
// the field is skipped if empty.
// Note the leading comma.
Field int `json:",omitempty"`
*/

/**

conn, err := pool.Acquire()
if err != nil {
	return err
}

_, err = conn.SetLogLevel(pgx.LogLevelTrace)
if err != nil {
	return err
}

*/

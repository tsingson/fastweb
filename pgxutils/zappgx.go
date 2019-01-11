package pgxutils

import (
	"github.com/jackc/pgx/log/zapadapter"
	"go.uber.org/zap"
)

func (postgresConfig *PostgresConfig) SetZapLogger(logger *zap.Logger) {
	// 	postgresConfig := &config.CmsPostgresConfig
	// 	zerolog.TimFieldFormat = ""
	// 	log = zerolog.New(os.Stdout)
	postgresConfig.Logger = zapadapter.NewLogger(logger)
}

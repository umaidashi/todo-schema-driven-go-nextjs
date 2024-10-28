package db

import (
	"database/sql"
	"server/common/config"
	"server/pkg/ent"

	entsql "entgo.io/ent/dialect/sql"

	"entgo.io/ent/dialect"
)

func Init() *ent.Client {
	db, err := sql.Open("pgx", config.Config.DB_DSN)
	if err != nil {
		panic(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

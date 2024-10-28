package db

import (
	"context"
	"database/sql"
	"server/common/config"
	"server/pkg/ent"

	entsql "entgo.io/ent/dialect/sql"

	"entgo.io/ent/dialect"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Init() *ent.Client {
	db, err := sql.Open("pgx", config.Config.DB_DSN)
	if err != nil {
		panic(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(drv))

	if err := client.Schema.Create(context.Background()); err != nil {
		panic(err)
	}

	return client
}

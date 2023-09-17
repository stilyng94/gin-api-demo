package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/gin-api-demo/ent"
	_ "modernc.org/sqlite"
)

func connectToDriver(dsn string) (*entsql.Driver, error) {
	fmt.Println(dsn)
	dbConn, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err

	}
	drv := entsql.OpenDB(dialect.SQLite, dbConn)
	return drv, nil

}

// OpenDB - Connect to database and run migrations
func OpenDB(dsn string) *ent.Client {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	drv, err := connectToDriver(dsn)
	if err != nil {
		log.Fatalf("connectToDriver: %v", err)
	}
	client := ent.NewClient(ent.Driver(drv))
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx, schema.WithDropIndex(true),
		schema.WithDropColumn(true),
		schema.WithForeignKeys(true),
		schema.WithGlobalUniqueID(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

package infrastructure

import (
	"context"

	"entgo.io/ent/dialect"
	"github.com/Ras96/go-clean-architecture-template/internal/interfaces/repository/ent"
	"github.com/Ras96/go-clean-architecture-template/internal/interfaces/repository/ent/migrate"
	"github.com/Ras96/go-clean-architecture-template/pkg/errors"

	// import mysql driver
	_ "github.com/go-sql-driver/mysql"
)

func SetupEntClient() (*ent.Client, error) {
	c, err := ent.Open(dialect.MySQL, "root:password@tcp(localhost:3306)/myapp?parseTime=true")
	if err != nil {
		return nil, errors.Wrap(err, "failed to open database")
	}

	if err := c.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(true),
	); err != nil {
		return nil, errors.Wrap(err, "failed to create schema resources")
	}

	return c, nil
}

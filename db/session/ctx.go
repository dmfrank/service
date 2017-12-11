package session

import (
	"context"

	"github.com/go-pg/pg"
)

const (
	keyRequest int = iota
	keyDatabase
	keyUser
)

func WithDatabase(
	ctx context.Context,
	database *pg.DB) context.Context {
	return context.WithValue(ctx, keyDatabase, database)
}

func Database(ctx context.Context) *pg.DB {
	return ctx.Value(keyDatabase).(*pg.DB)
}

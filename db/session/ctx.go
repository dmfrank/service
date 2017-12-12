package session

import (
	"context"

	"github.com/go-pg/pg"
)

type key struct {
	int
}

// WithDatabase sets context value for database
func WithDatabase(ctx context.Context, database *pg.DB) context.Context {
	return context.WithValue(ctx, key{}, database)
}

// Database constructs context value for database
func Database(ctx context.Context) *pg.DB {
	return ctx.Value(key{}).(*pg.DB)
}

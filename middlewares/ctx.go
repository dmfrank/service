package middlewares

import (
	"context"

	"github.com/dmfrank/service/db/session"
	"github.com/dmfrank/service/models"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
)

const keyContext = "context"
const keyConfig = "config"

func WithDatabase(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx = session.WithDatabase(ctx, db)
		c.Set(keyContext, ctx)
		c.Next()
	}
}

func WithConfig(c *gin.Context, config *models.Config) {
	ctx := Context(c)
	ctx = context.WithValue(ctx, keyConfig, config)
	c.Set(keyContext, ctx)
}

func Context(c *gin.Context) context.Context {
	return c.MustGet(keyContext).(context.Context)
}

func Config(c *gin.Context) *models.Config {
	ctx := Context(c)
	return ctx.Value(keyConfig).(*models.Config)
}

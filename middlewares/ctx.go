package middlewares

import (
	"context"

	"github.com/dmfrank/service/db/session"
	"github.com/dmfrank/service/models"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
)

const keyCfg = "cfg"
const keyCtx = "ctx"

// WithDatabase sets database session to Context
func WithDatabase(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(keyCtx, session.WithDatabase(c.Request.Context(), db))
		c.Next()
	}
}

// WithConfig sets config to Context
func WithConfig(c *gin.Context, config *models.Config) {
	c.Set(keyCtx, context.WithValue(Context(c), keyCfg, config))
}

func Context(c *gin.Context) context.Context {
	return c.MustGet(keyCtx).(context.Context)
}

func Config(c *gin.Context) *models.Config {
	return Context(c).Value(keyCfg).(*models.Config)
}

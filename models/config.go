package models

import (
	"context"
	"time"

	"github.com/dmfrank/service/db/session"
)

// Config entity
type Config struct {
	Id        string
	Type      string
	Data      string
	Host      string
	Port      string
	Database  string
	User      string
	Password  string
	Schema    string
	VHost     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// FindConfigByTypeAndData fetches configuration data by `type` & `data` fields
func FindConfigByTypeAndData(ctx context.Context, typ, data string) (*Config, error) {
	var config Config
	_, err := session.Database(ctx).
		QueryOne(&config, `SELECT * FROM configurations WHERE t = ? AND data = ?`, typ, data)
	return &config, err
}

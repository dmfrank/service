package views

import (
	"testing"

	"github.com/dmfrank/service/models"

	"github.com/stretchr/testify/assert"
)

func TestBuildConfigView(t *testing.T) {
	mc := models.Config{}
	assert.Equal(t, ConfigView{}, BuildConfigView(&mc))
	mc.Host = "host"
	mc.Port = "port"
	mc.Vhost = "vhost"
	mc.Db = "database"
	mc.U = "user"
	mc.Pass = "password"
	mc.Sch = "schema"
	assert.Equal(
		t,
		ConfigView{
			Host:     "host",
			Port:     "port",
			VHost:    "vhost",
			Database: "database",
			User:     "user",
			Password: "password",
			Schema:   "schema"}, BuildConfigView(&mc))
}

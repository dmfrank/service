package views

import "github.com/dmfrank/service/models"

// ConfigView response entity
type ConfigView struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	VHost    string `json:"virtualhost,omitempty"`
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
	Schema   string `json:"schema"`
}

// BuildConfigView constructs response
func BuildConfigView(config *models.Config) ConfigView {
	return ConfigView{
		Host:     config.Host,
		Port:     config.Port,
		VHost:    config.Vhost,
		Database: config.Db,
		User:     config.U,
		Password: config.Pass,
		Schema:   config.Sch,
	}
}

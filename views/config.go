package views

import "github.com/dmfrank/service/models"

type ConfigView struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	VHost    string `json:"virtualhost, noempty"`
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
	Schema   string `json:"schema"`
}

func BuildConfigView(config *models.Config) ConfigView {
	return ConfigView{
		Host:     config.Host,
		Port:     config.Port,
		VHost:    config.VHost,
		Database: config.Database,
		User:     config.User,
		Password: config.Password,
		Schema:   config.Schema,
	}
}

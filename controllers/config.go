package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dmfrank/service/middlewares"
	"github.com/dmfrank/service/models"
	"github.com/dmfrank/service/views"
)

type config struct{}

type configRequest struct {
	Type string `json:"Type"`
	Data string `json:"Data"`
}

// Config sets configuration endpoint
func Config(router *gin.Engine) {
	c := &config{}
	router.GET("/", c.index)
	router.POST("/config", c.show)
}

func (cfg *config) index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

func (cfg *config) show(c *gin.Context) {
	var body configRequest
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else if config, e := models.FindConfigByTypeAndData(
		middlewares.Context(c),
		body.Type,
		body.Data); config == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
	} else {
		c.JSON(http.StatusOK, views.BuildConfigView(config))
	}
}

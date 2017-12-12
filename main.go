package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dmfrank/service/controllers"
	"github.com/dmfrank/service/middlewares"
	"github.com/dmfrank/service/models"
	"github.com/gin-gonic/gin"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	ip   = "0.0.0.0"
	port = "8080"
)

func main() {
	router := gin.Default()
	router.Use(middlewares.WithDatabase(models.InitDb()))

	controllers.Config(router)

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", ip, port),
		Handler: router,
	}
	log.Fatal(server.ListenAndServe())
}

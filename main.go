package main

import (
	"log"
	"net/http"

	"github.com/dmfrank/service/controllers"
	"github.com/dmfrank/service/middlewares"
	"github.com/dmfrank/service/models"
	"github.com/gin-gonic/gin"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	router := gin.Default()

	router.Use(middlewares.WithDatabase(models.InitDb()))
	router.GET("/", index)

	controllers.Config(router)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
}

func index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "hello world"})
	log.Println("hello index")
}

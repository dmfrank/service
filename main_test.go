package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dmfrank/service/controllers"
	"github.com/dmfrank/service/middlewares"
	"github.com/dmfrank/service/models"
	"github.com/dmfrank/service/views"
	"github.com/gin-gonic/gin"
)

func getRequest(t *testing.T) {
	urlStr := fmt.Sprintf("http://%s:%s", ip, port)
	req, err := http.NewRequest("GET", urlStr, nil)

	assert.Nil(t, err)

	client := &http.Client{}
	resp, err := client.Do(req)

	assert.Nil(t, err)

	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)
}

type Params struct {
	postVar string
}

func postRequest(t *testing.T, data string) (*views.ConfigView, error) {
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("http://%s:%s/config", ip, port),
		bytes.NewBufferString(data))
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	assert.Nil(t, err)

	defer resp.Body.Close()

	if assert.Equal(t, 200, resp.StatusCode) {
		w := new(views.ConfigView)
		data, err := ioutil.ReadAll(resp.Body)
		assert.Nil(t, err)
		if assert.Nil(t, json.Unmarshal(data, w)) {
			return w, nil
		}
	}
	return nil, err
}

func TestService(t *testing.T) {
	go func() {
		router := gin.Default()
		router.Use(middlewares.WithDatabase(models.InitDb()))

		controllers.Config(router)

		server := &http.Server{
			Addr:    fmt.Sprintf("%s:%s", ip, port),
			Handler: router,
		}
		log.Fatal(server.ListenAndServe())
	}()

	getRequest(t)

	rawReqA := fmt.Sprint(`
		{
			"Type": "Test.vpn",
			"Data": "Rabbit.log"
		}`)
	rawReqB := fmt.Sprint(`
		{
			"Type": "Develop.mr_robot",
			"Data": "Database.processing"
		}`)

	respA := &views.ConfigView{
		Host:     "10.0.5.42",
		Port:     "5671",
		VHost:    "/",
		Database: "devdb",
		User:     "guest",
		Password: "guest",
		Schema:   "public",
	}
	respB := &views.ConfigView{
		Host:     "localhost",
		Port:     "5432",
		VHost:    "",
		Database: "devdb",
		User:     "mr_robot",
		Password: "secret",
		Schema:   "public"}

	resp, e := postRequest(t, rawReqA)
	assert.Nil(t, e)
	assert.Equal(t, respA, resp)

	resp, e = postRequest(t, rawReqB)
	assert.Nil(t, e)
	assert.Equal(t, respB, resp)
}

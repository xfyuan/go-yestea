package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/xfyuan/go-yestea/cmd/yestea/app"
	"github.com/xfyuan/go-yestea/cmd/yestea/gspec"
	"github.com/xfyuan/go-yestea/cmd/yestea/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTodo(t *testing.T) {
	app.DB = gspec.ResetDB()

	expectTitle := "Ruby"
	expectDescription := "Beautiful Language"
	app.DB.Create(&models.Todo{
		Title:       expectTitle,
		Description: expectDescription,
	})

	router := gspec.NewRouter()
	router.Handle("GET", "/todos/:id", GetTodo)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/todos/1", bytes.NewBufferString(""))
	router.ServeHTTP(res, req)

	body, _ := ioutil.ReadAll(res.Body)
	var expect map[string]interface{}
	if err := json.Unmarshal(body, &expect); err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, expectTitle, expect["title"])
	assert.Equal(t, expectDescription, expect["description"])
}

func TestGetTodo_NotFound(t *testing.T) {
	app.DB = gspec.ResetDB()

	expectTitle := "Ruby"
	expectDescription := "Beautiful Language"
	app.DB.Create(&models.Todo{
		Title:       expectTitle,
		Description: expectDescription,
	})

	router := gspec.NewRouter()
	router.Handle("GET", "/todos/:id", GetTodo)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/todos/9999", bytes.NewBufferString(""))
	router.ServeHTTP(res, req)

	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, http.StatusNotFound, res.Code)
	assert.Empty(t, string(body))
}

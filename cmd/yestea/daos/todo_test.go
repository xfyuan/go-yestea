package daos

import (
	"github.com/stretchr/testify/assert"
	"github.com/xfyuan/go-yestea/cmd/yestea/app"
	"github.com/xfyuan/go-yestea/cmd/yestea/gspec"
	"github.com/xfyuan/go-yestea/cmd/yestea/models"
	"testing"
)

func TestTodoDAO_Get(t *testing.T) {
	app.DB = gspec.ResetDB()
	app.DB.Create(&models.Todo{
		Title:       "Golang",
		Description: "A programming language",
	})
	dao := NewTodoDAO()

	todo, err := dao.Get(1)

	assert.Nil(t, err)
	assert.Equal(t, "Golang", todo.Title)
	assert.Equal(t, "A programming language", todo.Description)
}

func TestTodoDAO_Get_NotFound(t *testing.T) {
	app.DB = gspec.ResetDB()
	dao := NewTodoDAO()

	todo, err := dao.Get(99999)

	assert.NotNil(t, err)
	assert.Equal(t, "", todo.Title)
	assert.Equal(t, "", todo.Description)
}

package daos

import (
	"github.com/xfyuan/go-yestea/cmd/yestea/app"
	"github.com/xfyuan/go-yestea/cmd/yestea/models"
)

type TodoDAO struct{}

func NewTodoDAO() *TodoDAO {
	return &TodoDAO{}
}

func (dao *TodoDAO) Get(id uint) (*models.Todo, error) {
	var todo models.Todo

	err := app.DB.Where("id = ?", id).
		First(&todo).
		Error

	return &todo, err
}

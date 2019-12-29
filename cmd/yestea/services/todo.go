package services

import "github.com/xfyuan/go-yestea/cmd/yestea/models"

type todoDAO interface {
	Get(id uint) (*models.Todo, error)
}

type TodoService struct {
	dao todoDAO
}

func NewTodoService(dao todoDAO) *TodoService {
	return &TodoService{dao}
}

func (s *TodoService) Get(id uint) (*models.Todo, error) {
	return s.dao.Get(id)
}

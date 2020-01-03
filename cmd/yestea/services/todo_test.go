package services

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/xfyuan/go-yestea/cmd/yestea/models"
	"testing"
)

func TestNewTodoService(t *testing.T) {
	dao := newMockTodoDAO()
	s := NewTodoService(dao)
	assert.Equal(t, dao, s.dao)
}

func TestTodoService_Get(t *testing.T) {
	s := NewTodoService(newMockTodoDAO())
	todo, err := s.Get(2)

	assert.Nil(t, err)
	assert.NotNil(t, todo)
	assert.Equal(t, "Elixir", todo.Title)
	assert.Equal(t, "Greatest Language", todo.Description)
}

func TestTodoService_GetNotFound(t *testing.T) {
	s := NewTodoService(newMockTodoDAO())
	_, err := s.Get(9999)

	assert.NotNil(t, err)
}

type mockTodoDAO struct {
	records []models.Todo
}

func (m *mockTodoDAO) Get(id uint) (*models.Todo, error) {
	for _, record := range m.records {
		if record.ID == id {
			return &record, nil
		}
	}
	return nil, errors.New("record not found")
}

func newMockTodoDAO() todoDAO {
	return &mockTodoDAO{
		records: []models.Todo{
			{Model: gorm.Model{ID: 1}, Title: "Ruby", Description: "Beautiful Language"},
			{Model: gorm.Model{ID: 2}, Title: "Elixir", Description: "Greatest Language"},
		},
	}
}

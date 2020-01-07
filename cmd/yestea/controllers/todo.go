package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xfyuan/go-yestea/cmd/yestea/daos"
	"github.com/xfyuan/go-yestea/cmd/yestea/services"
	"log"
	"net/http"
	"strconv"
)

// GetUser godoc
// @Summary Retrieves todo based on given ID
// @Produce json
// @Param id path integer true "Todo ID"
// @Success 200 {object} models.Todo
// @Router /users/{id} [get]
// @Security ApiKeyAuth
func GetTodo(c *gin.Context) {
	s := services.NewTodoService(daos.NewTodoDAO())
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if todo, err := s.Get(uint(id)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

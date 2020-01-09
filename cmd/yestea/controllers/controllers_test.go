package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xfyuan/go-yestea/cmd/yestea/app"
	"github.com/xfyuan/go-yestea/cmd/yestea/gspec"
	"github.com/xfyuan/go-yestea/cmd/yestea/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Controllers", func() {
	var (
		router *gin.Engine
	)

	BeforeEach(func() {
		router = gspec.NewRouter()
		router.Use(gspec.SetAuthHeader())

		app.DB = gspec.ResetDB()
	})

	Describe("#GetTodo", func() {
		Context("with record", func() {
			BeforeEach(func() {
				app.DB.Create(&models.Todo{
					Title:       "Golang",
					Description: "Google's Language",
				})
			})

			It("should get a todo successfully", func() {
				router.Handle("GET", "/todos/:id", GetTodo)
				res := httptest.NewRecorder()
				req, _ := http.NewRequest("GET", "/todos/1", nil)
				router.ServeHTTP(res, req)

				body, _ := ioutil.ReadAll(res.Body)
				var expect map[string]interface{}
				if err := json.Unmarshal(body, &expect); err != nil {
					panic(err)
				}

				Expect(res.Code).To(Equal(http.StatusOK))
				Expect(expect["title"]).To(Equal("Golang"))
				Expect(expect["description"]).To(Equal("Google's Language"))
			})
		})

		Context("without record", func() {
			It("should return not found", func() {
				router.Handle("GET", "/todos/:id", GetTodo)
				res := httptest.NewRecorder()
				req, _ := http.NewRequest("GET", "/todos/1", nil)
				router.ServeHTTP(res, req)

				body, _ := ioutil.ReadAll(res.Body)

				Expect(res.Code).To(Equal(http.StatusNotFound))
				Expect(body).To(BeEmpty())
			})
		})
	})
})

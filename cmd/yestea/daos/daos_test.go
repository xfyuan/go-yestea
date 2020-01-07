package daos

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xfyuan/go-yestea/cmd/yestea/app"
	"github.com/xfyuan/go-yestea/cmd/yestea/gspec"
	"github.com/xfyuan/go-yestea/cmd/yestea/models"
)

var _ = Describe("Daos", func() {
	var (
		dao  *TodoDAO
		todo *models.Todo
		err  error
	)

	BeforeEach(func() {
		app.DB = gspec.ResetDB()
	})

	Describe("Daos with todo records", func() {
		Context("when exists", func() {
			BeforeEach(func() {
				app.DB.Create(&models.Todo{
					Title:       "Golang",
					Description: "A programming language",
				})
				dao = NewTodoDAO()
			})

			JustBeforeEach(func() {
				todo, err = dao.Get(1)
			})

			It("should has no error", func() {
				Expect(err).To(BeNil())
			})

			It("should has correct record", func() {
				Expect(todo.Title).To(Equal("Golang"))
				Expect(todo.Description).To(Equal("A programming language"))
			})
		})

		Context("when not exists", func() {
			JustBeforeEach(func() {
				todo, err = dao.Get(9999)
			})

			It("should has error", func() {
				Expect(err).To(HaveOccurred())
			})

			It("should has empty record", func() {
				Expect(todo.Title).To(BeEmpty())
				Expect(todo.Description).To(BeEmpty())
			})
		})
	})
})

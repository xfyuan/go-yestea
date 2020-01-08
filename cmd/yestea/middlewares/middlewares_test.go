package middlewares

import (
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xfyuan/go-yestea/cmd/yestea/gspec"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Middlewares", func() {
	var (
		router *gin.Engine
	)

	BeforeEach(func() {
		router = gspec.NewRouter()
	})

	Describe("Auth middleware", func() {
		Context("without authorization header", func() {
			It("should not pass", func() {
				router.GET("/", Auth())
				res := httptest.NewRecorder()
				req, _ := http.NewRequest("GET", "/", nil)
				router.ServeHTTP(res, req)

				Expect(res.Code).To(Equal(http.StatusUnauthorized))
			})
		})

		Context("with api key in authorization header", func() {
			It("should pass", func() {
				router.GET("/", gspec.SetAuthHeader(), Auth())
				res := httptest.NewRecorder()
				req, _ := http.NewRequest("GET", "/", nil)
				router.ServeHTTP(res, req)

				Expect(res.Code).To(Equal(http.StatusOK))
			})
		})
	})
})

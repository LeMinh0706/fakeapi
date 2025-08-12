package benchmark

import (
	"github.com/LeMinh0706/fakeapi/pkg/response"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.RouterGroup, s *BenchmarkService) {
	handler := &BenchmarkHandler{s: s}
	r.GET("", handler.GetBenchmark)
	r.GET("/results", handler.GetResults)
}

type BenchmarkHandler struct {
	s *BenchmarkService
}

func (h *BenchmarkHandler) GetBenchmark(c *gin.Context) {
	response.SuccessResponse(c, 20000, benchmarkData)
}

func (h *BenchmarkHandler) GetResults(c *gin.Context) {
	results, err := h.s.GetResults(c)
	if err != nil {
		response.ErrorNonKnow(c, 50000, err.Error())
		return
	}
	response.SuccessResponse(c, 20000, results)
}

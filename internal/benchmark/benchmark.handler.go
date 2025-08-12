package benchmark

import (
	"strconv"

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
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		response.ErrorResponse(c, 40000)
		return
	}
	res1, res2 := BenchmarkResult()
	if page > 1 {
		response.SuccessResponse(c, 20000, res1)
		return
	}
	response.SuccessResponse(c, 20000, res2)
}

func (h *BenchmarkHandler) GetResults(c *gin.Context) {
	results, err := h.s.GetResults(c)
	if err != nil {
		response.ErrorNonKnow(c, 50000, err.Error())
		return
	}
	response.SuccessResponse(c, 20000, results)
}

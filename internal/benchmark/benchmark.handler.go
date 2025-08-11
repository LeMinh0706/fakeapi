package benchmark

import (
	"github.com/LeMinh0706/fakeapi/pkg/response"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.RouterGroup) {
	handler := &BenchmarkHandler{}
	r.GET("", handler.GetBenchmark)
}

type BenchmarkHandler struct {
}

func (h *BenchmarkHandler) GetBenchmark(c *gin.Context) {
	response.SuccessResponse(c, 20000, benchmarkData)
}

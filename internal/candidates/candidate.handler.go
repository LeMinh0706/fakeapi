package candidates

import (
	"github.com/LeMinh0706/fakeapi/pkg/response"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.RouterGroup, s *CandidateService) {
	handler := &CandidateHandler{s: s}
	r.GET("", handler.GetCandidate)
}

type CandidateHandler struct {
	s *CandidateService
}

func (h *CandidateHandler) GetCandidate(c *gin.Context) {
	candidates, err := h.s.GetCandidates(c)
	if err != nil {
		response.ErrorNonKnow(c, 50000, err.Error())
		return
	}
	response.SuccessResponse(c, 20000, candidates)
}

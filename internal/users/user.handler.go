package users

import (
	"github.com/LeMinh0706/fakeapi/pkg/response"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.RouterGroup, s *UserService) {
	handler := &UserHandler{s: s}
	r.POST("/register", handler.RegisterUser)
}

type UserHandler struct {
	s *UserService
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var req *UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, 40000)
		return
	}
	if err := h.s.CreateUser(c, req.Username, req.Password); err != nil {
		response.ErrorResponse(c, 40000)
		return
	}
	response.SuccessResponse(c, 20000, gin.H{})
}

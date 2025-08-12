package users

import (
	"github.com/LeMinh0706/fakeapi/config"
	"github.com/LeMinh0706/fakeapi/pkg/response"
	"github.com/LeMinh0706/fakeapi/pkg/token/token"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.RouterGroup, s *UserService, jwt token.Maker, pas token.Maker) {
	handler := &UserHandler{s: s, jwt: jwt, pas: pas}
	r.POST("/register", handler.RegisterUser)
	r.POST("/login", handler.LoginUser)
	r.GET("/gmail", handler.GetGmailUser)
}

type UserHandler struct {
	s   *UserService
	jwt token.Maker
	pas token.Maker
}

func (h *UserHandler) GetGmailUser(c *gin.Context) {
	response.SuccessResponse(c, 20000, gin.H{
		"users": []GetEmail{
			{ID: 1, Email: "huudat25072003@gmail.com"},
			{ID: 2, Email: "qbulord0706@gmail.com"},
		},
	})
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var req *UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, 40000)
		return
	}

	if err := h.s.CreateUser(c, req.Username, req.Password); err != nil {
		response.ErrorNonKnow(c, 50000, err.Error())
		return
	}

	response.SuccessResponse(c, 20000, gin.H{})
}

func (h *UserHandler) LoginUser(c *gin.Context) {
	var req *UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, 40000)
		return
	}

	id, err := h.s.GetUserByUsername(c, req.Username, req.Password)
	if err != nil {
		response.ErrorResponse(c, 40401)
		return
	}

	access, err := h.jwt.CreateToken(id, config.Cfg.Token.AccessTTL)
	if err != nil {
		response.ErrorResponse(c, 40000)
		return
	}

	refresh, err := h.pas.CreateToken(id, config.Cfg.Token.RefreshTTL)
	if err != nil {
		response.ErrorResponse(c, 40000)
		return
	}

	response.SuccessResponse(c, 20000, &UserResponse{
		AccessToken:  access,
		RefreshToken: refresh,
	})
}

package initialize

import (
	"github.com/LeMinh0706/fakeapi/internal/users"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	conn := InitDb()
	router := gin.Default()
	s := NewService(conn)

	// Initialize user service
	users.RegisterRouter(router.Group("/auth"), s.user, s.jwt, s.pas)

	return router
}

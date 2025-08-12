package initialize

import (
	"github.com/LeMinh0706/fakeapi/internal/benchmark"
	"github.com/LeMinh0706/fakeapi/internal/candidates"
	"github.com/LeMinh0706/fakeapi/internal/users"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	conn := InitDb()
	router := gin.Default()
	s := NewService(conn)

	// Initialize user service
	users.RegisterRouter(router.Group("/auth"), s.user, s.jwt, s.pas)
	benchmark.RegisterRouter(router.Group("/benchmark"), s.benchmark)
	candidates.RegisterRouter(router.Group("/candidates"), s.candidate)

	return router
}

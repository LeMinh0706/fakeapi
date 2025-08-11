package initialize

import (
	"database/sql"
	"log"

	"github.com/LeMinh0706/fakeapi/config"
	"github.com/LeMinh0706/fakeapi/internal/candidates"
	"github.com/LeMinh0706/fakeapi/internal/db"
	"github.com/LeMinh0706/fakeapi/internal/users"
	"github.com/LeMinh0706/fakeapi/pkg/token/token"
)

type Service struct {
	user      *users.UserService
	candidate *candidates.CandidateService
	jwt       token.Maker
	pas       token.Maker
}

func NewService(conn *sql.DB) *Service {
	q := db.New(conn)
	jwt, _ := token.NewJwtMaker(config.Cfg.Token.AccessKey)
	pas, err := token.NewPasetoMaker([]byte(config.Cfg.Token.RefreshKey))
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &Service{
		user:      users.NewUserService(q),
		candidate: candidates.NewCandidateService(q),
		jwt:       jwt,
		pas:       pas,
	}
}

package initialize

import (
	"database/sql"

	"github.com/LeMinh0706/fakeapi/internal/db"
	"github.com/LeMinh0706/fakeapi/internal/users"
)

type Service struct {
	user *users.UserService
}

func NewService(conn *sql.DB) *Service {
	q := db.New(conn)
	return &Service{
		user: users.NewUserService(q),
	}
}

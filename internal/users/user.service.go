package users

import (
	"context"
	"log"

	"github.com/LeMinh0706/fakeapi/internal/db"
	"github.com/LeMinh0706/fakeapi/utils"
	"github.com/google/uuid"
)

type UserService struct {
	q *db.Queries
}

func NewUserService(q *db.Queries) *UserService {
	return &UserService{
		q: q,
	}
}

func (s *UserService) CreateUser(ctx context.Context, username, password string) error {
	hash, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	return s.q.CreateUser(ctx, db.CreateUserParams{
		Username: username,
		Password: hash,
	})
}

func (s *UserService) GetUserByUsername(ctx context.Context, username, password string) (uuid.UUID, error) {

	user, err := s.q.GetUserByUsername(ctx, username)
	if err != nil {
		return uuid.Nil, err
	}

	err = utils.CheckPassword(password, user.Password)
	if err != nil {
		log.Println("error here", user.Password, "<>", password)
		return uuid.Nil, err
	}
	return user.ID, nil
}

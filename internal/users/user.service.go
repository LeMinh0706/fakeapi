package users

import (
	"context"
	"log"

	"github.com/LeMinh0706/fakeapi/internal/db"
	"github.com/LeMinh0706/fakeapi/utils"
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
		Email:          username,
		FullName:       "admin",
		HashedPassword: hash,
	})
}

func (s *UserService) GetUserByUsername(ctx context.Context, username, password string) (int32, error) {

	user, err := s.q.GetUserByUsername(ctx, username)
	log.Println(user)
	if err != nil {
		return 0, err
	}
	err = utils.CheckPassword(password, user.HashedPassword)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

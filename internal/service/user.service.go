package service

import (
	"context"
	"time"

	"go-mongo/internal/model"
	"go-mongo/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	InsertOne(ctx context.Context, user model.User) (model.User, error)
	FindOne(ctx context.Context, id string) (model.User, error)
	Find(ctx context.Context) ([]model.User, error)
	UpdateOne(ctx context.Context, id string, user model.User) (model.User, error)
	DeleteOne(ctx context.Context, id string) (model.User, error)
}

type userService struct {
	r repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{
		r: r,
	}
}

func (s *userService) InsertOne(ctx context.Context, user model.User) (model.User, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.ID = primitive.NewObjectIDFromTimestamp(user.CreatedAt)

	res, err := s.r.InsertOne(ctx, user)
	return res, err
}

func (s *userService) Find(ctx context.Context) ([]model.User, error) {
	panic("unimplemented")
}

func (s *userService) FindOne(ctx context.Context, id string) (model.User, error) {
	res, err := s.r.FindOne(ctx, id)
	return res, err
}

func (s *userService) UpdateOne(ctx context.Context, id string, user model.User) (model.User, error) {
	panic("unimplemented")
}

func (s *userService) DeleteOne(ctx context.Context, id string) (model.User, error) {
	panic("unimplemented")
}

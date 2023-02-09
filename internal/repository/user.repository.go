package repository

import (
	"context"

	"go-mongo/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	InsertOne(ctx context.Context, user model.User) (model.User, error)
	Find(ctx context.Context) ([]model.User, error)
	FindOne(ctx context.Context, id string) (model.User, error)
	UpdateOne(ctx context.Context, id string, user model.User) (model.User, error)
	DeleteOne(ctx context.Context, id string) (model.User, error)
}

type userRepositry struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database, collectionName string) UserRepository {
	return &userRepositry{
		db:         db,
		collection: db.Collection(collectionName),
	}
}

func (r *userRepositry) InsertOne(ctx context.Context, user model.User) (model.User, error) {
	_, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepositry) Find(ctx context.Context) ([]model.User, error) {
	panic("unimplemented")
}

func (r *userRepositry) FindOne(ctx context.Context, id string) (model.User, error) {
	var user model.User
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepositry) UpdateOne(ctx context.Context, id string, user model.User) (model.User, error) {
	panic("unimplemented")
}

func (r *userRepositry) DeleteOne(ctx context.Context, id string) (model.User, error) {
	panic("unimplemented")
}

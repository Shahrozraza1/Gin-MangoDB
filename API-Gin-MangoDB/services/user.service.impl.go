package services

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(usercollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}

func (u *UserServiceImpl) CreateUser(user *models.User) error {
	_, err := u.usercollection.InsertOne(u.ctx, user)
	return err
}
func (u *UserServiceImpl) GetUser(name *string) (*model.User, error) {
	var user *models.User
	query := bson.D{bson.E{Key: "name", Value: name}}
	err := u.usercollection.FindOne(u.ctx, query).Decode(&user)
	u.usercollection.
	return user, err
}
func (u *UserserviceImpl) GetAll() ([]*models.User, error) {
	return nil, nil
}

func (u *UserserviceImpl) UpdateUser(user *models.User) error{
	return nil
}

func (u *UserserviceImpl) DeleteUser(ctx *gin.Context) {
	return nil
}

func (u *UserserviceImpl) RegisterUserRoutes(rg *gin.Context) {
	return nil
}

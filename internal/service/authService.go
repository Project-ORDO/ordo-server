package service

import (
	"context"
	"errors"
	"time"

	models "github.com/Project-ORDO/ORDO-backEnd/internal/model"
	"github.com/Project-ORDO/ORDO-backEnd/internal/model/request"
	interfaces "github.com/Project-ORDO/ORDO-backEnd/internal/repository/interface"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo interfaces.UserRepository
}

func NewUserService(repo interfaces.UserRepository) *UserService {
	return &UserService{UserRepo: repo}
}

func (s *UserService) SignUp(ctx context.Context, user *models.User) error {
	existingUser, err := s.UserRepo.FindByEmail(ctx, user.Email)

	if err == nil && existingUser != nil {
		// User found successfully → already exists
		return errors.New("user already exists with this email")
	}

	if err != nil && err.Error() != "mongo: no documents in result" {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	now := time.Now()
	active := true

	user.Password = string(hashedPassword)
	user.UUID = uuid.New().String()
	user.CreatedAt = &now
	user.IsActive = &active

	return s.UserRepo.CreateUser(ctx, user)
}

func (s *UserService) LoginUser(ctx context.Context,req *request.LoginRequest) (*models.User,error){
	user,err:=s.UserRepo.FindByEmail(ctx,req.Email)
	if err!=nil || user==nil{
		return nil,errors.New("user not found")
	}

	err= bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(req.Password))
	if err!=nil{
		return nil,errors.New("invalid credentials")
	}

	now :=time.Now()
	err = s.UserRepo.UpdateUser(ctx,user.ID,map[string]interface{}{"lastlogin":now})
	if err!=nil{
		return nil,errors.New("could not update last login")
	}
	user.LastLogin = &now

	return user,nil
}

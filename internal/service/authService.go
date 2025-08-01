package service

import (
	"context"
	"errors"
	"time"

	models "github.com/Project-ORDO/ORDO-backEnd/internal/model"
	"github.com/Project-ORDO/ORDO-backEnd/internal/model/request"
	interfaces "github.com/Project-ORDO/ORDO-backEnd/internal/repository/interface"
	"github.com/Project-ORDO/ORDO-backEnd/internal/utils"
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
	existingUser, _ := s.UserRepo.FindByEmail(ctx, user.Email)
	if existingUser != nil {
		return errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	now := time.Now()
	token := uuid.New().String()

	verification := &models.UserVerification{
		Name:         user.Name,
		Email:        user.Email,
		Password:    string(hashedPassword),
		Token:        token,
		AttemptCount: 0,
		ExpiresAt:    now.Add(15 * time.Minute),
		CreatedAt:    now,
		VerifiedAt:   nil,
	}

	err = s.UserRepo.SaveVerification(ctx, verification)
	if err != nil {
		return err
	}

	return utils.SendVerificationEmail(user.Email, token)
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

func (s *UserService) VerifyEmail(ctx context.Context, token string) error {
	verification, err := s.UserRepo.GetVerificationByToken(ctx, token)
	if err != nil {
		return errors.New("invalid or expired token")
	}

	if time.Now().After(verification.ExpiresAt) {
		return errors.New("token expired")
	}

	newUser := &models.User{
		Name:      verification.Name,
		Email:     verification.Email,
		Password:  verification.Password,
		UUID:      uuid.New().String(),
		CreatedAt: ptrTime(time.Now()),
		IsActive:  ptrBool(true),
	}

	if err := s.UserRepo.CreateUser(ctx, newUser); err != nil {
		return err
	}

	// cleanup token
	return s.UserRepo.DeleteVerificationByToken(ctx, token)
}

func ptrTime(t time.Time) *time.Time {
	return &t
}

func ptrBool(b bool) *bool {
	return &b
}


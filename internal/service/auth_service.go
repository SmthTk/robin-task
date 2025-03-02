package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"robin-task/internal/auth"
	"robin-task/internal/model"
	"robin-task/internal/repository"
)

type AuthService struct {
	userRepo   repository.UserRepository
	jwtService *auth.JWTService
}

func NewAuthService(userRepo repository.UserRepository, jwtService *auth.JWTService) *AuthService {
	return &AuthService{userRepo: userRepo, jwtService: jwtService}
}

func (s *AuthService) Register(username, password, role, avatar string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user := &model.User{
		Username: username,
		Password: string(hashedPassword),
		Role:     role,
		Avatar:   avatar,
	}

	if err := s.userRepo.CreateUser(user); err != nil {
		return "", err
	}

	return s.jwtService.GenerateToken(user.ID, user.Role)
}

func (s *AuthService) Login(username, password string) (*model.User, string, error) {
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return nil, "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, "", errors.New("incorrect password")
	}

	token, err := s.jwtService.GenerateToken(user.ID, user.Role)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (s *AuthService) GetUserByUsername(username string) (*model.User, error) {
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

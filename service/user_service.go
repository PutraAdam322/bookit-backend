package service

import (
	"errors"

	"bookit.com/model"
	"bookit.com/utils/auth"

	"github.com/golang-jwt/jwt/v4"
)

type UserRepository interface {
	Create(user *model.User) error
	Update(user *model.User) error
	GetByID(id uint) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
}

type JWTService interface {
	GenerateToken(userID uint) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
	GetUserByTokenID(token string) (uint, error)
}

type userService struct {
	jwtService     JWTService
	userRepository UserRepository
}

func NewUserService(jwtService JWTService, userRepository UserRepository) *userService {
	return &userService{
		jwtService:     jwtService,
		userRepository: userRepository,
	}
}

func (s *userService) Register(user *model.User) (*model.User, error) {
	existing, err := s.userRepository.GetByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	if existing != nil {
		return nil, errors.New("user with this email has already exist")
	}

	passHash, err := auth.HashAndSalt(user.Password)
	if err != nil {
		return nil, err
	}

	user.PasswordHash = passHash
	user.IsAdmin = false

	err = s.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (s *userService) Login(user *model.User) (string, error) {
	existing, err := s.userRepository.GetByEmail(user.Email)
	if err != nil {
		return "", err
	}

	if existing == nil {
		return "", errors.New("user does not exist")
	}

	res, err := auth.ComparePassword(existing.PasswordHash, []byte(user.Password))

	if !res {
		return "", errors.New("password does not match")
	}

	if err != nil {
		return "", err
	}

	token, err := s.jwtService.GenerateToken(existing.ID)
	if err != nil {
		return "", err
	}

	return token, err
}

func (s *userService) AdminLogin(user *model.User) (string, error) {
	existing, err := s.userRepository.GetByEmail(user.Email)
	if err != nil {
		return "", err
	}

	if !existing.IsAdmin {
		return "", errors.New("user is not admin")
	}

	if existing == nil {
		return "", errors.New("user does not exist")
	}

	res, err := auth.ComparePassword(existing.Password, []byte(user.Password))
	if err != nil {
		return "", err
	}

	if !res {
		return "", errors.New("password does not match")
	}

	token, err := s.jwtService.GenerateToken(existing.ID)
	if err != nil {
		return "", err
	}

	return token, err
}

func (s *userService) GetByID(id uint) (*model.User, error) {
	return s.userRepository.GetByID(id)
}

func (s *userService) Update(user *model.User) (*model.User, error) {
	passHash, err := auth.HashAndSalt(user.Password)
	if err != nil {
		return nil, err
	}

	user.PasswordHash = passHash

	err = s.userRepository.Update(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

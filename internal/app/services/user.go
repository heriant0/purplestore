package services

import (
	"errors"
	"fmt"

	"github.com/heriant0/purplestore/internal/app/models"
	"github.com/heriant0/purplestore/internal/app/schemas"
	"github.com/heriant0/purplestore/internal/pkg/utils"
	log "github.com/sirupsen/logrus"
)

type UserRepository interface {
	Register(data models.User) error
	GetByEmail(email string) (*models.User, error)
}

type UserService struct {
	repository UserRepository
	secretKey  []byte
}

func NewUserService(repository UserRepository, secretKey []byte) *UserService {
	return &UserService{
		repository: repository,
		secretKey:  secretKey,
	}
}

func (s *UserService) Register(req schemas.RegisterRequest) error {
	// check email
	_, err := s.repository.GetByEmail(req.Email)
	if err == nil {
		return errors.New("user with emaill" + req.Email + "already exists")
	}

	passwordHashed := utils.HashPassword(req.Password)
	data := models.User{
		Email:    req.Email,
		Password: passwordHashed,
	}

	err = s.repository.Register(data)
	if err != nil {
		errMsg := fmt.Errorf("user service - err create : %w", err)
		log.Error(errMsg)
		return err
	}

	return nil
}

func (s *UserService) Login(req schemas.LoginRequest) (*schemas.LoginResponse, error) {

	data, err := s.repository.GetByEmail(req.Email)
	if err != nil {
		errMsg := fmt.Errorf("user not found : %w", err)
		log.Error(errMsg)
		return nil, errMsg
	}

	if !utils.VerifiedPassword(req.Password, data.Password) {
		errMsg := fmt.Errorf("password does not match : %w", err)
		log.Error(errMsg)
		return nil, errMsg
	}

	accessTooken, err := utils.BuildJWT(req.Email, s.secretKey)
	if err != nil {
		errMsg := fmt.Errorf("jwt failed : %w", err)
		log.Error(errMsg)
		return nil, err
	}
	response := schemas.LoginResponse{
		Message:     "Login success",
		AccessToken: accessTooken,
	}

	return &response, nil
}

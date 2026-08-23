package authorization

import (
	"jwt/models"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepo interface {
	GetUser(d LoginInRequestDTO) (*models.User, error)
	CreateUser(d SignUpRequestDTO) error
}

type AuthService struct {
	AuthRepo AuthRepo
}

func NewAuthService(authRepo AuthRepo) *AuthService {
	return &AuthService{AuthRepo: authRepo}
}

func (s *AuthService) SignUp(d SignUpRequestDTO) error {
	// log the request
	log.Info("AuthService.SignUp new request received")

	// validate the request
	if err := validateSignUpRequest(&d); err != nil {
		log.Error("AuthService.SignUp.validateSignUpRequest failed: ", err)
		return err
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(d.Password), 10)
	if err != nil {
		log.Error("AuthService.SignUp.bcrypt.GenerateFromPassword failed: ", err)
		return err
	}
	d.Password = string(hashPass)

	// send request to repo
	if err := s.AuthRepo.CreateUser(d); err != nil {
		log.Error("AuthService.SignUp.CreateUser failed: ", err)
		return err
	}

	// return response
	log.Info("AuthService.SignUp successfully created user")
	return nil
}

func (s *AuthService) LoginIn(d LoginInRequestDTO) (*models.User, error) {
	// log request
	log.Info("AuthService.LoginIn new request received")

	// validate request
	if err := validateLoginInRequest(&d); err != nil {
		log.Error("AuthService.LoginIn.validateLoginInRequest failed: ", err)
		return nil, err
	}

	// send request to repo
	user, err := s.AuthRepo.GetUser(d)
	if err != nil {
		log.Error("AuthService.LoginIn.GetUser failed: ", err)
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(d.Password)); err != nil {
		log.Error("AuthService.LoginIn.CompareHashAndPassword failed: ", err)
		return nil, err
	}

	// return response
	log.Info("AuthService.LoginIn successfully logged in user")
	return user, nil
}

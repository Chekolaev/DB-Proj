package service

import (
	"Library/internal/repository"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	tokenKey = "iUI(D*(HDh87ahd87h7G79h7y6g*&^G&^g7db7as6dg8df7ig7&D^gsdsd"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserUUID string
}

type Service struct {
	repos *repository.Repository
}

func Init(repos *repository.Repository) *Service {
	return &Service{repos: repos}
}

func (s *Service) AddNewUser(name, surname, password, login string) error {
	return s.repos.RegistrateUser(name, surname, login, password)
}

func (s *Service) GenerateJWT(login, password string) (string, error) {
	user, err := s.repos.CheckUser(login, password)

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user["uuid"],
	})

	return token.SignedString([]byte(tokenKey))
}

func (s *Service) ParseToken(token string) (string, error) {

	ttoken, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(tttoken *jwt.Token) (interface{}, error) {
		if _, ok := tttoken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(tokenKey), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := ttoken.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserUUID, nil
}

func (s *Service) GetUserByInterface(value interface{}) (map[string]string, error) {
	userData, err := s.repos.Db.GetUserByInterface(value)

	if err != nil {
		return nil, err
	}

	return userData, nil
}

func (s *Service) GetAllBooks() ([]map[string]string, error) {
	books, err := s.repos.Db.GetAllBooks()

	if err != nil {
		return nil, err
	}

	return books, nil
}

func (s *Service) GetBookByUUID(uuidBook string) (map[string]string, error) {
	book, err := s.repos.Db.GetBookByUUID(uuidBook)

	if err != nil {
		return nil, err
	}

	return book, nil
}

func (s *Service) RentBookByUUID(uuidBook string, UUID interface{}) (string, error) {
	status, err := s.repos.Db.GetBookByUUID(uuidBook)

	if err != nil {
		return "", err
	}

	if status["uuid"] != "" {
		resp, err := s.repos.Db.SetNewBookStatus(uuidBook, "2", UUID)

		if err != nil {
			return "", err
		}

		return resp, nil
	}

	return "Rent no permission!", nil
}

func (s *Service) GetRequests() ([]map[string]string, error) {
	reqs, err := s.repos.Db.ShowRequests()

	if err != nil {
		return nil, err
	}

	return reqs, nil
}

func (s *Service) ChangeRequest(uuidBook, statusNum string) (string, error) {
	resp, err := s.repos.Db.SetNewBookStatus(uuidBook, statusNum, "")

	if err != nil {
		return "", err
	}

	return resp, nil
}

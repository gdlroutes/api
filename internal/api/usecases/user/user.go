package user

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gdlroutes/api/internal/api/models"
)

type useCases struct {
	storage        Storage
	tokenGenerator TokenGenerator
}

var _ UseCases = &useCases{}

// New creates a default set of geodata use cases
func New(storage Storage, tokenGenerator TokenGenerator) (UseCases, error) {
	if storage == nil {
		return nil, errors.New("nil storage")
	}

	return &useCases{storage: storage, tokenGenerator: tokenGenerator}, nil
}

// CreateUserAndToken creates a user and returns an access token
func (u *useCases) CreateUserAndToken(user *models.User) (*models.Token, error) {
	if storedUser, err := u.storage.GetUserByEmail(user.Email); err != nil {
		return nil, err
	} else if storedUser != nil {
		return nil, models.ConflictError(fmt.Sprintf("email %s is already registered", user.Email))
	}

	userID, err := u.storage.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return u.tokenGenerator.GenerateToken(userID, time.Now())
}

// CreateToken validates a user's credentials and returns an access token
func (u *useCases) CreateToken(user *models.User) (*models.Token, error) {
	storedUser, err := u.storage.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	} else if storedUser == nil {
		log.Printf("Login attempt for non-existent user %s\n", user.Email)
		return nil, models.InvalidCredentialsError("invalid email and/or password")
	}

	if isPasswordCorrect, err := u.storage.IsPasswordCorrect(user.Email, user.Password); err != nil {
		return nil, err
	} else if !isPasswordCorrect {
		log.Printf("Login attempt with bad credentials for user %s\n", user.Email)
		return nil, models.InvalidCredentialsError("invalid email and/or password")
	}

	return u.tokenGenerator.GenerateToken(storedUser.Email, time.Now())
}

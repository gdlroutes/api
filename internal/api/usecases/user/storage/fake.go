package storage

import (
	"github.com/gdlroutes/api/internal/api/models"
	"github.com/gdlroutes/api/internal/api/usecases/user"
)

type fakeStorage struct {
	users map[string]*models.User
}

var _ user.Storage = &fakeStorage{}

// NewFake returns a new fake storage
func NewFake() (user.Storage, error) {
	return &fakeStorage{
		users: map[string]*models.User{},
	}, nil
}

func (s *fakeStorage) GetUserByEmail(email string) (*models.User, error) {
	return s.users[email], nil
}

func (s *fakeStorage) CreateUser(user *models.User) (string, error) {
	s.users[user.Email] = user
	return user.Email, nil
}

func (s *fakeStorage) IsPasswordCorrect(email, password string) (bool, error) {
	user, ok := s.users[email]
	if !ok {
		return false, nil
	}

	return user.Password == password, nil
}

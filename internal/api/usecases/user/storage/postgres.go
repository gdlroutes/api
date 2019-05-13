package storage

import (
	"database/sql"

	"github.com/gdlroutes/api/internal/api/models"
	"github.com/gdlroutes/api/internal/api/usecases/user"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type postgresStorage struct {
	db *sql.DB
}

var _ user.Storage = &postgresStorage{}

// NewPostgres returns a new Postgres storage
func NewPostgres(db *sql.DB) (user.Storage, error) {
	return &postgresStorage{db}, nil
}

func (s *postgresStorage) GetUserByEmail(email string) (*models.User, error) {
	const query = `
	SELECT id, email, username
	FROM users
	WHERE email = $1;
	`
	user := &models.User{}
	switch err := s.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Username); err {
	case nil:
		return user, nil
	case sql.ErrNoRows:
		return nil, nil
	default:
		return nil, err
	}
}

func (s *postgresStorage) CreateUser(user *models.User) (string, error) {
	id := uuid.New().String()
	const query = `
	INSERT INTO users(id, email, username, password)
	VALUES($1, $2, $3, $4);
	`
	if _, err := s.db.Exec(query, id, user.Email, user.Username, hashPassword(user.Password)); err != nil {
		return "", err
	}

	return id, nil
}

func (s *postgresStorage) IsPasswordCorrect(email, password string) (bool, error) {
	const query = `
	SELECT password
	FROM users
	WHERE email = $1;
	`
	var hashedPassword string
	if err := s.db.QueryRow(query, email).Scan(&hashedPassword); err != nil {
		return false, err
	}

	return doPasswordsMatch(password, hashedPassword), nil
}

func hashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hash)
}

func doPasswordsMatch(password, hashedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

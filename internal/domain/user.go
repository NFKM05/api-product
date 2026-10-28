package domain

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID         uuid.UUID
	Name       string
	Email      string
	Password   string
	IsActive   bool
	InactiveDt time.Time
}

func (u *User) Normalize() {
	u.Name = strings.TrimSpace(u.Name)
	u.Name = strings.ToUpper(u.Name)
	u.Email = strings.TrimSpace(u.Email)
	u.Email = strings.ToLower(u.Email)
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("name is required")
	}
	if u.Email == "" {
		return errors.New("email is required")
	}
	if u.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

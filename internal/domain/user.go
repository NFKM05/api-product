package domain

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Struct que reflete os campos da tabela no banco de dados
type User struct {
	ID         uuid.UUID
	Name       string
	Email      string
	Password   string
	IsActive   bool
	InactiveDt time.Time
}

// Função de regra de negócio que padroniza os campos para um formato apenas
func (u *User) Normalize() {
	u.Name = strings.TrimSpace(u.Name)
	u.Name = strings.ToUpper(u.Name)
	u.Email = strings.TrimSpace(u.Email)
	u.Email = strings.ToLower(u.Email)
}

// Função de regra de negócio que valida os campos obrigatórios
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

// Função para gravar a senha do usuário com um hash no banco de dados
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

// Função utilizada para checar a senha inserida pelo usuário comparando a salva em hash no banco
func CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

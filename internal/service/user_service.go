package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"io.github.com/api-product/internal/domain"
	"io.github.com/api-product/internal/utils"
)

// Define o contrato de acesso a dados do UserService
type UserRepository interface {
	Create(ctx context.Context, name string, email string, password string) (domain.User, error)
	UpdateName(ctx context.Context, id uuid.UUID, name string) (domain.User, error)
	UpdateEmail(ctx context.Context, id uuid.UUID, email string) (domain.User, error)
	UpdatePassword(ctx context.Context, id uuid.UUID, password string) (domain.User, error)
	ListAllUsers(ctx context.Context) ([]domain.User, error)
	FindUserById(ctx context.Context, id uuid.UUID) (domain.User, error)
	UserByEmail(ctx context.Context, email string) (domain.User, error)
	DesactivateUser(ctx context.Context, id uuid.UUID) (domain.User, error)
}

// Concentra as regras de negócio da aplicação e geração de token JWT
type UserService struct {
	repository UserRepository
	jwtService *utils.JwtService
}

// Injeta a dependência do repository e o serviço do JWT no service
func NewUserService(repo UserRepository, jwtService *utils.JwtService) *UserService {
	return &UserService{
		repository: repo,
		jwtService: jwtService,
	}
}

// Função que une as duas funções para validação dos campos e garante que a senha tenha hash
func (s *UserService) prepare(u *domain.User) error {
	u.Normalize()
	if err := u.Validate(); err != nil {
		return err
	}

	hash, err := domain.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hash

	return nil
}

// Create valida, normaliza e cadastra um novo usuário no banco de dados
func (s *UserService) Create(ctx context.Context, name string, email string, password string) (domain.User, error) {
	user := domain.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	if err := s.prepare(&user); err != nil {
		return domain.User{}, err
	}

	return s.repository.Create(ctx, user.Name, user.Email, user.Password)
}

// UpdateName altera o nome do usuário cadastrado restrito ao mesmo usuário de criação
func (s *UserService) UpdateName(ctx context.Context, id uuid.UUID, name string, requesterID uuid.UUID) (domain.User, error) {

	if requesterID != id {
		return domain.User{}, errors.New("the request user must be the same as the account user.")
	}

	return s.repository.UpdateName(ctx, id, name)
}

// UpdateEmail altera o nome do email cadastrado restrito ao mesmo usuário de criação
func (s *UserService) UpdateEmail(ctx context.Context, id uuid.UUID, email string, requesterID uuid.UUID) (domain.User, error) {

	if requesterID != id {
		return domain.User{}, errors.New("the request user must be the same as the account user.")
	}

	return s.repository.UpdateEmail(ctx, id, email)
}

// UpdatePassword altera a senha do usuário cadastrado restrito ao mesmo usuário de criação
func (s *UserService) UpdatePassword(ctx context.Context, id uuid.UUID, password string, newPass string, requesterID uuid.UUID) (domain.User, error) {
	if id != requesterID {
		return domain.User{}, errors.New("the request user must be the same as the account user")
	}

	user, err := s.repository.FindUserById(ctx, id)
	if err != nil {
		return domain.User{}, errors.New("user not found")
	}

	if ok := domain.CheckPassword(password, user.Password); !ok {
		return domain.User{}, errors.New("incorrect password")
	}

	hash, err := domain.HashPassword(newPass)
	if err != nil {
		return domain.User{}, err
	}
	return s.repository.UpdatePassword(ctx, id, hash)
}

func (s *UserService) Login(ctx context.Context, email, senha string) (string, error) {
	user, err := s.repository.UserByEmail(ctx, email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if ok := domain.CheckPassword(senha, user.Password); !ok {
		return "", errors.New("invalid credentials")
	}

	if !user.IsActive {
		return "", errors.New("the user is inactive")
	}

	return s.jwtService.GenerateJwtToken(user.Email, user.ID)
}

// ListAllUsers lista todos os usuários cadastrados na base
func (s *UserService) ListAllUsers(ctx context.Context) ([]domain.User, error) {
	return s.repository.ListAllUsers(ctx)
}

// FindUserById consulta usuários na base por um identificador
func (s *UserService) FindUserById(ctx context.Context, id uuid.UUID) (domain.User, error) {
	return s.repository.FindUserById(ctx, id)
}

// DesactivateUser desativa o cadastro do usuário para acesso
func (s *UserService) DesactivateUser(ctx context.Context, id uuid.UUID, requesterID uuid.UUID) (domain.User, error) {

	if id != requesterID {
		return domain.User{}, errors.New("the request user must be the same as the account user.")
	}

	return s.repository.DesactivateUser(ctx, id)
}

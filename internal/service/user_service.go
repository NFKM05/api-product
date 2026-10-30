package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"io.github.com/api-product/internal/domain"
	"io.github.com/api-product/internal/utils"
)

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

type UserService struct {
	repository UserRepository
	jwtService *utils.JwtService
}

func NewUserService(repo UserRepository, jwtService *utils.JwtService) *UserService {
	return &UserService{
		repository: repo,
		jwtService: jwtService,
	}
}

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

func (s *UserService) Create(ctx context.Context, name string, email string, password string) (domain.User, error) {
	user := domain.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	if err := s.prepare(&user); err != nil {
		return domain.User{}, err
	}

	return s.repository.Create(ctx, name, email, password)
}

func (s *UserService) UpdateName(ctx context.Context, id uuid.UUID, name string, requesterID uuid.UUID) (domain.User, error) {

	if requesterID != id {
		return domain.User{}, errors.New("the request user must be the same as the account user.")
	}

	return s.repository.UpdateName(ctx, id, name)
}

func (s *UserService) UpdateEmail(ctx context.Context, id uuid.UUID, email string, requesterID uuid.UUID) (domain.User, error) {

	if requesterID != id {
		return domain.User{}, errors.New("the request user must be the same as the account user.")
	}

	return s.repository.UpdateEmail(ctx, id, email)
}

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

	return s.jwtService.GenerateJwtToken(user.Email, user.ID)
}

func (s *UserService) ListAllUsers(ctx context.Context) ([]domain.User, error) {
	return s.repository.ListAllUsers(ctx)
}

func (s *UserService) FindUserById(ctx context.Context, id uuid.UUID) (domain.User, error) {
	return s.repository.FindUserById(ctx, id)
}

func (s *UserService) DesactivateUser(ctx context.Context, id uuid.UUID, requesterID uuid.UUID) (domain.User, error) {

	if id != requesterID {
		return domain.User{}, errors.New("the request user must be the same as the account user.")
	}

	return s.repository.DesactivateUser(ctx, id)
}

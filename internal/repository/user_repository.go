package repository

import (
	"context"

	"github.com/google/uuid"
	"io.github.com/api-product/internal/db"
	"io.github.com/api-product/internal/domain"
	"io.github.com/api-product/internal/mapper"
)

// Encapsula o acesso a dados de produtos
type UserRepository struct {
	queries *db.Queries
}

// Injeta as queries geradas pelo sqlc no repository
func NewUserRepository(queries *db.Queries) *UserRepository {
	return &UserRepository{queries: queries}
}

// Create insere um novo usuário no banco de dados e devolve o usuário convertido para domain.User
func (r *UserRepository) Create(ctx context.Context, name string, email string, password string) (domain.User, error) {
	row, err := r.queries.CreateUser(ctx, db.CreateUserParams{
		Name:     name,
		Email:    email,
		Password: password,
	})

	if err != nil {
		return domain.User{}, err
	}

	return mapper.ToDomainUser(row), nil
}

// UpdateName altera o nome de um usuário existente pelo seu identificador
func (r *UserRepository) UpdateName(ctx context.Context, id uuid.UUID, name string) (domain.User, error) {
	row, err := r.queries.UpdateName(ctx, db.UpdateNameParams{
		ID:   id,
		Name: name,
	})

	if err != nil {
		return domain.User{}, err
	}

	return mapper.ToDomainUser(row), nil
}

// UpdateEmail altera o email de um usuário existente pelo seu identificador
func (r *UserRepository) UpdateEmail(ctx context.Context, id uuid.UUID, email string) (domain.User, error) {
	row, err := r.queries.UpdateEmail(ctx, db.UpdateEmailParams{
		ID:    id,
		Email: email,
	})

	if err != nil {
		return domain.User{}, err
	}

	return mapper.ToDomainUser(row), nil
}

// UpdatePassword altera a senha hasheada de um usuário existente
func (r *UserRepository) UpdatePassword(ctx context.Context, id uuid.UUID, password string) (domain.User, error) {
	row, err := r.queries.UpdateSenha(ctx, db.UpdateSenhaParams{
		ID:       id,
		Password: password,
	})

	if err != nil {
		return domain.User{}, err
	}

	return mapper.ToDomainUser(row), nil
}

// ListAllUsers lista todos os usuários cadastrados no banco de dados
func (r *UserRepository) ListAllUsers(ctx context.Context) ([]domain.User, error) {
	rows, err := r.queries.ListAllUsers(ctx)

	if err != nil {
		return nil, err
	}

	users := make([]domain.User, 0, len(rows))

	for _, row := range rows {
		users = append(users, mapper.ToDomainUser(row))
	}

	return users, nil
}

// FindUserById busca apenas um usuário pelo seu identificador
func (r *UserRepository) FindUserById(ctx context.Context, id uuid.UUID) (domain.User, error) {
	row, err := r.queries.UserById(ctx, id)

	if err != nil {
		return domain.User{}, err
	}

	return mapper.ToDomainUser(row), nil
}

// UserByEmail busca apenas um usuário pelo email cadastrado
func (r *UserRepository) UserByEmail(ctx context.Context, email string) (domain.User, error) {
	row, err := r.queries.UserByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}

	return mapper.ToDomainUser(row), nil
}

// DesactivateUser desativa o acesso de um usuário sem removê-lo do banco
func (r *UserRepository) DesactivateUser(ctx context.Context, id uuid.UUID) (domain.User, error) {
	row, err := r.queries.DesactiveUser(ctx, id)
	if err != nil {
		return domain.User{}, err
	}

	return mapper.ToDomainUser(row), nil
}

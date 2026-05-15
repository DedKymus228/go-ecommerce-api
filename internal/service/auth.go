package service

import (
	"context"
	apperrors "e-commerce-api/internal/errors"
	"e-commerce-api/internal/infrastructure/auth"
	db "e-commerce-api/internal/repository/db/sqlc"
	"e-commerce-api/internal/server/dto"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type Auth interface {
	Register(ctx context.Context, req dto.RegisterRequest) (uuid.UUID, error)
	Login(ctx context.Context, req dto.LoginRequest) (string, error)
}

type AuthService struct {
	repo         db.Querier
	tokenManager auth.TokenManager
	tokenTTl     time.Duration
}

func NewAuthService(repo db.Querier, tokenManager auth.TokenManager, tokenTTl time.Duration) *AuthService {
	return &AuthService{
		repo:         repo,
		tokenManager: tokenManager,
		tokenTTl:     tokenTTl,
	}
}

func (s *AuthService) Register(ctx context.Context, req dto.RegisterRequest) (uuid.UUID, error) {
	_, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err == nil {
		return uuid.Nil, apperrors.ErrUserAlreadyExists
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "failed to hash password")
	}

	arg := db.CreateUserParams{
		Email:        req.Email,
		PasswordHash: string(hash),
		FirstName:    pgtype.Text{String: req.FirstName, Valid: req.FirstName != ""},
		LastName:     pgtype.Text{String: req.LastName, Valid: req.LastName != ""},
	}

	user, err := s.repo.CreateUser(ctx, arg)
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "failed to create user in db")
	}

	return user.ID, nil
}

func (s *AuthService) Login(ctx context.Context, req dto.LoginRequest) (string, error) {
	user, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return "", apperrors.ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return "", apperrors.ErrInvalidCredentials
	}

	token, err := s.tokenManager.GenerateToken(user.ID.String())
	if err != nil {
		return "", errors.Wrap(err, "failed to generate token")
	}

	return token, nil
}

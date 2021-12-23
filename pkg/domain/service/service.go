package service

import (
	"context"

	"github.com/opencars/auth/pkg/domain"
	"github.com/opencars/auth/pkg/domain/command"
	"github.com/opencars/auth/pkg/domain/model"
	"github.com/opencars/auth/pkg/domain/query"
)

type UserService struct {
	repo domain.TokenRepository
}

func NewUserService(repo domain.TokenRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateToken(ctx context.Context, c *command.CreateToken) (*model.Token, error) {
	if err := command.Process(c); err != nil {
		return nil, err
	}

	token := model.NewToken(c.UserID, c.Name)

	if err := s.repo.Create(ctx, token); err != nil {
		return nil, err
	}

	return token, nil
}

// TODO: Wrap with transactions.
func (s *UserService) ResetToken(ctx context.Context, c *command.ResetToken) (*model.Token, error) {
	if err := command.Process(c); err != nil {
		return nil, err
	}

	token, err := s.repo.FindByID(ctx, c.TokenID)
	if err != nil {
		return nil, err
	}

	token.ResetSecret()

	if err = s.repo.Update(ctx, token); err != nil {
		return nil, err
	}

	return token, err

}

func (s *UserService) DeleteToken(ctx context.Context, c *command.DeleteToken) error {
	if err := command.Process(c); err != nil {
		return err
	}

	if err := s.repo.DeleteByID(ctx, c.TokenID); err != nil {
		return err
	}

	return nil
}

func (s *UserService) ListTokens(ctx context.Context, q *query.ListTokens) ([]model.Token, error) {
	if err := query.Process(q); err != nil {
		return nil, err
	}

	return s.repo.List(ctx, q)
}

func (s *UserService) TokenDetails(ctx context.Context, q *query.TokenDetails) (*model.Token, error) {
	if err := query.Process(q); err != nil {
		return nil, err
	}

	return s.repo.FindByID(ctx, q.TokenID)
}

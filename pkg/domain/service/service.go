package service

import (
	"context"
	"errors"
	"time"

	"github.com/opencars/auth/pkg/domain"
	"github.com/opencars/auth/pkg/domain/command"
	"github.com/opencars/auth/pkg/domain/model"
	"github.com/opencars/auth/pkg/domain/query"
	"github.com/opencars/auth/pkg/eventapi"
	"github.com/opencars/httputil"
)

type UserService struct {
	tokens    domain.TokenRepository
	blacklist domain.BlackListRepository
	publisher eventapi.Publisher
}

func NewUserService(tokens domain.TokenRepository, blacklist domain.BlackListRepository) *UserService {
	return &UserService{
		tokens:    tokens,
		blacklist: blacklist,
	}
}

func (s *UserService) CreateToken(ctx context.Context, c *command.CreateToken) (*model.Token, error) {
	if err := command.Process(c); err != nil {
		return nil, err
	}

	token := model.NewToken(c.UserID, c.Name)

	if err := s.tokens.Create(ctx, token); err != nil {
		return nil, err
	}

	return token, nil
}

// TODO: Wrap with transactions.
func (s *UserService) ResetToken(ctx context.Context, c *command.ResetToken) (*model.Token, error) {
	if err := command.Process(c); err != nil {
		return nil, err
	}

	token, err := s.tokens.FindByID(ctx, c.TokenID)
	if err != nil {
		return nil, err
	}

	token.ResetSecret()

	if err = s.tokens.Update(ctx, token); err != nil {
		return nil, err
	}

	return token, err

}

func (s *UserService) DeleteToken(ctx context.Context, c *command.DeleteToken) error {
	if err := command.Process(c); err != nil {
		return err
	}

	if err := s.tokens.DeleteByID(ctx, c.TokenID); err != nil {
		return err
	}

	return nil
}

func (s *UserService) ListTokens(ctx context.Context, q *query.ListTokens) ([]model.Token, error) {
	if err := query.Process(q); err != nil {
		return nil, err
	}

	tokens, err := s.tokens.List(ctx, q)
	if err != nil {
		return nil, err
	}

	for i := range tokens {
		tokens[i].ClearSecret()
	}

	return tokens, nil
}

func (s *UserService) TokenDetails(ctx context.Context, q *query.TokenDetails) (*model.Token, error) {
	if err := query.Process(q); err != nil {
		return nil, err
	}

	token, err := s.tokens.FindByID(ctx, q.TokenID)
	if err != nil {
		return nil, err
	}

	token.ClearSecret()

	return token, nil
}

func (s *UserService) VerifyToken(ctx context.Context, c *command.VerifyToken) (*model.Token, error) {
	auth := model.Authorization{
		Status: model.AuthStatusSucceed,
		IP:     c.IP,
		Time:   time.Now().UTC(),
		Token: model.Token{
			ID: c.Secret,
		},
	}

	token, err := s.tokens.FindBySecret(ctx, c.Secret)
	if errors.Is(err, model.ErrTokenNotFound) {
		return nil, model.ErrInvalidToken
	}

	auth.Token = *token
	if !token.Enabled {
		return nil, model.ErrTokenRevoked
	}

	item, err := s.blacklist.FindByIPv4(auth.IP)
	if err == nil && item.Enabled {
		return nil, model.ErrAccessDenied
	}

	if err != nil && !errors.Is(err, model.ErrBlacklistRecordNotFound) {
		return nil, err
	}

	s.publisher.Publish(c.Event())

	return token, nil
}

func (h *UserService) result(auth *model.Authorization, httpErr *httputil.Error) error {
	if httpErr != nil {
		auth.Status = model.AuthStatusFailed
		auth.Error = new(string)
		*auth.Error = (*httpErr).Error()
	}

	event, err := eventapi.NewEvent(eventapi.EventAuthorizationKind, &auth)
	if err != nil {
		return err
	}

	if err := h.publisher.Publish(event); err != nil {
		return err
	}

	if httpErr != nil {
		return *httpErr
	}

	return nil
}

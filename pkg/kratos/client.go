package kratos

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/opencars/auth/pkg/domain/model"
	"github.com/opencars/auth/pkg/logger"
	kratos "github.com/ory/kratos-client-go"
)

type Client struct {
	host string
	c    *kratos.APIClient
}

func NewClient(host string) (*Client, error) {
	url, err := url.Parse(host)
	if err != nil {
		return nil, err
	}

	return &Client{
		host: url.String(),
		c: kratos.NewAPIClient(&kratos.Configuration{
			Host:      host,
			UserAgent: "opencars/1.0.0/go",
			HTTPClient: &http.Client{
				Timeout: 5 * time.Second,
			},
		}),
	}, nil
}

func (c *Client) CheckSession(ctx context.Context, cookie string) (*model.User, error) {
	logger.Debugf("seneding to session: %#v", cookie)

	req := c.c.V0alpha2Api.ToSession(ctx).Cookie(cookie)

	session, resp, err := c.c.V0alpha2Api.ToSessionExecute(req)
	if err != nil {
		return nil, err
	}

	logger.Debugf("response: %#v", resp)
	logger.Debugf("session: %#v", session)

	return &model.User{
		ID: session.Identity.Id,
	}, nil
}

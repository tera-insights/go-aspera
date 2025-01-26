package aspera

import (
	"context"
	"net/http"
)

type AuthenticateService struct {
	Client *Client
}

func NewAuthenticateService(client *Client) *AuthenticateService {
	return &AuthenticateService{
		Client: client,
	}
}

func (a *AuthenticateService) Authenticate(ctx context.Context, authSpec *AuthSpec) error {

	req, err := a.Client.NewRequest(http.MethodPost, "authenticate", authSpec)
	if err != nil {
		return err
	}

	return a.Client.Do(ctx, req, nil)
}

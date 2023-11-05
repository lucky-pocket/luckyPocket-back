package gauth

import (
	"net/http"
	"strings"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/onee-only/gauth-go"
)

type client struct {
	client *gauth.Client
}

func NewClient(clientID, clientSecret, redirectURI string) domain.GAuthClient {
	opts := gauth.ClientOpts{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURI:  redirectURI,
	}
	return &client{
		client: gauth.NewClient(&http.Client{}, opts),
	}
}

func (c *client) IssueToken(code string) (access, refresh string, err error) {
	return c.client.IssueToken(code)
}

func (c *client) GetUserInfo(accessToken string) (*domain.GAuthUser, error) {
	info, err := c.client.GetUserInfo(accessToken)
	if err != nil {
		return nil, err
	}

	role, _ := strings.CutPrefix(string(info.Role), "ROLE_")

	user := &domain.GAuthUser{
		Email:      info.Email,
		Name:       info.Name,
		Grade:      info.Grade,
		ClassNum:   info.ClassNum,
		Num:        info.Num,
		Gender:     constant.Gender(info.Gender),
		ProfileURL: info.ProfileURL,
		Role:       constant.UserType(role),
	}

	return user, nil
}

package internal

import (
	"time"

	apiclient "github.com/bjerkio/tripletex-go/client"
	"github.com/bjerkio/tripletex-go/client/session"
	"github.com/cobraz/jira-to-tripletex/config"
)

var token string

// GetToken retrieves a new tripletex token
func GetToken() (string, error) {

	if token != "" {
		return token, nil
	}

	ttl := 24 * time.Hour
	client := apiclient.Default

	cnf, err := config.GetConfig()

	sessionReq := &session.TokenSessionCreateCreateParams{
		ConsumerToken:  cnf.ConsumerToken,
		EmployeeToken:  cnf.EmployeeToken,
		ExpirationDate: time.Now().Add(ttl).Format("2006-01-02"),
	}

	res, err := client.Session.TokenSessionCreateCreate(sessionReq.WithTimeout(10 * time.Second))

	if err != nil {
		return "", err
	}

	token = res.Payload.Value.Token

	return token, nil
}

package jira

import (
	"github.com/cobraz/jira-to-tripletex/internal/pkg/config"
	jira "gopkg.in/andygrunwald/go-jira.v1"
)

// New represents new Jira Config
func New() (*jira.Client, error) {
	cnf, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	tp := jira.BasicAuthTransport{
		Username: cnf.JiraEmail,
		Password: cnf.JiraToken,
	}

	return jira.NewClient(tp.Client(), cnf.Host)
}

package internal

import (
	"github.com/cobraz/jira-to-tripletex/config"
	jira "gopkg.in/andygrunwald/go-jira.v1"
)

func InitJira() (*jira.Client, error) {
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

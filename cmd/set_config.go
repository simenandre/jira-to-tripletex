package cmd

import (
	"github.com/cobraz/jira-to-tripletex/config"
	"github.com/urfave/cli/v2"
)

func SetConfig(c *cli.Context) error {
	err := config.SetConfig(config.Config{
		Host:          c.String("host"),
		JiraToken:     c.String("jiraToken"),
		JiraEmail:     c.String("jiraEmail"),
		EmployeeToken: c.String("employeeToken"),
		ConsumerToken: c.String("consumerToken"),
		ActivityCode:  c.String("activityCode"),
	})
	if err != nil {
		return err
	}
	return nil
}

package cmd

import (
	"github.com/cobraz/jira-to-tripletex/config"
	"github.com/urfave/cli/v2"
)

func GetConfig(c *cli.Context) error {
	config.Print()
	return nil
}

package jtt

import (
	"github.com/cobraz/jira-to-tripletex/internal/pkg/config"
	"github.com/urfave/cli/v2"
)

// GetConfig prints current configuration
func GetConfig(c *cli.Context) error {
	config.Print()
	return nil
}

package main

import (
	"log"
	"os"

	"github.com/cobraz/jira-to-tripletex/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "Jira-To-Tripletex",
		Description: "Transport Jira issues to Tripletex",
		Commands: []*cli.Command{
			&cli.Command{
				Name:    "get:config",
				Aliases: []string{"gc"},
				Action:  cmd.GetConfig,
			},
			&cli.Command{
				Name:    "set:config",
				Aliases: []string{"cnf"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "host",
						Required: true,
						Usage:    "Jira server host",
						// Destination: &language,
					},
					&cli.StringFlag{
						Name:     "jiraEmail",
						Required: true,
						Usage:    "Jira Email",
						// Destination: &language,
					},
					&cli.StringFlag{
						Name:     "jiraToken",
						Required: true,
						Usage:    "Jira API token",
						// Destination: &language,
					},
					&cli.StringFlag{
						Name:     "consumerToken",
						Required: true,
						Usage:    "Tripletex ConsumerToken",
						// Destination: &language,
					},
					&cli.StringFlag{
						Name:     "employeeToken",
						Required: true,
						Usage:    "Tripletex EmployeeToken",
						// Destination: &language,
					},
					&cli.StringFlag{
						Name:     "activityCode",
						Required: true,
						Usage:    "Tripletex Activity Id",
						// Destination: &language,
					},
				},
				Action: cmd.SetConfig,
			},
			&cli.Command{
				Name:        "list:activities",
				Description: "Lists all activities in Tripletex",
				Action:      cmd.GetActivities,
			},
			&cli.Command{
				Name:        "copy",
				Description: "Copy project to Tripletex",
				Aliases:     []string{"c"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "projectId",
						Required: true,
						Usage:    "Tripletex Project ID",
					},
					&cli.StringFlag{
						Name:     "key",
						Required: true,
						Usage:    "Jira Project Key",
					},
				},
				Action: cmd.CopyToTripletex,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

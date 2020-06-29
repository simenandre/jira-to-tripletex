package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/bjerkio/tripletex-go/client/project_activity"
	"github.com/bjerkio/tripletex-go/models"
	"github.com/cobraz/jira-to-tripletex/internal"
	"github.com/urfave/cli/v2"
)

func CopyToTripletex(c *cli.Context) error {

	tripletexClient, authInfo, err := internal.TripletexClient()
	if err != nil {
		log.Fatal(err)
		return err
	}

	jiraClient, err := internal.InitJira()
	if err != nil {
		log.Fatal(err)
		return err
	}

	jql := fmt.Sprintf("project = %s AND issuetype = Story", c.String("key"))
	issue, _, err := jiraClient.Issue.Search(jql, nil)
	if err != nil {
		log.Fatal(err)
		return err
	}

	projectID := int32(c.Int("projectId"))
	project := models.Project{
		ID: projectID,
	}

	for _, i := range issue {
		activityName := fmt.Sprintf("(%s) %s", i.Key, i.Fields.Summary)
		newActivity := models.Activity{
			Name:         activityName,
			ActivityType: "PROJECT_SPECIFIC_ACTIVITY",
		}

		newProjectActivity := models.ProjectActivity{
			Project:  &project,
			Activity: &newActivity,
		}

		req := project_activity.ProjectProjectActivityPostParams{
			Body: &newProjectActivity,
		}

		_, err := tripletexClient.ProjectActivity.ProjectProjectActivityPost(req.WithTimeout(10*time.Second), authInfo)
		if err != nil {
			log.Fatal(err)
			return err
		}

	}
	return nil
}

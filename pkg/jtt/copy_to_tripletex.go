package jtt

import (
	"fmt"
	"log"
	"time"

	"github.com/bjerkio/tripletex-go/client/project_activity"
	"github.com/bjerkio/tripletex-go/models"
	j "github.com/cobraz/jira-to-tripletex/internal/pkg/jira"
	tx "github.com/cobraz/jira-to-tripletex/internal/pkg/tripletex"
)

// CopyToTripletex copies tasks from Jira to Tripletex
func CopyToTripletex(key string, projectID int32) error {

	tripletexClient, authInfo, err := tx.New()
	if err != nil {
		return err
	}

	jiraClient, err := j.New()
	if err != nil {
		return err
	}

	jql := fmt.Sprintf("project = %s AND issuetype = Story", key)
	issue, _, err := jiraClient.Issue.Search(jql, nil)
	if err != nil {
		log.Fatal(err)
		return err
	}

	yes := true

	project := models.Project{
		ID: projectID,
	}

	for _, i := range issue {
		activityName := fmt.Sprintf("(%s) %s", i.Key, i.Fields.Summary)
		newActivity := models.Activity{
			Name:         activityName,
			ActivityType: "PROJECT_SPECIFIC_ACTIVITY",
			IsChargeable: &yes,
		}

		var budget float64
		if i.Fields.AggregateTimeOriginalEstimate > 0 {
			budget = float64(i.Fields.AggregateTimeOriginalEstimate) / 3600
		} else {
			budget = 0
		}
		newProjectActivity := models.ProjectActivity{
			Project:     &project,
			Activity:    &newActivity,
			BudgetHours: budget,
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

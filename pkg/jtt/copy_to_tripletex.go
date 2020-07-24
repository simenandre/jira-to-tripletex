// Copyright 2020 Simen A. W. Olsen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

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

	"github.com/bjerkio/tripletex-go/client/activity"
	tx "github.com/cobraz/jira-to-tripletex/internal/pkg/tripletex"
)

// GetActivities prints activities from Tripletex
func GetActivities() error {

	tripletexClient, authInfo, err := tx.New()
	if err != nil {
		log.Fatal(err)
		return err
	}

	yes := true
	req := activity.ActivitySearchParams{
		IsProjectActivity: &yes,
		IsGeneral:         &yes,
	}
	res, err := tripletexClient.Activity.ActivitySearch(req.WithTimeout(10*time.Second), authInfo)
	if err != nil {
		log.Fatal(err)
		return err
	}

	for _, c := range res.Payload.Values {
		line := fmt.Sprintf("%d: %s", c.ID, c.Name)
		log.Println(line)
	}

	return nil
}

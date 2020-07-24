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

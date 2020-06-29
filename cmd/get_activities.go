package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/bjerkio/tripletex-go/client/activity"
	"github.com/cobraz/jira-to-tripletex/internal"
	"github.com/urfave/cli/v2"
)

func GetActivities(c *cli.Context) error {

	tripletexClient, authInfo, err := internal.TripletexClient()
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

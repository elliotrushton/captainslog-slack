package slacker

import (
	"errors"
	"os"
	"time"

	"github.com/slack-go/slack"
)

type Slacker struct {
	apiToken string
	api      *slack.Client
}

func (s Slacker) GetTimezoneForUser(userID string) (*time.Location, error) {
	// Call the users.info method to get user details
	userInfo, err := s.api.GetUserInfo(userID)
	if err != nil {
		return nil, err
	}

	// Check if the user has a timezone set
	if userInfo.TZ == "" {
		return nil, errors.New("No timezone set for Slack user")
	}

	loc, err := time.LoadLocation(userInfo.TZ)
	return loc, nil

}

func NewSlacker() *Slacker {
	s := &Slacker{apiToken: os.Getenv("SLACK_TOKEN")}
	s.api = slack.New(s.apiToken)

	return s
}

package actions

import (
	"fmt"

	"github.com/elliotrushton/captainslog-slack/logbook"
	"github.com/elliotrushton/captainslog-slack/slacker"
)

func Show(captainsLog *logbook.Log, slackClient *slacker.Slacker, userID string) string {
	tz, err := slackClient.GetTimezoneForUser(userID)
	if err != nil {
		fmt.Println("Error getting Slack users tz: ", err)
	}

	return captainsLog.String(tz)
}

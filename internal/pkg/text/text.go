package text

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

// Entities ...
type Entities struct {
	Username string
	Tweet    string
	DateTime time.Time
}

var usernameRegex = regexp.MustCompile(`@(\w{1,15})\b`)
var dateTimeRegex = regexp.MustCompile(`((1[0-2]|0?[1-9]):([0-5][0-9]) ?([AaPp][Mm]))\s-\s\d{1,2}\s\w+\s\d{4}`)
var alphanumRegex = regexp.MustCompile(`[^A-Za-z0-9]+`)

const stdDateTimeLayout = "03:04 PM - 02 Jan 2006"

// Parse ...
func Parse(rawTweet string) (*Entities, error) {
	username := string(usernameRegex.FindString(rawTweet))
	dateTimeStr := string(dateTimeRegex.FindString(rawTweet))
	uIndex := usernameRegex.FindStringIndex(rawTweet)
	dIndex := dateTimeRegex.FindStringIndex(rawTweet)
	if uIndex == nil || dIndex == nil {
		return nil, errors.New("datetime not found")
	}
	tweetStr := rawTweet[uIndex[1]:dIndex[0]]
	dateTime, err := time.Parse(stdDateTimeLayout, strings.TrimSpace(dateTimeStr))
	if err != nil {
		return nil, err
	}
	return &Entities{
		Username: username,
		DateTime: dateTime,
		Tweet:    tweetStr,
	}, nil
}

package search

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/notAI-tech/verifytweet-go/configs"
	"github.com/notAI-tech/verifytweet-go/internal/pkg/text"
)

type Tweet struct {
	ConversationID string   `json:"conversation_id,omitempty"`
	Datestamp      string   `json:"datestamp,omitempty"`
	Datetime       int64    `json:"datetime,omitempty"`
	Hashtags       []string `json:"hashtags,omitempty"`
	ID             int64    `json:"id"`
	LikesCount     string   `json:"likes_count,omitempty"`
	Link           string   `json:"link"`
	Mentions       []string `json:"mentions,omitempty"`
	Name           string   `json:"name"`
	Photos         []string `json:"photos,omitempty"`
	Place          string   `json:"place,omitempty"`
	RepliesCount   string   `json:"replies_count,omitempty"`
	Retweet        bool     `json:"retweet,omitempty"`
	RetweetsCount  string   `json:"retweets_count,omitempty"`
	Timestamp      string   `json:"timestamp,omitempty"`
	Timezone       string   `json:"timezone,omitempty"`
	Tweet          string   `json:"tweet"`
	Urls           []string `json:"urls,omitempty"`
	UserID         int      `json:"user_id,omitempty"`
	Username       string   `json:"username"`
	Video          int      `json:"video,omitempty"`
}

func SelfHosted(data *text.Entities) ([]Tweet, error) {
	host := configs.Init().Search.Host
	endpoint := configs.Init().Search.URL
	byteData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(fmt.Sprintf("%s%s", host, endpoint), "application/json", bytes.NewBuffer(byteData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	tweets := []Tweet{}
	err = json.Unmarshal(body, &tweets)
	if err != nil {
		return nil, err
	}
	return tweets, nil
}

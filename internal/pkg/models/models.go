package models

import "time"

type Tweet struct {
	ConversationID string    `json:"conversation_id,omitempty"`
	Datestamp      string    `json:"datestamp,omitempty"`
	Datetime       int64     `json:"datetime,omitempty"`
	Hashtags       []string  `json:"hashtags,omitempty"`
	ID             int64     `json:"id"`
	LikesCount     string    `json:"likes_count,omitempty"`
	Link           string    `json:"link"`
	Mentions       []string  `json:"mentions,omitempty"`
	Name           string    `json:"name"`
	ParsedText     []string  `json:"parsedText"`
	Photos         []string  `json:"photos,omitempty"`
	Place          string    `json:"place,omitempty"`
	RepliesCount   string    `json:"replies_count,omitempty"`
	Retweet        bool      `json:"retweet,omitempty"`
	RetweetsCount  string    `json:"retweets_count,omitempty"`
	Similarity     []float64 `json:"similarity"`
	Timestamp      string    `json:"timestamp,omitempty"`
	Timezone       string    `json:"timezone,omitempty"`
	Tweet          string    `json:"tweet"`
	Urls           []string  `json:"urls,omitempty"`
	UserID         int       `json:"user_id,omitempty"`
	Username       string    `json:"username"`
	Vector         []float64 `json:"vector"`
	Video          int       `json:"video,omitempty"`
}

// Entities ...
type Entities struct {
	Username string    `json:"username"`
	Tweet    string    `json:"tweet"`
	DateTime time.Time `json:"datetime"`
}

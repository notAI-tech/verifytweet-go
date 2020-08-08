package search

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/notAI-tech/verifytweet-go/configs"
	"github.com/notAI-tech/verifytweet-go/internal/pkg/models"
)

func SelfHosted(data *models.Entities) ([]models.Tweet, error) {
	confHandler := configs.Init()
	host := confHandler.Search.Host
	endpoint := confHandler.Search.URL
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
	tweets := []models.Tweet{}
	err = json.Unmarshal(body, &tweets)
	if err != nil {
		return nil, err
	}
	return tweets, nil
}

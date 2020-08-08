package text

import (
	"errors"
	"math"
	"regexp"
	"strings"
	"time"

	"github.com/notAI-tech/verifytweet-go/internal/pkg/models"
)

var usernameRegex = regexp.MustCompile(`@(\w{1,15})\b`)
var dateTimeRegex = regexp.MustCompile(`((1[0-2]|0?[1-9]):([0-5][0-9]) ?([AaPp][Mm]))\s-\s\d{1,2}\s\w+\s\d{4}`)
var alphanumRegex = regexp.MustCompile(`[^A-Za-z0-9]+`)
var smallWordsRegex = regexp.MustCompile(`\W*\b\w{1,2}\b`)

const stdDateTimeLayout = "3:04 PM - 02 Jan 2006"

// Parse ...
func Parse(rawTweet string) (*models.Entities, error) {
	username := string(usernameRegex.FindString(rawTweet))
	dateTimeStr := string(dateTimeRegex.FindString(rawTweet))
	uIndex := usernameRegex.FindStringIndex(rawTweet)
	dIndex := dateTimeRegex.FindStringIndex(rawTweet)
	if uIndex == nil || dIndex == nil {
		return nil, errors.New("datetime not found")
	}
	rawTweetStr := rawTweet[uIndex[1]:dIndex[0]]
	tweetStr := Sanitize(rawTweetStr)
	dateTime, err := time.Parse(stdDateTimeLayout, strings.TrimSpace(dateTimeStr))
	if err != nil {
		return nil, err
	}
	return &models.Entities{
		Username: username,
		DateTime: dateTime,
		Tweet:    tweetStr,
	}, nil
}

// CalculateSimilarityMatrix ...
func CalculateSimilarityMatrix(tweetObjs []models.Tweet, entities *models.Entities) models.Tweet {
	tweetObjs = ExtractTweets(tweetObjs)
	tweetObjs = append(tweetObjs, createDoc(entities.Tweet))
	vocabulary := WordMap(tweetObjs)
	tweetObjs = CreateSparseMatrix(vocabulary, tweetObjs)
	for i := 0; i < len(tweetObjs); i++ {
		simMatrix := make([]float64, 0)
		for j := 0; j < len(tweetObjs); j++ {
			s := ConsineSimilarity(tweetObjs[i].Vector, tweetObjs[j].Vector)
			simMatrix = append(simMatrix, s)
		}
		tweetObjs[i].Similarity = simMatrix
	}
	parsedTweetIdx := len(tweetObjs) - 1
	idx, _ := findMaxIndex(tweetObjs[parsedTweetIdx].Similarity)
	return tweetObjs[idx]
}

// Sanitize ...
func Sanitize(rawTweet string) string {
	t := alphanumRegex.ReplaceAllString(rawTweet, " ")
	r := smallWordsRegex.ReplaceAllString(t, "")
	c := strings.TrimSpace(r)
	l := strings.ToLower(c)
	return l
}

// ExtractTweets ...
func ExtractTweets(tweetObjs []models.Tweet) []models.Tweet {
	for idx := range tweetObjs {
		tweetObjs[idx].ParsedText = strings.Split(Sanitize(tweetObjs[idx].Tweet), " ")
	}
	return tweetObjs
}

// WordMap ...
func WordMap(documents []models.Tweet) map[string]int {
	wordmap := make(map[string]int, 0)
	for idx := range documents {
		for _, word := range documents[idx].ParsedText {
			if _, ok := wordmap[word]; !ok {
				wordmap[word] = 1
			} else {
				wordmap[word]++
			}
		}
	}
	return wordmap
}

// CreateSparseMatrix ...
func CreateSparseMatrix(vocabulary map[string]int, documents []models.Tweet) []models.Tweet {
	for idx := range documents {
		wordOccurence := make([]float64, len(vocabulary))
		for _, word := range documents[idx].ParsedText {
			if _, ok := vocabulary[word]; ok {
				wordOccurence = wordOccurence[1:]
				wordOccurence = append(wordOccurence, 1)
			}
		}
		documents[idx].Vector = wordOccurence
	}
	return documents
}

// ConsineSimilarity ...
func ConsineSimilarity(x, y []float64) float64 {
	var dotproduct float64
	var powX float64
	var powY float64
	for idx := range x {
		dotproduct += x[idx] * y[idx]
		powX = math.Pow(x[idx], 2)
		powY = math.Pow(y[idx], 2)
	}
	cosineSim := dotproduct / (math.Sqrt(powX) * math.Sqrt(powY))
	return cosineSim
}

func findMaxIndex(arr []float64) (int, float64) {
	idx := 0
	max := math.Inf(-1)
	for i := range arr {
		if arr[i] > max {
			max = arr[i]
			idx = i
		}
	}
	return idx, max
}

func createDoc(tweet string) models.Tweet {
	doc := models.Tweet{}
	doc.Tweet = tweet
	doc.ParsedText = strings.Split(Sanitize(tweet), " ")
	return doc
}

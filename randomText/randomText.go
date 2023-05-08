package randomtext

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type Response struct {
	Words []string `json:"words"`
}

func GenerateText(wordCount int) string {
	words, err := getRandomWords()
	if err != nil {
		return "Unable to generate random words"
	}
	words = pickRandomWords(words, wordCount)

	return strings.Join(words, " ")
}

func getRandomWords() ([]string, error) {
	response, err := http.Get("https://monkeytype.com/languages/english_1k.json")
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if len(bodyBytes) == 0 {
		return nil, fmt.Errorf("body was empty")
	}

	var returnData Response
	err = json.Unmarshal(bodyBytes, &returnData)
	if err != nil {
		return nil, err
	}

	return returnData.Words, nil
}

func pickRandomWords(words []string, wordCount int) []string {
	var returnData []string
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for _, i := range r.Perm(wordCount) {
		returnData = append(returnData, words[i])
	}

	return returnData
}

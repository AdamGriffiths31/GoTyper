package randomtext

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type RandomGenerator struct {
	logger *log.Logger
}

type Response struct {
	Words []string `json:"words"`
}

func NewRandomGenerator(logger *log.Logger) *RandomGenerator {
	return &RandomGenerator{logger: logger}
}

func (r RandomGenerator) GenerateText(wordCount int) string {
	words, err := r.getRandomWords()
	if err != nil {
		r.logger.Println("Error:", err)
		return "Unable to generate random words"
	}
	words = pickRandomWords(words, wordCount)

	return strings.Join(words, " ")
}

func (r RandomGenerator) getRandomWords() ([]string, error) {
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

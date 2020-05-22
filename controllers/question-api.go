package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dryairship/messenger-quiz-bot/models"
)

var QUESTION_API_URL string = "https://opentdb.com/api.php?amount=1"

func GetQuestion() (models.Question, error) {
	var apiResponseBody models.QuestionAPIResponseBody

	response, err := http.Get(QUESTION_API_URL)
	if err != nil {
		log.Println("[ERROR] Cannnot send GET request to Question API. Error: ", err)
		return models.Question{}, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("[ERROR] Cannnot read Question API Response. Error: ", err)
		return models.Question{}, err
	}

	if response.StatusCode != http.StatusOK {
		log.Println("[ERROR] Invalid response from Question API. Response: ", string(body))
		return models.Question{}, errors.New(string(body))
	}

	err = json.Unmarshal(body, &apiResponseBody)
	if err != nil {
		log.Println("[ERROR] Cannnot parse Question API Response. Error: ", err)
		return models.Question{}, err
	}

	return apiResponseBody.Results[0], nil
}

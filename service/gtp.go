package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/869413421/wechatbot/model"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const BASEURL = "https://api.openai.com/v1/chat/completions"

// ChatGPTResponseBody 请求体
type ChatGPTResponseBody struct {
	ID      string                   `json:"id"`
	Object  string                   `json:"object"`
	Created int                      `json:"created"`
	Model   string                   `json:"model"`
	Choices []map[string]interface{} `json:"choices"`
	Usage   map[string]interface{}   `json:"usage"`
}

type ChoiceItem struct {
}

// ChatGPTRequestBody 响应体
type ChatGPTRequestBody struct {
	Model            string  `json:"model"`
	Prompt           string  `json:"prompt"`
	MaxTokens        int     `json:"max_tokens"`
	Temperature      float32 `json:"temperature"`
	TopP             int     `json:"top_p"`
	FrequencyPenalty int     `json:"frequency_penalty"`
	PresencePenalty  int     `json:"presence_penalty"`
}


func Completions(msg string) (string, error) {
	start := time.Now().Unix()
	apiURL := "https://api.openai.com/v1/chat/completions"
	messages := make([]model.Message,0)
	messages = append(messages,model.Message{
		Role:    "user",
		Content: msg,
	})
	data := &model.OpenAIRequest{
		Model:    "gpt-3.5-turbo",
		Messages: messages,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "",err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "",err
	}
	secret := os.Getenv("OPENAI_API_KEY")
	if secret == "" {
		log.Println(" empty secret")
		return "",nil
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", secret))

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	end := time.Now().Unix()
	log.Printf("time cost : %v",end-start)
	log.Println()
	log.Println()
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	var result model.OpenAIResponse
	err = json.Unmarshal(body, &result)
	log.Println(result)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	if len(result.Choices) == 0 {
		return "", err
	}
	return strings.TrimSpace(result.Choices[0].Message.Content),nil
}

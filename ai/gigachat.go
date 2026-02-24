package ai

import (
	"ai-api/env"

	gigachat "github.com/tigusigalpa/gigachat-go"
	"ai-api/data"
)

func RequestToGigaChat(request *data.Request) (string, error) {
	key, err := env.GetAPIKey(request.Provider) 
	if err != nil {
		return "", err
	}

	tokenManager := gigachat.NewTokenManager(key)
	client := gigachat.NewClient(tokenManager)

	messages := []gigachat.Message{
		{Role: "user", Content: request.Message},
	}

	response, err := client.Chat(messages, gigachat.WithModel(request.Model))
	if err != nil {
		return "", err
	}

	res := response.Choices[0].Message.Content

	return res, err
}

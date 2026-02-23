package ai

import (
	"log"

	"ai-api/env"

	gigachat "github.com/tigusigalpa/gigachat-go"
	"ai-api/data"
)

func RequestToGigaChat(request *data.Request) (string, error) {
	key, err := env.LoadEnv(request.Provider)
	if err != nil {
		log.Fatal(err)
	}

	tokenManager := gigachat.NewTokenManager(key)
	client := gigachat.NewClient(tokenManager)

	messages := []gigachat.Message{
		{Role: "user", Content: request.Message},
	}

	response, err := client.Chat(messages, gigachat.WithModel(request.Model))
	if err != nil {
		log.Fatal(err)
	}

	res := response.Choices[0].Message.Content

	return res, err
}

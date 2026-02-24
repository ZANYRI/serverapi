package ai

import (
	"context"

	"google.golang.org/genai"

	"ai-api/env"
	"ai-api/data"
)

func RequestToGemini(request *data.Request) (string, error) {

	key, err := env.GetAPIKey(request.Provider)
	if err != nil {
		return "", err
	}

	ctx := context.Background()

	cc := genai.ClientConfig{
		APIKey: key,
	}

	client, err := genai.NewClient(ctx, &cc)
	if err != nil {
		return "", err
	}

	result, err := client.Models.GenerateContent(
		ctx,
		request.Model,
		genai.Text(request.Message),
		nil,
	)
	if err != nil {
		return "", err
	}

	res := result.Text()

	return res, err
}

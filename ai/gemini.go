package ai

import (
	"context"
	"log"

	"google.golang.org/genai"

	"ai-api/env"
	"ai-api/data"
)

func RequestToGemini(request *data.Request) (string, error) {



	key, err := env.LoadEnv(request.Provider)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	cc := genai.ClientConfig{
		APIKey: key,
	}

	client, err := genai.NewClient(ctx, &cc)
	if err != nil {
		log.Fatal(err)
	}

	result, err := client.Models.GenerateContent(
		ctx,
		request.Model,
		genai.Text(request.Message),
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	res := result.Text()

	return res, err
}

package request

import (
	"context"
	"log"

	"google.golang.org/genai"

	"ai-api/env"
)

func RequestToGemini(name_key, message string) string {

	key, _ := env.LoadEnv(name_key) 

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
        "gemini-3-flash-preview",
        genai.Text(message),
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }

	res := result.Text()

	return res
}

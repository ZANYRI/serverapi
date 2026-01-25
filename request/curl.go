package request

import (
	"io"
	"net/http"
	"strings"

	"ai-api/env"
)

func TestFunc(name_key string) string{
    key, _ := env.LoadEnv("GEMINI_API_KEY")

    body := `{
    "contents": [
      {
        "parts": [
          {
            "text": "Что делаешь?"
          }
        ]
      }
    ]
  }`

    req, err := http.NewRequest("POST", "https://generativelanguage.googleapis.com/v1beta/models/gemini-3-flash-preview:generateContent", 
        strings.NewReader(body))
    if err != nil {
        panic(err)
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("x-goog-api-key", key)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    respBody, _ := io.ReadAll(resp.Body)
    q := string(respBody)

	return q
}
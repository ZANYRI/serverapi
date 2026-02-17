package request

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"ai-api/env"
)

func TestFunc(name_key, message string) string {
	key, _ := env.LoadEnv(name_key)

	body := new(Content)
	body.Contents = []Part{
		{
			Parts: []Text{
				{Texts: message},
			},
		},
	}
	fbody, _ := json.Marshal(body)

	req, err := http.NewRequest("POST", "https://generativelanguage.googleapis.com/v1beta/models/gemini-3-flash-preview:generateContent",
		bytes.NewReader(fbody))
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

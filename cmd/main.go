package main

import (
	"ai-api/request"
	"fmt"
)

func main() {
	fmt.Print(request.RequestToGemini("GEMINI_API_KEY", "Хай"))
}
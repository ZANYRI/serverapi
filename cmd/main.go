package main

import (
	"ai-api/request"
	"fmt"
)

func main() {
	fmt.Print(request.TestFunc("GEMINI_API_KEY"))
}
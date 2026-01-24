package main

import (
	"ai-api/env"
	"fmt"
)

func main() {
	fmt.Print(env.LoadEnv("GEMINI_KEY"))
}
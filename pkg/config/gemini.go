package config

import (
	"context"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"log"
	"os"
)

type Gemini struct {
	Key string
}

var gemini = &Gemini{}

func GeminiModel() (*genai.Client, *context.Context) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(gemini.Key))
	if err != nil {
		log.Fatal(err)
	}

	return client, &ctx
}

func loadGeminiCfg() {
	gemini.Key = os.Getenv("GEMINI_API_KEY")
}

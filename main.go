package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"

	"cloud.google.com/go/vertexai/genai"
	"google.golang.org/api/option"
)

func main() {
	err := tryGemini(os.Stdout, "vetexai-demo", "us-central1", "gemini-1.5-pro", "./vetexai-demo-21cff856d7f5.json")
	if err != nil {
		slog.Error(err.Error())
		return
	}
}

func tryGemini(w io.Writer, projectID string, location string, modelName string, credentialsFile string) error {
	// location := "us-central1"
	// modelName := "gemini-1.0-pro-vision-001"

	ctx := context.Background()
	client, err := genai.NewClient(ctx, projectID, location, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		return fmt.Errorf("error creating client: %w", err)
	}
	gemini := client.GenerativeModel(modelName)

	// img := genai.FileData{
	// 	MIMEType: "image/jpeg",
	// 	FileURI:  "gs://generativeai-downloads/images/scones.jpg",
	// }
	// prompt := genai.Text("What is in this image?")

	// resp, err := gemini.GenerateContent(ctx, img, prompt)
	prompt := genai.Text("天空为什么是蓝色的？")
	resp, err := gemini.GenerateContent(ctx, prompt)
	if err != nil {
		return fmt.Errorf("error generating content: %w", err)
	}
	rb, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return fmt.Errorf("json.MarshalIndent: %w", err)
	}
	fmt.Fprintln(w, string(rb))
	return nil
}

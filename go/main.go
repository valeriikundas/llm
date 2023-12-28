package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/replicate/replicate-go"
)

func main() {
	log.SetFlags(log.Ltime)

	err := godotenv.Load("../.env")
	must(err)

	prompt := os.Args[1]

	client, err := replicate.NewClient(replicate.WithTokenFromEnv())
	must(err)

	input := replicate.PredictionInput{"prompt": prompt}

	prediction, err := client.CreatePredictionWithModel(context.Background(), "mistralai", "mixtral-8x7b-instruct-v0.1", input, nil, false)
	must(err)

	err = client.Wait(context.Background(), prediction)
	must(err)

	log.Print(prediction.Output)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"os"

	"ice-breaker/pkg/helpers"

	"github.com/tmc/langchaingo/llms/openai"
)

func main() {
	_ = helpers.LoadDotEnv(".env")
	openai_api_key, ok := os.LookupEnv("OPENAI_API_KEY")
	if !ok {
		panic("OPENAI_API_KEY environment variable not set")
	}

	llm, err := openai.New(openai.WithToken(openai_api_key))
	if err != nil {
		panic(err)
	}
	completion, err := llm.Call(nil, "Hello, world!")
	if err != nil {
		panic(err)
	}
	fmt.Println(completion)

}

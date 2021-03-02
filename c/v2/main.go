package main

import (
	"fmt"

	env "github.com/caarlos0/env/v5"
)

type environment struct {
	Home string `env:"HOME"`
}

func main() {
	// #1
	environment := environment{}
	if err := env.Parse(&environment); err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Printf("%+v\n", environment)
}

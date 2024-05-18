package main

import (
	"log"

	"github.com/opoccomaxao/go-galaxy-experiment/pkg/game"
)

func main() {
	game := game.New()

	err := game.Run()
	if err != nil {
		log.Fatal(err)
	}
}

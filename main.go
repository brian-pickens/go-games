package main

import (
	"log"

	"github.com/brian-pickens/go-games/wordle"
)

func main() {
	err := wordle.StartGame()
	if err != nil {
		log.Fatal(err)
	}
}
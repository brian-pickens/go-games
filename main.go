package main

import (
	"log"

	"github.com/brian-pickens/hello-world-go-game/wordle"
)

func main() {
	err := wordle.StartGame()
	if err != nil {
		log.Fatal(err)
	}
}



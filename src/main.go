package main

import (
	"fmt"
	"math/rand"
	"time"
)

// -------------------- Main --------------------

func main() {
	rand.Seed(time.Now().UnixNano())
	go playSong()
	introduction()

	player := ChooseCharacter()
	fmt.Println("\n✅ Tu as choisi ton héros !")
	player.DisplayInfo()

	player.StartMenu()
}

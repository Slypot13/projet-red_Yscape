package main

import (
	"fmt"
	"time"
)

// Liste des codes ANSI correspondant aux couleurs et à la mise en forme
const (
	Reset     = "\033[0m"
	Red       = "\033[31m"
	Green     = "\033[32m"
	Yellow    = "\033[33m"
	Blue      = "\033[34m"
	Magenta   = "\033[35m"
	Cyan      = "\033[36m"
	Bold      = "\033[1m"
	Underline = "\033[4m"
)

// typeWriter affiche une chaîne de caractères par caractère,
// avec un délai delay entre chaque caractère.
func typeWriter(str string, delay time.Duration) {
	for _, r := range str {
		fmt.Printf("%c", r)
		time.Sleep(delay)
	}
	fmt.Println()
}

func main() {
	// Exemple de chaîne de caractères avec de la couleur & mise en forme
	texte := Blue + Bold + "Bienvenue dans le RPG CLI!\n" + Reset +
		Green + "Préparez-vous à une aventure...\n" + Reset +
		Yellow + Underline + "Que la quête commence !" + Reset

	// On affiche le texte avec effet typeWriter
	typeWriter(texte, 50*time.Millisecond)

	// Ensuite, un message simple sans mise en forme ni couleur
	typeWriter("Tapez 1 pour continuer, 0 pour quitter.", 30*time.Millisecond)
}

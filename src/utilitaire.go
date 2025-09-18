package main

import (
	"fmt"
	"time"
)

// -------------------- Utilitaires Texte --------------------

func typewriterPrint(text string, delay time.Duration, color string) {
	reset := "\033[0m"
	fmt.Print(color)
	for _, c := range text {
		fmt.Printf("%c", c)
		time.Sleep(delay)
	}
	fmt.Println(reset)
}

func colorPrintln(color, s string) {
	reset := "\033[0m"
	fmt.Print(color)
	fmt.Println(s)
	fmt.Print(reset)
}

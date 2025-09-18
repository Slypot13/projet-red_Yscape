package main

import (
	"fmt"
	"strings"
	"unicode"
)

// -------------------- Gestion Nom perso --------------------

func characterCreation(p Character) Character {
	var nom string
	for {
		fmt.Print("✏️  Choisis un nom pour ton personnage : ")
		fmt.Scan(&nom)

		valid := true
		for _, r := range nom {
			if !unicode.IsLetter(r) {
				valid = false
				break
			}
		}
		if !valid || len(nom) == 0 {
			fmt.Println("❌ Le nom ne doit contenir que des lettres.")
			continue
		}
		nom = strings.ToLower(nom)
		nom = strings.Title(nom)
		p.Nom = nom
		break
	}
	return p
}

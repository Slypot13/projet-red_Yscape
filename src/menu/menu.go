package ui

import (
	"fmt"
	"mon-projet/src/character"
	"os"
)

func StartMenu(p *character.Character) {
	for {
		fmt.Println("=== Menu principal ===")
		fmt.Printf("\t1 - Afficher les informations du personnage\n")
		fmt.Printf("\t2 - Accéder au contenu de l’inventaire\n")
		fmt.Printf("\t0 - Quitter\n")
		fmt.Println("Sélectionner un choix (1,2 ou 0) :")
		var userChose int
		fmt.Scan(&userChose)

		switch userChose {
		case 1:
			p.DisplayInfo()
		case 2:
			p.MenuInventory()
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Erreur : choix non valide")
		}
	}
}

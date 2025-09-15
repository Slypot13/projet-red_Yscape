package character

import (
	"fmt"
)

func (p *Character) StartMenu() {
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
			//p.MenuInventory()
		case 3:
			p.MerchantMenu()
		case 0:
			fmt.Println("À bientôt !")
			return
		default:
			fmt.Println("Choix invalide.")
		}

		if p.IsDead() {
			fmt.Println("💀 Vous êtes mort ! Fin du jeu.")
			return
		}
	}
}

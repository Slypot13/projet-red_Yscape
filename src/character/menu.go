package character

import (
	"fmt"
)

func (p *Character) MenuInventory() {
	for {
		p.AccessInventory()
		fmt.Println("=== Menu inventaire ===")
		fmt.Println("\t1 - Utiliser une potion de vie")
		fmt.Println("\t2 - Utiliser une potion de poison")
		fmt.Println("\t0 - Retour")
		fmt.Print("Choix : ")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			p.TakePotion()
		case 2:
			p.UsePoisonPotion()
		case 0:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

package character

import "fmt"

func (p *Character) MerchantMenu() {
	for {
		fmt.Println("=== Marchand ===")
		fmt.Println("\t1 - Acheter une Potion de vie (gratuit)")
		fmt.Println("\t0 - Quitter")
		fmt.Print("Choix : ")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			p.AddInventory("Potion de vie", 1)
			fmt.Println("Vous avez reçu : Potion de vie x1")
		case 0:
			return
		default:
			fmt.Println("Choix invalide")
		}
	}
}

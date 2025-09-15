package character

import "fmt"

//---------------------------------Affiche le menu marchand---------------------------------
func (p *Character) MerchantMenu() {
	for {
		fmt.Println("=== Marchand ===")
		fmt.Println("\t1 - Potion de vie (gratuit)")
		fmt.Println("\t2 - Potion de poison (gratuit)")
		fmt.Println("\t0 - Quitter")
		fmt.Print("Choix : ")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			p.AddInventory("Potion de vie", 1)
			fmt.Println("Vous avez reçu : Potion de vie x1")
		case 2:
			p.AddInventory("Potion de poison", 1)
			fmt.Println("Vous avez reçu : Potion de poison x1")
		case 0:
			return
		default:
			fmt.Println("Choix invalide")
		}
	}
}

package main

import "fmt"

// -------------------- Menu Principal --------------------

func (p *Character) StartMenu() {
	const (
		yellowBold = "\033[1;33m"
		green      = "\033[32m"
		reset      = "\033[0m"
		white      = "\033[97m"
		purpleBold = "\033[1;35m"
	)
	for {
		fmt.Println()
		fmt.Print(purpleBold)
		fmt.Println("╔══════════════════════════════════╗")
		fmt.Println("║         Menu Principal           ║")
		fmt.Println("╚══════════════════════════════════╝")
		fmt.Print(reset)

		fmt.Print(white)
		fmt.Println("\t1 - Afficher infos personnage")
		fmt.Println("\t2 - Inventaire")
		fmt.Println("\t3 - Marchand du Campus")
		fmt.Println("\t4 - Forgeron")
		fmt.Println("\t5 - Combat pour ton Diplome!")
		fmt.Println("\t0 - Quitter")
		fmt.Print(reset)

		fmt.Print("Choix : ")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			p.DisplayInfo()
		case 2:
			p.AccessInventoryMenu()
		case 3:
			p.MerchantMenu()
		case 4:
			p.BlacksmithMenu()
		case 5:
			TrainingFight(p)
		case 0:
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide")
		}
	}
}

package main

import "fmt"

// -------------------- Menus : Marchand & Forgeron --------------------

func (p *Character) MerchantMenu() {
	for {
		fmt.Println()
		fmt.Print("\033[1;33m")
		fmt.Println("┏━━━━━━┛  Marchand du Campus  ┗━━━━━━┓")
		fmt.Print("\033[0m")

		fmt.Print("\033[32m")
		fmt.Println("\t1 - RedBull (+50 PV) - 10£")
		fmt.Println("\t2 - Coca bien frais Chakal (+10% attaque / 1 tour) - 15£")
		fmt.Println("\t3 - Café dilué au Ciao Kambucha (Poison -10PV/s 3s) - 20£")
		fmt.Println("\t4 - Sac à dos perdu (+5 places inventaire) - 30£")
		fmt.Println("\t0 - Retour")
		fmt.Print("\033[0m")

		fmt.Printf("💰 Argent : %d£\n", p.Argent)
		fmt.Print("Choix : ")
		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			if p.Argent < 10 {
				fmt.Println("❌ Pas assez d'argent pour acheter RedBull.")
			} else {
				if ok := p.AddInventory("RedBull", 1); ok {
					p.Argent -= 10
					fmt.Println("✅ Achat : RedBull (ajouté à l'inventaire).")
				} else {
					fmt.Println("❌ Inventaire plein, impossible d'acheter.")
				}
			}
		case 2:
			if p.Argent < 15 {
				fmt.Println("❌ Pas assez d'argent pour acheter le Coca.")
			} else {
				if ok := p.AddInventory("Coca bien frais Chakal", 1); ok {
					p.Argent -= 15
					fmt.Println("✅ Achat : Coca bien frais Chakal (ajouté à l'inventaire).")
				} else {
					fmt.Println("❌ Inventaire plein, impossible d'acheter.")
				}
			}
		case 3:
			if p.Argent < 20 {
				fmt.Println("❌ Pas assez d'argent pour acheter le Café dilué.")
			} else {
				if ok := p.AddInventory("Café dilué au Ciao Kambucha", 1); ok {
					p.Argent -= 20
					fmt.Println("✅ Achat : Café dilué au Ciao Kambucha (ajouté à l'inventaire).")
				} else {
					fmt.Println("❌ Inventaire plein, impossible d'acheter.")
				}
			}
		case 4:
			if p.Backpack {
				fmt.Println("🎒 Tu possèdes déjà le Sac à dos perdu.")
			} else if p.Argent < 30 {
				fmt.Println("❌ Pas assez d'argent pour acheter le Sac à dos perdu.")
			} else {
				p.Argent -= 30
				p.Backpack = true
				p.MaxInv = 10
				fmt.Println("✅ Achat : Sac à dos perdu. Inventaire étendu à 10 emplacements.")
			}
		case 0:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func (p *Character) BlacksmithMenu() {
	for {
		fmt.Println()
		fmt.Print("\033[1;36m")
		fmt.Println("┏━━━━━━┛  Forgeron  ┗━━━━━━┓")
		fmt.Print("\033[0m")

		fmt.Print("\033[32m")
		fmt.Println("\t1 - Casquette Gucci (60 pièces) [+20% attaque à partir du 3e tour]")
		fmt.Println("\t2 - Asics Kayano (60 pièces) [empêche le monstre de jouer 1 tour]")
		fmt.Println("\t0 - Retour")
		fmt.Print("\033[0m")

		fmt.Printf("⚠️ Pour acheter et équiper, tu dois posséder l'objet 'flow du contrôleur RATP'.\n")
		fmt.Printf("💰 Pièces : %d\n", p.Pieces)
		fmt.Print("Choix : ")
		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			if p.Pieces < 60 {
				fmt.Println("❌ Pas assez de pièces pour la Casquette Gucci.")
				continue
			}
			if !p.HasItem("flow du contrôleur RATP") {
				fmt.Println("❌ Il te manque 'flow du contrôleur RATP' pour acheter/équiper cet équipement.")
				continue
			}
			if ok := p.AddInventory("Casquette Gucci", 1); !ok {
				fmt.Println("❌ Inventaire plein, impossible d'ajouter la Casquette Gucci.")
				continue
			}
			p.Pieces -= 60
			fmt.Println("✅ Tu as fabriqué la Casquette Gucci (ajoutée à l'inventaire).")
		case 2:
			if p.Pieces < 60 {
				fmt.Println("❌ Pas assez de pièces pour les Asics Kayano.")
				continue
			}
			if !p.HasItem("flow du contrôleur RATP") {
				fmt.Println("❌ Il te manque 'flow du contrôleur RATP' pour acheter/équiper cet équipement.")
				continue
			}
			if ok := p.AddInventory("Asics Kayano", 1); !ok {
				fmt.Println("❌ Inventaire plein, impossible d'ajouter Asics Kayano.")
				continue
			}
			p.Pieces -= 60
			fmt.Println("✅ Tu as fabriqué les Asics Kayano (ajoutées à l'inventaire).")
		case 0:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

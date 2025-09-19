package main

import "fmt"

// -------------------- Menus : Marchand & Forgeron --------------------

func (p *Character) MerchantMenu() {
	for {
		fmt.Println()
		fmt.Print("\033[1;33m") // Jaune gras
		fmt.Println("┏━━━━━━┛  Marchand du Campus  ┛━━━━━━┓")
		fmt.Print("\033[0m") // Reset couleur

		// Options en blanc
		fmt.Print("\033[97m") // Blanc
		fmt.Println("\t1 - RedBull (+50 PV) - 10£")
		fmt.Println("\t2 - Coca bien frais Chakal (+10% attaque / 1 tour) - 15£")
		fmt.Println("\t3 - Café dilué au Ciao Kambucha (Poison -10PV/s 3s) - 20£")
		fmt.Println("\t4 - Sacoche perdu (+5 places inventaire) - 30£")
		fmt.Println("\t5 - Épée Légendaire (inflige +100 dégâts en combat) - 120£") // <- nouvel objet ajouté
		fmt.Println("\t0 - Retour")
		fmt.Print("\033[0m") // Reset couleur

		// Infos supplémentaires
		fmt.Printf("💰 Argent : %d£\n", p.Argent)
		fmt.Print("Choix : ")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			if p.Argent < 10 {
				fmt.Println("❌ Pas assez d'argent pour acheter une RedBull.")
			} else {
				if ok := p.AddInventory("RedBull", 1); ok {
					p.Argent -= 10
					fmt.Println("✅ Achat : RedBull (ajoutée à l'inventaire).")
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
				fmt.Println("🎒 Tu possèdes déjà la Sacoche perdue.")
			} else if p.Argent < 30 {
				fmt.Println("❌ Pas assez d'argent pour acheter la Sacoche perdue.")
			} else {
				p.Argent -= 30
				p.Backpack = true
				p.MaxInv = 10
				fmt.Println("✅ Achat : Sacoche perdue. Inventaire étendu à 10 emplacements.")
			}
		case 5: // Épée Légendaire
			if p.Argent < 150 {
				fmt.Println("❌ Pas assez d'argent pour acheter l'Épée Légendaire.")
			} else {
				if ok := p.AddInventory("Épée Légendaire", 1); ok {
					p.Argent -= 150
					fmt.Println("✅ Achat : Épée Légendaire (ajoutée à l'inventaire).")
				} else {
					fmt.Println("❌ Inventaire plein, impossible d'acheter.")
				}
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
		fmt.Print("\033[1;33m") // Jaune
		fmt.Println("┏━━━━━━┛  Forgeron  ┛━━━━━━┓")
		fmt.Print("\033[0m")

		// Options en blanc
		fmt.Print("\033[97m") // Blanc
		fmt.Println("\t1 - Casquette Gucci (60 pièces) [+20% attaque à partir du 3e tour]")
		fmt.Println("\t2 - Asics Kayano (60 pièces) [empêche le monstre de jouer 1 tour]")

		// Ajouter l'option pour la Flûte de gasba (débloquée après la victoire contre les frères Khabils)
		if p.HasItem("La puissance DZ") {
			fmt.Println("\t3 - Flûte de gasba (30% de dégâts supplémentaires) - 80 pièces")
		}

		fmt.Println("\t0 - Retour")
		fmt.Print("\033[0m")

		// Infos supplémentaires
		fmt.Printf("⚠️ Pour acheter et équiper, tu dois posséder l'objet 'Flow du Contrôleur RATP'.\n")
		fmt.Printf("💰 Argent : %d\n", p.Argent)
		fmt.Print("Choix : ")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			if p.Argent < 60 {
				fmt.Println("❌ Pas assez de pièces pour la Casquette Gucci.")
				continue
			}
			if !p.HasItem("Flow du Contrôleur RATP") {
				fmt.Println("❌ Il te manque 'flow du contrôleur RATP' pour acheter/équiper cet équipement.")
				continue
			}
			if ok := p.AddInventory("Casquette Gucci", 1); !ok {
				fmt.Println("❌ Inventaire plein, impossible d'ajouter la Casquette Gucci.")
				continue
			}
			p.Argent -= 60
			fmt.Println("✅ Tu as fabriqué la Casquette Gucci (ajoutée à l'inventaire).")
		case 2:
			if p.Argent < 60 {
				fmt.Println("❌ Pas assez de pièces pour les Asics Kayano.")
				continue
			}
			if !p.HasItem("Flow du Contrôleur RATP") {
				fmt.Println("❌ Il te manque 'flow du contrôleur RATP' pour acheter/équiper cet équipement.")
				continue
			}
			if ok := p.AddInventory("Asics Kayano", 1); !ok {
				fmt.Println("❌ Inventaire plein, impossible d'ajouter Asics Kayano.")
				continue
			}
			p.Argent -= 60
			fmt.Println("✅ Tu as fabriqué les Asics Kayano (ajoutées à l'inventaire).")
		case 3: // Flûte de gasba
			if p.Argent < 80 {
				fmt.Println("❌ Pas assez de pièces pour la Flûte de gasba.")
				continue
			}
			if !p.HasItem("La puissance DZ") {
				fmt.Println("❌ Tu dois avoir 'La puissance DZ' pour acheter la Flûte de gasba.")
				continue
			}
			if ok := p.AddInventory("Flûte de gasba", 1); !ok {
				fmt.Println("❌ Inventaire plein, impossible d'ajouter la Flûte de gasba.")
				continue
			}
			p.Argent -= 80
			fmt.Println("✅ Tu as fabriqué la Flûte de gasba (ajoutée à l'inventaire).")
		case 0:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

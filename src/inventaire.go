package main

import "fmt"

// -------------------- Inventaire --------------------

func (p *Character) DisplayInventory() {
	fmt.Println("\n🎒 Inventaire :")
	if len(p.Inventaire) == 0 {
		fmt.Println("  (vide)")
	} else {
		for i, item := range p.Inventaire {
			fmt.Printf(" %d - %s (x%d)\n", i+1, item.Nom, item.Quantite)
		}
	}
	fmt.Printf("Capacité utilisée : %d/%d\n", len(p.Inventaire), p.MaxInv)
}

func (p *Character) AddItem(nom string) {
	if len(p.Inventaire) >= p.MaxInv {
		fmt.Println("❌ Ton inventaire est plein !")
		return
	}
	for i, item := range p.Inventaire {
		if item.Nom == nom {
			p.Inventaire[i].Quantite++
			fmt.Printf("✅ %s ajouté à l'inventaire (x%d)\n", nom, p.Inventaire[i].Quantite)
			return
		}
	}
	p.Inventaire = append(p.Inventaire, Item{Nom: nom, Quantite: 1})
	fmt.Printf("✅ %s ajouté à l'inventaire !\n", nom)
}

// -------------------- Inventaire : affichage et utilisation --------------------

func (p *Character) AccessInventoryMenu() {
	fmt.Println()
	fmt.Print("\033[1;33m")
	fmt.Println("╭─━━━━━─╯  Inventaire  ╰─━━━━━─╮")
	fmt.Print("\033[0m")

	if len(p.Inventaire) == 0 {
		fmt.Println("\tInventaire vide")
		return
	}
	for i, it := range p.Inventaire {
		fmt.Printf("\t%d - %s x%d\n", i+1, it.Nom, it.Quantite)
	}
	fmt.Printf("\t0 - Retour\n")

	var choix int
	fmt.Print("Choisis un objet (numéro) pour l'utiliser / équiper : ")
	fmt.Scan(&choix)

	if choix == 0 {
		return
	}
	if choix < 1 || choix > len(p.Inventaire) {
		fmt.Println("Choix invalide.")
		return
	}
	name := p.Inventaire[choix-1].Nom

	switch name {
	case "RedBull":
		p.UseRedBull()
	case "Coca bien frais Chakal":
		p.UseCoca()
	case "Café dilué au Ciao Kambucha":
		// on applique poison sur le joueur (ou sur ennemi si en combat, ici hors combat on montre l'effet)
		fmt.Println("⚠️ Le Café dilué est toxique à consommer hors combat. Ne l'utilise pas maintenant.")
	case "Bouteille de Kambucha alcoolisé à 2%":
		p.UseKambucha()
	case "Casquette Gucci":
		p.EquipFromInventory("Casquette Gucci")
	case "Asics Kayano":
		p.EquipFromInventory("Asics Kayano")
	case "flow du contrôleur RATP":
		fmt.Println("Ceci est un composant spécial pour le Forgeron.")
	default:
		fmt.Println("Objet non reconnu / utilisation non implémentée.")
	}
}

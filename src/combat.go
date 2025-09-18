package main

import (
	"fmt"
	"math/rand"
)

// -------------------- Combat : Goblin (ClasseInfobugé) --------------------

func InitGoblin() Monster {
	return Monster{Nom: "ClasseInfobugé", PvMax: 115, Pv: 115, Attaque: 20}
}

func GoblinPattern(monstre *Monster, player *Character, tour int) {
	attack := monstre.Attaque
	if tour%3 == 0 {
		attack = int(float64(monstre.Attaque) * 1.2)
	}
	cris := []string{"Lancé de carte graphique", "Lancé de souris", "Lancé de clavier"}
	cri := cris[rand.Intn(len(cris))]
	fmt.Printf("\n%s crie \"%s\" !\n", monstre.Nom, cri)

	// si joueur a Asics effect (empêche une action), handle in CharacterTurn when player triggers it.
	player.Pv -= attack
	if player.Pv < 0 {
		player.Pv = 0
	}
	fmt.Printf("%s inflige à %s %d dégâts !\n", monstre.Nom, player.Nom, attack)
	fmt.Printf("%s PV : %d/%d\n", player.Nom, player.Pv, player.PvMax)
}

func CharacterTurn(monstre *Monster, player *Character, tour *int, monsterSkipped *bool) {
	fmt.Println("\n--- Ton tour ---")
	fmt.Println("1 - Attaquer")
	fmt.Println("2 - Inventaire")
	fmt.Println("3 - Utiliser Asics (si équipé) [bloque l'ennemi 1 tour]")
	fmt.Print("Choix : ")
	var choix int
	fmt.Scan(&choix)
	switch choix {
	case 1:
		damage := player.Attaque
		fmt.Printf("%s utilise Attaque basique et inflige %d dégâts !\n", player.Nom, damage)
		monstre.Pv -= damage
		if monstre.Pv < 0 {
			monstre.Pv = 0
		}

		// Affichage barre de vie monstre
		bar := DisplayHPBar(monstre.Pv, monstre.PvMax, 20)
		fmt.Printf("%s PV : [%s] %d/%d\n", monstre.Nom, bar, monstre.Pv, monstre.PvMax)

	case 2:
		player.AccessInventoryCombat()
	case 3:
		if player.Equipement.Pieds == "Asics Kayano" && player.HasAsicsEffect {
			*monsterSkipped = true
			player.HasAsicsEffect = false
			fmt.Println("👟 Tu actives Asics Kayano : le monstre est bloqué pour 1 tour !")
		} else {
			fmt.Println("❌ Tu n'as pas Asics Kayano équipées ou l'effet n'est pas disponible.")
		}
	default:
		fmt.Println("Choix invalide, tu perds ton action.")
	}
	fmt.Println("-----------------------")
}

func (p *Character) AccessInventoryCombat() {
	if len(p.Inventaire) == 0 {
		fmt.Println("Inventaire vide.")
		return
	}
	fmt.Println("Inventaire :")
	for i, it := range p.Inventaire {
		fmt.Printf("%d - %s x%d\n", i+1, it.Nom, it.Quantite)
	}
	fmt.Print("Choisis un objet (numéro) : ")
	var choix int
	fmt.Scan(&choix)
	if choix < 1 || choix > len(p.Inventaire) {
		fmt.Println("Choix invalide.")
		return
	}
	switch p.Inventaire[choix-1].Nom {
	case "RedBull":
		p.UseRedBull()
	case "Bouteille de Kambucha alcoolisé à 2%":
		p.UseKambucha()
	case "Coca bien frais Chakal":
		p.UseCoca()
	case "Café dilué au Ciao Kambucha":
		fmt.Println("☠️ Tu lances le Café dilué sur l'ennemi (poison) ! (implémentation simplifiée)")
		// appliquer poison de test sur le monstre via un petit loop: pour simplifier, on applique direct damage here
		// but since monstre passed by pointer in caller, we cannot access it here; in combat we handle using AccessInventoryCombat from CharacterTurn when relevant.
	default:
		fmt.Println("Objet non utilisable en combat.")
	}
}

// TrainingFight : le combat d'entraînement contre la ClasseInfobugé
func TrainingFight(player *Character) {
	monstre := InitGoblin()
	tour := 1
	monsterSkipped := false
	fmt.Println("\n⚔️ Début du combat contre", monstre.Nom, "!")
	for player.Pv > 0 && monstre.Pv > 0 {
		fmt.Printf("\n======== Tour %d ========\n", tour)

		// Si casquette équipée et pas encore active, décrémente et active si nécessaire
		if player.Equipement.Tete == "Casquette Gucci" && !player.CasquetteActive {
			if player.CasquetteDelay > 0 {
				fmt.Printf("(Casquette Gucci : %d tours avant activation)\n", player.CasquetteDelay)
				player.CasquetteDelay--
				if player.CasquetteDelay == 0 {
					// activer l'effet : +20% attaque
					bonus := int(float64(player.Attaque) * 0.20)
					if bonus < 1 {
						bonus = 1
					}
					player.Attaque += bonus
					player.CasquetteActive = true
					fmt.Printf("🧢 Casquette Gucci s'active ! Attaque augmentée de +%d (Attaque = %d)\n", bonus, player.Attaque)
				}
			}
		}

		// Tour du joueur
		CharacterTurn(&monstre, player, &tour, &monsterSkipped)
		if monstre.Pv <= 0 {
			break
		}

		// Tour du monstre (sauf si joueur a appliqué Asics -> monsterSkipped true)
		if monsterSkipped {
			fmt.Println("\nLe monstre est bloqué ce tour, il ne peut pas attaquer.")
			monsterSkipped = false
		} else {
			GoblinPattern(&monstre, player, tour)
		}

		// Après l'attaque du monstre, gérer la durée du boost Coca
		if player.TempBoostTurns > 0 {
			player.TempBoostTurns--
			if player.TempBoostTurns == 0 {
				// revert boost
				player.Attaque -= player.TempAttackBoost
				fmt.Printf("🥤 Effet Coca terminé. Attaque revenue à %d\n", player.Attaque)
				player.TempAttackBoost = 0
			}
		}

		// vérifier si joueur mort
		if player.Pv <= 0 {
			break
		}
		tour++
	}

	// Résultat du combat
	if player.Pv <= 0 {
		fmt.Println("\n❌ Tu es vaincu... Retour au menu principal.")
		// optionnel : restaurer PV partiellement ou renvoyer au menu
	} else {
		fmt.Println("\n🎉 Yesss mon gaté c'est gagné ! EZ la classe")
		player.Argent += 15
		added := player.AddInventory("Bouteille de Kambucha alcoolisé à 2%", 1)
		if added {
			fmt.Println("Récompense : +15 pièces et 1x Bouteille de Kambucha alcoolisé à 2% ajouté à l'inventaire (soigne 30PV).")
		} else {
			// si inventaire plein, on donne argent à la place
			player.Argent += 0
			fmt.Println("Ton inventaire était plein : la récompense 'Bouteille de Kambucha' n'a pas pu être ajoutée.")
			fmt.Println("Tu as quand même reçu +15 pièces.")
		}
	}
}

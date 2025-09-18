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
		// Attaque basique - tenir compte si casquette active => l'effet casquette s'applique via CasquetteActive field
		damage := player.Attaque
		fmt.Printf("%s utilise Attaque basique et inflige %d dégâts !\n", player.Nom, damage)
		monstre.Pv -= damage
		if monstre.Pv < 0 {
			monstre.Pv = 0
		}
		fmt.Printf("%s PV : %d/%d\n", monstre.Nom, monstre.Pv, monstre.PvMax)
	case 2:
		// Accès à l'inventaire (utilisation d'objets en combat)
		player.AccessInventoryCombat()
	case 3:
		// utiliser Asics si équipé
		if player.Equipement.Pieds == "Asics Kayano" && player.HasAsicsEffect {
			*monsterSkipped = true
			player.HasAsicsEffect = false // effet consommé (ou garder selon design)
			fmt.Println("👟 Tu actives Asics Kayano : le monstre est bloqué pour 1 tour !")
		} else {
			fmt.Println("❌ Tu n'as pas Asics Kayano équipées ou l'effet n'est pas disponible.")
		}
	default:
		fmt.Println("Choix invalide, tu perds ton action.")
	}
	// fin du turn du joueur, on décrémente les tours boost si actifs (mais revert après le tour du monstre)
	// CasquetteDelay est gérée dans boucle de combat
	_ = tour
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
	fmt.Println("\n⚔️ Début du combat d'entraînement contre", monstre.Nom, "!")
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

	// Fin du combat
	if player.Pv <= 0 {
		fmt.Println("\n❌ Tes vaincu, tu flop trop... Retour au menu principal.")
	} else {
		fmt.Println("\n🎉 Yesss mon gaté c'est gagné ! EZ la classe")
		player.Pieces += 15
		added := player.AddInventory("Bouteille de Kambucha alcoolisé à 2%", 1)
		if added {
			fmt.Println("Récompense : +15 pièces et 1x Bouteille de Kambucha alcoolisé à 2% ajouté à l'inventaire (soigne 30PV).")
		} else {
			player.Argent += 0
			fmt.Println("Ton inventaire était plein : la récompense 'Bouteille de Kambucha' n'a pas pu être ajoutée.")
			fmt.Println("Tu as quand même reçu +15 pièces.")
			// === SUITE ===

		}
	}
	fmt.Println("\n🚪 Une nouvelle salle s'ouvre devant toi...")
	fmt.Println("👉 Tu peux maintenant accéder au niveau B2 !")
	Boss2Fight(player)

	// ========================
	// BOSS 2 : CONTRÔLEUR RATP
	// ========================
}
func InitControleurRATP() Monster {
	return Monster{Nom: "Contrôleur RATP", PvMax: 150, Pv: 150, Attaque: 30}
}

func ControleurRATPPattern(monstre *Monster, player *Character, tour int) {
	attack := monstre.Attaque
	if tour%2 == 0 {
		attack = int(float64(monstre.Attaque) * 1.3)
		fmt.Printf("\n🚆 %s te jette un train entier !\n", monstre.Nom)
	} else {
		fmt.Printf("\n🚇 %s te matraque comme un fraudeur !\n", monstre.Nom)
	}
	player.Pv -= attack
	if player.Pv < 0 {
		player.Pv = 0
	}
	fmt.Printf("%s inflige %d dégâts à %s !\n", monstre.Nom, attack, player.Nom)
}

func Boss2PreFight(player *Character) {
	fmt.Println("\n🚪 Trois portes devant toi :")
	fmt.Println("1 - Salle 201")
	fmt.Println("2 - Salle 202")
	fmt.Println("3 - Salle 203")
	fmt.Print("Choix : ")

	var choix int
	fmt.Scan(&choix)

	switch choix {
	case 1:
		fmt.Println("\n👷 Une Archi t'aide : +10 dégâts à tes attaques contre le Contrôleur RATP.")
		player.TempAttackBoost += 10
	case 2:
		fmt.Println("\n🥤 Salle vide, juste une bouteille d’eau vide...")
	case 3:
		fmt.Println("\n📜 Tu trouves un Passe Navigo 2013 !")
		player.AddInventory("Passe Navigo 2013", 1)
	default:
		fmt.Println("❌ Choix invalide, aucun bonus.")
	}
}

func Boss2Fight(player *Character) {
	Boss2PreFight(player)

	monstre := InitControleurRATP()
	tour := 1
	monsterSkipped := false
	fmt.Println("\n⚔️ Le Contrôleur RATP apparaît avec son gilet fluorescent !")

	for player.Pv > 0 && monstre.Pv > 0 {
		fmt.Printf("\n======== Tour %d ========\n", tour)
		CharacterTurn(&monstre, player, &tour, &monsterSkipped)
		if monstre.Pv <= 0 {
			break
		}
		if monsterSkipped {
			fmt.Println("\n🚷 Le Contrôleur est bloqué ce tour !")
			monsterSkipped = false
		} else {
			ControleurRATPPattern(&monstre, player, tour)
		}
		if player.Pv <= 0 {
			break
		}
		tour++
	}

	if player.Pv <= 0 {
		fmt.Println("\n❌ Le Contrôleur t'a recalé sans ticket...")
	} else {
		fmt.Println("\n🎉 Victoire ! Tu as battu le Contrôleur RATP.")
		player.Pieces += 30
		player.AddInventory("Flow du Contrôleur RATP", 1)
	}
}

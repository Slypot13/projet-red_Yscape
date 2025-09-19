package main

import (
	"fmt"
	"math/rand"
)

// État de progression des boss
var bossProgression = [5]bool{false, false, false, false, false}

// -------------------- Combat principal --------------------
func TrainingFight(p *Character) {
	for {
		// Si le dernier boss a été vaincu, on arrête le jeu et on affiche la narration
		if bossProgression[4] { // Le boss Marocain Start-Up est vaincu
			// Affichage de la narration finale
			fmt.Println("\n🎉 Félicitations !")
			fmt.Println("Vous avez réussi à vaincre le Marocain Start-Up, le dernier obstacle de l'école.")
			fmt.Println("Le système de l'école est maintenant débugué, et vous avez le diplôme en main.")
			fmt.Println("Vous accedez au toit de l'école, où un hélicoptère vous attend.")
			fmt.Println("Avec votre diplôme, vous êtes prêt à affronter de nouveaux défis, à débuguer le monde entier !")
			fmt.Println("Merci d'avoir joué ! Le jeu est terminé.")
			return // On arrête ici la fonction et on ne retourne pas au menu des boss
		}

		// Affichage du menu des boss
		fmt.Println("\n🏆 Choisis ton boss :")
		if !bossProgression[0] {
			fmt.Println("1 - Classe Infobugé (niveau 1)")
		} else {
			fmt.Println("1 - Classe Infobugé ✅")
		}

		if bossProgression[0] && !bossProgression[1] {
			fmt.Println("2 - Le Contrôleur RATP (niveau 2)")
		} else if bossProgression[1] {
			fmt.Println("2 - Le Contrôleur RATP ✅")
		}

		if bossProgression[1] && !bossProgression[2] {
			fmt.Println("3 - Les deux frères Khabil (niveau 3)")
		} else if bossProgression[2] {
			fmt.Println("3 - Les deux frères Khabil ✅")
		}

		if bossProgression[2] && !bossProgression[3] {
			fmt.Println("4 - Lucas et les Archi (niveau 4)")
		} else if bossProgression[3] {
			fmt.Println("4 - Lucas et les Archi ✅")
		}

		if bossProgression[3] && !bossProgression[4] {
			fmt.Println("5 - Marocain Start-Up (niveau 5)")
		} else if bossProgression[4] {
			fmt.Println("5 - Marocain Start-Up ✅")
		}

		fmt.Println("0 - Retour")
		fmt.Print("Choix : ")
		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			if bossProgression[0] {
				fmt.Println("✅ Boss déjà vaincu.")
			} else {
				startCombat(p, Monster{"Classe Infobugé", 120, 120, 10}, 0, 50)
			}
		case 2:
			if !bossProgression[0] {
				fmt.Println("❌ Tu dois d'abord vaincre le boss précédent.")
			} else if bossProgression[1] {
				fmt.Println("✅ Boss déjà vaincu.")
			} else {
				startCombat(p, Monster{"Le Contrôleur RATP", 200, 200, 20}, 1, 100)
			}
		case 3:
			if !bossProgression[1] {
				fmt.Println("❌ Tu dois d'abord vaincre le boss précédent.")
			} else if bossProgression[2] {
				fmt.Println("✅ Boss déjà vaincu.")
			} else {
				startCombat(p, Monster{"Les deux frères Khabil", 350, 350, 30}, 2, 150)
			}
		case 4:
			if !bossProgression[2] {
				fmt.Println("❌ Tu dois d'abord vaincre le boss précédent.")
			} else if bossProgression[3] {
				fmt.Println("✅ Boss déjà vaincu.")
			} else {
				startCombat(p, Monster{"Lucas et les Archi", 500, 500, 40}, 3, 200)
			}
		case 5:
			if !bossProgression[3] {
				fmt.Println("❌ Tu dois d'abord vaincre le boss précédent.")
			} else if bossProgression[4] {
				fmt.Println("✅ Boss déjà vaincu.")
			} else {
				startCombat(p, Monster{"Marocain Start-Up", 700, 700, 50}, 4, 300)
			}
		case 0:
			return
		default:
			fmt.Println("❌ Choix invalide.")
		}
	}
}

// Fonction qui affiche une barre de vie esthétique
func afficherBarreDeVie(nom string, pv, pvMax int) {
	barre := "["
	percent := (pv * 100) / pvMax
	for i := 0; i < 20; i++ {
		if i < (percent / 5) {
			barre += "█"
		} else {
			barre += " "
		}
	}
	barre += "]"
	fmt.Printf("%s %s %d/%d %s\n", nom, barre, pv, pvMax, fmt.Sprintf("(%d%%)", percent))
}

// Fonction qui lance le combat avec un boss
func startCombat(p *Character, m Monster, bossIndex int, reward int) {
	fmt.Printf("\n⚔️  Combat contre %s ! (%d PV, %d ATK)\n", m.Nom, m.Pv, m.Attaque)

	p.Tour = 1
	if p.HasAsicsEffect {
		fmt.Println("👟 Effet Asics activé : Le boss est paralysé ce tour.")
	} else {
		p.Tour = 0
	}

	// Début du combat
	for p.Pv > 0 && m.Pv > 0 {
		// Affichage de la barre de vie esthétique pour le joueur et le monstre
		afficherBarreDeVie(p.Nom, p.Pv, p.PvMax)
		afficherBarreDeVie(m.Nom, m.Pv, m.PvMax)

		// Si on combat les frères Khabil, il y a une chance sur 4 de déclencher la flûte de Gasba
		if m.Nom == "Les deux frères Khabil" && rand.Intn(4) == 1 {
			fmt.Println("🎶 Flûte de Gasba activée ! Tu es endormi pendant un tour...")
			p.Tour = 0 // Le joueur est endormi pendant ce tour et ne pourra pas attaquer
		}

		// Affichage du statut de combat
		fmt.Printf("\nTour %d - %s: %d/%d PV | %s: %d/%d PV\n", p.Tour, p.Nom, p.Pv, p.PvMax, m.Nom, m.Pv, m.PvMax)
		fmt.Println("1 - Attaquer")
		fmt.Println("2 - Inventaire")
		fmt.Println("3 - Abandonner")
		fmt.Print("Choix : ")
		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			// Appliquer les effets des capacités spéciales avant de calculer les dégâts
			dmg := p.Attaque

			// Effet du Trackeur (1 chance sur 3 de doubler les dégâts)
			if p.Nom == "Le Trackeur" && rand.Intn(3) == 1 {
				dmg *= 2
				fmt.Println("⚡ KAMEAMEAMEAMEA ! Attaque doublée !")
			}

			// Bonus de l'Ingénieur
			if p.Nom == "L'Ingénieur" {
				dmg += 10
			}

			// Casquette Gucci active ?
			if p.Equipement.Tete == "Casquette Gucci" {
				if !p.CasquetteActive {
					p.CasquetteDelay--
					if p.CasquetteDelay <= 0 {
						p.CasquetteActive = true
						fmt.Println("🧢 Casquette Gucci activée ! +20% d'attaque.")
					}
				}
				if p.CasquetteActive {
					bonus := int(float64(dmg) * 0.20)
					dmg += bonus
				}
			}

			// Calcul des dégâts
			fmt.Printf("🗡️  Tu infliges %d dégâts à %s !\n", dmg, m.Nom)
			m.Pv -= dmg
			if m.Pv < 0 {
				m.Pv = 0
			}
		case 2:
			p.AccessInventoryMenu()
			continue
		case 3:
			fmt.Println("🚪 Tu as abandonné le combat.")
			return
		default:
			fmt.Println("❌ Choix invalide.")
			continue
		}

		// Application du poison (si applicable)
		if p.TempBoostTurns > 0 {
			p.TempBoostTurns--
			if p.TempBoostTurns == 0 {
				p.Attaque -= p.TempAttackBoost
				p.TempAttackBoost = 0
				fmt.Println("🧪 Fin de l'effet du Coca.")
			}
		}

		// Tour du boss
		if p.HasAsicsEffect && p.Tour == 0 {
			fmt.Println("⏳ Le boss reprend ses esprits...")
		} else if m.Pv > 0 {
			fmt.Printf("💥 %s t'attaque et inflige %d dégâts.\n", m.Nom, m.Attaque)
			p.Pv -= m.Attaque
			if p.Pv < 0 {
				p.Pv = 0
			}
		}
		p.Tour++
	}

	if p.Pv <= 0 {
		fmt.Println("💀 Tu as été vaincu...")
	} else {
		fmt.Printf("🎉 Tu as vaincu %s !\n", m.Nom)
		bossProgression[bossIndex] = true

		// Récompense en argent
		p.Argent += reward
		fmt.Printf("💰 Tu as gagné %d£ !\n", reward)

		// Récompense spéciale pour boss 2
		if m.Nom == "Le Contrôleur RATP" {
			if !p.HasItem("Flow du Contrôleur RATP") {
				p.AddItem("Flow du Contrôleur RATP")
				fmt.Println("🎁 Tu as obtenu le Flow du Contrôleur RATP !")
			}
		}
	}
}

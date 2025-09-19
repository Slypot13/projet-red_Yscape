package main

import (
	"fmt"
)

// État de progression des boss
var bossProgression = [5]bool{false, false, false, false, false}

// -------------------- Combat principal --------------------
func TrainingFight(p *Character) {
	for {
		// Si le dernier boss est vaincu -> fin du jeu
		if bossProgression[4] {
			fmt.Println("\n🎉 Félicitations !")
			fmt.Println("Vous avez réussi à vaincre le Marocain Start-Up, le dernier obstacle de l'école.")
			fmt.Println("Le système est maintenant débugué, et vous avez le diplôme en main.")
			fmt.Println("Un hélico vous attend sur le toit...")
			fmt.Println("Merci d'avoir joué ! Le jeu est terminé.")
			return
		}

		// Menu des boss
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
			fmt.Println("3 - Les deux frères Kabyle (niveau 3)")
		} else if bossProgression[2] {
			fmt.Println("3 - Les deux frères Kabyle ✅")
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
			if !bossProgression[0] {
				startCombat(p, Monster{"Classe Infobugé", 100, 100, 10}, 0, 50)
			} else {
				fmt.Println("✅ Boss déjà vaincu.")
			}
		case 2:
			if bossProgression[0] && !bossProgression[1] {
				startCombat(p, Monster{"Le Contrôleur RATP", 150, 150, 20}, 1, 100)
			} else if bossProgression[1] {
				fmt.Println("✅ Boss déjà vaincu.")
			} else {
				fmt.Println("❌ Tu dois d'abord battre le boss précédent.")
			}
		case 3:
			if bossProgression[1] && !bossProgression[2] {
				startCombat(p, Monster{"Les deux frères Kabyle", 200, 200, 30}, 2, 150)
			} else if bossProgression[2] {
				fmt.Println("✅ Boss déjà vaincu.")
			} else {
				fmt.Println("❌ Tu dois d'abord battre le boss précédent.")
			}
		case 4:
			if bossProgression[2] && !bossProgression[3] {
				startCombat(p, Monster{"Lucas et les Archi", 250, 250, 40}, 3, 200)
			} else if bossProgression[3] {
				fmt.Println("✅ Boss déjà vaincu.")
			} else {
				fmt.Println("❌ Tu dois d'abord battre le boss précédent.")
			}
		case 5:
			if bossProgression[3] && !bossProgression[4] {
				startCombat(p, Monster{"Marocain Start-Up", 300, 300, 50}, 4, 300)
			} else if bossProgression[4] {
				fmt.Println("✅ Boss déjà vaincu.")
			} else {
				fmt.Println("❌ Tu dois d'abord battre le boss précédent.")
			}
		case 0:
			return
		default:
			fmt.Println("❌ Choix invalide.")
		}
	}
}

// -------------------- Fonction de combat --------------------
func startCombat(p *Character, m Monster, bossIndex int, reward int) {
	fmt.Printf("\n⚔️ Combat contre %s ! (%d PV, %d ATK)\n", m.Nom, m.Pv, m.Attaque)

	// Bonus Archi (après Lucas)
	if p.BonusNext > 0 {
		fmt.Printf("🔥 Bonus Archi : +%d dégâts pour ce combat !\n", p.BonusNext)
	}

	// Début du combat
	for p.Pv > 0 && m.Pv > 0 {
		fmt.Printf("\nTour %d - %s: %d/%d PV | %s: %d/%d PV\n",
			p.Tour, p.Nom, p.Pv, p.PvMax, m.Nom, m.Pv, m.PvMax)
		fmt.Println("1 - Attaquer")
		fmt.Println("2 - Inventaire")
		fmt.Println("3 - Abandonner")
		fmt.Print("Choix : ")
		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			// Dégâts du joueur
			dmg := p.Attaque
			if p.BonusNext > 0 {
				dmg += p.BonusNext
			}
			fmt.Printf("🗡️ Tu infliges %d dégâts à %s !\n", dmg, m.Nom)
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

		// Attaque du boss si encore vivant
		if m.Pv > 0 {
			fmt.Printf("💥 %s t'attaque et inflige %d dégâts.\n", m.Nom, m.Attaque)
			p.Pv -= m.Attaque
			if p.Pv < 0 {
				p.Pv = 0
			}
		}
		p.Tour++
	}

	// Fin du combat
	if p.Pv <= 0 {
		fmt.Println("💀 Tu as été vaincu...")
	} else {
		fmt.Printf("🎉 Tu as vaincu %s !\n", m.Nom)
		bossProgression[bossIndex] = true

		// Récompense en argent
		p.Argent += reward
		fmt.Printf("💰 Tu as gagné %d£ !\n", reward)

		// Récompense en XP
		gainedXP := reward / 2
		p.XP += gainedXP
		fmt.Printf("⭐ +%d XP (total %d/%d)\n", gainedXP, p.XP, p.NextXP)
		if p.XP >= p.NextXP {
			p.Level++
			p.XP = 0
			p.NextXP += 50
			p.PvMax += 20
			p.Attaque += 5
			p.Pv = p.PvMax
			fmt.Printf("⬆️ Niveau %d atteint ! Stats augmentées.\n", p.Level)
		}

		// Cas spécial : Boss 3
		if m.Nom == "Les deux frères Kabyle" {
			fmt.Println("🎁 Les frères Kabyle t'offrent leur Flûte Gasba !")
			p.AddItem("Flûte Gasba")
		}

		// Cas spécial : Boss 4
		if m.Nom == "Lucas et les Archi" {
			fmt.Println("📢 Lucas te dit : 'Je te donne mon équipe d'archi (+10 dégâts au prochain combat)'")
			p.BonusNext = 10
			fmt.Println("Deux portes apparaissent :")
			fmt.Println("201 - Épée du seigneur d'architecte (175 dégâts bonus permanents)")
			fmt.Println("202 - Papier avec numéro (inutile)")
			fmt.Print("Choix : ")
			var porte int
			fmt.Scan(&porte)
			if porte == 201 {
				fmt.Println("⚔️ Tu obtiens l'Épée du seigneur d'architecte !")
				p.Attaque += 175
			} else {
				fmt.Println("📜 Tu obtiens un papier inutile avec un numéro écrit dessus.")
				p.AddItem("Papier avec numéro d'architecte")
			}
		}
	}
}

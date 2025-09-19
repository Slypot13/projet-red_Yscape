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
		// Si le dernier boss a été vaincu
		if bossProgression[4] {
			fmt.Println("\n🎉 Félicitations !")
			fmt.Println("Vous avez vaincu le Marocain Start-Up, le dernier obstacle de l'école.")
			fmt.Println("Le système est maintenant débugué, et vous avez le diplôme en main.")
			fmt.Println("Vous êtes prêt à affronter de nouveaux défis, à débuguer le monde entier !")
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
			if bossProgression[0] {
				fmt.Println("✅ Boss déjà vaincu.")
			} else {
				startCombat(p, Monster{"Classe Infobugé", 120, 120, 10}, 0, 50, "🌌 Tu gagnes +1000 d'Aura !")
			}
		case 2:
			if !bossProgression[0] {
				fmt.Println("❌ Tu dois d'abord vaincre le boss précédent.")
			} else if bossProgression[1] {
				fmt.Println("✅ Boss déjà vaincu.")
			} else {
				startCombat(p, Monster{"Le Contrôleur RATP", 250, 250, 20}, 1, 100, "🌌 Tu gagnes +2000 d'Aura !")
			}
		case 3:
			if !bossProgression[1] {
				fmt.Println("❌ Tu dois d'abord vaincre le boss précédent.")
			} else if bossProgression[2] {
				fmt.Println("✅ Boss déjà vaincu.")
			} else {
				startCombat(p, Monster{"Les deux frères Kabyle", 350, 350, 30}, 2, 150, "🌌 Tu gagnes +3000 d'Aura !")
			}
		case 4:
			if !bossProgression[2] {
				fmt.Println("❌ Tu dois d'abord vaincre le boss précédent.")
			} else if bossProgression[3] {
				fmt.Println("✅ Boss déjà vaincu.")
			} else {
				startCombat(p, Monster{"Lucas et les Archi", 550, 550, 40}, 3, 200, "🌌 Tu gagnes +4000 d'Aura !")
			}
		case 5:
			if !bossProgression[3] {
				fmt.Println("❌ Tu dois d'abord vaincre le boss précédent.")
			} else if bossProgression[4] {
				fmt.Println("✅ Boss déjà vaincu.")
			} else {
				go playSong2()
				startCombat(p, Monster{"Marocain Start-Up", 1000, 1000, 50}, 4, 300, "🌌 Tu gagnes +9999999 d'Aura !Adib te remercie de l'avoir libéré, et fait une danse Fortnite avec toi, puis s'en va sans rien te dire")
			}
		case 0:
			return
		default:
			fmt.Println("❌ Choix invalide.")
		}
	}
}

// -------------------- Fonction de combat --------------------
func startCombat(p *Character, m Monster, bossIndex int, reward int, auraMessage string) {
	fmt.Printf("\n⚔️  Combat contre %s ! (%d PV, %d ATK)\n", m.Nom, m.Pv, m.Attaque)

	p.Tour = 1
	if p.HasAsicsEffect {
		fmt.Println("👟 Effet Asics activé : Le boss est paralysé ce tour.")
	} else {
		p.Tour = 0
	}

	enemyAsleepTurns := 0
	playerAsleep := false

	for p.Pv > 0 && m.Pv > 0 {
		// Les frères Kabyle ont une chance d'endormir
		if m.Nom == "Les deux frères Kabyle" && rand.Intn(4) == 1 {
			fmt.Println("🎶 Les frères utilisent leur Flûte Gasba sur toi ! Tu es endormi pendant un tour...")
			playerAsleep = true
		}

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
			if playerAsleep {
				fmt.Println("😴 Tu es endormi et ne peux pas attaquer ce tour.")
				playerAsleep = false
			} else {
				dmg := p.Attaque
				if p.Nom == "Le Trackeur" && rand.Intn(3) == 1 {
					dmg *= 2
					fmt.Println("⚡ KAMEAMEAMEAMEA ! Attaque doublée !")
				}
				if p.Nom == "L'Ingénieur" {
					dmg += 10
				}
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
				fmt.Printf("🗡️  Tu infliges %d dégâts à %s !\n", dmg, m.Nom)
				m.Pv -= dmg
				if m.Pv < 0 {
					m.Pv = 0
				}
			}

		case 2:
			p.AccessInventoryMenuCombat(&m, &enemyAsleepTurns)
			continue
		case 3:
			fmt.Println("🚪 Tu as abandonné le combat.")
			return
		default:
			fmt.Println("❌ Choix invalide.")
			continue
		}

		if p.TempBoostTurns > 0 {
			p.TempBoostTurns--
			if p.TempBoostTurns == 0 {
				p.Attaque -= p.TempAttackBoost
				p.TempAttackBoost = 0
				fmt.Println("🧪 Fin de l'effet du Coca.")
			}
		}

		if enemyAsleepTurns > 0 {
			fmt.Println("😴 L'ennemi est endormi et ne peut pas attaquer ce tour.")
			enemyAsleepTurns--
		} else if p.HasAsicsEffect && p.Tour == 0 {
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

	// ---------- Fin du combat ----------
	if p.Pv <= 0 {
		fmt.Println("💀 Tu as été vaincu...")
	} else {
		fmt.Printf("🎉 Tu as vaincu %s !\n", m.Nom)
		bossProgression[bossIndex] = true

		// Récompense argent
		p.Argent += reward
		fmt.Printf("💰 Tu as gagné %d£ !\n", reward)

		// Récompense Aura (juste un message narratif)
		fmt.Println(auraMessage)

		// Boss 2 : objet spécial
		if m.Nom == "Le Contrôleur RATP" {
			if !p.HasItem("Flow du Contrôleur RATP") {
				p.AddItem("Flow du Contrôleur RATP")
				fmt.Println("🎁 Tu as obtenu le Flow du Contrôleur RATP !")
			}
		}

		// Boss 3 : Flûte Gasba
		if m.Nom == "Les deux frères Kabyle" {
			fmt.Println("🎁 Les frères Kabyle t'offrent leur Flûte Gasba !")
			p.AddItem("Flûte Gasba")
		}

		// Boss 4 : Bonus + portes
		if m.Nom == "Lucas et les Archi" {
			fmt.Println("📢 Lucas te dit : 'Je te donne mon équipe d'archi, ils t'aideront au prochain combat (+10 dégâts)'")

			fmt.Println("Deux portes apparaissent devant toi :")
			fmt.Println("201 ")
			fmt.Println("202 ")
			fmt.Print("Choix (201 ou 202) : ")
			var porte int
			fmt.Scan(&porte)

			if porte == 201 {
				fmt.Println("⚔️ Tu obtiens l'Épée du seigneur de la meilleur architecte ! (+175 ATK permanent)")
				p.Attaque += 175
				p.AddItem("Épée du seigneur d'architecte")
			} else {
				fmt.Println("📜 Tu obtiens un papier inutile... mais stylé.")
				p.AddItem("Papier avec le numéro de la plus belle architecte de France.")
			}
		}
	}
}

// -------------------- Inventaire en combat --------------------
func (p *Character) AccessInventoryMenuCombat(m *Monster, enemyAsleepTurns *int) {
	fmt.Println()
	fmt.Print("\033[1;33m")
	fmt.Println("╭─━━━━━─╯  Inventaire (combat) ╰─━━━━━─╮")
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
	fmt.Print("Choisis un objet (numéro) : ")
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
		fmt.Println("⚠️ Pas utilisable en combat.")
	case "Bouteille de Kambucha alcoolisé à 2%":
		p.UseKambucha()
	case "Casquette Gucci":
		p.EquipFromInventory("Casquette Gucci")
	case "Asics Kayano":
		p.EquipFromInventory("Asics Kayano")
	case "flow du contrôleur RATP":
		fmt.Println("Ceci est un composant spécial pour le Forgeron.")
	case "Flûte Gasba":
		p.UseFluteGasba(enemyAsleepTurns)
	default:
		fmt.Println("Objet non reconnu.")
	}
}

func (p *Character) UseFluteGasba(enemyAsleepTurns *int) {
	if !p.HasItem("Flûte Gasba") {
		fmt.Println("⚠️ Tu n'as pas la Flûte Gasba.")
		return
	}
	if *enemyAsleepTurns > 0 {
		fmt.Println("L'ennemi dort déjà.")
		return
	}
	*enemyAsleepTurns = 1
	fmt.Println("🎶 Tu joues de la Flûte Gasba — l'ennemi s'endort pour 1 tour !")
}

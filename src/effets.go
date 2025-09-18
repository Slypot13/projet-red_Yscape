package main

import (
	"fmt"
	"time"
)

// -------------------- Effets / Utilisation objets --------------------

func (p *Character) UseRedBull() {
	if !p.RemoveInventory("RedBull", 1) {
		fmt.Println("❌ Tu n'as pas de RedBull.")
		return
	}
	heal := 50
	if p.Pv+heal > p.PvMax {
		p.Pv = p.PvMax
	} else {
		p.Pv += heal
	}
	fmt.Printf("💖 RedBull consommé, la tie bien ! PV : %d/%d\n", p.Pv, p.PvMax)
}

func (p *Character) UseCoca() {
	if !p.RemoveInventory("Coca bien frais Chakal", 1) {
		fmt.Println("❌ Tu n'as pas de Coca bien frais Chakal.")
		return
	}
	bonus := int(float64(p.Attaque) * 0.10)
	if bonus < 1 {
		bonus = 1
	}
	p.TempAttackBoost += bonus
	p.Attaque += bonus
	p.TempBoostTurns = 1 // valable 1 tour de combat (sera décrémenté dans boucle de combat)
	fmt.Printf("🥤 Coca utilisé ! Attaque boostée de +%d pour 1 tour (Attaque = %d)\n", bonus, p.Attaque)
}

func (p *Character) UseKambucha() {
	if !p.RemoveInventory("Bouteille de Kambucha alcoolisé à 2%", 1) {
		fmt.Println("❌ Tu n'as pas de Kambucha.")
		return
	}
	heal := 30
	if p.Pv+heal > p.PvMax {
		p.Pv = p.PvMax
	} else {
		p.Pv += heal
	}
	fmt.Printf("🍹 Kambucha utilisé ! PV : %d/%d\n", p.Pv, p.PvMax)
}

func (p *Character) PoisonPot() {
	// applique un effet poison sur le joueur (utilisé pour test)
	fmt.Println("☠️ Poison activé...")
	for sec := 1; sec <= 3; sec++ {
		time.Sleep(1 * time.Second)
		p.Pv -= 10
		if p.Pv < 0 {
			p.Pv = 0
		}
		fmt.Printf("💀 Poison - %ds : PV %d/%d\n", sec, p.Pv, p.PvMax)
	}
}

// equiper un équipement (depuis inventaire) : retire de l'inventaire et update Equipement
func (p *Character) EquipFromInventory(nom string) {
	switch nom {
	case "Casquette Gucci":
		if !p.RemoveInventory(nom, 1) {
			fmt.Println("❌ Tu n'as pas de Casquette Gucci dans l'inventaire.")
			return
		}
		if p.Equipement.Tete != "" {
			fmt.Println("❌ Tu as déjà un équipement en tête (retire-le d'abord).")
			return
		}
		p.Equipement.Tete = nom
		p.CasquetteDelay = 3 // s'activera après 3 tours de combat
		p.CasquetteActive = false
		fmt.Println("✅ Casquette Gucci équipée ! Effet : +20% d'attaque à partir du 3e tour.")
	case "Asics Kayano":
		if !p.RemoveInventory(nom, 1) {
			fmt.Println("❌ Tu n'as pas d'Asics Kayano dans l'inventaire.")
			return
		}
		if p.Equipement.Pieds != "" {
			fmt.Println("❌ Tu as déjà un équipement aux pieds.")
			return
		}
		p.Equipement.Pieds = nom
		p.HasAsicsEffect = true
		fmt.Println("✅ Asics Kayano équipées aux pieds ! Effet : bloque l’ennemi 1 tour (à déclencher en combat).")
	default:
		fmt.Println("❌ Cet objet n'est pas équipable.")
	}
}

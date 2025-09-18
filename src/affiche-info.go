package main

import "fmt"

// -------------------- Affichage infos --------------------

func (p Character) DisplayInfo() {
	fmt.Println("\n\033[1;35m☆━━━━━━☆ Information du Personnage ☆━━━━━━☆\033[0m")
	fmt.Printf("\t- Nom       : %s\n", p.Nom)
	fmt.Printf("\t- Classe    : %s\n", p.Classe)
	fmt.Printf("\t- Niveau    : %d\n", p.Niveau)
	fmt.Printf("\t- PV        : %d/%d\n", p.Pv, p.PvMax)
	fmt.Printf("\t- Attaque   : %d\n", p.Attaque)
	if p.TempAttackBoost != 0 {
		fmt.Printf("\t  (Boost temporaire actuel : +%d, tours restants : %d)\n", p.TempAttackBoost, p.TempBoostTurns)
	}
	if p.CasquetteActive {
		fmt.Println("\t  (Casquette Gucci active : +20% attaque)")
	}
	fmt.Printf("\t- Argent    : %d£\n", p.Argent)
	fmt.Printf("\t- Pièces    : %d\n", p.Pieces)
	fmt.Printf("\t- Capacité  : %s\n", p.Capacite)
	fmt.Printf("\t- Inventaire: %d/%d\n", len(p.Inventaire), p.MaxInv)
	fmt.Printf("\t- Équipement: Tête=%s, Pieds=%s\n", p.Equipement.Tete, p.Equipement.Pieds)
	if len(p.Inventaire) == 0 {
		fmt.Println("\t   (Vide)")
	} else {
		for _, it := range p.Inventaire {
			fmt.Printf("\t   - %s x%d\n", it.Nom, it.Quantite)
		}
	}
}

// ajoute un item à l'inventaire (respecte la capacité), retourne true si ajouté
func (p *Character) AddInventory(nom string, quantite int) bool {
	for i := range p.Inventaire {
		if p.Inventaire[i].Nom == nom {
			p.Inventaire[i].Quantite += quantite
			return true
		}
	}
	if len(p.Inventaire) >= p.MaxInv {
		return false
	}
	p.Inventaire = append(p.Inventaire, Item{Nom: nom, Quantite: quantite})
	return true
}

// retire un item, retourne true si réussi
func (p *Character) RemoveInventory(nom string, quantite int) bool {
	for i := range p.Inventaire {
		if p.Inventaire[i].Nom == nom {
			if p.Inventaire[i].Quantite >= quantite {
				p.Inventaire[i].Quantite -= quantite
				if p.Inventaire[i].Quantite == 0 {
					p.Inventaire = append(p.Inventaire[:i], p.Inventaire[i+1:]...)
				}
				return true
			}
			return false
		}
	}
	return false
}

func (p *Character) HasItem(nom string) bool {
	for _, it := range p.Inventaire {
		if it.Nom == nom {
			return true
		}
	}
	return false
}

package character

import (
	"fmt"
	"yscape-game/src/items"
)

// Affiche l'inventaire
func (p Character) AccessInventory() {
	fmt.Println("=== Inventaire du personnage ===")
	if len(p.Inventaire) == 0 {
		fmt.Println("\tInventaire vide")
		return
	}
	for _, item := range p.Inventaire {
		fmt.Printf("\t- %s x%d\n", item.Nom, item.Quantite)
	}
}

// Ajoute un item
func (p *Character) AddInventory(nom string, quantite int) {
	for i := range p.Inventaire {
		if p.Inventaire[i].Nom == nom {
			p.Inventaire[i].Quantite += quantite
			return
		}
	}
	p.Inventaire = append(p.Inventaire, items.Item{Nom: nom, Quantite: quantite})
}

// Retire un item
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
			break
		}
	}
	return false
}

// Vérifie si le personnage est mort
func (p Character) IsDead() bool {
	return p.Pv <= 0
}

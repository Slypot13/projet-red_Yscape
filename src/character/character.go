package character

import (
	"fmt"
	"yscape-game/src/items"
)

type Character struct {
	Nom        string
	Classe     string
	Niveau     int
	Pv         int
	PvMax      int
	Inventaire []items.Item
}

func (p *Character) InitCharacter() {
	*p = Character{
		Nom:    "cyril",
		Classe: "Elfe",
		Niveau: 1,
		Pv:     40,
		PvMax:  100,
		Inventaire: []items.Item{
			{"Item", 1},
			{"Potion de vie", 2},
		},
	}
}

func (p Character) DisplayInfo() {
	fmt.Println("=== Information du personnage ===")
	fmt.Printf("\t - Nom : %s\n", p.Nom)
	fmt.Printf("\t - Classe : %s\n", p.Classe)
	fmt.Printf("\t - Niveau : %d\n", p.Niveau)
	fmt.Printf("\t - Pv : %d\n", p.Pv)
	fmt.Printf("\t - PvMax : %d\n", p.PvMax)
}

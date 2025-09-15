package character

import (
	"fmt"
	"time"
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

func (p *Character) PoisonPot() {
	fmt.Println("Potion de poison utilisée !")

	for i := 1; i <= 3; i++ {
		time.Sleep(1 * time.Second)
		p.Pv -= 10
		if p.Pv < 0 {
			p.Pv = 0
		}
		fmt.Printf("⚠️  Poison - Dégâts %d/3 : %d PV / %d PV Max\n", i, p.Pv, p.PvMax)
		if p.IsDead() {
			fmt.Println("💀 Vous êtes mort à cause du poison !")
			return
		}
	}
}

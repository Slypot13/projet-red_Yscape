package main

import (
	"fmt"
	"time"
)

// -------------------- Items --------------------

type Item struct {
	Nom      string
	Quantite int
}

// -------------------- Character --------------------

type Character struct {
	Nom        string
	Classe     string
	Niveau     int
	Pv         int
	PvMax      int
	Inventaire []Item
}

// Affiche le texte en mode machine à écrire, en gras cyan
func typewriterPrint(text string, delay time.Duration) {
	startStyle := "\033[1;36m" // gras + cyan
	resetStyle := "\033[0m"
	fmt.Print(startStyle)
	for _, c := range text {
		fmt.Printf("%c", c)
		time.Sleep(delay)
	}
	fmt.Println(resetStyle)
}

// Initialise le personnage
func (p *Character) InitCharacter() {
	*p = Character{
		Nom:    "cyril",
		Classe: "Elfe",
		Niveau: 1,
		Pv:     40,
		PvMax:  100,
		Inventaire: []Item{
			{Nom: "Potion magique", Quantite: 1},
			{Nom: "Potion de vie", Quantite: 2},
		},
	}
}

// Affiche les infos du personnage
func (p Character) DisplayInfo() {
	fmt.Println("=== Information du personnage ===")
	fmt.Printf("\t- Nom : %s\n", p.Nom)
	fmt.Printf("\t- Classe : %s\n", p.Classe)
	fmt.Printf("\t- Niveau : %d\n", p.Niveau)
	fmt.Printf("\t- Pv : %d\n", p.Pv)
	fmt.Printf("\t- PvMax : %d\n", p.PvMax)
}

// Vérifie si le personnage est mort
func (p *Character) IsDead() bool {
	return p.Pv <= 0
}

// Affiche l’inventaire
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

// Ajoute un item à l’inventaire
func (p *Character) AddInventory(nom string, quantite int) {
	for i := range p.Inventaire {
		if p.Inventaire[i].Nom == nom {
			p.Inventaire[i].Quantite += quantite
			return
		}
	}
	p.Inventaire = append(p.Inventaire, Item{Nom: nom, Quantite: quantite})
}

// Retire un item de l’inventaire
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

// Effet poison (diminue PV 3 fois)
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

// Utiliser potion de poison si disponible
func (p *Character) UsePoisonPotion() {
	for i := range p.Inventaire {
		if p.Inventaire[i].Nom == "Potion de poison" && p.Inventaire[i].Quantite > 0 {
			p.Inventaire[i].Quantite--
			if p.Inventaire[i].Quantite == 0 {
				p.Inventaire = append(p.Inventaire[:i], p.Inventaire[i+1:]...)
			}
			p.PoisonPot()
			return
		}
	}
	fmt.Println("⚠️  Vous n'avez pas de potion de poison.")
}

// Utiliser potion de vie
func (p *Character) TakePotion() {
	for i := range p.Inventaire {
		if p.Inventaire[i].Nom == "Potion de vie" && p.Inventaire[i].Quantite > 0 {
			p.Inventaire[i].Quantite--
			if p.Inventaire[i].Quantite == 0 {
				p.Inventaire = append(p.Inventaire[:i], p.Inventaire[i+1:]...)
			}
			p.Pv += 30
			if p.Pv > p.PvMax {
				p.Pv = p.PvMax
			}
			fmt.Printf("💖 Potion de vie utilisée ! PV restaurés : %d/%d\n", p.Pv, p.PvMax)
			return
		}
	}
	fmt.Println("⚠️  Vous n'avez pas de potion de vie.")
}

// Menu inventaire avec couleur
func (p *Character) MenuInventory() {
	const (
		yellowBold = "\033[1;33m"
		green      = "\033[32m"
		reset      = "\033[0m"
	)

	for {
		p.AccessInventory()
		fmt.Println()
		fmt.Print(yellowBold)
		fmt.Println("=== Menu inventaire ===")
		fmt.Print(reset)

		fmt.Print(green)
		fmt.Println("\t1 - Utiliser une potion de vie")
		fmt.Println("\t2 - Utiliser une potion de poison")
		fmt.Println("\t0 - Retour")
		fmt.Print(reset)

		fmt.Print("Choix : ")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			p.TakePotion()
		case 2:
			p.UsePoisonPotion()
		case 0:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

// Menu marchand avec couleur
func (p *Character) MerchantMenu() {
	const (
		yellowBold = "\033[1;33m"
		green      = "\033[32m"
		reset      = "\033[0m"
	)

	for {
		fmt.Println()
		fmt.Print(yellowBold)
		fmt.Println("=== Marchand ===")
		fmt.Print(reset)

		fmt.Print(green)
		fmt.Println("\t1 - Potion de vie (gratuit)")
		fmt.Println("\t2 - Potion de poison (gratuit)")
		fmt.Println("\t0 - Quitter")
		fmt.Print(reset)

		fmt.Print("Choix : ")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			p.AddInventory("Potion de vie", 1)
			fmt.Println("Vous avez reçu : Potion de vie x1")
		case 2:
			p.AddInventory("Potion de poison", 1)
			fmt.Println("Vous avez reçu : Potion de poison x1")
		case 0:
			return
		default:
			fmt.Println("Choix invalide")
		}
	}
}

// Menu principal du jeu avec titre jaune gras et choix verts
func (p *Character) StartMenu() {
	const (
		yellowBold = "\033[1;33m"
		green      = "\033[32m"
		reset      = "\033[0m"
	)
	for {
		fmt.Println()
		fmt.Print(yellowBold)
		fmt.Println("=== Menu Principal ===")
		fmt.Print(reset)

		fmt.Print(green)
		fmt.Println("\t1 - Afficher infos personnage")
		fmt.Println("\t2 - Inventaire")
		fmt.Println("\t3 - Marchand")
		fmt.Println("\t0 - Quitter")
		fmt.Print(reset)

		fmt.Print("Choix : ")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			p.DisplayInfo()
		case 2:
			p.MenuInventory()
		case 3:
			p.MerchantMenu()
		case 0:
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide")
		}
	}
}

// -------------------- Main --------------------

func main() {
	intro := `Tu arrives dans l'école Ynov, chaque étage correspond à un niveau, il y en a 5, pour pouvoir t'échapper.
Affronte les personnalités de l'école, les B1/B2/B3/M1/M2. Tu commences par l'accueil, où tu choisis ton personnage.
Durant ton voyage, tu devras faire des combats avec des Boss de différents niveaux.`

	typewriterPrint(intro, 50*time.Millisecond)

	p1 := Character{}
	p1.InitCharacter()
	p1.StartMenu()
}

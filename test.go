package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"
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
	Attaque    int
	Argent     int
	Inventaire []Item
	MaxInv     int
	Capacite   string
	Backpack   bool
}

// -------------------- Fonctions générales --------------------

func typewriterPrint(text string, delay time.Duration, color string) {
	reset := "\033[0m"
	fmt.Print(color)
	for _, c := range text {
		fmt.Printf("%c", c)
		time.Sleep(delay)
	}
	fmt.Println(reset)
}

// -------------------- Création des personnages --------------------

func CreateCharacters() []Character {
	return []Character{
		{
			Nom:        "Le Codeur",
			Classe:     "Élève",
			Niveau:     1,
			Pv:         75,
			PvMax:      200,
			Attaque:    75,
			Argent:     20,
			Inventaire: make([]Item, 0, 5),
			MaxInv:     5,
			Capacite:   "Objet aléatoire : Carte graphique (-10 PV) ou Chargeur iPhone 16 (-30 PV).",
		},
		{
			Nom:        "L'Ingénieur",
			Classe:     "Élève",
			Niveau:     1,
			Pv:         100,
			PvMax:      200,
			Attaque:    50,
			Argent:     20,
			Inventaire: make([]Item, 0, 5),
			MaxInv:     5,
			Capacite:   "Inflige +10 dégâts supplémentaires à chaque combat.",
		},
		{
			Nom:        "Le Trackeur",
			Classe:     "Élève",
			Niveau:     1,
			Pv:         50,
			PvMax:      200,
			Attaque:    100,
			Argent:     20,
			Inventaire: make([]Item, 0, 5),
			MaxInv:     5,
			Capacite:   "1 chance sur 3 de doubler son attaque (⚡ KAMEAMEAMEAMEA !).",
		},
	}
}

// -------------------- Création nom custom --------------------

func characterCreation(p Character) Character {
	var nom string
	for {
		fmt.Print("✏️  Choisis un nom pour ton personnage : ")
		fmt.Scan(&nom)

		valid := true
		for _, r := range nom {
			if !unicode.IsLetter(r) {
				valid = false
				break
			}
		}
		if !valid {
			fmt.Println("❌ Le nom ne doit contenir que des lettres.")
			continue
		}

		nom = strings.ToLower(nom)
		nom = strings.Title(nom)
		break
	}

	p.Nom = nom
	return p
}

// -------------------- Inventaire --------------------

func (p *Character) DisplayInventory() {
	fmt.Println("\n🎒 Inventaire :")
	if len(p.Inventaire) == 0 {
		fmt.Println("  (vide)")
	} else {
		for i, item := range p.Inventaire {
			fmt.Printf(" %d - %s (x%d)\n", i+1, item.Nom, item.Quantite)
		}
	}
	fmt.Printf("Capacité utilisée : %d/%d\n", len(p.Inventaire), p.MaxInv)
}

func (p *Character) AddItem(nom string) {
	if len(p.Inventaire) >= p.MaxInv {
		fmt.Println("❌ Ton inventaire est plein !")
		return
	}
	for i, item := range p.Inventaire {
		if item.Nom == nom {
			p.Inventaire[i].Quantite++
			fmt.Printf("✅ %s ajouté à l'inventaire (x%d)\n", nom, p.Inventaire[i].Quantite)
			return
		}
	}
	p.Inventaire = append(p.Inventaire, Item{Nom: nom, Quantite: 1})
	fmt.Printf("✅ %s ajouté à l'inventaire !\n", nom)
}

// -------------------- Affichage infos --------------------

func (p Character) DisplayInfo() {
	fmt.Println("\n☆━━━━━━☆ Information du Personnage ☆━━━━━━☆")
	fmt.Printf("\t- Nom : %s\n", p.Nom)
	fmt.Printf("\t- Classe : %s\n", p.Classe)
	fmt.Printf("\t- Niveau : %d\n", p.Niveau)
	fmt.Printf("\t- Pv : %d/%d\n", p.Pv, p.PvMax)
	fmt.Printf("\t- Attaque : %d\n", p.Attaque)
	fmt.Printf("\t- Argent : %d£\n", p.Argent)
	fmt.Printf("\t- Capacité : %s\n", p.Capacite)
	fmt.Printf("\t- Inventaire : %d/%d\n", len(p.Inventaire), p.MaxInv)
	if len(p.Inventaire) == 0 {
		fmt.Println("\t   (Vide)")
	}
}

// -------------------- Marchand --------------------

func (p *Character) MerchantMenu() {
	for {
		fmt.Println("\n🏪 Bienvenue au Marchand du Campus !")
		fmt.Println("\t1 - RedBull (10£)")
		fmt.Println("\t2 - Coca bien frais Chakal (15£)")
		fmt.Println("\t3 - Café dilué au Ciao Kambucha (10£)")
		fmt.Println("\t4 - Sac à dos perdu (30£) [permanent]")
		fmt.Println("\t0 - Quitter")
		fmt.Printf("💰 Votre argent : %d£\n", p.Argent)

		var choix int
		fmt.Print("Choix : ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			if p.Argent >= 10 {
				p.Argent -= 10
				p.AddItem("RedBull")
			} else {
				fmt.Println("❌ Pas assez d'argent...")
			}
		case 2:
			if p.Argent >= 15 {
				p.Argent -= 15
				p.AddItem("Coca bien frais Chakal")
			} else {
				fmt.Println("❌ Pas assez d'argent...")
			}
		case 3:
			if p.Argent >= 10 {
				p.Argent -= 10
				p.AddItem("Café dilué au Ciao Kambucha")
			} else {
				fmt.Println("❌ Pas assez d'argent...")
			}
		case 4:
			if p.Backpack {
				fmt.Println("🎒 Tu possèdes déjà le Sac à dos perdu.")
			} else if p.Argent >= 30 {
				p.Argent -= 30
				p.MaxInv = 10
				p.Backpack = true
				fmt.Println("✅ Tu as acheté le Sac à dos perdu ! Inventaire max = 10 objets.")
			} else {
				fmt.Println("❌ Pas assez d'argent...")
			}
		case 0:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

// -------------------- Choix personnage (horizontal) --------------------

func ChooseCharacter() Character {
	personnages := CreateCharacters()

	fmt.Println("\n━━━━━━━━━━━ Présentation des personnages ━━━━━━━━━━━\n")

	lines := make([][]string, 3)
	for i, p := range personnages {
		lines[i] = []string{
			fmt.Sprintf("╭───────────── %d ─────────────╮", i+1),
			fmt.Sprintf("│ %-28s │", p.Nom),
			fmt.Sprintf("│ Classe : %-19s │", p.Classe),
			fmt.Sprintf("│ Niveau : %-19d │", p.Niveau),
			fmt.Sprintf("│ PV : %-23d │", p.Pv),
			fmt.Sprintf("│ Attaque : %-18d │", p.Attaque),
			fmt.Sprintf("│ Capacité :                     │"),
			fmt.Sprintf("│ %-28s │", truncate(p.Capacite, 28)),
			"╰──────────────────────────────╯",
		}
	}

	for i := 0; i < len(lines[0]); i++ {
		fmt.Printf("%s   %s   %s\n", lines[0][i], lines[1][i], lines[2][i])
	}

	var choix int
	for {
		fmt.Print("\n➡️  Entre le numéro de ton personnage : ")
		fmt.Scan(&choix)
		if choix >= 1 && choix <= len(personnages) {
			break
		}
		fmt.Println("❌ Choix invalide.")
	}

	return characterCreation(personnages[choix-1])
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-3] + "..."
}

// -------------------- Menu Principal --------------------

func (p *Character) StartMenu() {
	for {
		fmt.Println("\n▁ ▂ ▄ ▅ ▆ ▇ █  Menu Principal  █ ▇ ▆ ▅ ▄ ▂ ▁")
		fmt.Println("\t1 - Afficher infos personnage")
		fmt.Println("\t2 - Inventaire")
		fmt.Println("\t3 - Marchand du Campus")
		fmt.Println("\t0 - Quitter")

		var choix int
		fmt.Print("Choix : ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			p.DisplayInfo()
		case 2:
			p.DisplayInventory()
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
	rand.Seed(time.Now().UnixNano())

	fmt.Println("\033[1;33m")
	fmt.Println(`
          ██████████████████████████████████████████
          █  []  []  []  []  []  []  []  []  []  []█
          █  []  []  []  []  []  []  []  []  []  []█
          █  []  []  []  []  []  []  []  []  []  []█
          █  []  []  []  []  []  []  []  []  []  []█
          █  []  []  []  []  []  []  []  []  []  []█
          █  []  []  []  []  []  []  []  []  []  []█
          ██████████████████████████████████████████
          ██████████████████████████████████████████
                 ┌──────────────┐
                 │      🚪      │
                 │     (👤)     │
                 └──────────────┘
	`)
	fmt.Println("\033[0m")

	intro := `L'école Ynov s'étend devant toi, six étages à franchir...
À l’accueil, deux étudiantes t’attendent : Marie et Lisa.
Elles te fixent avec un regard déterminé.

「 Marie 」: Enfin te voilà ! On n’a pas beaucoup de temps...
「 Lisa 」: Le campus est rempli de pièges et de boss terrifiants.
Tu devras gravir les 6 étages pour t’échapper.

Choisis ton personnage...`
	typewriterPrint(intro, 25*time.Millisecond, "\033[36m")

	player := ChooseCharacter()
	fmt.Println("\n✅ Tu as choisi ton héros !")
	player.DisplayInfo()

	player.StartMenu()
}

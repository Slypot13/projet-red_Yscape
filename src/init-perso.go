package main

import (
	"fmt"
	"time"
)

// -------------------- Initialisation personnages --------------------

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
			Pieces:     0,
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
			Pieces:     0,
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
			Pieces:     0,
			Inventaire: make([]Item, 0, 5),
			MaxInv:     5,
			Capacite:   "1 chance sur 3 de doubler son attaque (⚡ KAMEAMEAMEAMEA !).",
		},
	}
}

// -------------------- Présentation des persos côte à côte --------------------

func ChooseCharacter() Character {
	personnages := CreateCharacters()

	typewriterPrint("\n👩 Marie : Enfin ! Te voilà à l'accueil.", 40*time.Millisecond, "\033[1;35m")
	typewriterPrint("👩 Lisa : Choisis vite ton personnage avant qu'il ne soit trop tard !", 40*time.Millisecond, "\033[1;34m")

	fmt.Println("\n━━━━━━━━━━━ Présentation des personnages ━━━━━━━━━━━")
	// ligne de noms (quadrillage simple)
	fmt.Printf("+-----------------+-----------------+-----------------+\n")
	fmt.Printf("| %-15s | %-15s | %-15s |\n", personnages[0].Nom, personnages[1].Nom, personnages[2].Nom)
	fmt.Printf("| %-15s | %-15s | %-15s |\n", "Classe: "+personnages[0].Classe, "Classe: "+personnages[1].Classe, "Classe: "+personnages[2].Classe)
	fmt.Printf("| PV: %-10d | PV: %-10d | PV: %-10d |\n", personnages[0].Pv, personnages[1].Pv, personnages[2].Pv)
	fmt.Printf("| ATK: %-8d | ATK: %-8d | ATK: %-8d |\n", personnages[0].Attaque, personnages[1].Attaque, personnages[2].Attaque)
	fmt.Printf("+-----------------+-----------------+-----------------+\n")

	// Afficher un par un en détails (défilement)
	for i, perso := range personnages {
		typewriterPrint(fmt.Sprintf("\nDécouvrons le personnage [%d] : %s", i+1, perso.Nom), 30*time.Millisecond, "\033[35m")
		perso.DisplayInfo()
		fmt.Println("------------------------------------------------")
		time.Sleep(900 * time.Millisecond)
	}

	var choix int
	for {
		fmt.Print("➡️  Entre le numéro (1/2/3) de ton personnage : ")
		fmt.Scan(&choix)
		if choix >= 1 && choix <= len(personnages) {
			break
		}
		fmt.Println("❌ Choix invalide.")
	}

	// permet au joueur de choisir un nom
	return characterCreation(personnages[choix-1])
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-3] + "..."
}

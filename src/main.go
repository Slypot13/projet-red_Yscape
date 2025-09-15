package main

import (
	"fmt"
	"os"
)

// Tâche 1 : Création du personnage
// Définition d'une structure Character représentant un personnage
type Character struct {
	Nom        string
	Classe     string
	Niveau     int
	Pv         int
	PvMax      int
	Inventaire []Item
}

// Définition d'une structure Item représentant un objet
type Item struct {
	Nom      string
	Quantite int
}

// Tâche 2 : Initialisation du personnage
// Méthode pour initialiser un personnage avec des valeurs par défaut
func (player *Character) initCharacter() {
	*player = Character{
		Nom:    "cyril",
		Classe: "Elfe",
		Niveau: 1,
		Pv:     40,
		PvMax:  100,
		Inventaire: []Item{
			{"Item", 1},
			{"Potion de vie", 2},
		},
	}
}

// Tâche 3 : Affichage des informations du personnage
// Affiche les informations générales du personnage
func (player Character) displayInfo() {
	fmt.Println("=== Information du personnage ===")
	fmt.Printf("\t - Nom : %s\n", player.Nom)
	fmt.Printf("\t - Classe : %s\n", player.Classe)
	fmt.Printf("\t - Niveau : %d\n", player.Niveau)
	fmt.Printf("\t - Pv : %d\n", player.Pv)
	fmt.Printf("\t - PvMax : %d\n", player.PvMax)
}

// Tâche 4 : Accès à l’inventaire
// Affiche le contenu de l’inventaire
func (player Character) accessInventory() {
	fmt.Println("=== Inventaire du personnage ===")
	if len(player.Inventaire) == 0 {
		fmt.Println("\tInventaire vide")
	} else {
		for _, item := range player.Inventaire {
			fmt.Printf("\t- %s x %d\n", item.Nom, item.Quantite)
		}
	}
}

// Menu permettant d’interagir avec l’inventaire
func (player *Character) MenuInventory() {
	for true {
		player.accessInventory() // Affiche l’inventaire
		fmt.Println("=== Menu inventaire ===")
		fmt.Printf("\t1 - Utiliser une potion de vie\n")
		fmt.Printf("\t0 - Retour\n")
		fmt.Println("Sélectionner un choix (1 ou 0) :")
		var userChose int
		fmt.Scan(&userChose)

		switch userChose {
		case 1:
			player.takePot() // Utilisation d’une potion
		case 0:
			return
		default:
			fmt.Println("Erreur : choix non valide")
		}
	}
}

// Tâche 5 : Potion de vie
// Méthode pour utiliser une potion de vie si disponible
func (player *Character) takePot() {
	for index := range player.Inventaire {
		if player.Inventaire[index].Nom == "Potion de vie" && player.Inventaire[index].Quantite > 0 {
			// Soigne le personnage de 50 Pv
			player.Pv += 50
			if player.Pv > player.PvMax {
				player.Pv = player.PvMax
			}
			fmt.Println("Potion de vie utilisée (quantité -1)")
			fmt.Printf("Nouveau Pv : %d\n", player.Pv)

			// Réduit la quantité de potion
			player.Inventaire[index].Quantite -= 1
			// Supprime l’item si la quantité est 0
			if player.Inventaire[index].Quantite <= 0 {
				player.Inventaire = append(player.Inventaire[:index], player.Inventaire[index+1:]...)
			}
			return
		}
	}
	fmt.Println("Utilisation impossible : potion de vie manquante")
}

// Tâche 6 : Création du menu principal
// Menu principal pour interagir avec le personnage
func (player *Character) MainMenu() {
	for true {
		fmt.Println("=== Menu principal ===")
		fmt.Printf("\t1 - Afficher les informations du personnage\n")
		fmt.Printf("\t2 - Accéder au contenu de l’inventaire\n")
		fmt.Printf("\t0 - Quitter\n")
		fmt.Println("Sélectionner un choix (1,2 ou 0) :")
		var userChose int
		fmt.Scan(&userChose)

		switch userChose {
		case 1:
			player.displayInfo()
		case 2:
			player.MenuInventory()
		case 0:
			// Quitte le programme
			os.Exit(0)
			return
		default:
			fmt.Println("Erreur : choix non valide")
		}
	}
}

// Point d’entrée du programme
func main() {
	p1 := Character{}  // Création d’un personnage vide
	p1.initCharacter() // Initialisation avec des valeurs par défaut
	p1.MainMenu()      // Lancement du menu principal
}

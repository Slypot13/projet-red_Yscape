package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

// -------------------- Types --------------------

type Item struct {
	Nom      string
	Quantite int
}

type Equipment struct {
	Tete  string
	Pieds string
}

type Character struct {
	Nom             string
	Classe          string
	Niveau          int
	Pv              int
	PvMax           int
	Attaque         int
	Argent          int // argent «£»
	Pieces          int // pièces d'or (pour le forgeron)
	Inventaire      []Item
	MaxInv          int
	Capacite        string
	Backpack        bool
	Equipement      Equipment
	TempAttackBoost int  // montant ajouté temporairement (ex: Coca)
	TempBoostTurns  int  // tours restants pour le boost temporaire
	CasquetteDelay  int  // tours restants avant activation de la casquette
	CasquetteActive bool // si l'effet casquette est déjà appliqué
	HasAsicsEffect  bool // si Asics est équipé (effet à déclencher en combat)
	Tour            int
}

type Monster struct {
	Nom     string
	PvMax   int
	Pv      int
	Attaque int
}

// -------------------- Utilitaires Texte --------------------

func typewriterPrint(text string, delay time.Duration, color string) {
	reset := "\033[0m"
	fmt.Print(color)
	for _, c := range text {
		fmt.Printf("%c", c)
		time.Sleep(delay)
	}
	fmt.Println(reset)
}

func colorPrintln(color, s string) {
	reset := "\033[0m"
	fmt.Print(color)
	fmt.Println(s)
	fmt.Print(reset)
}

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

// -------------------- Building ASCII --------------------

func printCampusBig() {
	fmt.Println("\033[1;33m") // jaune gras
	fmt.Println(`
████████████████████████████████████████████████████████████
█  []  []  []  []  []  []  []  []  []  []  []  []  []  []  █
█  []  []  []  []  []  []  []  []  []  []  []  []  []  []  █
█  []  []  []  []  []  []  []  []  []  []  []  []  []  []  █
█  []  []  []  []  []  []  []  []  []  []  []  []  []  []  █
█  []  []  []  []  []  []  []  []  []  []  []  []  []  []  █
█  []  []  []  []  []  []  []  []  []  []  []  []  []  []  █
████████████████████████████████████████████████████████████
████████████████████████████████████████████████████████████
████████████████████████████████████████████████████████████
           ┌────────────────────────────┐
           │           [  DOOR  ]        │
           │          (élève attend)    │
           └────────────────────────────┘
`)
	fmt.Println("\033[0m")
}

// -------------------- Narration & Intro --------------------

func introduction() {
	printCampusBig()
	typewriterPrint("Tu as cinq étages à franchir, cinq niveaux pour t’échapper.", 30*time.Millisecond, "\033[36m")
	typewriterPrint("Les personnalités de l'école — B1, B2, B3, M1, M2 — t’attendent.", 30*time.Millisecond, "\033[36m")
	typewriterPrint("Tout commence à l'accueil, où tu choisis ton personnage...", 30*time.Millisecond, "\033[36m")
	typewriterPrint("Choisis-le bien et gare à toi !", 30*time.Millisecond, "\033[36m")
	typewriterPrint("Les 5 monstres que tu vas affronter sont d’anciens élèves bloqués dans le passé,", 30*time.Millisecond, "\033[36m")
	typewriterPrint("à cause de la faille spatio-temporelle créée lorsque Cyril et Bastien", 30*time.Millisecond, "\033[36m")
	typewriterPrint("ont fusionné leurs PC et déclenché une boucle à remonter le temps.", 30*time.Millisecond, "\033[36m")
	typewriterPrint("Il ne te reste que quelques heures pour récupérer ton Saint Diplôme à temps et sauver l'humanité.", 30*time.Millisecond, "\033[36m")

	// laisser la phrase finale EXACTE comme demandé
	typewriterPrint("Nous comptons sur toi jeune Skylanders.. Euhh éleve de Ynov !!!", 40*time.Millisecond, "\033[1;31m")
}

// -------------------- Gestion Nom perso --------------------

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
		if !valid || len(nom) == 0 {
			fmt.Println("❌ Le nom ne doit contenir que des lettres.")
			continue
		}
		nom = strings.ToLower(nom)
		nom = strings.Title(nom)
		p.Nom = nom
		break
	}
	return p
}

// -------------------- Affichage & Inventaire général --------------------

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
	fmt.Printf("💖 RedBull consommé ! PV : %d/%d\n", p.Pv, p.PvMax)
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

// -------------------- Menus : Marchand & Forgeron --------------------

func (p *Character) MerchantMenu() {
	for {
		fmt.Println()
		fmt.Print("\033[1;33m")
		fmt.Println("┏━━━━━━┛  Marchand du Campus  ┗━━━━━━┓")
		fmt.Print("\033[0m")

		fmt.Print("\033[32m")
		fmt.Println("\t1 - RedBull (+50 PV) - 10£")
		fmt.Println("\t2 - Coca bien frais Chakal (+10% attaque / 1 tour) - 15£")
		fmt.Println("\t3 - Café dilué au Ciao Kambucha (Poison -10PV/s 3s) - 20£")
		fmt.Println("\t4 - Sac à dos perdu (+5 places inventaire) - 30£")
		fmt.Println("\t0 - Retour")
		fmt.Print("\033[0m")

		fmt.Printf("💰 Argent : %d£\n", p.Argent)
		fmt.Print("Choix : ")
		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			if p.Argent < 10 {
				fmt.Println("❌ Pas assez d'argent pour acheter RedBull.")
			} else {
				if ok := p.AddInventory("RedBull", 1); ok {
					p.Argent -= 10
					fmt.Println("✅ Achat : RedBull (ajouté à l'inventaire).")
				} else {
					fmt.Println("❌ Inventaire plein, impossible d'acheter.")
				}
			}
		case 2:
			if p.Argent < 15 {
				fmt.Println("❌ Pas assez d'argent pour acheter le Coca.")
			} else {
				if ok := p.AddInventory("Coca bien frais Chakal", 1); ok {
					p.Argent -= 15
					fmt.Println("✅ Achat : Coca bien frais Chakal (ajouté à l'inventaire).")
				} else {
					fmt.Println("❌ Inventaire plein, impossible d'acheter.")
				}
			}
		case 3:
			if p.Argent < 20 {
				fmt.Println("❌ Pas assez d'argent pour acheter le Café dilué.")
			} else {
				if ok := p.AddInventory("Café dilué au Ciao Kambucha", 1); ok {
					p.Argent -= 20
					fmt.Println("✅ Achat : Café dilué au Ciao Kambucha (ajouté à l'inventaire).")
				} else {
					fmt.Println("❌ Inventaire plein, impossible d'acheter.")
				}
			}
		case 4:
			if p.Backpack {
				fmt.Println("🎒 Tu possèdes déjà le Sac à dos perdu.")
			} else if p.Argent < 30 {
				fmt.Println("❌ Pas assez d'argent pour acheter le Sac à dos perdu.")
			} else {
				p.Argent -= 30
				p.Backpack = true
				p.MaxInv = 10
				fmt.Println("✅ Achat : Sac à dos perdu. Inventaire étendu à 10 emplacements.")
			}
		case 0:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func (p *Character) BlacksmithMenu() {
	for {
		fmt.Println()
		fmt.Print("\033[1;36m")
		fmt.Println("┏━━━━━━┛  Forgeron  ┗━━━━━━┓")
		fmt.Print("\033[0m")

		fmt.Print("\033[32m")
		fmt.Println("\t1 - Casquette Gucci (60 pièces) [+20% attaque à partir du 3e tour]")
		fmt.Println("\t2 - Asics Kayano (60 pièces) [empêche le monstre de jouer 1 tour]")
		fmt.Println("\t0 - Retour")
		fmt.Print("\033[0m")

		fmt.Printf("⚠️ Pour acheter et équiper, tu dois posséder l'objet 'flow du contrôleur RATP'.\n")
		fmt.Printf("💰 Pièces : %d\n", p.Pieces)
		fmt.Print("Choix : ")
		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			if p.Pieces < 60 {
				fmt.Println("❌ Pas assez de pièces pour la Casquette Gucci.")
				continue
			}
			if !p.HasItem("flow du contrôleur RATP") {
				fmt.Println("❌ Il te manque 'flow du contrôleur RATP' pour acheter/équiper cet équipement.")
				continue
			}
			if ok := p.AddInventory("Casquette Gucci", 1); !ok {
				fmt.Println("❌ Inventaire plein, impossible d'ajouter la Casquette Gucci.")
				continue
			}
			p.Pieces -= 60
			fmt.Println("✅ Tu as fabriqué la Casquette Gucci (ajoutée à l'inventaire).")
		case 2:
			if p.Pieces < 60 {
				fmt.Println("❌ Pas assez de pièces pour les Asics Kayano.")
				continue
			}
			if !p.HasItem("flow du contrôleur RATP") {
				fmt.Println("❌ Il te manque 'flow du contrôleur RATP' pour acheter/équiper cet équipement.")
				continue
			}
			if ok := p.AddInventory("Asics Kayano", 1); !ok {
				fmt.Println("❌ Inventaire plein, impossible d'ajouter Asics Kayano.")
				continue
			}
			p.Pieces -= 60
			fmt.Println("✅ Tu as fabriqué les Asics Kayano (ajoutées à l'inventaire).")
		case 0:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

// -------------------- Inventaire : affichage et utilisation --------------------

func (p *Character) AccessInventoryMenu() {
	fmt.Println()
	fmt.Print("\033[1;33m")
	fmt.Println("╭─━━━━━─╯  Inventaire  ╰─━━━━━─╮")
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
	fmt.Print("Choisis un objet (numéro) pour l'utiliser / équiper : ")
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
		// on applique poison sur le joueur (ou sur ennemi si en combat, ici hors combat on montre l'effet)
		fmt.Println("⚠️ Le Café dilué est toxique à consommer hors combat. Ne l'utilise pas maintenant.")
	case "Bouteille de Kambucha alcoolisé à 2%":
		p.UseKambucha()
	case "Casquette Gucci":
		p.EquipFromInventory("Casquette Gucci")
	case "Asics Kayano":
		p.EquipFromInventory("Asics Kayano")
	case "flow du contrôleur RATP":
		fmt.Println("Ceci est un composant spécial pour le Forgeron.")
	default:
		fmt.Println("Objet non reconnu / utilisation non implémentée.")
	}
}

// -------------------- Combat : Goblin (ClasseInfobugé) --------------------

func InitGoblin() Monster {
	return Monster{Nom: "ClasseInfobugé", PvMax: 115, Pv: 115, Attaque: 20}
}

func GoblinPattern(monstre *Monster, player *Character, tour int) {
	attack := monstre.Attaque
	if tour%3 == 0 {
		attack = int(float64(monstre.Attaque) * 1.2)
	}
	cris := []string{"Lancé de carte graphique", "Lancé de souris", "Lancé de clavier"}
	cri := cris[rand.Intn(len(cris))]
	fmt.Printf("\n%s crie \"%s\" !\n", monstre.Nom, cri)

	// si joueur a Asics effect (empêche une action), handle in CharacterTurn when player triggers it.
	player.Pv -= attack
	if player.Pv < 0 {
		player.Pv = 0
	}
	fmt.Printf("%s inflige à %s %d dégâts !\n", monstre.Nom, player.Nom, attack)
	fmt.Printf("%s PV : %d/%d\n", player.Nom, player.Pv, player.PvMax)
}

func CharacterTurn(monstre *Monster, player *Character, tour *int, monsterSkipped *bool) {
	fmt.Println("\n--- Ton tour ---")
	fmt.Println("1 - Attaquer")
	fmt.Println("2 - Inventaire")
	fmt.Println("3 - Utiliser Asics (si équipé) [bloque l'ennemi 1 tour]")
	fmt.Print("Choix : ")
	var choix int
	fmt.Scan(&choix)
	switch choix {
	case 1:
		// Attaque basique - tenir compte si casquette active => l'effet casquette s'applique via CasquetteActive field
		damage := player.Attaque
		fmt.Printf("%s utilise Attaque basique et inflige %d dégâts !\n", player.Nom, damage)
		monstre.Pv -= damage
		if monstre.Pv < 0 {
			monstre.Pv = 0
		}
		fmt.Printf("%s PV : %d/%d\n", monstre.Nom, monstre.Pv, monstre.PvMax)
	case 2:
		// Accès à l'inventaire (utilisation d'objets en combat)
		player.AccessInventoryCombat()
	case 3:
		// utiliser Asics si équipé
		if player.Equipement.Pieds == "Asics Kayano" && player.HasAsicsEffect {
			*monsterSkipped = true
			player.HasAsicsEffect = false // effet consommé (ou garder selon design)
			fmt.Println("👟 Tu actives Asics Kayano : le monstre est bloqué pour 1 tour !")
		} else {
			fmt.Println("❌ Tu n'as pas Asics Kayano équipées ou l'effet n'est pas disponible.")
		}
	default:
		fmt.Println("Choix invalide, tu perds ton action.")
	}
	// fin du turn du joueur, on décrémente les tours boost si actifs (mais revert après le tour du monstre)
	// CasquetteDelay est gérée dans boucle de combat
	_ = tour
}

func (p *Character) AccessInventoryCombat() {
	if len(p.Inventaire) == 0 {
		fmt.Println("Inventaire vide.")
		return
	}
	fmt.Println("Inventaire :")
	for i, it := range p.Inventaire {
		fmt.Printf("%d - %s x%d\n", i+1, it.Nom, it.Quantite)
	}
	fmt.Print("Choisis un objet (numéro) : ")
	var choix int
	fmt.Scan(&choix)
	if choix < 1 || choix > len(p.Inventaire) {
		fmt.Println("Choix invalide.")
		return
	}
	switch p.Inventaire[choix-1].Nom {
	case "RedBull":
		p.UseRedBull()
	case "Bouteille de Kambucha alcoolisé à 2%":
		p.UseKambucha()
	case "Coca bien frais Chakal":
		p.UseCoca()
	case "Café dilué au Ciao Kambucha":
		fmt.Println("☠️ Tu lances le Café dilué sur l'ennemi (poison) ! (implémentation simplifiée)")
		// appliquer poison de test sur le monstre via un petit loop: pour simplifier, on applique direct damage here
		// but since monstre passed by pointer in caller, we cannot access it here; in combat we handle using AccessInventoryCombat from CharacterTurn when relevant.
	default:
		fmt.Println("Objet non utilisable en combat.")
	}
}

// TrainingFight : le combat d'entraînement contre la ClasseInfobugé
func TrainingFight(player *Character) {
	monstre := InitGoblin()
	tour := 1
	monsterSkipped := false
	fmt.Println("\n⚔️ Début du combat d'entraînement contre", monstre.Nom, "!")
	for player.Pv > 0 && monstre.Pv > 0 {
		fmt.Printf("\n======== Tour %d ========\n", tour)

		// Si casquette équipée et pas encore active, décrémente et active si nécessaire
		if player.Equipement.Tete == "Casquette Gucci" && !player.CasquetteActive {
			if player.CasquetteDelay > 0 {
				fmt.Printf("(Casquette Gucci : %d tours avant activation)\n", player.CasquetteDelay)
				player.CasquetteDelay--
				if player.CasquetteDelay == 0 {
					// activer l'effet : +20% attaque
					bonus := int(float64(player.Attaque) * 0.20)
					if bonus < 1 {
						bonus = 1
					}
					player.Attaque += bonus
					player.CasquetteActive = true
					fmt.Printf("🧢 Casquette Gucci s'active ! Attaque augmentée de +%d (Attaque = %d)\n", bonus, player.Attaque)
				}
			}
		}

		// Tour du joueur
		CharacterTurn(&monstre, player, &tour, &monsterSkipped)
		if monstre.Pv <= 0 {
			break
		}

		// Tour du monstre (sauf si joueur a appliqué Asics -> monsterSkipped true)
		if monsterSkipped {
			fmt.Println("\nLe monstre est bloqué ce tour, il ne peut pas attaquer.")
			monsterSkipped = false
		} else {
			GoblinPattern(&monstre, player, tour)
		}

		// Après l'attaque du monstre, gérer la durée du boost Coca
		if player.TempBoostTurns > 0 {
			player.TempBoostTurns--
			if player.TempBoostTurns == 0 {
				// revert boost
				player.Attaque -= player.TempAttackBoost
				fmt.Printf("🥤 Effet Coca terminé. Attaque revenue à %d\n", player.Attaque)
				player.TempAttackBoost = 0
			}
		}

		// vérifier si joueur mort
		if player.Pv <= 0 {
			break
		}
		tour++
	}

	// Résultat du combat
	if player.Pv <= 0 {
		fmt.Println("\n❌ Tu es vaincu... Retour au menu principal.")
		// optionnel : restaurer PV partiellement ou renvoyer au menu
	} else {
		fmt.Println("\n🎉 Yesss mon gaté c'est gagné ! EZ la classe")
		player.Pieces += 15
		added := player.AddInventory("Bouteille de Kambucha alcoolisé à 2%", 1)
		if added {
			fmt.Println("Récompense : +15 pièces et 1x Bouteille de Kambucha alcoolisé à 2% ajouté à l'inventaire (soigne 30PV).")
		} else {
			// si inventaire plein, on donne argent à la place
			player.Argent += 0
			fmt.Println("Ton inventaire était plein : la récompense 'Bouteille de Kambucha' n'a pas pu être ajoutée.")
			fmt.Println("Tu as quand même reçu +15 pièces.")
		}
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

// -------------------- Menu Principal --------------------

func (p *Character) StartMenu() {
	const (
		yellowBold = "\033[1;33m"
		green      = "\033[32m"
		reset      = "\033[0m"
	)
	for {
		fmt.Println()
		fmt.Print(yellowBold)
		fmt.Println("▁ ▂ ▄ ▅ ▆ ▇ █  Menu Principal  █ ▇ ▆ ▅ ▄ ▂ ▁")
		fmt.Print(reset)

		fmt.Print(green)
		fmt.Println("\t1 - Afficher infos personnage")
		fmt.Println("\t2 - Inventaire")
		fmt.Println("\t3 - Marchand du Campus")
		fmt.Println("\t4 - Forgeron")
		fmt.Println("\t5 - Combat d'entraînement (ClasseInfobugé)")
		fmt.Println("\t0 - Quitter")
		fmt.Print(reset)

		fmt.Print("Choix : ")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			p.DisplayInfo()
		case 2:
			p.AccessInventoryMenu()
		case 3:
			p.MerchantMenu()
		case 4:
			p.BlacksmithMenu()
		case 5:
			TrainingFight(p)
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
	introduction()

	player := ChooseCharacter()
	fmt.Println("\n✅ Tu as choisi ton héros !")
	player.DisplayInfo()

	player.StartMenu()
}

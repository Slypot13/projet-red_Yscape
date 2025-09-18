package main 
// Barre de vie à afficher dans le terminal
func DisplayHPBar(current, max int, width int) string {
	percentage := float64(current) / float64(max)
	filled := int(percentage * float64(width))
	empty := width - filled

	bar := ""
	for i := 0; i < filled; i++ {
		bar += "█"
	}
	for i := 0; i < empty; i++ {
		bar += "░"
	}

	return bar
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
		damage := player.Attaque
		fmt.Printf("%s utilise Attaque basique et inflige %d dégâts !\n", player.Nom, damage)
		monstre.Pv -= damage
		if monstre.Pv < 0 {
			monstre.Pv = 0
		}

		// Affichage barre de vie monstre
		bar := DisplayHPBar(monstre.Pv, monstre.PvMax, 20)
		fmt.Printf("%s PV : [%s] %d/%d\n", monstre.Nom, bar, monstre.Pv, monstre.PvMax)
	}
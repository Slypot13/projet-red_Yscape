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

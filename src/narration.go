package main

import "time"

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
	typewriterPrint("Nous comptons sur toi jeune teteeeeeeeeeeeeeee de Neuille.. Euhh éleve de Ynov !!!", 40*time.Millisecond, "\033[1;31m")
}

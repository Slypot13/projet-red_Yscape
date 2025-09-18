package main

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

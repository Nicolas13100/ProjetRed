package game

var GoblinAttackIncreased bool

func MonsterPattern(P1 *Personnage, Monstre1 *Monstre) {
	// Check if the goblin's attack is currently increased
	if GoblinAttackIncreased {
		// Deactivate the increased attack
		Goblin.Atk = Goblin.Atk / 2
		GoblinAttackIncreased = false
	}

	// Reset goblin turn count every 3 turns
	if Tours%3 == 0 {
		Goblin.Atk = Goblin.Atk * 2
		GoblinAttackIncreased = true
	}
}

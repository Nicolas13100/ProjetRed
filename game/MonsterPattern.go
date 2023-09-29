package game

var GoblinAttackIncreased bool

func MonsterPattern(Monstre1 *Monstre) {
	if GoblinAttackIncreased {
		Goblin.Atk = Goblin.Atk / 2
		GoblinAttackIncreased = false
	}
	if Tours%3 == 0 {
		Goblin.Atk = Goblin.Atk * 2
		GoblinAttackIncreased = true
	}
}

package game

var (
	Goblin = Monstre{
		Name:  "Goblin",
		HpMax: 50,
		Hp:    50,
		Mana:  20,
		Type:  "Goblin",
	}

	Orc = Monstre{
		Name:  "Orc",
		HpMax: 75,
		Hp:    75,
		Mana:  30,
		Type:  "Orc",
	}

	HamsterGeant = Monstre{
		Name:  "HamsterGeant",
		HpMax: 100,
		Hp:    100,
		Mana:  40,
		Type:  "Boss",
	}
)

package game

var (
	Goblin = Monstre{
		Name:            "Goblin",
		Atk:             10,
		HpMax:           50,
		Hp:              50,
		Mana:            20,
		Type:            "Goblin",
		AlreadyDefeated: false,
		Defense:         0,
		Initiative:      0,
		XpDrop:          20,
	}

	Orc = Monstre{
		Name:            "Orc",
		Atk:             15,
		HpMax:           75,
		Hp:              75,
		Mana:            30,
		Type:            "Orc",
		AlreadyDefeated: false,
		Defense:         10,
		Initiative:      0,
		XpDrop:          20,
	}

	HamsterGeant = Monstre{
		Name:            "HamsterGeant",
		Atk:             30,
		HpMax:           100,
		Hp:              100,
		Mana:            40,
		Type:            "Boss",
		AlreadyDefeated: false,
		Defense:         20,
		Initiative:      0,
		XpDrop:          20,
	}
	Loup = Monstre{
		Name:            "Loup",
		Atk:             8,
		HpMax:           40,
		Hp:              50,
		Mana:            20,
		AlreadyDefeated: false,
		Defense:         0,
		Initiative:      0,
		XpDrop:          20,
		ItemDrop:        "Fourrue de Loup",
	}

	Troll = Monstre{
		Name:            "Troll",
		Atk:             15,
		HpMax:           75,
		Hp:              75,
		Mana:            30,
		AlreadyDefeated: false,
		Defense:         0,
		Initiative:      0,
		XpDrop:          20,
		ItemDrop:        "Peau de Troll",
	}

	Corbeau = Monstre{
		Name:            "Corbeau",
		Atk:             4,
		HpMax:           10,
		Hp:              10,
		Mana:            40,
		AlreadyDefeated: false,
		Defense:         0,
		Initiative:      0,
		XpDrop:          20,
		ItemDrop:        "Plume de Corbeau",
	}
	Sanglier = Monstre{
		Name:            "Sanglier",
		Atk:             8,
		HpMax:           30,
		Hp:              30,
		Mana:            40,
		AlreadyDefeated: false,
		Defense:         0,
		Initiative:      0,
		XpDrop:          20,
		ItemDrop:        "Cuir de Sanglier",
	}
)

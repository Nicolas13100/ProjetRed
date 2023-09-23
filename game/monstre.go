package game

import "fmt"

var (
	Goblin = Monstre{
		Name:       "Goblin",
		Atk:        10,
		HpMax:      50,
		Hp:         50,
		Mana:       20,
		Type:       "Goblin",
		Defense:    0,
		Initiative: 0,
		XpDrop:     20,
	}

	Orc = Monstre{
		Name:       "Orc",
		Atk:        15,
		HpMax:      75,
		Hp:         75,
		Mana:       30,
		Type:       "Orc",
		Defense:    10,
		Initiative: 0,
		XpDrop:     20,
	}

	HamsterGeant = Monstre{
		Name:       "HamsterGeant",
		Atk:        30,
		HpMax:      100,
		Hp:         100,
		Mana:       40,
		Type:       "Boss",
		Defense:    20,
		Initiative: 0,
		XpDrop:     20,
	}

	Loup = Monstre{
		Name:       "Loup",
		Atk:        10,
		HpMax:      50,
		Hp:         50,
		Mana:       20,
		Defense:    0,
		Initiative: 0,
		XpDrop:     20,
		ItemDrop: map[string]int{
			"Fourrure de Loup": 1,
		},
	}

	Loup2 = Monstre{
		Name:       "Loup ténébreux",
		Atk:        16,
		HpMax:      70,
		Hp:         70,
		Mana:       40,
		Defense:    10,
		Initiative: 10,
		XpDrop:     40,
		ItemDrop: map[string]int{
			"Fourrure de Loup ténébreux": 1,
			"Fourrure de Loup":           2,
		},
	}

	Renard = Monstre{
		Name:       "Maître Renard",
		Atk:        16,
		HpMax:      70,
		Hp:         70,
		Mana:       40,
		Defense:    10,
		Initiative: 10,
		XpDrop:     40,
		ItemDrop: map[string]int{
			"Fourrure de Renard": 1,
			"Crocs acérés":       1,
		},
	}

	Troll = Monstre{
		Name:       "Troll",
		Atk:        15,
		HpMax:      75,
		Hp:         75,
		Mana:       30,
		Defense:    10,
		Initiative: 0,
		XpDrop:     20,
		ItemDrop: map[string]int{
			"Peau de Troll": 1,
		},
	}

	Troll2 = Monstre{
		Name:       "Troll des montagnes",
		Atk:        35,
		HpMax:      150,
		Hp:         150,
		Mana:       60,
		Defense:    20,
		Initiative: 15,
		XpDrop:     80,
		ItemDrop: map[string]int{
			"Peau dure de Troll": 1,
			"Peau de Troll":      2,
			"Morceau de Gourdin": 1,
		},
	}

	Corbeau = Monstre{
		Name:       "Corbeau",
		Atk:        4,
		HpMax:      10,
		Hp:         10,
		Mana:       40,
		Defense:    0,
		Initiative: 0,
		XpDrop:     20,
		ItemDrop: map[string]int{
			"Plume de Corbeau": 1,
		},
	}

	Corbeau2 = Monstre{
		Name:       "Maître Corbeau",
		Atk:        8,
		HpMax:      20,
		Hp:         20,
		Mana:       40,
		Defense:    5,
		Initiative: 15,
		XpDrop:     30,
		ItemDrop: map[string]int{
			"Plume de Corbeau": 2,
			"Fromage":          1,
		},
	}

	Sanglier = Monstre{
		Name:       "Sanglier",
		Atk:        8,
		HpMax:      30,
		Hp:         30,
		Mana:       40,
		Defense:    0,
		Initiative: 0,
		XpDrop:     20,
		ItemDrop: map[string]int{
			"Fourrure de Loup": 1,
		},
	}

	Sanglier2 = Monstre{
		Name:       "Sanglier enragé",
		Atk:        16,
		HpMax:      60,
		Hp:         60,
		Mana:       80,
		Defense:    20,
		Initiative: 12,
		XpDrop:     20,
		ItemDrop: map[string]int{
			"Cuir supérieur de Sanglier": 1,
			"Cuir de Sanglier":           2,
		},
	}
)

func NewCorbeau() Monstre {
	return Monstre{
		Name:       "Corbeau",
		Atk:        4,
		HpMax:      10,
		Hp:         10,
		Mana:       40,
		Defense:    0,
		Initiative: 0,
		XpDrop:     20,
		ItemDrop: map[string]int{
			"Plume de Corbeau": 1,
		},
	}
}

func NewCorbeau2() Monstre {
	return Monstre{
		Name:       "Maître Corbeau",
		Atk:        8,
		HpMax:      20,
		Hp:         20,
		Mana:       40,
		Defense:    5,
		Initiative: 15,
		XpDrop:     30,
		ItemDrop: map[string]int{
			"Plume de Corbeau": 2,
			"Fromage":          1,
		},
	}
}

func NewSanglier() Monstre {
	return Monstre{
		Name:       "Sanglier",
		Atk:        8,
		HpMax:      30,
		Hp:         30,
		Mana:       40,
		Defense:    0,
		Initiative: 0,
		XpDrop:     20,
		ItemDrop: map[string]int{
			"Cuir de Sanglier": 1,
		},
	}
}

func NewSanglier2() Monstre {
	return Monstre{
		Name:       "Sanglier enragé",
		Atk:        16,
		HpMax:      60,
		Hp:         60,
		Mana:       80,
		Defense:    20,
		Initiative: 12,
		XpDrop:     20,
		ItemDrop: map[string]int{
			"Cuir supérieur de Sanglier": 1,
			"Cuir de Sanglier":           2,
		},
	}
}

func NewLoup() Monstre {
	return Monstre{
		Name:       "Loup",
		Atk:        10,
		HpMax:      40,
		Hp:         50,
		Mana:       20,
		Defense:    0,
		Initiative: 0,
		XpDrop:     20,
		ItemDrop: map[string]int{
			"Fourrure de Loup": 1,
		},
	}
}

func NewLoup2() Monstre {
	return Monstre{
		Name:       "Loup ténébreux",
		Atk:        16,
		HpMax:      70,
		Hp:         70,
		Mana:       40,
		Defense:    10,
		Initiative: 10,
		XpDrop:     40,
		ItemDrop: map[string]int{
			"Fourrure de Loup ténébreux": 1,
			"Fourrure de Loup":           2,
		},
	}
}

func NewRenard() Monstre {
	return Monstre{
		Name:       "Renard",
		Atk:        16,
		HpMax:      70,
		Hp:         70,
		Mana:       40,
		Defense:    10,
		Initiative: 10,
		XpDrop:     40,
		ItemDrop: map[string]int{
			"Fourrure de Renard": 2,
			"Crocs acérés":       1,
		},
	}
}

func NewTroll() Monstre {
	return Monstre{
		Name:       "Troll",
		Atk:        15,
		HpMax:      75,
		Hp:         75,
		Mana:       30,
		Defense:    0,
		Initiative: 0,
		XpDrop:     20,
		ItemDrop: map[string]int{
			"Peau de Troll": 1,
		},
	}
}

func NewTroll2() Monstre {
	return Monstre{
		Name:       "Troll des montagnes",
		Atk:        35,
		HpMax:      150,
		Hp:         150,
		Mana:       60,
		Defense:    20,
		Initiative: 15,
		XpDrop:     80,
		ItemDrop: map[string]int{
			"Peau dure de Troll": 1,
			"Peau de Troll":      2,
			"Morceau de Gourdin": 1,
		},
	}
}
func (Monstre1 *Monstre) DeadMonstre(P1 *Personnage) {
	if Monstre1.Hp <= 0 {
		fmt.Printf("%s a vaincu un(e) %s !\n", P1.Name, Monstre1.Name)
		fmt.Println("Fin du combat")
		GagnerXp(P1, Monstre1)
		DropsToInventory(P1, Monstre1.ItemDrop)
		for itemName, quantity := range Monstre1.ItemDrop {
			fmt.Printf("Objet obtenu: %s, Quantité: %d\n", itemName, quantity)
		}

	}
}
func waitForUserInput(message string) {
	fmt.Println(message)
	fmt.Print("Press Enter to continue...")
	fmt.Scan()
}

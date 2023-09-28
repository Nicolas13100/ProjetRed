package game

import "fmt"

var (
	Spayder = Monstre{
		Name:       "Spayder",
		Type:       "Araignée",
		Atk:        20,
		HpMax:      60,
		Hp:         60,
		Mana:       30,
		Defense:    10,
		Initiative: 0,
		XpDrop:     40,
		ItemDrop: map[string]int{
			"Fil de Spayder": 1,
		},
	}

	Lareinegnée = Monstre{
		Name:       "La reinegnée",
		Type:       "Araignée",
		Atk:        20,
		HpMax:      60,
		Hp:         60,
		Mana:       30,
		Defense:    10,
		Initiative: 0,
		XpDrop:     40,
		ItemDrop: map[string]int{
			"Toile de la reignée": 1,
			"Venin royal":         1,
			"Fil de Spayder":      2,
		},
	}

	Jarpent = Monstre{
		Name:       "Jarpent",
		Atk:        12,
		HpMax:      30,
		Hp:         30,
		Mana:       20,
		Type:       "Jarpent",
		Defense:    0,
		Initiative: 0,
		XpDrop:     20,
		ItemDrop: map[string]int{
			"Mûe de Jarpent": 1,
			"Bout d'argile":  1,
		},
	}

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
		Type:       "Loup",
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
		Type:       "Loup",
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
		Type:       "Renard",
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
		Type:       "Troll",
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
		Type:       "Troll",
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
		Type:       "Corbeau",
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
		Type:       "Corbeau",
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
		Type:       "Sanglier",
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

	Sanglier2 = Monstre{
		Name:       "Sanglier enragé",
		Type:       "Sanglier",
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

func (Monstre1 *Monstre) DeadMonstre(P1 *Personnage) {
	if Monstre1.Hp <= 0 {
		fmt.Printf("%s a vaincu un(e) %s !\n", P1.Name, Monstre1.Name)
		fmt.Println("Fin du combat")
		GagnerXp(P1, Monstre1)
		DropsToInventory(P1, Monstre1.ItemDrop)
		for itemName, quantity := range Monstre1.ItemDrop {
			fmt.Printf("Objet obtenu: %s, Quantité: %d\n", itemName, quantity)
		}
		waitForUserInput("Entrer 0 pour continuer ...")
	}
}

func waitForUserInput(message string) {
	var choice int
	fmt.Print(message)
	fmt.Scan(&choice)
	switch choice {
	case 0:
		break
	default:
		break
	}
}

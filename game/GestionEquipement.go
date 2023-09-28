package game

import "fmt"

func NewEquipement(P1 *Personnage) *Equipment {
	return &Equipment{
		Weapon:   false,
		Head:     false,
		Armor:    false,
		Hands:    false,
		Legs:     false,
		Boots:    false,
		Rings1:   false,
		Ring2:    false,
		Necklace: false,
	}
}

var (
	Chapeau1 = Equipment{
		Name:            "Chapeau de l'aventurier",
		Type:            "Head",
		HPBonus:         5,
		DefBonus:        3,
		InitiativeBonus: 3,
		Durability:      30,
		DurabilityMax:   30,
		Materials: map[string]int{
			"Plume de Corbeau": 1,
			"Cuir de Sanglier": 1,
		},
	}
	Tunique1 = Equipment{
		Name:            "Tunique de l'aventurier",
		Type:            "Armor",
		HPBonus:         7,
		DefBonus:        7,
		InitiativeBonus: 2,
		Durability:      30,
		DurabilityMax:   30,
		Materials: map[string]int{
			"Peau de Troll":    1,
			"Cuir de Sanglier": 2,
			"Fourrure de Loup": 1,
		},
	}

	Jambe1 = Equipment{
		Name: "Jambiere de l'aventurier",
		Type: "Legs",

		HPBonus:         6,
		DefBonus:        4,
		InitiativeBonus: 3,
		Durability:      30,
		DurabilityMax:   30,
		Materials: map[string]int{
			"Cuir de Sanglier": 2,
			"Fourrure de Loup": 1,
		},
	}

	Bottes1 = Equipment{
		Name: "Bottes de l'aventurier",
		Type: "Boots",

		HPBonus:         4,
		DefBonus:        2,
		InitiativeBonus: 3,
		Durability:      30,
		DurabilityMax:   30,
		Materials: map[string]int{
			"Cuir de Sanglier": 1,
			"Fourrure de Loup": 1,
		},
	}
	HeaumeCuir = Equipment{
		Name:            "Heaume en cuir",
		Type:            "Head",
		HPBonus:         10,
		DefBonus:        6,
		InitiativeBonus: 6,
		Durability:      50,
		DurabilityMax:   50,
		Materials: map[string]int{
			"Cuir de Sanglier": 5,
			"Fourrure de Loup": 2,
			"Peau de Troll":    1,
		},
	}
	PlastronCuir = Equipment{
		Name:            "Plastron en cuir",
		Type:            "Armor",
		HPBonus:         15,
		DefBonus:        15,
		InitiativeBonus: 4,
		Durability:      50,
		DurabilityMax:   50,
		Materials: map[string]int{
			"Cuir de Sanglier": 6,
			"Fourrure de Loup": 3,
			"Peau de Troll":    2,
		},
	}
	PantalonCuir = Equipment{
		Name: "Pantalon en cuir",
		Type: "Legs",

		HPBonus:         12,
		DefBonus:        8,
		InitiativeBonus: 5,
		Durability:      50,
		DurabilityMax:   50,
		Materials: map[string]int{
			"Cuir de Sanglier": 5,
			"Fourrure de Loup": 2,
			"Peau de Troll":    1,
		},
	}
	BottesCuir = Equipment{
		Name: "Bottes en cuir",
		Type: "Boots",

		HPBonus:         8,
		DefBonus:        5,
		InitiativeBonus: 6,
		Durability:      50,
		DurabilityMax:   50,
		Materials: map[string]int{
			"Cuir de Sanglier": 3,
			"Fourrure de Loup": 1,
		},
	}
	HeaumeMetal = Equipment{
		Name:            "Heaume en métal",
		Type:            "Head",
		HPBonus:         20,
		DefBonus:        12,
		InitiativeBonus: 9,
		Durability:      100,
		DurabilityMax:   100,
		Materials: map[string]int{

			"Fourrure de Loup": 2,
			"Peau de Troll":    1,
			"Acier":            5,
		},
	}
	PlastronMetal = Equipment{
		Name:            "Plastron en métal",
		Type:            "Armor",
		HPBonus:         30,
		DefBonus:        30,
		InitiativeBonus: 5,
		Durability:      100,
		DurabilityMax:   100,
		Materials: map[string]int{
			"Cuir de Sanglier": 2,
			"Fourrure de Loup": 3,
			"Peau de Troll":    2,
			"Acier":            8,
		},
	}
	PantalonMetal = Equipment{
		Name: "Pantalon en métal",
		Type: "Legs",

		HPBonus:         25,
		DefBonus:        16,
		InitiativeBonus: 6,
		Durability:      100,
		DurabilityMax:   100,
		Materials: map[string]int{
			"Cuir de Sanglier": 1,
			"Fourrure de Loup": 2,
			"Peau de Troll":    1,
			"Acier":            6,
		},
	}
	BottesMetal = Equipment{
		Name: "Bottes en métal",
		Type: "Boots",

		HPBonus:         16,
		DefBonus:        10,
		InitiativeBonus: 7,
		Durability:      100,
		DurabilityMax:   100,
		Materials: map[string]int{
			"Fourrure de Loup": 1,
			"Acier":            4,
		},
	}

	Epée1 = Equipment{
		Name:          "Epée de l'aventurier",
		Type:          "Weapon",
		AtkBonus:      5,
		Durability:    30,
		DurabilityMax: 30,
		Materials: map[string]int{
			"Acier":            2,
			"Cuir de Sanglier": 1,
		},
	}
	Katana = Equipment{
		Name:          "Katana",
		Type:          "Weapon",
		AtkBonus:      10,
		Durability:    50,
		DurabilityMax: 50,
		Materials: map[string]int{
			"Acier":            4,
			"Cuir de Sanglier": 2,
			"Peau de Troll":    1,
		},
	}
	Epéerenforcé = Equipment{
		Name:          "Epée renforcé",
		Type:          "Weapon",
		AtkBonus:      20,
		Durability:    100,
		DurabilityMax: 100,
		Materials: map[string]int{
			"Acier renforcé":   2,
			"Cuir de Sanglier": 2,
			"Peau de Troll":    1,
		},
	}

	KatanaSup = Equipment{
		Name:          "Katana supérieur",
		Type:          "Weapon",
		AtkBonus:      40,
		Durability:    120,
		DurabilityMax: 120,
		Materials: map[string]int{
			"Acier renforcé":   4,
			"Cuir de Sanglier": 2,
			"Peau de Troll":    2,
		},
	}

	ZoroBlade = Equipment{
		Name:            "Wadô Ichimonji",
		Type:            "Weapon",
		AtkBonus:        60,
		InitiativeBonus: 15,
		Durability:      150,
		DurabilityMax:   150,
		Materials: map[string]int{
			"Acier renforcé":   6,
			"Cuir de Sanglier": 3,
			"Peau de Troll":    3,
		},
	}
	CouronneSphinx = Equipment{
		Name:            "Couronne du Sphinx",
		Type:            "Couronne",
		Head:            true,
		AtkBonus:        30,
		InitiativeBonus: 25,
		HPBonus:         30,
		Durability:      90,
		DurabilityMax:   90,
	}

	Arc1 = Equipment{
		Name:            "Arc de l'aventurier",
		Type:            "Weapon",
		AtkBonus:        10,
		InitiativeBonus: 2,
		Durability:      30,
		DurabilityMax:   30,
		Materials: map[string]int{
			"Cuir de Sanglier": 1,
			"Fil de Spayder":   1,
		},
	}
	ArcSup = Equipment{
		Name:            "Arc supérieur",
		Type:            "Weapon",
		AtkBonus:        20,
		InitiativeBonus: 4,
		Durability:      50,
		DurabilityMax:   50,
		Materials: map[string]int{
			"Cuir de Sanglier":    3,
			"Toile de la reignée": 1,
			"Fil de Spayder":      2,
		},
	}
	Failure = Equipment{
		Name:            "Failure",
		Type:            "Weapon",
		AtkBonus:        100,
		InitiativeBonus: -5,
		Durability:      200,
		DurabilityMax:   200,
		Materials: map[string]int{
			"Squelette de Rouquin": 1,
			"Acier renforcé":       6,
			"Peau de Troll":        2,
		},
	}
)

func (p *Personnage) EquipItem(item Equipment) {
	switch item.Type {
	case "Head":
		if p.Equipement.Head {
			p.UnequipItem(p.Head)
		}
		p.Head[item.Name] = item
		p.Equipement.Head = true
		p.EquippedHead = item.Name
	case "Armor":
		if p.Equipement.Armor {
			p.UnequipItem(p.Armor)
		}
		p.Armor[item.Name] = item
		p.Equipement.Armor = true
		p.EquippedArmor = item.Name
	case "Legs":
		if p.Equipement.Legs {
			p.UnequipItem(p.Legs)
		}
		p.Legs[item.Name] = item
		p.Equipement.Legs = true
		p.Equippedlegs = item.Name
	case "Boots":
		if p.Equipement.Boots {
			p.UnequipItem(p.Feets)
		}
		p.Feets[item.Name] = item
		p.Equipement.Boots = true
		p.EquippedFeets = item.Name
	case "Weapon":
		if p.Equipement.Weapon {
			p.UnequipItem(p.Weapon)
		}
		p.Weapon[item.Name] = item
		p.Equipement.Weapon = true
		p.EquippedWeapon = item.Name
	}

	// Update character stats based on the equipped item
	p.HpMax += item.HPBonus
	p.Atk += item.AtkBonus
	p.Defense += item.DefBonus
	p.Initiative += item.InitiativeBonus
	// ... add similar lines for other stats

	// Update inventory if necessary (remove item from inventory)
	delete(p.Inventory, item.Name)

	// Add item to EquipementMap
	p.EquipementMap[item.Name] = item
}

func (p *Personnage) ListEquipableItems() []string {
	equipableItems := []string{}

	for item := range p.Inventory {
		switch item {
		case "Failure", "Wadô Ichimonji", "Chapeau de l'aventurier", "Tunique de l'aventurier", "Jambiere de l'aventurier", "Bottes de l'aventurier", "Epée de l'aventurier", "Arc de l'aventurier", "Heaume en cuir", "Plastron en cuir", "Pantalon en cuir", "Bottes en cuir", "Heaume en métal", "Plastron en métal", "Pantalon en métal", "Bottes en métal", "Katana", "Epée renforcé", "Katana supérieur", "Couronne du Sphinx", "Arc supérieur":
			equipableItems = append(equipableItems, item)
		}
	}

	return equipableItems
}

func (p *Personnage) GetItemByName(name string) Equipment {
	switch name {
	case "Chapeau de l'aventurier":
		return Chapeau1
	case "Tunique de l'aventurier":
		return Tunique1
	case "Jambiere de l'aventurier":
		return Jambe1
	case "Bottes de l'aventurier":
		return Bottes1
	case "Heaume en cuir":
		return HeaumeCuir
	case "Plastron en cuir":
		return PlastronCuir
	case "Pantalon en cuir":
		return PantalonCuir
	case "Bottes en cuir":
		return BottesCuir
	case "Heaume en métal":
		return HeaumeMetal
	case "Plastron en métal":
		return PlastronMetal
	case "Pantalon en métal":
		return PantalonMetal
	case "Bottes en métal":
		return BottesMetal
	case "Epée de l'aventurier":
		return Epée1
	case "Wadô Ichimonji":
		return ZoroBlade
	case "Arc de l'aventurier":
		return Arc1
	case "Failure":
		return Failure

	default:
		return Equipment{}
	}
}

func (p *Personnage) EquipItemFromInventory() {
	// List equipable items
	equipableItems := p.ListEquipableItems()

	// Print the list of equipable items
	for i, item := range equipableItems {
		fmt.Printf("%d. %s\n", i+1, item)
	}

	// Prompt user for selection
	fmt.Print("Enter the number of the item you want to equip: ")
	var selection int
	_, err := fmt.Scan(&selection)
	if err != nil || selection < 1 || selection > len(equipableItems) {
		fmt.Println("Invalid input. Please try again.")
		return
	}

	// Get the selected item
	selectedItemName := equipableItems[selection-1]
	selectedItem := p.GetItemByName(selectedItemName)

	// Equip the selected item
	p.EquipItem(selectedItem)

	fmt.Printf("%s equipped!\n", selectedItemName)
}

func (p *Personnage) UnequipItem(item map[string]Equipment) {
	for itemName, equippedItem := range item {
		switch equippedItem.Type {
		case "Head":
			if p.Equipement.Head {
				deletedItem := itemName
				delete(p.Head, deletedItem)
				p.Equipement.Head = false
				// Remove item from EquipementMap
				delete(p.EquipementMap, deletedItem)
			}
		case "Armor":
			if p.Equipement.Armor {
				deletedItem := itemName
				delete(p.Armor, deletedItem)
				p.Equipement.Armor = false
				// Remove item from EquipementMap
				delete(p.EquipementMap, deletedItem)
			}
		case "Legs":
			if p.Equipement.Legs {
				deletedItem := itemName
				delete(p.Legs, deletedItem)
				p.Equipement.Legs = false
				// Remove item from EquipementMap
				delete(p.EquipementMap, deletedItem)
			}
		case "Boots":
			if p.Equipement.Boots {
				deletedItem := itemName
				delete(p.Feets, deletedItem)
				p.Equipement.Boots = false
				// Remove item from EquipementMap
				delete(p.EquipementMap, deletedItem)
			}
		case "Weapon":
			if p.Equipement.Weapon {
				deletedItem := itemName
				delete(p.Weapon, deletedItem)
				p.Equipement.Weapon = false
				// Remove item from EquipementMap
				delete(p.EquipementMap, deletedItem)
			}
		}

		// Update character stats based on the unequipped item
		p.HpMax -= equippedItem.HPBonus
		p.Atk -= equippedItem.AtkBonus
		p.Defense -= equippedItem.DefBonus
		p.Initiative -= equippedItem.InitiativeBonus
		// ... add similar lines for other stats

		// Update inventory if necessary (add item back to inventory)
		p.Inventory[itemName]++
	}
}
func GetMaterialsFromEquipment(equipment Equipment) map[string]int {
	return equipment.Materials
}
func GetitemFromInventory(P1 Personnage) map[string]int {
	return P1.Inventory
}
func (P1 *Personnage) RemoveItem(itemName string) {
	if P1.Inventory[itemName] > 1 {
		P1.Inventory[itemName]--
	} else {
		// Si la quantité est de 1, supprimez complètement l'entrée de l'inventaire
		delete(P1.Inventory, itemName)
	}
}
func (P1 *Personnage) AddItem(itemName string) {
	if quantity, ok := P1.Inventory[itemName]; ok {
		P1.Inventory[itemName] = quantity + 1
	} else {
		P1.Inventory[itemName] = 1
	}
}

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
		HPBonus:         4,
		DefBonus:        2,
		InitiativeBonus: 2,
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
		Durability:      50,
		DurabilityMax:   50,
		Materials: map[string]int{
			"Peau de Troll":    1,
			"Cuir de Sanglier": 2,
			"Fourrure de Loup": 1,
		},
	}

	Jambe1 = Equipment{
		Name: "Jambiere de l'aventurier",
		Type: "Legs",

		HPBonus:         5,
		DefBonus:        2,
		InitiativeBonus: 2,
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

		HPBonus:         5,
		DefBonus:        2,
		InitiativeBonus: 2,
		Durability:      30,
		DurabilityMax:   30,
		Materials: map[string]int{
			"Cuir de Sanglier": 1,
			"Fourrure de Loup": 1,
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

	ZoroBlade = Equipment{
		Name:          "Wadô Ichimonji",
		Type:          "Weapon",
		AtkBonus:      60,
		Durability:    90,
		DurabilityMax: 90,
		Materials: map[string]int{
			"Acier renforcé": 1,
			"Peau de Troll":  2,
		},
	}
	CouronneSphinx = Equipment{
		Name:            "Couronne du Sphinx",
		Type:            "Couronne",
		Necklace:        true,
		AtkBonus:        30,
		InitiativeBonus: 25,
		HPBonus:         20,
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
	Failure = Equipment{
		Name:            "Failure",
		Type:            "Weapon",
		AtkBonus:        100,
		InitiativeBonus: -5,
		Durability:      200,
		DurabilityMax:   200,
		Materials: map[string]int{
			"Squelette de Rouquin": 1,
			"Acier renforcé":       2,
		},
	}
	Headtest = Equipment{
		Name:            "Head Test",
		Type:            "Head",
		HPBonus:         4,
		DefBonus:        2,
		InitiativeBonus: 2,
		Durability:      3,
		DurabilityMax:   3,
	}
	Headtest2 = Equipment{
		Name:            "Head Test2",
		Type:            "Head",
		HPBonus:         4,
		DefBonus:        2,
		InitiativeBonus: 2,
		Durability:      3,
		DurabilityMax:   3,
	}
	Weapontest = Equipment{
		Name:          "Test",
		Type:          "Weapon",
		AtkBonus:      1,
		Durability:    3,
		DurabilityMax: 3}
)

func (p *Personnage) EquipItem(item Equipment) {
	switch item.Type {
	case "Head":
		if p.Equipement.Head {
			p.UnequipItem(p.Head)
		}
		p.Head[item.Name] = item
		p.Equipement.Head = true
	case "Armor":
		if p.Equipement.Armor {
			p.UnequipItem(p.Armors)
		}
		p.Armors[item.Name] = item
		p.Equipement.Armor = true
	case "Legs":
		if p.Equipement.Legs {
			p.UnequipItem(p.Legs)
		}
		p.Legs[item.Name] = item
		p.Equipement.Legs = true
	case "Boots":
		if p.Equipement.Boots {
			p.UnequipItem(p.Feets)
		}
		p.Feets[item.Name] = item
		p.Equipement.Boots = true
	case "Weapon":
		if p.Equipement.Weapon {
			p.UnequipItem(p.Weapons)
		}
		p.Weapons[item.Name] = item
		p.Equipement.Weapon = true
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
		case "Test", "Head Test2", "Head Test", "Failure", "Wadô Ichimonji", "Chapeau de l'aventurier", "Tunique de l'aventurier", "Jambiere de l'aventurier", "Bottes de l'aventurier", "Epée de l'aventurier", "Arc de l'aventurier":
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
	case "Epée de l'aventurier":
		return Epée1
	case "Wadô Ichimonji":
		return ZoroBlade
	case "Arc de l'aventurier":
		return Arc1
	case "Failure":
		return Failure
	case "Head Test":
		return Headtest
	case "Head Test2":
		return Headtest2
	case "Test":
		return Weapontest
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
				delete(p.Armors, deletedItem)
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
				delete(p.Weapons, deletedItem)
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

func (p *Personnage) UpdateDurability(item map[string]Equipment) {
	for itemName, equipment := range item {
		// Assuming 'Equipment' struct has a field named 'Durability'
		if equipment.Durability > 0 {
			equipment.Durability--
			fmt.Printf("%s durability reduced to %d\n", itemName, equipment.Durability)
		} else {
			fmt.Printf("%s is already broken\n", itemName)
		}
	}
}

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
	}
	Tunique1 = Equipment{
		Name:            "Tunique de l'aventurier",
		Type:            "Armor",
		HPBonus:         7,
		DefBonus:        7,
		InitiativeBonus: 2,
		Durability:      50,
		DurabilityMax:   50,
	}

	Jambe1 = Equipment{
		Name: "Jambiere de l'aventurier",
		Type: "Legs",

		HPBonus:         5,
		DefBonus:        2,
		InitiativeBonus: 2,
		Durability:      30,
		DurabilityMax:   30,
	}

	Bottes1 = Equipment{
		Name: "Bottes de l'aventurier",
		Type: "Boots",

		HPBonus:         5,
		DefBonus:        2,
		InitiativeBonus: 2,
		Durability:      30,
		DurabilityMax:   30,
	}

	Epée1 = Equipment{
		Name:          "Epée de l'aventurier",
		Type:          "Weapon",
		AtkBonus:      5,
		Durability:    30,
		DurabilityMax: 30,
	}

	ZoroBlade = Equipment{
		Name:          "Wadô Ichimonji",
		Type:          "Weapon",
		AtkBonus:      60,
		Durability:    90,
		DurabilityMax: 90,
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
	}
	Failure = Equipment{
		Name:            "Failure",
		Type:            "Weapon",
		AtkBonus:        100,
		InitiativeBonus: -5,
		Durability:      200,
		DurabilityMax:   200,
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
	fmt.Println(item)
	switch item.Type {
	case "Head":
		if p.Equipement.Head {
			p.UnequipItem(item)
		}
		p.Head[item.Name] = item
		p.Equipement.Head = true
	case "Armor":
		if p.Equipement.Armor {
			p.UnequipItem(item)
		}
		p.Armors[item.Name] = item
		p.Equipement.Armor = true
	case "Legs":
		if p.Equipement.Legs {
			p.UnequipItem(item)
		}
		p.Legs[item.Name] = item
		p.Equipement.Legs = true
	case "Boots":
		if p.Equipement.Boots {
			p.UnequipItem(item)
		}
		p.Feets[item.Name] = item
		p.Equipement.Boots = true
	case "Weapon":
		if p.Equipement.Weapon {
			p.UnequipItem(item)
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

func (p *Personnage) UnequipItem(item Equipment) {
	fmt.Println(item)
	switch item.Type {
	case "Head":
		if p.Equipement.Head {
			deletedItem := p.Head[item.Name]
			delete(p.Head, item.Name)
			p.Equipement.Head = false
			// Remove item from EquipementMap
			delete(p.EquipementMap, deletedItem.Name)
		}
	case "Armor":
		if p.Equipement.Armor {
			deletedItem := p.Armors[item.Name]
			delete(p.Armors, item.Name)
			p.Equipement.Armor = false
			// Remove item from EquipementMap
			delete(p.EquipementMap, deletedItem.Name)
		}
	case "Legs":
		if p.Equipement.Legs {
			deletedItem := p.Legs[item.Name]
			delete(p.Legs, item.Name)
			p.Equipement.Legs = false
			// Remove item from EquipementMap
			delete(p.EquipementMap, deletedItem.Name)
		}
	case "Boots":
		if p.Equipement.Boots {
			deletedItem := p.Feets[item.Name]
			delete(p.Feets, item.Name)
			p.Equipement.Boots = false
			// Remove item from EquipementMap
			delete(p.EquipementMap, deletedItem.Name)
		}
	case "Weapon":
		if p.Equipement.Weapon {
			deletedItem := p.Weapons[item.Name]
			delete(p.Weapons, item.Name)
			p.Equipement.Weapon = false
			// Remove item from EquipementMap
			delete(p.EquipementMap, deletedItem.Name)
		}
	}

	// Update character stats based on the unequipped item
	p.HpMax -= item.HPBonus
	p.Atk -= item.AtkBonus
	p.Defense -= item.DefBonus
	p.Initiative -= item.InitiativeBonus
	// ... add similar lines for other stats

	// Update inventory if necessary (add item back to inventory)
	p.Inventory[item.Name]++
}

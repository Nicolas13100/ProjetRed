package game

import "fmt"

func NewEquipement(P1 *Personnage) *Equipment {
	return &Equipment{
		Head: false,
		Body: false,
		Leg:  false,
	}
}

var (
	Chapeau1 = Equipment{
		Name:            "Chapeau de l'aventurier",
		Type:            "Head",
		Head:            true,
		HPBonus:         4,
		DefBonus:        2,
		InitiativeBonus: 2,
	}
	Tunique1 = Equipment{
		Name:            "Tunique de l'aventurier",
		Type:            "Body",
		Body:            true,
		HPBonus:         7,
		DefBonus:        7,
		InitiativeBonus: 2,
	}
	Jambe1 = Equipment{
		Name:            "Jambiere de l'aventurier",
		Type:            "Legs",
		Leg:             true,
		HPBonus:         5,
		DefBonus:        2,
		InitiativeBonus: 2,
	}
	Bottes1 = Equipment{
		Name:            "Bottes de l'aventurier",
		Type:            "Boots",
		Leg:             true,
		HPBonus:         5,
		DefBonus:        2,
		InitiativeBonus: 2,
	}

	Epée1 = Equipment{
		Name:          "Epée de l'aventurier",
		Type:          "1HandWeapon",
		OneHandWeapon: true,
		AtkBonus:      5,
	}
	ZoroBlade = Equipment{
		Name:          "Wadô Ichimonji",
		Type:          "1HandWeapon",
		OneHandWeapon: true,
		AtkBonus:      60,
	}

	Arc1 = Equipment{
		Name:            "Arc de l'aventurier",
		Type:            "2HandWeapon",
		TwoHandWeapon:   true,
		AtkBonus:        10,
		InitiativeBonus: 2,
	}
	Failure = Equipment{
		Name:            "Failure",
		Type:            "2HandWeapon",
		TwoHandWeapon:   true,
		AtkBonus:        100,
		InitiativeBonus: -5,
	}
	Chapeau2 = Equipment{
		Name:            "Chapeau de l'aventurier",
		Type:            "Head",
		Head:            true,
		HPBonus:         4,
		DefBonus:        2,
		InitiativeBonus: 2,
	}
)

func IsEquipable(itemName string) bool {
	switch itemName {
	case "Failure", "Wadô Ichimonji", "Chapeau de l'aventurier", "Tunique de l'aventurier", "Bottes de l'aventurier", "Epée de l'aventurier", "Arc de l'aventurier":
		return true
	default:
		return false
	}
}

func GetEquipmentByName(itemName string) Equipment {
	switch itemName {
	case "Chapeau de l'aventurier":
		return Chapeau1
	case "Tunique de l'aventurier":
		return Tunique1
	case "Bottes de l'aventurier":
		return Bottes1
	case "Epée de l'aventurier":
		return Epée1
	case "Arc de l'aventurier":
		return Arc1
	case "Failure":
		return Failure
	case "Wadô Ichimonji":
		return ZoroBlade
	default:
		// Handle the case where the item is not found
		return Equipment{} // Assuming an empty Equipment struct here
	}
}

func ListEquipableItems(inventory map[string]int) []Equipment {
	equipableItems := []Equipment{}

	for itemName, quantity := range inventory {
		if IsEquipable(itemName) {
			for i := 0; i < quantity; i++ {
				equipableItems = append(equipableItems, GetEquipmentByName(itemName))
			}
		}
	}

	return equipableItems
}

func (p *Personnage) EquipItemFromInventory(itemName string) {
	itemToEquip := GetEquipmentByName(itemName)

	// Check if the item is equipable
	if !IsEquipable(itemName) {
		fmt.Println("Item is not equipable")
		return
	}

	// Check if there's already an item equipped of the same type/key
	if existingItem, ok := p.EquipementMap[itemToEquip.Type]; ok {
		// If equipping a one-handed weapon, remove the existing one
		if itemToEquip.Type == "1HandWeapon" {
			// Place the existing one-handed weapon back in the inventory
			p.Inventory[existingItem.Name]++
		}
	}

	// If equipping a two-handed weapon, remove all weapons from hand
	if itemToEquip.Type == "2HandWeapon" {
		for _, weapon := range p.WeaponInHand {
			p.Inventory[weapon.Name]++
		}
		p.WeaponInHand = make(map[int]Equipment)
	}

	// Equip the item
	p.EquipementMap[itemToEquip.Type] = itemToEquip

	// Remove the item from inventory
	if p.Inventory[itemName] > 0 {
		p.Inventory[itemName]--
	} else {
		fmt.Println("Item not found in inventory")
	}
}

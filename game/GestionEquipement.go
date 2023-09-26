package game

import "fmt"

func NewEquipement(P1 *Personnage) *Equipment {
	return &Equipment{
		Head: false,
		Body: false,
		Leg:  false,
	}
}

func (p *Personnage) Equip(item Equipment) {
	switch item.Type {
	case "Head", "Body", "Legs", "Boots":
		p.Equipement = item
		p.Inventory[item.Name]--
		p.EquipementMap[item.Type] = item
		p.RemoveZeroValueItems()
	case "Weapon":

		fmt.Println("Selectioné dans qu'elle main metre l'arme : (1. Droite/2. Gauche)")
		var choice int
		fmt.Scan(&choice)

		if choice == 1 {
			if p.Equipement.TwoHandWeapon {
				p.RemoveEquipment(p.Equipement.Name)
			}
			if p.Equipement.Weapon1 {
				p.RemoveEquipment(p.Equipement.Name)
				p.Equipement.Weapon1 = false
			}
			p.Equipement = item
			p.Inventory[item.Name]--
			p.EquipementMap[item.Type] = item
			p.Equipement.Weapon1 = true
			p.RemoveZeroValueItems()
		} else if choice == 2 {
			if p.Equipement.TwoHandWeapon {
				p.RemoveEquipment(p.Equipement.Name)
			}
			if p.Equipement.Weapon2 {
				p.RemoveEquipment(p.Equipement.Name)
				p.Equipement.Weapon2 = false
			}
			p.Equipement = item
			p.Inventory[item.Name]--
			p.EquipementMap[item.Type] = item
			p.Equipement.Weapon2 = true
			p.RemoveZeroValueItems()
		} else {
			fmt.Println("Invalid choice. No changes made.")
		}

	case "2HandWeapon":
		if p.Equipement.Weapon1 {
			p.RemoveEquipment(p.Equipement.Name)
			p.Equipement.Weapon1 = false
		}
		if p.Equipement.Weapon2 {
			p.RemoveEquipment(p.Equipement.Name)
			p.Equipement.Weapon2 = false
		}
		if p.Equipement.TwoHandWeapon {
			p.RemoveEquipment(p.Equipement.Name)
		}
		p.Equipement = item
		p.Inventory[item.Name]--
		p.EquipementMap[item.Type] = item
		p.Equipement.Weapon2 = true
		p.RemoveZeroValueItems()
	default:
		fmt.Println("Invalid equipment type")
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
		Type:          "Weapon",
		OneHandWeapon: true,
		AtkBonus:      5,
	}

	Arc1 = Equipment{
		Name:            "Arc de l'aventurier",
		Type:            "2HandWeapon",
		TwoHandWeapon:   true,
		AtkBonus:        10,
		InitiativeBonus: 2,
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
	case "Chapeau de l'aventurier", "Tunique de l'aventurier", "Bottes de l'aventurier", "Epée de l'aventurier", "Arc de l'aventurier":
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

func (p *Personnage) RemoveEquipment(itemName string) {
	if equippedItem, ok := p.EquipementMap[itemName]; ok {
		p.Inventory[itemName]++
		delete(p.EquipementMap, itemName)

		switch equippedItem.Type {
		case "Head":
			p.Equipement.Head = false
		case "Body":
			p.Equipement.Body = false
		case "Legs":
			p.Equipement.Leg = false
		case "Boots":
			p.Equipement.Boots = false
		case "Weapon":
			p.Equipement.OneHandWeapon = false
		case "2HandWeapon":
			p.Equipement.TwoHandWeapon = false
		default:
			// Handle any other equipment types if needed
		}
	}
}

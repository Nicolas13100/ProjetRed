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
	case "Head":
		p.Equipement = item
	case "Body":
		p.Equipement = item
	case "Legs":
		p.Equipement = item
	case "Boots":
		p.Equipement = item
	case "Weapon":
		p.Equipement = item
	default:
		fmt.Println("Invalid equipment type")
	}
}

func (P1 *Personnage) DesequiperHead() {
	if P1.Equipement.Head {
		P1.Inventory["Chapeau de l'aventurier"]++
		P1.Equipement.HPBonus -= 10
		P1.Equipement.Head = false
	}
}

func (P1 *Personnage) DesequiperBody() {
	if P1.Equipement.Body {
		P1.Inventory["Tunique de l'aventurier"]++
		P1.Equipement.HPBonus -= 25
		P1.Equipement.Body = false
	}
}

func (P1 *Personnage) DesequiperLeg() {
	if P1.Equipement.Leg {
		P1.Inventory["Bottes de l'aventurier"]++
		P1.Equipement.HPBonus -= 15
		P1.Equipement.Leg = false
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
		Name:     "Epée de l'aventurier",
		Type:     "Weapon",
		RHWeapon: true,
		AtkBonus: 5,
	}

	Arc1 = Equipment{
		Name:            "Arc de l'aventurier",
		Type:            "Weapon",
		RHWeapon:        true,
		LHWeapon:        true,
		AtkBonus:        5,
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

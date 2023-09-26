package game

import "fmt"

func NewEquipement(P1 *Personnage) *Equipment {
	return &Equipment{
		Head: false,
		Body: false,
		Leg:  false,
	}
}

func (P1 *Personnage) Equip(equipment Equipment) {
	for i, equipped := range P1.Equipements {
		if equipped.Type == equipment.Type {
			P1.Unequip(i)
			break
		}
	}

	P1.Equipements = append(P1.Equipements, equipment)

	P1.HpMax += equipment.HPBonus
	P1.Atk += equipment.AtkBonus
	P1.Defense += equipment.DefBonus
	P1.Initiative += equipment.InitiativeBonus
}

func (P1 *Personnage) Unequip(index int) {
	if index < 0 || index >= len(P1.Equipements) {
		fmt.Println("Equipement incorrect")
		return
	}

	equipmentToRemove := P1.Equipements[index]

	P1.HpMax -= equipmentToRemove.HPBonus
	P1.Atk -= equipmentToRemove.AtkBonus
	P1.Defense -= equipmentToRemove.DefBonus
	P1.Initiative -= equipmentToRemove.InitiativeBonus

	P1.Equipements = append(P1.Equipements[:index], P1.Equipements[index+1:]...)
}

/*func (P1 *Personnage) EquiperHead() {
	if P1.Inventory["Chapeau de l'aventurier"] > 0 {
		if P1.Equipement.Head {
			// Si le personnage a déjà un chapeau équipé, le remettre dans l'inventaire
			P1.Inventory["Chapeau de l'aventurier"]++
		}
		P1.Equipement.Head = true
		P1.Equipement.HPBonus += 10
		delete(P1.Inventory, "Chapeau de l'aventurier")
	}
}

func (P1 *Personnage) EquiperBody() {
	if P1.Inventory["Tunique de l'aventurier"] > 0 {
		if P1.Equipement.Body {
			P1.Inventory["Tunique de l'aventurier"]++
		}
		P1.Equipement.Body = true
		P1.Equipement.HPBonus += 25
		delete(P1.Inventory, "Tunique de l'aventurier")
	}
}

func (P1 *Personnage) EquiperLeg() {
	if P1.Inventory["Bottes de l'aventurier"] > 0 {
		if P1.Equipement.Leg {
			P1.Inventory["Bottes de l'aventurier"]++
		}
		P1.Equipement.Leg = true
		P1.Equipement.HPBonus += 15
		delete(P1.Inventory, "Bottes de l'aventurier")
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
*/
var (
	Chapeau = Equipment{
		Name:            "Chapeau de l'aventurier",
		Type:            "Chapeau",
		Head:            true,
		HPBonus:         4,
		DefBonus:        2,
		InitiativeBonus: 2,
	}
	Tunique = Equipment{
		Name:            "Tunique de l'aventurier",
		Type:            "Tunique",
		Body:            true,
		HPBonus:         7,
		DefBonus:        7,
		InitiativeBonus: 2,
	}
	Bottes = Equipment{
		Name:            "Bottes de l'aventurier",
		Type:            "Bottes",
		Leg:             true,
		HPBonus:         5,
		DefBonus:        2,
		InitiativeBonus: 2,
	}

	Epée = Equipment{
		Name:     "Epée de l'aventurier",
		Weapon:   true,
		AtkBonus: 5,
		DefBonus: 2,
	}

	Arc = Equipment{
		Name:            "Arc de l'aventurier",
		Weapon:          true,
		AtkBonus:        5,
		InitiativeBonus: 2,
	}
)

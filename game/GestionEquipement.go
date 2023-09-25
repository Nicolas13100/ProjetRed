package game

import "fmt"

func NewEquipement(P1 *Personnage) *Equipment {
	return &Equipment{
		Head: false,
		Body: false,
		Leg:  false,
	}
}

func (P1 *Personnage) ShowEquipBonus(index int) {
	equipment, found := P1.Equipements[index]
	if !found {
		fmt.Println("Équipement non trouvé.")
		return
	}

	fmt.Printf("Équipement choisi : %s\n", equipment.Name)
	fmt.Printf("- Bonus d'attaque : %d\n", equipment.AtkBonus)
	fmt.Printf("- Bonus de défense : %d\n", equipment.DefBonus)
	fmt.Printf("- Bonus de santé : %d\n", equipment.HPBonus)
	fmt.Printf("- Bonus d'initiative : %d\n", equipment.InitiativeBonus)
}

func (P1 *Personnage) EquiperHead() {
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

package game

func NewEquipement(personnage *Personnage) *Equipment {
	return &Equipment{
		Head: false,
		Body: false,
		Leg:  false,
	}
}

func (p1 *Personnage) EquiperHead() {
	if p1.Inventory["Chapeau de l'aventurier"] > 0 {
		if p1.Equipement.Head {
			// Si le personnage a déjà un chapeau équipé, le remettre dans l'inventaire
			p1.Inventory["Chapeau de l'aventurier"]++
		}
		p1.Equipement.Head = true
		p1.Equipement.HPBonus += 10
		delete(p1.Inventory, "Chapeau de l'aventurier")
	}
}

func (p1 *Personnage) EquiperBody() {
	if p1.Inventory["Tunique de l'aventurier"] > 0 {
		if p1.Equipement.Body {
			p1.Inventory["Tunique de l'aventurier"]++
		}
		p1.Equipement.Body = true
		p1.Equipement.HPBonus += 25
		delete(p1.Inventory, "Tunique de l'aventurier")
	}
}

func (p1 *Personnage) EquiperLeg() {
	if p1.Inventory["Bottes de l'aventurier"] > 0 {
		if p1.Equipement.Leg {
			p1.Inventory["Bottes de l'aventurier"]++
		}
		p1.Equipement.Leg = true
		p1.Equipement.HPBonus += 15
		delete(p1.Inventory, "Bottes de l'aventurier")
	}
}

func (p1 *Personnage) DesequiperHead() {
	if p1.Equipement.Head {
		p1.Inventory["Chapeau de l'aventurier"]++
		p1.Equipement.HPBonus -= 10
		p1.Equipement.Head = false
	}
}

func (p1 *Personnage) DesequiperBody() {
	if p1.Equipement.Body {
		p1.Inventory["Tunique de l'aventurier"]++
		p1.Equipement.HPBonus -= 25
		p1.Equipement.Body = false
	}
}

func (p1 *Personnage) DesequiperLeg() {
	if p1.Equipement.Leg {
		p1.Inventory["Bottes de l'aventurier"]++
		p1.Equipement.HPBonus -= 15
		p1.Equipement.Leg = false
	}
}

package game

import "fmt"

func (p1 *Personnage) FightInventory(m1 *Monstre, spells *Spell) {
	for {
		ClearConsole()
		fmt.Println("\nInventaire:")
		fmt.Println("Cash :", p1.Gold)
		for key, value := range p1.Inventory {
			fmt.Printf("%s: %d \n", key, value)

		}
		fmt.Println("\nQue voulez-vous faire ?")
		fmt.Println("1. Sélectionner une potion de soin")
		fmt.Println("2. Sélectionner une potion de poison")
		fmt.Println("3. Sélectionner une potion de mana")
		fmt.Println("4. Sélectionnez un sort à utiliser")
		fmt.Println("5. Equipements")
		fmt.Println("0. Retourner en arrière")

		var input int
		fmt.Scan(&input)

		switch input {
		case 1:
			if count, ok := p1.Inventory["Potion de soin"]; ok && count > 0 {
				TakePot(p1)
			} else {
				fmt.Println("Vous n'avez pas de Potion de soin dans votre inventaire.")
			}
		case 2:
			if count, ok := p1.Inventory["Potion de poison"]; ok && count > 0 {
				poisonPot(p1, m1)
			} else {
				fmt.Println("Vous n'avez pas de Potion de poison dans votre inventaire.")
			}
		case 3:
			if count, ok := p1.Inventory["Potion de mana"]; ok && count > 0 {
				ManaPot(p1)
			} else {
				fmt.Println("Vous n'avez pas de Potion de mana dans votre inventaire.")
			}
		case 4:
			InitSpells()

		case 5:
			fmt.Println("Voulez-vous 1.équiper ou 2.désequiper un equipement ? (3. retour)")
			var input int
			fmt.Scan(&input)
			switch input {
			case 1:
				fmt.Println("Que souhaitez-vous équiper ? 1. Casque / 2. Armure / 3. Pied / 4. retour ")
				var input int
				fmt.Scan(&input)
				switch input {
				case 1:
					p1.EquiperHead()
				case 2:
					p1.EquiperBody()
				case 3:
					p1.EquiperLeg()
				case 4:
					continue

				}
			case 2:
				fmt.Println("Que souhaitez-vous désequiper ? 1. Casque / 2. Armure / 3. Pied / 4. retour ")
				var input int
				fmt.Scan(&input)
				switch input {
				case 1:
					p1.DesequiperHead()
				case 2:
					p1.DesequiperBody()
				case 3:
					p1.DesequiperLeg()
				case 4:
					continue
				}
			case 3:
				continue
			}
		case 0:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func (p1 *Personnage) BaseInventory(m1 *Monstre, spells *Spell) {
	for {
		ClearConsole()
		fmt.Println("\nInventaire:")
		fmt.Println("Cash :", p1.Gold)
		for key, value := range p1.Inventory {
			fmt.Printf("%s: %d \n", key, value)

		}
		fmt.Println("\nQue voulez-vous faire ?")
		fmt.Println("1. Sélectionner une potion de soin")
		fmt.Println("2. Sélectionner une potion de poison")
		fmt.Println("3. Sélectionner une potion de mana")
		fmt.Println("4. Consulter les sorts acquis")
		fmt.Println("5. Equipements")
		fmt.Println("0. Retourner en arrière")

		var input int
		fmt.Scan(&input)

		switch input {
		case 1:
			if count, ok := p1.Inventory["Potion de soin"]; ok && count > 0 {
				TakePot(p1)
			} else {
				fmt.Println("Vous n'avez pas de Potion de soin dans votre inventaire.")
			}
		case 2:
			if count, ok := p1.Inventory["Potion de poison"]; ok && count > 0 {
				poisonPot(p1, m1)
			} else {
				fmt.Println("Vous n'avez pas de Potion de poison dans votre inventaire.")
			}
		case 3:
			if count, ok := p1.Inventory["Potion de mana"]; ok && count > 0 {
				ManaPot(p1)
			} else {
				fmt.Println("Vous n'avez pas de Potion de mana dans votre inventaire.")
			}
		case 4:
			p1.ShowSpells()

		case 5:
			fmt.Println("Voulez-vous 1.équiper ou 2.désequiper un equipement ? (3. retour)")
			var input int
			fmt.Scan(&input)
			switch input {
			case 1:
				fmt.Println("Que souhaitez-vous équiper ? 1. Casque / 2. Armure / 3. Pied / 4. retour ")
				var input int
				fmt.Scan(&input)
				switch input {
				case 1:
					p1.EquiperHead()
				case 2:
					p1.EquiperBody()
				case 3:
					p1.EquiperLeg()
				case 4:
					continue

				}
			case 2:
				fmt.Println("Que souhaitez-vous déquiper ? 1. Casque / 2. Armure / 3. Pied / 4. retour ")
				var input int
				fmt.Scan(&input)
				switch input {
				case 1:
					p1.DesequiperHead()
				case 2:
					p1.DesequiperBody()
				case 3:
					p1.DesequiperLeg()
				case 4:
					continue
				}
			case 3:
				continue
			}
		case 0:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

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

package game

import (
	"fmt"
)

func TakePot(p1 *Personnage) {
	var choice string
	fmt.Println("Voulez-vous prendre une potion ? (Oui/Non) ")
	fmt.Scan(&choice)

	if choice == "Oui" {
		if p1.Hp >= p1.HpMax {
			fmt.Println("Vos points de vie sont déjà pleins.")
		} else if count, ok := p1.Inventory["Potion de soin"]; ok && count > 0 {
			p1.Inventory["Potion de soin"]--
			p1.Hp += 50
			if p1.Hp > p1.HpMax {
				surplus := p1.Hp - p1.HpMax
				fmt.Printf("Cette potion vous rendra %d PV de trop, continuer (Oui/Non)?\n", surplus)
				fmt.Scan(&choice)
				if choice == "Oui" {
					p1.Hp = p1.HpMax
				} else if choice == "Non" {
					fmt.Printf("Vous n'avez pas pris une Potion de soin. Vous avez %d PV\n", p1.Hp)
					return
				}
			}
			fmt.Printf("Vous avez pris une Potion de soin. Vous avez %d PV\n", p1.Hp)
		} else {
			fmt.Println("Vous n'avez pas de Potion de soin dans votre inventaire.")
		}
	} else if choice == "Non" {
		return
	}
}

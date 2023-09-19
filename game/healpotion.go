package game

import (
	"fmt"
)

func TakePot(p1 *Personnage) {
	var choice int
	fmt.Println("Voulez-vous prendre une potion ? (1.Oui/2.Non) ")
	fmt.Scan(&choice)

	switch choice {

	case 1:
		if p1.Hp >= p1.HpMax {
			fmt.Println("Vos points de vie sont déjà pleins.")
		} else if count, ok := p1.Inventory["Potion de soin"]; ok && count > 0 {
			p1.Inventory["Potion de soin"]--
			p1.RemoveZeroValueItems()
			p1.Hp += 50
			if p1.Hp > p1.HpMax {
				surplus := p1.Hp - p1.HpMax
				fmt.Printf("Cette potion vous rendra %d PV de trop, continuer (1.Oui/2.Non)?\n", surplus)
				fmt.Scan(&choice)
				if choice == 1 {
					p1.Hp = p1.HpMax
				} else if choice == 2 {
					fmt.Printf("Vous n'avez pas pris une Potion de soin. Vous avez %d PV\n", p1.Hp)
					return
				}
			}
			fmt.Printf("Vous avez pris une Potion de soin. Vous avez %d PV\n", p1.Hp)
		} else {
			fmt.Println("Vous n'avez pas de Potion de soin dans votre inventaire.")
		}
	case 2:
		return
	}
}

package game

import (
	"fmt"
	"time"
)

func poisonPot(p1 *Personnage) {
	var choice string
	fmt.Println("Voulez-vous prendre une potion ? (Oui/Non) ")
	fmt.Scan(&choice)

	if choice == "Oui" {
		if count, ok := p1.Inventory["Potion de poison"]; ok && count > 0 {
			p1.Inventory["Potion de poison"]--
			fmt.Println("Vous avez utilisé une potion de poison.")
			duration := 3 * time.Second
			fmt.Printf("Vous perdez 10 PV chaque seconde pendant %s\n", duration)
			for start := time.Now(); time.Since(start) < duration; {
				p1.Hp -= 10
				if p1.Hp < 0 {
					p1.Dead()
					return
				}
				fmt.Printf("Vous avez %d PV restants sur %d PV\n", p1.Hp, p1.HpMax)
				time.Sleep(time.Second)
			}
		} else {
			fmt.Println("Vous n'avez pas de Potion de soin dans votre inventaire.")
		}

	} else if choice == "Non" {
		return
	}
	if p1.Hp <= 0 {
		p1.Dead()
	}

}

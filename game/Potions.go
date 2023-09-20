package game

import (
	"fmt"
	"time"
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
func poisonPot(p1 *Personnage, m1 *Monstre) {
	var choice int
	fmt.Println("Voulez-vous utiliser une Potion de poison ? (1.Oui/2.Non) ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		if count, ok := p1.Inventory["Potion de poison"]; ok && count > 0 {
			p1.Inventory["Potion de poison"]--
			p1.RemoveZeroValueItems()
			fmt.Println("Vous avez utilisé une potion de poison.")
			duration := 3 * time.Second
			fmt.Printf("Vous perdez 10 PV chaque seconde pendant %s\n", duration)
			for start := time.Now(); time.Since(start) < duration; {
				m1.Hp -= 10
				fmt.Printf("%d PV restants sur %d PV\n", m1.Hp, m1.HpMax)
				time.Sleep(time.Second)
			}
		} else {
			fmt.Println("Vous n'avez pas de Potion de poison dans votre inventaire.")
		}
	case 2:
		return
	}
}
func ManaPot(p1 *Personnage) {
	var choice int
	fmt.Println("Voulez-vous prendre une potion de mana ? (1.Oui/2.Non) ")
	fmt.Scan(&choice)

	switch choice {

	case 1:
		if p1.Mana >= p1.ManaMax {
			fmt.Println("Vos points de mana sont déjà pleins.")
		} else if count, ok := p1.Inventory["Potion de mana"]; ok && count > 0 {
			p1.Inventory["Potion de mana"]--
			p1.RemoveZeroValueItems()
			p1.Mana += 50
			if p1.Mana > p1.ManaMax {
				surplusM := p1.Mana - p1.ManaMax
				fmt.Printf("Cette potion vous rendra %d PV de trop, continuer (1.Oui/2.Non)?\n", surplusM)
				fmt.Scan(&choice)
				if choice == 1 {
					p1.Mana = p1.ManaMax
				} else if choice == 2 {
					fmt.Printf("Vous n'avez pas pris une Potion de soin. Vous avez %d PV\n", p1.Mana)
					return
				}
			}
			fmt.Printf("Vous avez pris une Potion de soin. Vous avez %d PV\n", p1.Mana)
		} else {
			fmt.Println("Vous n'avez pas de Potion de soin dans votre inventaire.")
		}
	case 2:
		return
	}
}

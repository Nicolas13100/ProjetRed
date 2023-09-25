package game

import (
	"fmt"
	"time"
)

func TakePot(P1 *Personnage) {
	var choice int
	fmt.Println("Voulez-vous prendre une potion de soin? (1.Oui/2.Non) ")
	fmt.Scan(&choice)

	switch choice {

	case 1:
		if P1.Hp >= P1.HpMax {
			fmt.Println("Vos points de vie sont déjà pleins.")
		} else if count, ok := P1.Inventory["Potion de soin"]; ok && count > 0 {
			P1.Inventory["Potion de soin"]--
			P1.RemoveZeroValueItems()
			P1.Hp += 50
			if P1.Hp > P1.HpMax {
				surplus := P1.Hp - P1.HpMax
				fmt.Printf("Cette potion vous rendra %d PV de trop, continuer (1.Oui/2.Non)?\n", surplus)
				fmt.Scan(&choice)
				if choice == 1 {
					P1.Hp = P1.HpMax
				} else if choice == 2 {
					fmt.Printf("Vous n'avez pas pris une Potion de soin. Vous avez %d PV\n", P1.Hp)
					return
				}
			}
			fmt.Printf("Vous avez pris une Potion de soin. Vous avez %d PV\n", P1.Hp)
		} else {
			fmt.Println("Vous n'avez pas de Potion de soin dans votre inventaire.")
		}
	case 2:
		return
	}
}

func poisonPot(P1 *Personnage) {
	var choice int
	fmt.Println("Voulez-vous utiliser une Potion de poison ? (1.Oui/2.Non) ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		if count, ok := P1.Inventory["Potion de poison"]; ok && count > 0 && Monstre1.Hp > 0 {
			P1.Inventory["Potion de poison"]--
			P1.RemoveZeroValueItems()
			fmt.Println("Vous avez utilisé une potion de poison.")
			duration := 3 * time.Second
			fmt.Printf("L'ennemi touché perd 10 PV chaque seconde pendant %s\n", duration)
			for start := time.Now(); time.Since(start) < duration; {
				Monstre1.Hp -= 10
				fmt.Printf("%d PV restants sur %d PV\n", Monstre1.Hp, Monstre1.HpMax)
				time.Sleep(time.Second)
			}
		} else {
			fmt.Println("Vous n'avez pas de Potion de poison dans votre inventaire.")
		}
	case 2:
		return
	}
}

func ManaPot(P1 *Personnage) {
	var choice int
	fmt.Println("Voulez-vous prendre une potion de mana ? (1.Oui/2.Non) ")
	fmt.Scan(&choice)

	switch choice {

	case 1:
		if P1.Mana >= P1.ManaMax {
			fmt.Println("Vos points de mana sont déjà pleins.")
		} else if count, ok := P1.Inventory["Potion de mana"]; ok && count > 0 {
			P1.Inventory["Potion de mana"]--
			P1.RemoveZeroValueItems()
			P1.Mana += 50
			if P1.Mana > P1.ManaMax {
				surplusM := P1.Mana - P1.ManaMax
				fmt.Printf("Cette potion vous rendra %d PV de trop, continuer (1.Oui/2.Non)?\n", surplusM)
				fmt.Scan(&choice)
				if choice == 1 {
					P1.Mana = P1.ManaMax
				} else if choice == 2 {
					fmt.Printf("Vous n'avez pas pris une Potion de soin. Vous avez %d PV\n", P1.Mana)
					return
				}
			}
			fmt.Printf("Vous avez pris une Potion de soin. Vous avez %d PV\n", P1.Mana)
		} else {
			fmt.Println("Vous n'avez pas de Potion de soin dans votre inventaire.")
		}
	case 2:
		return
	}
}

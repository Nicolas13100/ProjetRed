package game

import (
	"fmt"
	"time"
)

func poisonPot(p1 *Personnage, m1 *Monstre) {
	var choice int
	fmt.Println("Voulez-vous utiliser une potion ? (1.Oui/2.Non) ")
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

package game

import "fmt"

func Campain(P1 *Personnage) {
	if P1.Level < 5 {
		fmt.Println("Vous entrez dans une fôret remplie de monstres et d'animaux sauvages, souhaitez-vous combattre ou faire demi-tour ? (1 : Combattre / 0 : Partir)")
		var choix int
		fmt.Scan(&choix)
		switch choix {
		case 1:
			fmt.Println("Quel monstre ou animal souhaitez vous affronter ? (1 : Corbeau / 2 : Sanglier / 3 : Loup / 4 : Troll)")
			var mechant int
			fmt.Scan(&mechant)
			switch mechant {
			case 1:
				mechant := Corbeau
				Fight(P1, &mechant)
				return
			case 2:
				mechant := Sanglier
				Fight(P1, &mechant)
				return
			case 3:
				mechant := Loup
				Fight(P1, &mechant)
				return
			case 4:
				mechant := Troll
				Fight(P1, &mechant)
				return
			}
		case 2:
			return
		}
	}
}

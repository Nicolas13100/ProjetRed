package game

import "fmt"

func Campain(P1 *Personnage) {
	if P1.Level < 5 {
		fmt.Println("Vous entrez un fôret remplie de monstres et d'animaux sauvages, continuer ? (1 : Oui / 2 : Non)")
		var choix int
		fmt.Scan(&choix)
		switch choix {
		case 1:
			fmt.Println("Quel monstre ou animal souhaitez vous affronter ? (1 : Corbeau / 2 : Sanglier / 3 : Loup / 4 : Troll)")
			var mechant int
			fmt.Scan(&mechant)
			switch mechant {
			case 1:
				Fight(P1, &Corbeau)
				return
			case 2:
				Fight(P1, &Sanglier)
			case 3:
				Fight(P1, &Loup)
			case 4:
				Fight(P1, &Troll)
			}
		case 2:
			return
		}
	} else if P1.Level >= 5 {
		fmt.Println("Vous avez les capacités nécessaires pour aller dans la montagne, continuer (1 : Oui / 2 : Non)")
		var choix int
		fmt.Scan(&choix)
		switch choix {
		case 1:
			fmt.Println("Quel monstre ou animal souhaitez vous affronter ? (1 : Maitre corbeau / 2 : Sanglier enragé  / 3 : Loup Garou / 4 : Troll'ourd)")
			var mechant int
			fmt.Scan(&mechant)
			switch mechant {
			case 1:
				Fight(P1, &Corbeau)
				return
			case 2:
				Fight(P1, &Sanglier)
			case 3:
				Fight(P1, &Loup)
			case 4:
				Fight(P1, &Troll)
			}
		}
	} else {
		return
	}

}

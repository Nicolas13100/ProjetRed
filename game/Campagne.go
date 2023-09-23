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
				mechant := NewCorbeau()
				Fight(P1, &mechant)
				return
			case 2:
				mechant := NewSanglier()
				Fight(P1, &mechant)
			case 3:
				mechant := NewLoup()
				Fight(P1, &mechant)
			case 4:
				mechant := NewTroll()
				Fight(P1, &mechant)
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
			fmt.Println("Quel monstre ou animal souhaitez vous affronter ? (1 : Maitre Corbeau / 2 : Sanglier enragé  / 3 : Loup ténébreux / 4 : Troll des montagnes)")
			var mechant int
			fmt.Scan(&mechant)
			switch mechant {
			case 1:
				fmt.Println("Vous avez choisi d'affronter le Maître Corbeau ! ")
				mechant := NewCorbeau2()
				Fight(P1, &mechant)
				return
			case 2:
				fmt.Println("Vous avez choisi d'affronter le Sanglier enragé ! ")
				mechant := NewSanglier2()
				Fight(P1, &mechant)
				return
			case 3:
				fmt.Println("Vous avez choisi d'affronter le Loup ténébreux ! ")
				mechant := NewLoup2()
				Fight(P1, &mechant)
				return
			case 4:
				fmt.Println("Vous avez choisi d'affronter le Troll des montagnes! ")
				mechant := NewTroll2()
				Fight(P1, &mechant)
				return
			}
		}
	} else {
		return
	}

}

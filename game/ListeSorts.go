package game

import (
	"fmt"
)

func Spell1(P1 *Personnage) {
	var choice int
	fmt.Println("Voulez-vous utiliser le sort Coup de poing vénère ? (1.Oui/2.Non) ")
	fmt.Scan(&choice)

	switch choice {

	case 1:
		if P1.Mana <= 0 {
			fmt.Println("Vous ne pouvez pas utiliser de sorts, vous n'avez pas de mana.")
		} else if count, ok := P1.Skills["Coup de point vénère"]; ok && count > 0 {
			Monstre1.Hp -= 8
		}
	case 2:
		{
			fmt.Printf("Vous n'avez pas utilisé de sort.")
			return
		}
	}
}

func Spell2(P1 *Personnage) {
	var choice int
	fmt.Println("Voulez-vous utiliser le sort boule de feu suprême ? (1.Oui/2.Non) ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		if P1.Mana <= 0 {
			fmt.Println("Vous ne pouvez pas utiliser de sorts, vous n'avez pas de mana.")
		} else if count, ok := P1.Skills["boule de feu suprême"]; ok && count > 0 {
			Monstre1.Hp -= 18
		}
	case 2:
		{
			fmt.Printf("Vous n'avez pas utilisé de sort.")
			return
		}
	}
}

func Spell3(P1 *Personnage) {
	var choice int
	fmt.Println("Voulez-vous utiliser le sort Hakaï ? (1.Oui/2.Non) ")
	fmt.Scan(&choice)

	switch choice {

	case 1:
		if P1.Mana <= 0 {
			fmt.Println("Vous ne pouvez pas utiliser de sorts, vous n'avez pas de mana.")
		} else if count, ok := P1.Skills["Hakaï"]; ok && count > 0 {
			Monstre1.Hp -= 9999
		}
	case 2:
		{
			fmt.Printf("Vous n'avez pas utilisé de sort.")
			return
		}
	}
}

func (P1 Personnage) ShowSpells() {
	fmt.Println("Souhaitez-vous continuer ou retourner en arriere (1 : Oui / 2 : Non)")
	var choix int
	fmt.Scan(&choix)
	switch choix {
	case 1:

		fmt.Println("Sort Disponible:")
		for key, value := range P1.Skills {
			fmt.Printf("%s. %d \n", key, value)
		}
	case 2:

		fmt.Println("Aucun sort disponible")
		return
	}
}

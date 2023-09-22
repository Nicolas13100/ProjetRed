package game

import (
	"fmt"
)

func FightSpell(P1 *Personnage, s1 Spell, s2 Spell, s3 Spell) {
	fmt.Println("Choisissez quel sort vous souhaitez utiliser :")
	fmt.Printf("Vous possédez %s comme sorts")
	var spell int
	fmt.Scan(&spell)

	switch spell {

	case 1:
		fmt.Printf("Vous avez utilisé %s et infligé %d dégâts à l'adversaire", s1.name, s1.damage)
		fmt.Printf("Vous avez utilisé %d de Mana", s1.useMana)
	case 2:
		fmt.Printf("Vous avez utilisé %s et infligé %d dégâts à l'adversaire", s2.name, s2.damage)
		fmt.Printf("Vous avez utilisé %d de Mana", s2.useMana)
	case 3:
		fmt.Printf("Vous avez utilisé %s et infligé %d dégâts à l'adversaire", s3.name, s3.damage)
		fmt.Printf("Vous avez utilisé %d de Mana", s3.useMana)
	}

}

func (P1 Personnage) ShowSpells() {
	fmt.Println("Souhaitez-vous continuer ou retourner en arriere (1 : Oui / 2 : Non)")
	var choix int
	fmt.Scan(&choix)
	switch choix {
	case 1:

		fmt.Println("Sort Disponible:")
		for _, spell := range P1.Skills {
			fmt.Printf("%d. %s\n", spell)
		}
	case 2:

		fmt.Println("Aucun sort disponible")
		return
	}
}

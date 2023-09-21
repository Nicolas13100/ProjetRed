package game

import (
	"fmt"
)

func InitSpell(P1 *Personnage, Spells Spell) string {
	s1 := Spell{
		"Boule de Feu Suprême", 18, 6,
	}

	s2 := Spell{
		"Coup de poing vénère", 8, 6,
	}
	s3 := Spell{
		"Hakaï", 999, 99,
	}

	spell1 := "Livre de sort : " + s1.name
	for _, skill := range P1.Spells {
		if skill == spell1 {
			fmt.Printf("Vous avez déjà appris le sort %s \n", s1.name)
			return spell1
		}
	}
	P1.Spells = append(P1.Spells, spell1)
	fmt.Println("Sort appris :", spell1)

	spell2 := "Livre de sort : " + s2.name
	for _, skill := range P1.Spells {
		if skill == spell2 {
			fmt.Printf("Vous avez déjà appris le %s \n", s2.name)
			return spell2
		}
	}
	P1.Spells = append(P1.Spells, spell2)
	fmt.Println("Sort appris :", spell2)

	spellD := "Livre de sort : " + s3.name
	for _, skill := range P1.Spells {
		if skill == spellD {
			fmt.Printf("Vous avez déjà appris le %s \n", s3.name)
			return spellD
		}
	}
	P1.Spells = append(P1.Spells, spellD)
	fmt.Println("Sort appris :", spellD)

	return (Spells.name)
}

func FightSpell(P1 *Personnage, s1 Spell, s2 Spell, s3 Spell) {
	fmt.Println("Choisissez quel sort vous souhaitez utiliser :")
	fmt.Printf("Vous possédez %s comme sorts", P1.Skills[0:])
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
	fmt.Println("Sort Disponible:")
	for i, spell := range P1.Spells {
		fmt.Printf("%d. %s\n", i+1, spell)
	}
}

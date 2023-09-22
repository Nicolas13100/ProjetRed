package game

import (
	"fmt"
)
S1 := Spell1
func Spell1(P1 *Personnage) {
	S1 := Spell{
		"Coup de poing vénère", 8, 6,
	}
	return
}

func Spell2(P1 *Personnage, Spells Spell) {
	s2 := Spell{
		"Boule de Feu Suprême", 18, 6,
	}

}

func Spell3(P1 *Personnage, Spells Spell) {
	s3 := Spell{
		"Hakaï", 999, 99,
	}
}

func InitSpell(P1 *Personnage, Spells Spell) string {
	Spell1()

	spell1 := "Livre de sort : " + s1.name
	for _, skill := range P1.Skills {
		if skill == spell1 {
			fmt.Printf("Vous avez déjà appris le sort %s \n", s1.name)
			return spell1
		}
	}
	P1.Spells = append(P1.Skills, spell1)
	fmt.Println("Sort appris :", spell1)

	spell2 := "Livre de sort : " + s2.name
	for _, skill := range P1.Skills {
		if skill == spell2 {
			fmt.Printf("Vous avez déjà appris le %s \n", s2.name)
			return spell2
		}
	}
	P1.Spells = append(P1.Skills, spell2)
	fmt.Println("Sort appris :", spell2)

	spellD := "Livre de sort : " + s3.name
	for _, skill := range P1.Skills {
		if skill == spellD {
			fmt.Printf("Vous avez déjà appris le %s \n", s3.name)
			return spellD
		}
	}
	P1.Spells = append(P1.Skills, spellD)
	fmt.Println("Sort appris :", spellD)

	return (Spells.name)
}

func FightSpell(P1 *Personnage, s1 Spell, s2 Spell, s3 Spell) {
	fmt.Println("Choisissez quel sort vous souhaitez utiliser :")
	fmt.Printf("Vous possédez 1.%s", P1.Skills[0:])
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
<<<<<<< HEAD
		for i, spell := range P1.Skills {
			fmt.Printf("%d. %s\n", i+1, spell)
=======
		for _, spell := range P1.Skills {
			fmt.Printf("%d. %s\n", spell)
>>>>>>> 178222bacf674eafa0ad8284ac9de0a21da04ddf
		}
	case 2:

		fmt.Println("Aucun sort disponible")
		return
	}
}

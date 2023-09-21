package game

import "fmt"

func SpellBook(P1 *Personnage, Spell *Spells) {
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
			return
		}
	}
	P1.Spells = append(P1.Spells, spell1)
	fmt.Println("Sort appris :", spell1)

	spell2 := "Livre de sort : " + s2.name
	for _, skill := range P1.Spells {
		if skill == spell2 {
			fmt.Printf("Vous avez déjà appris le %s \n", s2.name)
			return
		}
	}
	P1.Spells = append(P1.Spells, spell2)
	fmt.Println("Sort appris :", spell2)

	spellD := "Livre de sort : " + s3.name
	for _, skill := range P1.Spells {
		if skill == spellD {
			fmt.Printf("Vous avez déjà appris le %s \n", s3.name)
			return
		}
	}
	P1.Spells = append(P1.Spells, spellD)
	fmt.Println("Sort appris :", spellD)

}

func (P1 Personnage) ShowSpells() {
	fmt.Println("Sort Disponible:")
	for i, spell := range P1.Spells {
		fmt.Printf("%d. %s\n", i+1, spell)
	}
}

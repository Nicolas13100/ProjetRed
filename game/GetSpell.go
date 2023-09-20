package game

import "fmt"

func SpellBook(p1 *Personnage, s *Spell) {
	spell1 := "Livre de sort : Boule de feu"
	p1.Inventory[spell1]--
	p1.RemoveZeroValueItems()
	for _, skill := range p1.Skills {
		if skill == spell1 {
			fmt.Println("Vous avez déjà appris le sort Boule de feu")
			return
		}
	}
	p1.Skills = append(p1.Skills, spell1)
	fmt.Println("Sort appris :", spell1)

	spell2 := "Livre de sort : Coup de poing"
	p1.Inventory[spell2]--
	p1.RemoveZeroValueItems()
	for _, skill := range p1.Skills {
		if skill == spell2 {
			fmt.Println("Vous avez déjà appris le Coup de poing")
			return
		}
	}
	p1.Skills = append(p1.Skills, spell2)
	fmt.Println("Sort appris :", spell2)
}

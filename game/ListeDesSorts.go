package game

import "fmt"

func (p1 *Personnage) SpellBook() {
	spell := "Boule de feu"
	for _, skill := range p1.Skills {
		if skill == spell {
			fmt.Println("Vous avez déjà appris le sort Boule de feu")
			return
		}
	}
	p1.Skills = append(p1.Skills, spell)
	fmt.Println("Sort appris :", spell)
}

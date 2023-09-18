package game

import "fmt"

func (p1 *Personnage) SpellBook() {
	spell1 := "Boule de feu"
	for _, skill := range p1.Skills {
		if skill == spell1 {
			fmt.Println("Vous avez déjà appris le sort Boule de feu")
			return
		}
	}
	p1.Skills = append(p1.Skills, spell1)
	fmt.Println("Sort appris :", spell1)
}

package game

import "fmt"

type Spell struct {
	name    string
	damage  int
	useMana int
}

func InitSpells() *Spell {
	fmt.Println("Choisissez quel sort vous souhaitez utiliser :")
	fmt.Println("Sort n°1 : Boule de Feu Suprême (1)")
	fmt.Println("Sort n°2 : Coup de poing vénère (2)")
	fmt.Println("Sort n°3 : Hakaï (3)")
	var spell int
	fmt.Scan(&spell)
	switch spell {
	case 1:
		s1 := &Spell{
			name:    "Boule de Feu Suprême",
			damage:  18,
			useMana: 6,
		}
		return s1

	case 2:
		s2 := &Spell{
			name:    "Coup de poing vénère",
			damage:  10,
			useMana: 3,
		}
		return s2

	case 3:
		s3 := &Spell{
			name:    "Hakaï",
			damage:  9999,
			useMana: 99,
		}
		return s3

	default:
		fmt.Println("Sort invalide. Aucun sort sélectionné.")
		return nil
	}

}

func SpellBook(P1 *Personnage, Spell *Spell) {

	spell1 := "Livre de sort : " + Spell.name
	for _, skill := range P1.Spells {
		if skill == spell1 {
			fmt.Printf("Vous avez déjà appris le sort %s \n", Spell.name)
			return
		}
	}
	P1.Spells = append(P1.Spells, spell1)
	fmt.Println("Sort appris :", spell1)

	spell2 := "Livre de sort : " + Spell.name
	for _, skill := range P1.Spells {
		if skill == spell2 {
			fmt.Printf("Vous avez déjà appris le %s \n", Spell.name)
			return
		}
	}
	P1.Spells = append(P1.Spells, spell2)
	fmt.Println("Sort appris :", spell2)

	spellD := "Livre de sort : " + Spell.name
	for _, skill := range P1.Spells {
		if skill == spellD {
			fmt.Printf("Vous avez déjà appris le %s \n", Spell.name)
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

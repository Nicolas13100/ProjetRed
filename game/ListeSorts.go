package game

type Spell struct {
	name    string
	damage  int
	useMana int
}

func InitSpells() []*Spell {
	spells := []*Spell{
		{
			name:    "Boule de Feu Suprême",
			damage:  18,
			useMana: 6,
		},

		{
			name:    "Coup de poing vénère",
			damage:  10,
			useMana: 3,
		},
		{
			name:    "Hakaï",
			damage:  9999,
			useMana: 99,
		},
	}
	return spells
}

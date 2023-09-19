package game

type Monstre struct {
	name    string
	HpMax   int
	Hp      int
	Atk     int
	Defense int
}

func InitGoblin() *Monstre {

	m1 := &Monstre{
		name:    "Goblin",
		HpMax:   40,
		Hp:      40,
		Atk:     5,
		Defense: 0,
	}
	return m1
}

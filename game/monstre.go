package game

type Monstre struct {
	name       string
	HpMax      int
	Hp         int
	Atk        int
	Defense    int
	Initiative int
}

func InitGoblin() *Monstre {

	m1 := &Monstre{
		name:       "Goblin",
		HpMax:      40,
		Hp:         40,
		Atk:        5,
		Defense:    0,
		Initiative: 10,
	}
	return m1
}

func GoblinPattern(m1 *Monstre) bool {
	if GetTours()%3 == 0 {
		m1.Atk = m1.Atk * 2
		return true
	}
	return false
}

package game

func InitGoblin() *Monstre {

	Monstre1 := &Monstre{
		name:       "Goblin",
		HpMax:      40,
		Hp:         40,
		XpDrop:     110,
		Atk:        5,
		Defense:    0,
		Initiative: 10,
	}
	return Monstre1
}

func GoblinPattern(m1 *Monstre) bool {
	if GetTours()%3 == 0 || GetTours1()%3 == 0 {
		m1.Atk = m1.Atk * 2
		return true
	}
	return false
}

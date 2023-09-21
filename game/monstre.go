package game

func InitGoblin() *Monstre {

	Monstre1 := &Monstre{
		Name:       "Goblin",
		HpMax:      40,
		Hp:         40,
		XpDrop:     110,
		Atk:        5,
		Defense:    0,
		Initiative: 10,
	}
	return Monstre1
}

func GoblinPattern(Monstre1 *Monstre) bool {
	if GetTours()%3 == 0 || GetTours1()%3 == 0 {
		Monstre1.Atk = Monstre1.Atk * 2
		return true
	}
	return false
}

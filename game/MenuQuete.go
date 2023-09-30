package game

import "fmt"

func (p *Personnage) MenuQuetes() {
	fmt.Println("Bienvenue à la guilde des aventuriers, que puis-je pour vous ? \n(1.Chek rank\n2.Prendre,rendre une quete\n0.Rien merci)")
	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
	case 2:
	case 0:
	}
}

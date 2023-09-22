package game

import "fmt"

func Campain(P1 Personnage) {
	var choix int
	fmt.Println("Vous entrez un fôret remplie de monstre, continuer ? (1 : Oui / 2 : Non)")
	fmt.Scan(&choix)
	switch choix {
	case 1:
		fmt.Println("Cherchez le monstre que vous ")
	}
}

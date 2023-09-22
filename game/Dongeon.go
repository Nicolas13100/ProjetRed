package game

import "fmt"

func Dongeon(P1 *Personnage, Monstre1 *Monstre) {
	etagedongeon := 0
	if etagedongeon < 1 {
		fmt.Println("Vous entrez dans une forêt dense.")
		Fight(P1, Monstre1)
	}
}

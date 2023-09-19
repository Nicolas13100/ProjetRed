package game

import "fmt"

var Tours int

func TrainFight(p1 *Personnage, m1 *Monstre) {
	fmt.Println("Bienvenue dans le tutoriel")
	fmt.Println("Nous allons vous apprendre les bases du combat en tour par tour")
	fmt.Println("Vous allez vous battre contre un gobin basic")
	fmt.Println("Bonne chance")

	for p1.Hp > 0 && m1.Hp > 0 {
		fmt.Printf("Vous etes au tour %d", Tours)
		// Tour du Joueur 1
		attaqueJoueur1 := p1.Atk - m1.Defense
		if attaqueJoueur1 < 0 {
			attaqueJoueur1 = 0
		}
		m1.Hp -= attaqueJoueur1

		fmt.Printf("%s attaque %s et lui inflige %d points de dégâts.\n", p1.Name, m1.name, attaqueJoueur1)

		// Vérifier si le Joueur 2 est toujours en vie
		if m1.Hp <= 0 {
			fmt.Printf("%s est vaincu!\n", m1.name)
			break
		}

		// Tour du Joueur 2
		attaqueJoueur2 := m1.Atk - p1.Defense
		if attaqueJoueur2 < 0 {
			attaqueJoueur2 = 0
		}
		p1.Hp -= attaqueJoueur2

		fmt.Printf("%s attaque %s et lui inflige %d points de dégâts.\n", m1.name, p1.Name, attaqueJoueur2)
		Tours++

		// Vérifier si le Joueur 1 est toujours en vie
		if p1.Hp <= 0 {
			fmt.Printf("%s est vaincu!\n", p1.Name)
			break
		}
	}
}

func GetTours() int {
	return Tours
}

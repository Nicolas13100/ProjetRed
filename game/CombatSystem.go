package game

import "fmt"

var Tours int

func Tutorial() {
	fmt.Println("Bienvenue dans le tutoriel")
	fmt.Println("Nous allons vous apprendre les bases du combat en tour par tour")
	fmt.Println("Vous allez vous battre contre un golbin d'entrainement")
	fmt.Println("Bonne chance")
	trainFight()

}

func trainFight(p1 *Personnage, m1 *Monstre) {
	for p1.Hp > 0 && m1.Hp > 0 {
		fmt.Printf("Vous etes au tour %d", Tours)
		// Tour du Joueur 1
		charTurn(p1, m1)
		attaqueJoueur := p1.Atk - m1.Defense
		if attaqueJoueur < 0 {
			attaqueJoueur = 0
		}
		m1.Hp -= attaqueJoueur

		fmt.Printf("%s attaque %s et lui inflige %d points de dégâts.\n", p1.Name, m1.name, attaqueJoueur)

		// Vérifier si le Joueur 2 est toujours en vie
		if m1.Hp <= 0 {
			fmt.Printf("%s est vaincu!\n", m1.name)
			break
		}

		// Tour du Monstre
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

func charTurn(p1 *Personnage, m1 *Monstre) {
	fmt.Println("C'est votre tour, que voulez vous faire ?")
	fmt.Println("1 : Attaquer")
	fmt.Println("2 : Ouvrir l'inventaire")

	var choice int
	switch choice {
	case 1:

	case 2:
		p1.AccessInventory()
	}
}

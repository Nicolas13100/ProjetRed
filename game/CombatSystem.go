package game

import "fmt"

var Tours int

func Tutorial(p1 *Personnage, m1 *Monstre) {
	fmt.Println("Bienvenue dans le tutoriel")
	fmt.Println("Nous allons vous apprendre les bases du combat en tour par tour")
	fmt.Println("Vous allez vous battre contre un goblin d'entrainement")
	fmt.Println("Bonne chance")
	Tours = 1
	TrainFight(p1, m1)

}

func TrainFight(p1 *Personnage, m1 *Monstre) {
	for (p1.Hp > 0 && m1.Hp > 0 && Tours == 1 && p1.Initiative >= m1.Initiative) || (Tours > 1 && p1.Hp > 0 && m1.Hp > 0) {
		fmt.Printf("Vous etes au tour %d\n", Tours)
		goblinPatternActivated := GoblinPattern(m1)

		// Tour du Joueur 1
		charTurn(p1, m1)
		attaqueJoueur := p1.Atk - m1.Defense
		if attaqueJoueur < 0 {
			attaqueJoueur = 0
		}
		m1.Hp -= attaqueJoueur

		fmt.Printf("%s attaque %s et lui inflige %d points de dégâts.\n", p1.Name, m1.name, attaqueJoueur)

		// Vérifier si le Monstre est toujours en vie
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
		// Apply goblin's attack pattern if it was activated
		if goblinPatternActivated {
			// Reset goblin's attack back to its original value
			m1.Atk /= 2
		}

		// Vérifier si le Joueur est toujours en vie
		if p1.Hp <= 0 {
			fmt.Printf("%s est vaincu!\n", p1.Name)
			p1.Dead()
			break
		}
	}
	for (p1.Hp > 0 && m1.Hp > 0 && Tours == 1 && p1.Initiative < m1.Initiative) || (Tours > 1 && p1.Hp > 0 && m1.Hp > 0) {
		fmt.Printf("Vous etes au tour %d\n", Tours)
		goblinPatternActivated := GoblinPattern(m1)

		// Tour du Monstre
		attaqueJoueur2 := m1.Atk - p1.Defense
		if attaqueJoueur2 < 0 {
			attaqueJoueur2 = 0
		}
		p1.Hp -= attaqueJoueur2

		fmt.Printf("%s attaque %s et lui inflige %d points de dégâts.\n", m1.name, p1.Name, attaqueJoueur2)
		Tours++
		// Apply goblin's attack pattern if it was activated
		if goblinPatternActivated {
			// Reset goblin's attack back to its original value
			m1.Atk /= 2
		}

		// Vérifier si le Joueur est toujours en vie
		if p1.Hp <= 0 {
			fmt.Printf("%s est vaincu!\n", p1.Name)
			p1.Dead()
			break
		} // Tour du Joueur 1
		charTurn(p1, m1)
		attaqueJoueur := p1.Atk - m1.Defense
		if attaqueJoueur < 0 {
			attaqueJoueur = 0
		}
		m1.Hp -= attaqueJoueur

		fmt.Printf("%s attaque %s et lui inflige %d points de dégâts.\n", p1.Name, m1.name, attaqueJoueur)

		// Vérifier si le Monstre est toujours en vie
		if m1.Hp <= 0 {
			fmt.Printf("%s est vaincu!\n", m1.name)
			break
		}
	}
}

func GetTours() int {
	return Tours + 1
}

func charTurn(p1 *Personnage, m1 *Monstre) {
	fmt.Println("C'est votre tour, que voulez vous faire ?")
	fmt.Print("Vos point de vie :")
	fmt.Print(p1.Hp)
	fmt.Print(" / ")
	fmt.Print(p1.HpMax)
	fmt.Println("HP")
	fmt.Print("Goblin:")
	fmt.Print(m1.Hp)
	fmt.Print(" / ")
	fmt.Print(m1.HpMax)
	fmt.Println("HP")
	fmt.Println("1 : Attaquer")
	fmt.Println("2 : Ouvrir l'inventaire")
	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("Attaque basic")
	case 2:
		p1.AccessInventory(m1)
	}
}

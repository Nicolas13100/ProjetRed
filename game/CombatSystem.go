package game

import "fmt"

var Tours int

func Tutorial(p1 *Personnage, m1 *Monstre, spells *Spell) {
	fmt.Println("Bienvenue dans le tutoriel")
	fmt.Println("Nous allons vous apprendre les bases du combat en tour par tour")
	fmt.Println("Vous allez vous battre contre un goblin d'entrainement")
	fmt.Println("Bonne chance")
	Tours = 1
	TrainFight(p1, m1, spells)

}

func TrainFight(p1 *Personnage, m1 *Monstre, spells *Spell) {
	for (p1.Hp > 0 && m1.Hp > 0 && Tours == 1 && p1.Initiative >= m1.Initiative) || (Tours > 1 && p1.Hp > 0 && m1.Hp > 0) {
		fmt.Printf("Vous etes au tour %d\n", Tours)
		goblinPatternActivated := GoblinPattern(m1)

		// Tour du Joueur 1
		charTurn(p1, m1, spells)
		attaqueJoueur := p1.Atk - m1.Defense
		if attaqueJoueur < 0 {
			attaqueJoueur = 0
		}
		m1.Hp -= attaqueJoueur

		fmt.Printf("%s attaque %s et lui inflige %d points de dégâts.\n", p1.Name, m1.name, attaqueJoueur)

		// Vérifier si le Monstre est toujours en vie
		if m1.Hp <= 0 {
			fmt.Printf("%s est vaincu!\n", m1.name)
			GagnerXp(p1, m1)
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
		charTurn(p1, m1, spells)
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

func charTurn(p1 *Personnage, m1 *Monstre, spells *Spell) {
	fmt.Printf("Vous avez %d Pv / %d Pv \n", p1.Hp, p1.HpMax)
	fmt.Printf("Il reste %d Pv / %d Pv à l'ennemi \n", m1.Hp, m1.HpMax)

	fmt.Println("C'est votre tour, que voulez vous faire ?")
	fmt.Println("1 : Attaquer")
	fmt.Println("2 : Ouvrir l'inventaire")
	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("Attaque basic")
	case 2:
		p1.AccessInventory(m1, spells)
	}
}

func GagnerXp(p1 *Personnage, m1 *Monstre) {
	p1.Xp += m1.XpDrop
	fmt.Printf("Vous avez gagné %d points d'XP \n", m1.XpDrop)
	fmt.Printf("Vous avez %d points d'XP / %d points d'XP avant le prochain niveau \n", p1.Xp, p1.XpMax)
	if p1.Xp >= p1.XpMax {
		p1.Level += 1
		surplusxp := p1.Xp - p1.XpMax
		p1.Xp = 0
		p1.XpMax += 10
		p1.Xp += surplusxp
		fmt.Printf("Vous avez gagné un niveau et êtes au niveau : %d \n", p1.Level)
		fmt.Printf("Vous avez %d points d'XP / %d points d'XP avant le prochain niveau \n", p1.Xp, p1.XpMax)
	}
}

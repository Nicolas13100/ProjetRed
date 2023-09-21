package game

import "fmt"

var Tours1 int

func Fight(p1 *Personnage, m1 *Monstre, spells *Spell) bool {
	ClearConsole()
	CombatStarted := false
	Tours1 = 1

	var choice2 int
	fmt.Println("Voulez-vous commencer un entrainement ? (1 : Oui / 2 : Non)")
	fmt.Scan(&choice2)

	switch choice2 {
	case 1:
		fmt.Println("Début du combat d'entrainement")
		for (p1.Hp > 0 && m1.Hp > 0 && Tours1 == 1 && p1.Initiative >= m1.Initiative) || (Tours1 > 1 && p1.Hp > 0 && m1.Hp > 0) {
			CombatStarted = true
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
			fmt.Println("-----------------------------------------------------------------------------")

			// Vérifier si le Joueur est toujours en vie
			if p1.Hp <= 0 {
				fmt.Printf("%s est vaincu!\n", p1.Name)
				p1.Dead()
				break
			}
		}
		for (p1.Hp > 0 && m1.Hp > 0 && Tours1 == 1 && p1.Initiative < m1.Initiative) || (Tours1 > 1 && p1.Hp > 0 && m1.Hp > 0) {
			CombatStarted = true
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
			fmt.Println("-----------------------------------------------------------------------------")

			// Vérifier si le Joueur est toujours en vie
			if p1.Hp <= 0 {
				fmt.Printf("%s est vaincu!\n", p1.Name)
				p1.Dead()
				break
			}

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

		}

	case 2:
		fmt.Println("Retour au menu principal")
	}
	return CombatStarted
}
func GetTours1() int {
	return Tours1 + 1
}

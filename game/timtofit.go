package game

import "fmt"

var Tours1 int

func Fight(Personnag1 *Personnage, Monster1 *Monstre, spells *Spell) bool {
	ClearConsole()
	CombatStarted := false
	Tours1 = 1

	var choice2 int
	fmt.Println("Voulez-vous commencer un entrainement ? (1 : Oui / 2 : Non)")
	fmt.Scan(&choice2)

	switch choice2 {
	case 1:
		fmt.Println("Début du combat d'entrainement")
		for (Personnag1.Hp > 0 && Monster1.Hp > 0 && Tours1 == 1 && Personnag1.Initiative >= Monster1.Initiative) || (Tours1 > 1 && Personnag1.Hp > 0 && Monster1.Hp > 0) {
			CombatStarted = true
			fmt.Printf("Vous etes au tour %d\n", Tours)
			goblinPatternActivated := GoblinPattern(Monster1)
			// Tour du Joueur 1
			charTurn(Personnag1, Monster1, spells)
			attaqueJoueur := Personnag1.Atk - Monster1.Defense
			if attaqueJoueur < 0 {
				attaqueJoueur = 0
			}
			Monster1.Hp -= attaqueJoueur

			fmt.Printf("%s attaque %s et lui inflige %d points de dégâts.\n", Personnag1.Name, Monster1.Name, attaqueJoueur)

			// Vérifier si le Monstre est toujours en vie
			if Monster1.Hp <= 0 {
				fmt.Printf("%s est vaincu!\n", Monster1.Name)
				GagnerXp(Personnag1, Monster1)
				break
			}

			// Tour du Monstre
			attaqueJoueur2 := Monster1.Atk - Personnag1.Defense
			if attaqueJoueur2 < 0 {
				attaqueJoueur2 = 0
			}
			Personnag1.Hp -= attaqueJoueur2

			fmt.Printf("%s attaque %s et lui inflige %d points de dégâts.\n", Monster1.Name, Personnag1.Name, attaqueJoueur2)
			Tours++
			// Apply goblin's attack pattern if it was activated
			if goblinPatternActivated {
				// Reset goblin's attack back to its original value
				Monster1.Atk /= 2
			}
			fmt.Println("-----------------------------------------------------------------------------")

			// Vérifier si le Joueur est toujours en vie
			if Personnag1.Hp <= 0 {
				fmt.Printf("%s est vaincu!\n", Personnag1.Name)
				Personnag1.Dead()
				break
			}
		}
		for (Personnag1.Hp > 0 && Monster1.Hp > 0 && Tours1 == 1 && Personnag1.Initiative < Monster1.Initiative) || (Tours1 > 1 && Personnag1.Hp > 0 && Monster1.Hp > 0) {
			CombatStarted = true
			fmt.Printf("Vous etes au tour %d\n", Tours)
			goblinPatternActivated := GoblinPattern(Monster1)

			// Tour du Monstre
			attaqueJoueur2 := Monster1.Atk - Personnag1.Defense
			if attaqueJoueur2 < 0 {
				attaqueJoueur2 = 0
			}
			Personnag1.Hp -= attaqueJoueur2

			fmt.Printf("%s attaque %s et lui inflige %d points de dégâts.\n", Monster1.Name, Personnag1.Name, attaqueJoueur2)
			Tours++
			// Apply goblin's attack pattern if it was activated
			if goblinPatternActivated {
				// Reset goblin's attack back to its original value
				Monster1.Atk /= 2
			}
			fmt.Println("-----------------------------------------------------------------------------")

			// Vérifier si le Joueur est toujours en vie
			if Personnag1.Hp <= 0 {
				fmt.Printf("%s est vaincu!\n", Personnag1.Name)
				Personnag1.Dead()
				break
			}

			// Tour du Joueur 1
			charTurn(Personnag1, Monster1, spells)
			attaqueJoueur := Personnag1.Atk - Monster1.Defense
			if attaqueJoueur < 0 {
				attaqueJoueur = 0
			}
			Monster1.Hp -= attaqueJoueur

			fmt.Printf("%s attaque %s et lui inflige %d points de dégâts.\n", Personnag1.Name, Monster1.Name, attaqueJoueur)

			// Vérifier si le Monstre est toujours en vie
			if Monster1.Hp <= 0 {
				fmt.Printf("%s est vaincu!\n", Monster1.Name)
				GagnerXp(Personnag1, Monster1)
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

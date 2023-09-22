package game

import "fmt"

var Tours int

func Tutorial(P1 *Personnage, Monstre1 *Monstre) {
	ClearConsole()
	fmt.Println("Bienvenue dans le tutoriel")
	fmt.Println("Nous allons vous apprendre les bases du combat en tour par tour")
	fmt.Println("Vous allez vous battre contre un goblin d'entrainement")
	fmt.Println("Bonne chance")
	Tours = 1
	TrainFight(P1, Monstre1)
}

func TrainFight(P1 *Personnage, Monstre1 *Monstre) {
	fmt.Println("Début du combat d'entrainement")
	for (P1.Hp > 0 && Monstre1.Hp > 0 && Tours == 1 && P1.Initiative >= Monstre1.Initiative) || (Tours > 1 && P1.Hp > 0 && Monstre1.Hp > 0) {
		fmt.Printf("Vous etes au tour %d\n", Tours)
		goblinPatternActivated := GoblinPattern(Monstre1)

		// Tour du Joueur 1
		charTurn(P1, Monstre1)
		attaqueJoueur := P1.Atk - Monstre1.Defense
		if attaqueJoueur < 0 {
			attaqueJoueur = 0
		}
		Monstre1.Hp -= attaqueJoueur

		fmt.Printf("%s attaque %s et lui inflige %d points de dégâts.\n", P1.Name, Monstre1.Name, attaqueJoueur)

		// Vérifier si le Monstre est toujours en vie
		if Monstre1.Hp <= 0 {
			fmt.Printf("%s est vaincu!\n", Monstre1.Name)
			GagnerXp(P1, Monstre1)
			break
		}

		// Tour du Monstre
		attaqueJoueur2 := Monstre1.Atk - P1.Defense
		if attaqueJoueur2 < 0 {
			attaqueJoueur2 = 0
		}
		P1.Hp -= attaqueJoueur2

		fmt.Printf("%s attaque %s et lui inflige %d points de dégâts.\n", Monstre1.Name, P1.Name, attaqueJoueur2)
		Tours++
		// Apply goblin's attack pattern if it was activated
		if goblinPatternActivated {
			// Reset goblin's attack back to its original value
			Monstre1.Atk /= 2
		}
		fmt.Println("-----------------------------------------------------------------------------")

		// Vérifier si le Joueur est toujours en vie
		if P1.Hp <= 0 {
			fmt.Printf("%s est vaincu!\n", P1.Name)
			P1.Dead()
			break
		}
	}
	for (P1.Hp > 0 && Monstre1.Hp > 0 && Tours == 1 && P1.Initiative < Monstre1.Initiative) || (Tours > 1 && P1.Hp > 0 && Monstre1.Hp > 0) {
		fmt.Printf("Vous etes au tour %d\n", Tours)
		goblinPatternActivated := GoblinPattern(Monstre1)

		// Tour du Monstre
		attaqueJoueur2 := Monstre1.Atk - P1.Defense
		if attaqueJoueur2 < 0 {
			attaqueJoueur2 = 0
		}
		P1.Hp -= attaqueJoueur2

		fmt.Printf("%s attaque %s et lui inflige %d points de dégâts.\n", Monstre1.Name, P1.Name, attaqueJoueur2)
		Tours++
		// Apply goblin's attack pattern if it was activated
		if goblinPatternActivated {
			// Reset goblin's attack back to its original value
			Monstre1.Atk /= 2
		}

		// Vérifier si le Joueur est toujours en vie
		if P1.Hp <= 0 {
			fmt.Printf("%s est vaincu!\n", P1.Name)
			P1.Dead()
			break
		} // Tour du Joueur 1
		charTurn(P1, Monstre1)
		attaqueJoueur := P1.Atk - Monstre1.Defense
		if attaqueJoueur < 0 {
			attaqueJoueur = 0
		}
		Monstre1.Hp -= attaqueJoueur

		fmt.Printf("%s attaque %s et lui inflige %d points de dégâts.\n", P1.Name, Monstre1.Name, attaqueJoueur)

		// Vérifier si le Monstre est toujours en vie
		if Monstre1.Hp <= 0 {
			fmt.Printf("%s est vaincu!\n", Monstre1.Name)
			break
		}
	}
}

func GetTours() int {
	return Tours + 1
}

func charTurn(P1 *Personnage, Monstre1 *Monstre) {
	fmt.Printf("Vous avez %d Pv / %d Pv \n", P1.Hp, P1.HpMax)
	fmt.Printf("Il reste %d Pv / %d Pv à l'ennemi \n", Monstre1.Hp, Monstre1.HpMax)

	fmt.Println("C'est votre tour, que voulez vous faire ?")
	fmt.Println("1 : Attaquer")
	fmt.Println("2 : Ouvrir l'inventaire")
	fmt.Println("3 : Utiliser un sort")
	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("Attaque basic")
	case 2:
		P1.FightInventory(Monstre1)
	case 3:

	}
}

func GagnerXp(P1 *Personnage, Monstre1 *Monstre) {
	P1.Xp += Monstre1.XpDrop
	fmt.Printf("Vous avez gagné %d points d'XP \n", Monstre1.XpDrop)
	fmt.Printf("Vous avez %d points d'XP / %d points d'XP avant le prochain niveau \n", P1.Xp, P1.XpMax)
	if P1.Xp >= P1.XpMax {
		P1.Level += 1
		surplusxp := P1.Xp - P1.XpMax
		P1.Xp = 0
		P1.XpMax += 10
		P1.Xp += surplusxp
		fmt.Printf("Vous avez gagné un niveau et êtes au niveau : %d \n", P1.Level)
		fmt.Printf("Vous avez %d points d'XP / %d points d'XP avant le prochain niveau \n", P1.Xp, P1.XpMax)
	}
}

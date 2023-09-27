package game

import "fmt"

var Tours int
var PlayerTurnTaken = false

func Tutorial(P1 *Personnage) {
	ClearConsole()
	fmt.Println("Bienvenue dans le tutoriel")
	fmt.Println("Nous allons vous apprendre les bases du combat en tour par tour")
	fmt.Println("Vous allez vous battre contre un goblin d'entrainement")
	fmt.Println("Bonne chance")
	fmt.Println("-----------------------------------------------------------------------------")

	goblin := Monstre{
		Name:       "Goblin d'entrainement",
		HpMax:      30,
		Hp:         30,
		ManaMax:    10,
		Mana:       10,
		Atk:        8,
		Defense:    2,
		Initiative: 12,
	}
	Fight(P1, &goblin)
}

func Entrainement(P1 *Personnage) {
	ClearConsole()
	text := "Bienvenue dans la zone d'entrainement\nAmusez-vous a tapé le mannequin d'entrainement\n-----------------------------------------------------------------------------\n"
	centeredText := CenterText(text)
	fmt.Println(centeredText)

	mannequin := Monstre{
		Name:       "Mannequin d'entrainement",
		Niveau:     1,
		HpMax:      30,
		Hp:         30,
		Defense:    5,
		Initiative: 0,
		XpDrop:     20,
	}
	Fight(P1, &mannequin)

}

func Fight(P1 *Personnage, Monstre1 *Monstre) {
	Tours = 1
	fmt.Println("Début du combat")
	for P1.Hp > 0 && Monstre1.Hp > 0 {
		fmt.Printf("Vous êtes au tour %d\n", Tours)
		if P1.Initiative >= Monstre1.Initiative {
			if Tours == 1 {
				fmt.Println("Vous avez l'initiative et attaquez en premier !")
			}
			for { // Player's turn
				charTurn(P1, Monstre1)
				if PlayerTurnTaken {
					PlayerTurnTaken = false // Reset playerTurnTaken for the next turn
					break
				}
			}
			// Check if the Monster is defeated
			if Monstre1.Hp <= 0 {
				Monstre1.DeadMonstre(P1)
				break
			}

			attaqueMonstre := Monstre1.Atk - P1.Defense
			if attaqueMonstre < 0 {
				attaqueMonstre = 0
			}
			P1.Hp -= attaqueMonstre

			fmt.Printf("%s attaque %s et lui inflige %d points de dégâts.\n", Monstre1.Name, P1.Name, attaqueMonstre)

			// Check if the Player is defeated
			if P1.Hp <= 0 {
				P1.Dead()
				break
			}
		}
		if P1.Initiative < Monstre1.Initiative {
			if Tours == 1 {
				fmt.Println("Le monstre est plus rapide que vous et attaque en premier !")
			}
			// Monster's turn
			attaqueMonstre := Monstre1.Atk - P1.Defense
			if attaqueMonstre < 0 {
				attaqueMonstre = 0
			}
			P1.Hp -= attaqueMonstre

			fmt.Printf("%s attaque %s et lui inflige %d points de dégâts.\n", Monstre1.Name, P1.Name, attaqueMonstre)

			// Check if the Player is defeated
			if P1.Hp <= 0 {
				fmt.Printf("%s est vaincu!\n", P1.Name)
				P1.Dead()
				break
			}
			for { // Player's turn
				charTurn(P1, Monstre1)
				if PlayerTurnTaken {
					PlayerTurnTaken = false // Reset playerTurnTaken for the next turn
					break
				}
			}

			// Check if the Monster is defeated
			if Monstre1.Hp <= 0 {
				Monstre1.DeadMonstre(P1)
				break
			}

			Tours++
			fmt.Println("-----------------------------------------------------------------------------")
		}
	}
}

func GetTours() int {
	return Tours + 1
}

func charTurn(P1 *Personnage, Monstre1 *Monstre) {
	fmt.Printf("Vous avez %d Pv / %d Pv \n", P1.Hp, P1.HpMax)
	fmt.Printf("Vous avez %d de mana restant / %d mana \n", P1.Mana, P1.ManaMax)
	fmt.Printf("Il reste %d Pv / %d Pv à l'ennemi \n", Monstre1.Hp, Monstre1.HpMax)
	fmt.Printf("Il reste %d de mana restant / %d mana à l'ennemi \n", Monstre1.Mana, Monstre1.ManaMax)

	fmt.Println("C'est votre tour, que voulez vous faire ?")
	fmt.Println("1 : Attaquer")
	fmt.Println("2 : Ouvrir l'inventaire")
	fmt.Println("3 : Utiliser un sort")
	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		P1.AtkUsed = true
		fmt.Println("Attaque basique")
		attaqueJoueur := P1.Atk - Monstre1.Defense
		if attaqueJoueur < 0 {
			attaqueJoueur = 0
		}
		Monstre1.Hp -= attaqueJoueur
		fmt.Printf("%s attaque %s et lui inflige %d points de dégâts.\n", P1.Name, Monstre1.Name, attaqueJoueur)
		PlayerTurnTaken = true
	case 2:
		P1.FightInventory(Monstre1)

	case 3:
		fmt.Println("Sélectionnez un sort :")
		for i, spell := range P1.Spells {
			fmt.Printf("%d. %s (Type: %s, Dommages: %d, Coût en Mana: %d)\n", i+1, spell.Name, spell.Type, spell.Damage, spell.ManaCost)

		}
		var selectedSpellIndex int
		fmt.Print("Entrez le numéro du sort que vous souhaitez lancer : ")
		fmt.Scan(&selectedSpellIndex)
		if selectedSpellIndex >= 1 && selectedSpellIndex <= len(P1.Spells) {
			selectedSpell := P1.Spells[selectedSpellIndex-1]
			P1.CastSpell(selectedSpell, Monstre1)
			PlayerTurnTaken = true

		} else {
			fmt.Println("Sélection de sort invalide.")
			return
		}
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

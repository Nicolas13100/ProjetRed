package game

import (
	"fmt"
	"os"
)

var S Spell
var P1 Personnage
var Monstre1 Monstre
var Marchand1 Marchand

func Menu() {

	var startChoice int

	for { // Start menu
		ClearConsole()
		fmt.Println("Bienvenue dans le jeu RPG !")
		fmt.Println("1. Commencer le jeu")
		fmt.Println("2. Quitter")
		fmt.Print("Entrez votre choix : ")
		fmt.Scan(&startChoice)

		switch startChoice {
		case 1:
			P1 := CharCreation()
			Monstre1 := InitGoblin()

			Marchand1 := NewMarchand()
			var choice1 int
			fmt.Print("voulez-vous un tutoriel ? 1.Oui / 2.Non ")
			fmt.Scan(&choice1)
			switch choice1 {
			case 1:
				{
					Tutorial(P1, Monstre1)
				}
			case 2:
				break
			default:
				fmt.Println("Il faut répondre par 1 ou 2, je ne comprend en dehors de cette selection")
				continue
			}
			ClearConsole()
			for {
				fmt.Println("Que voulez-vous faire ?")
				fmt.Println("1. Afficher les informations du personnage")
				fmt.Println("2. Accéder au contenu de l'inventaire")
				fmt.Println("3. Marchand")
				fmt.Println("4. Forgeron")
				fmt.Println("5. Entrainement")
				fmt.Println("9. Qui sont-ils ? ")
				fmt.Println("0. Quitter")

				var choice int
				fmt.Print("Entrez votre choix : ")
				if _, err := fmt.Scan(&choice); err != nil {
					fmt.Println("Erreur lors de la saisie.")
					return
				}

				switch choice {
				case 1:
					P1.DisplayInfo()
				case 2:
					P1.BaseInventory(Monstre1)
				case 3:
					var merchantChoice int
					fmt.Print("Veux-tu acheter ou vendre un objet ? (1.Acheter/2.Vendre): ")
					fmt.Scan(&merchantChoice)
					switch merchantChoice {
					case 1:
						Buy(P1, Marchand1)
					case 2:
						Sell(P1, Marchand1)
					}

				case 4:

				case 5:
					Fight(P1, Monstre1)
				case 9:

				case 0:
					fmt.Println("Au revoir !")
					return
				default:
					fmt.Println("Choix invalide.")
				}
			}

		case 2:
			fmt.Println("Au revoir !")
			os.Exit(0)

		default:
			fmt.Println("Choix invalide.")
			fmt.Println("bug")
			return
		}
	}
}

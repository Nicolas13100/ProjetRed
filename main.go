package main

import (
	"ProjectRed/game"
	"ProjectRed/utility"
	"fmt"
)

func main() {
	var name, race string

	fmt.Println("Bienvenue dans le jeu RPG !")

	m := utility.NewMarchand(nil)

	// Start menu
	fmt.Println("1. Commencer le jeu")
	fmt.Println("2. Quitter")

	var startChoice int
	fmt.Print("Entrez votre choix : ")
	if _, err := fmt.Scan(&startChoice); err != nil {
		fmt.Println("Erreur lors de la saisie.")
		return
	}

	switch startChoice {
	case 1:
		fmt.Print("Entrez votre nom : ")
		if _, err := fmt.Scan(&name); err != nil {
			fmt.Println("Erreur lors de la saisie.")
			return
		}

		fmt.Print("Entrez votre classe : ")
		if _, err := fmt.Scan(&race); err != nil {
			fmt.Println("Erreur lors de la saisie.")
			return
		}

		personnage := game.InitPersonnage(name, race)

		for {
			fmt.Println("Que voulez-vous faire ?")
			fmt.Println("1. Afficher les informations du personnage")
			fmt.Println("2. Accéder au contenu de l'inventaire")
			fmt.Println("3. Marchand")
			fmt.Println("4. Quitter")

			var choice int
			fmt.Print("Entrez votre choix : ")
			if _, err := fmt.Scan(&choice); err != nil {
				fmt.Println("Erreur lors de la saisie.")
				return
			}

			switch choice {
			case 1:
				personnage.DisplayInfo()
			case 2:
				personnage.AccessInventory()
			case 3:
				var merchantChoice string
				fmt.Print("Do you want to buy or sell? (buy/sell): ")
				fmt.Scan(&merchantChoice)
				switch merchantChoice {
				case "buy":
					m.Buy(personnage)
				case "sell":
					// Handle selling to the merchant if needed
					fmt.Println("You chose to sell.")
				default:
					fmt.Println("Invalid choice for merchant.")
				}

			case 4:
				fmt.Println("Au revoir !")
				return
			default:
				fmt.Println("Choix invalide.")
			}
		}

	case 2:
		fmt.Println("Au revoir !")
		return

	default:
		fmt.Println("Choix invalide.")
	}
}

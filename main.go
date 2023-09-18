package main

import (
	"ProjectRed/game"
	"ProjectRed/utility"
	"bufio"
	"fmt"
	"os"
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

		fmt.Print("Entrez votre race : ")
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
				fmt.Print("Veux-tu acheter ou vendre un objet ? (Acheter/Vendre): ")
				fmt.Scan(&merchantChoice)
				switch merchantChoice {
				case "Acheter":
					m.Buy(personnage)
				case "Vendre":
					// Handle selling to the merchant if needed
					fmt.Println("Que veux-tu vendre ?")
				default:
					fmt.Println("Choix invalide.")
				}

				// Wait for user to press Enter to continue
				fmt.Print("Taper Entrer pour continuer...")
				bufio.NewReader(os.Stdin).ReadBytes('\n')

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

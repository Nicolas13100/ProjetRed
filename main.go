package main

import (
	"ProjectRed/game"
	"ProjectRed/utility"
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func clearConsole() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func main() {
	fmt.Println("Bienvenue dans le jeu RPG !")
	m := utility.NewMarchand(nil)
	var startChoice int
	var personnage *game.Personnage // Declare the personnage variable outside of the switch
	for {                           // Start menu
		clearConsole()
		fmt.Println("1. Commencer le jeu")
		fmt.Println("2. Quitter")
		fmt.Print("Entrez votre choix : ")
		fmt.Scan(&startChoice)

		switch startChoice {
		case 1:
			character := game.CharCreation()
			goblin := game.InitGoblin()
			if character == nil {
				fmt.Println("Mauvais choix de race")
			} else {
				personnage = character
				// Handle the character creation success
			}

			var choice int
			fmt.Print("voulez-vous un tutoriel ? 1.Oui / 2.Non ")
			if _, err := fmt.Scan(&choice); err != nil {
				fmt.Println("Erreur lors de la saisie.")
				return
			}
			switch choice {
			case 1:

				game.Tutorial(personnage, goblin)

			case 2:
				break
			}

			for {
				fmt.Println("Que voulez-vous faire ?")
				fmt.Println("1. Afficher les informations du personnage")
				fmt.Println("2. Accéder au contenu de l'inventaire")
				fmt.Println("3. Marchand")
				fmt.Println("4. Forgeron")
				fmt.Println("5. Entrainement")
				fmt.Println("6. Quitter")

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
					personnage.AccessInventory(goblin)
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
					utility.Forgeron(personnage)
				case 5:
					game.TrainFight(personnage, goblin)
				case 6:
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
			continue
		}
	}
}

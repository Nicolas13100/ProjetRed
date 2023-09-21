package main

import (
	"ProjectRed/game"
	"ProjectRed/npc"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func ClearConsole() {
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

	m := npc.NewMarchand(nil)
	var startChoice int
	var spells *game.Spell
	var personnage *game.Personnage // Declare the personnage variable outside of the switch
	for {                           // Start menu

		fmt.Println("1. Commencer le jeu")
		fmt.Println("2. Quitter")
		fmt.Print("Entrez votre choix : ")
		fmt.Scan(&startChoice)

		switch startChoice {
		case 1:
			character := game.CharCreation()
			goblin := game.InitGoblin()

			if character == nil {
				ClearConsole()
				fmt.Println("Mauvais choix de race")
			} else {
				personnage = character
				// Handle the character creation success
			}

			var choice int
			fmt.Print("voulez-vous un tutoriel ? 1.Oui / 2.Non ")
			switch choice {
			case 1:
				{
					game.Tutorial(personnage, goblin, spells)
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
					personnage.DisplayInfo()
				case 2:
					personnage.BaseInventory(goblin, spells)
				case 3:
					var merchantChoice int
					fmt.Print("Veux-tu acheter ou vendre un objet ? (1.Acheter/2.Vendre): ")
					fmt.Scan(&merchantChoice)
					switch merchantChoice {
					case 1:
						m.Buy(personnage)
					case 2:
						m.Sell(personnage, "")
					}

				case 4:
					npc.Forgeron(personnage)
				case 5:
					game.Fight(personnage, goblin, spells)
				case 9:
					npc.QuiSontIls()
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
			continue
		}
	}
}

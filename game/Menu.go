package game

import (
	"fmt"
	"os"
	"strings"

	"github.com/mattn/go-runewidth"
	"golang.org/x/term"
)

var P1 Personnage
var Monstre1 Monstre

func GetConsoleWidth() int {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic(err)
	}
	return width
}

func CenterText(text string) string {
	consoleWidth := GetConsoleWidth()

	lines := strings.Split(text, "\n")
	centeredLines := []string{}

	for _, line := range lines {
		padding := (consoleWidth - runewidth.StringWidth(line)) / 2
		centeredLine := fmt.Sprintf("%*s", padding+runewidth.StringWidth(line), line)
		centeredLines = append(centeredLines, centeredLine)
	}
	return strings.Join(centeredLines, "\n")
}

func Menu() {

	var startChoice int

	for { // Start menu
		ClearConsole()
		text := "Bienvenue dans le jeu RPG\n 1.Commencer le jeux\n 2.Quitter\n Entrer votre choix"
		centeredText := CenterText(text)
		fmt.Println(centeredText)

		fmt.Scan(&startChoice)

		switch startChoice {
		case 1:
			P1 := CharCreation()
			var choice1 int
			text := "Voulez-vous un tutoriel ? 1.Oui / 2.Non"
			centeredText := CenterText(text)
			fmt.Println(centeredText)
			fmt.Scan(&choice1)
			switch choice1 {
			case 1:
				{
					Tutorial(P1)
				}
			case 2:
				break
			default:
				text := "Il faut répondre par 1 ou 2, je ne comprend en dehors de cette selection"
				centeredText := CenterText(text)
				fmt.Println(centeredText)
				continue
			}
			ClearConsole()
			for {
				text := "Bienvenue à toi:%s\n\nQue voulez-vous faire ?\n1. Afficher les informations du personnage \n2. Accéder au contenu de l'inventaire\n3. Marchand \n4. Forgeron\n5. Entrainement \n6. Dongeon\n 9. Qui sont-ils ?\n0. Quitter\nEntrer votre choix :\n"
				centeredText := CenterText(text)
				fmt.Printf(centeredText, P1.Name)

				var choice int

				if _, err := fmt.Scan(&choice); err != nil {
					text := "Erreur lors de la saisie."
					centeredText := CenterText(text)
					fmt.Println(centeredText)
					return
				}

				switch choice {
				case 1:
					P1.DisplayInfo()
				case 2:
					P1.BaseInventory()
				case 3:
					var merchantChoice int
					text := "Veux-tu acheter ou vendre un objet ? (1.Acheter/2.Vendre)"
					centeredText := CenterText(text)
					fmt.Println(centeredText)
					fmt.Scan(&merchantChoice)
					switch merchantChoice {
					case 1:
						Buy(P1, &Marchand1)
					case 2:
						Sell(P1, &Marchand1)
					}
				case 4:
					Forgeron(P1)
				case 5:
					Entrainement(P1)
				case 6:
					InitMap("", "")
				case 7:
					Campain(P1)
				case 9:
					QuiSontIls()
				case 0:
					text := "Au revoir !"
					centeredText := CenterText(text)
					fmt.Println(centeredText)
					return
				default:
					text := "Choix invalide."
					centeredText := CenterText(text)
					fmt.Println(centeredText)
					return
				}
			}

		case 2:
			text := "Au revoir !"
			centeredText := CenterText(text)
			fmt.Println(centeredText)
			os.Exit(0)
		default:
			text := "Choix invalide."
			centeredText := CenterText(text)
			fmt.Println(centeredText)
			continue
		}
	}
}

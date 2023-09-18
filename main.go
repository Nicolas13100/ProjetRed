package main

import (
	"ProjectRed/game"
	"ProjectRed/utility"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CharCreation() *game.Personnage {
	var name, race string

	fmt.Println("Création de votre personnage :")

	// Demander à l'utilisateur de choisir son nom
	fmt.Print("Entrez votre nom : ")
	fmt.Scan(&name)

	// Vérifier si le nom contient uniquement des lettres
	if !OnlyLetters(name) {
		fmt.Println("Le nom doit contenir uniquement des lettres.")
		return nil
	}

	// Mettre la première lettre en majuscule et le reste en minuscule
	name = strings.Title(strings.ToLower(name))

	// Demander à l'utilisateur de choisir sa race (vous pouvez adapter cette partie selon vos besoins)
	fmt.Print("Choisissez votre race : ")
	fmt.Println("1. Humain : Vous commencez avec 100 PV Max")
	fmt.Println("2. Elfe : Vous commencez avec 80 PV Max")
	fmt.Println("3. Nain : Vous commencez avec 120 PV Max")
	fmt.Scan(&race)

	var choicerace int
	switch choicerace {
	case 1:

	}

	// Créer et initialiser le personnage
	p := game.InitPersonnage(name, race)

	fmt.Printf("Bienvenue à toi : %s\n", p.Name)
	return p
}

func OnlyLetters(s string) bool {
	for _, char := range s {
		if !('a' <= char && char <= 'z' || 'A' <= char && char <= 'Z') {
			return false
		}
	}
	return true
}
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
		CharCreation()

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
		return
	}
}

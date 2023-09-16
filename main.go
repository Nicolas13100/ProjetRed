package main

import "fmt"

type Personnage struct {
	name      string
	race      string
	level     int
	HpMax     int
	Hp        int
	inventory map[string]int
}

func Init(name string, race string) *Personnage {
	p1 := &Personnage{
		// Initialisation des paramètres du personnage
		name:  name,
		race:  race,
		level: 1,
		HpMax: 100,
		Hp:    40,
		inventory: map[string]int{
			"Potion de soin": 3,
		},
	}
	return p1
}

func main() {
	// Perso

	var name, race string

	// Nom et class
	fmt.Print("Entrez votre nom : ")
	_, err := fmt.Scan(&name)
	if err != nil {
		fmt.Println("Erreur de lecture de l'entrée :", err)
		return
	}
	fmt.Print("Entrez votre classe : ")
	_, errc := fmt.Scan(&race)
	if errc != nil {
		fmt.Println("Erreur de lecture de l'entrée :", err)
		return
	}
	p1 := Init(name, race)

	for {
		// Choix
		fmt.Println("Que voulez-vous faire ?")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder au contenu de l'inventaire")
		fmt.Println("3. Marchand")
		fmt.Println("4. Quitter")

		var choice int
		fmt.Print("Entrez votre choix : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			p1.displayInfo()
		case 2:
			p1.accessInventory()
		case 3:
			p1.Marchand()
		case 4:
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func (p1 *Personnage) displayInfo() {
	// Affichage des informations du Personnage p1

	for {
		fmt.Println("Nom:", p1.name)
		fmt.Println("Classe:", p1.race)
		fmt.Println("Niveau:", p1.level)
		fmt.Println("Points de vie maximum:", p1.HpMax)
		fmt.Println("Points de vie actuels:", p1.Hp)
		// Affichage de l'inventaire
		fmt.Println("\nType 'return' to come back to main menu")

		// Read user input
		var input string
		fmt.Scan(&input)

		// Check if the user wants to return to the main menu
		if input == "return" {
			break
		}
	}
}

func (p1 *Personnage) accessInventory() {
	for {
		fmt.Println("\nInventaire:")
		for key, value := range p1.inventory {
			fmt.Printf("%s: %d \n", key, value)
		}
		fmt.Println("\nQue voulez-vous faire ?")
		fmt.Println("1. Sélectionner une potion de soin")
		fmt.Println("2. Retourner en arrière")

		var input int
		fmt.Print("Entrez votre choix : ")
		fmt.Scan(&input)

		switch input {
		case 1:
			if count, ok := p1.inventory["Potion de soin"]; ok && count > 0 {
				p1.takePot()
			} else {
				fmt.Println("Vous n'avez pas de Potion de soin dans votre inventaire.")
			}
		case 2:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func (p1 *Personnage) takePot() {
	var choice string
	fmt.Println("Voulez-vous prendre une potion ? (Oui/Non) ")
	fmt.Scan(&choice)
	switch choice == "Oui" {
	case p1.Hp >= p1.HpMax:
		fmt.Println(" Vos points de vie sont déjà plein")
	case p1.Hp < p1.HpMax:
		if count, ok := p1.inventory["Potion de soin"]; ok && count > 0 {
			p1.inventory["Potion de soin"]--
			p1.Hp += 50
			if p1.Hp > p1.HpMax {
				surplu := p1.Hp - p1.HpMax
				fmt.Printf("Cette potion vous rendra %d PV de trop, continuer (Oui/Non)?\n", surplu)
				fmt.Scan(&choice)
				if choice == "Oui" {
					p1.Hp = p1.HpMax
				} else if choice == "Non" {
					fmt.Printf("Vous n'avez  pas pris une Potion de soin. Vous avez %d PV\n", p1.Hp)
					return
				}
			}
			fmt.Printf("Vous avez pris une Potion de soin. Vous avez %d PV\n", p1.Hp)
		}
	}
	if choice == "Non" {
		return
	}
}

func (p1 *Personnage) Marchand() {
	var choice string
	étale := []string{"potion de soin", "potion de soin"}

	// Display the items in étale before the user makes a choice
	fmt.Println("Les objets disponibles à l'achat sont :")
	for i := 0; i < len(étale); i++ {
		fmt.Println(i+1, ": ", étale[i])
	}

	fmt.Println("Bonjour voyageur, que puis-je pour vous aujourd'hui ? (Acheter / Vendre)")
	fmt.Scan(&choice)

	if choice == "Acheter" {
		scannb := 0
		for i := 0; i < len(étale); i++ {
			var choice int
			scannb++
			fmt.Println("Qu'est-ce-qui vous ferait plaisir ? (numero)")
			fmt.Scan(&choice)
			if choice > scannb {
				fmt.Println("Je n'ai pas cet objet, souhaitez-vouz autre chose ,")
				continue
			}
			if choice < 0 {
				fmt.Println("Je n'ai pas cet objet, souhaitez-vouz autre chose ,")
				continue
			}
		}
	}
	if choice == "Vendre" {
		fmt.Println("Je n'ai pas d'argent")
		return
	}
}

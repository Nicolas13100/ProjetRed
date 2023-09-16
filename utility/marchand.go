package utility

import (
	"ProjectRed/game"
	"fmt"
)

type Marchand struct {
	inventory map[string]int
}

func NewMarchand(personnage *game.Personnage) *Marchand {
	return &Marchand{
		inventory: map[string]int{
			"Potion de soin": 3,
		},
	}
}

func (m *Marchand) Buy(personnage *game.Personnage, item string, price int) {
	fmt.Println("Les objets disponibles à l'achat sont :")
	for key, value := range m.inventory {
		fmt.Println(key, ": ", value)
	}

	fmt.Println("Que voulez-vous acheter ? (Entrez le nom de l'objet)")
	var itemToBuy string
	fmt.Scan(&itemToBuy)

	fmt.Println("User input:", itemToBuy) // Debug statement
	for key := range m.inventory {
		fmt.Println("Inventory item:", key) // Debug statement
	}

	if quantity, ok := m.inventory[itemToBuy]; ok && quantity > 0 {
		if personnage.Gold >= price {
			m.inventory[itemToBuy]--
			personnage.Gold -= price
			if quantity > 0 {
				personnage.Inventory[itemToBuy]++
			} else {
				personnage.Inventory[itemToBuy] = 1
			}
			fmt.Println("Vous avez acheté une", itemToBuy)
		} else {
			fmt.Println("Vous n'avez pas assez d'argent pour acheter cet objet.")
		}
	} else {
		fmt.Println("Désolé, je n'ai pas cet objet en stock.")
	}
}
func (m *Marchand) Sell(Personnage *game.Personnage, item string, price int) {
	fmt.Println("Les objets disponibles à l'achat sont :")
	for key, value := range m.inventory {
		fmt.Println(key, ": ", value)
	}

	fmt.Println("Que voulez-vous acheter ? (Entrez le nom de l'objet)")
	var itemToBuy string
	if _, err := fmt.Scan(&itemToBuy); err != nil {
		fmt.Println("Erreur lors de la saisie.")
		return
	}

	if quantity, ok := m.inventory[itemToBuy]; ok && quantity > 0 {
		if Personnage.Gold >= price {
			m.inventory[itemToBuy]--
			Personnage.Gold -= price
			Personnage.Inventory[itemToBuy]++
			fmt.Println("Vous avez acheté une", itemToBuy)
		} else {
			fmt.Println("Vous n'avez pas assez d'argent pour acheter cet objet.")
		}
	} else {
		fmt.Println("Désolé, je n'ai pas cet objet en stock.")
	}
}

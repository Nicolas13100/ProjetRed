package utility

import (
	"ProjectRed/game"
	"fmt"
)

type Marchand struct {
	inventory map[string]int
	prices    map[string]int
}

func NewMarchand(personnage *game.Personnage) *Marchand {
	return &Marchand{
		inventory: map[string]int{
			"Augmentation d'inventaire":    3,
			"Potion de soin":               3,
			"Potion de poison":             3,
			"Livre de sort : Boule de feu": 1,
			"Fourrure de Loup":             4,
			"Peau de Troll":                1,
			"Cuir de Sanglier":             1,
			"Plume de Corbeau":             1,
		},
		prices: map[string]int{
			"Potion de soin":   3,
			"Potion de poison": 6,
			// Set the price of "Potion de soin" to 10
			// Add more items and their prices as needed
			"Livre de sort : Boule de feu": 25,
			"Augmentation d'inventaire":    30,
			"Fourrure de Loup":             4,
			"Peau de Troll":                7,
			"Cuir de Sanglier":             3,
			"Plume de Corbeau":             1,
		},
	}
}

func (m *Marchand) Buy(personnage *game.Personnage) {
	fmt.Println("Les objets disponibles à l'achat sont :")
	for key, value := range m.inventory {
		fmt.Println(key, ": ", value)
	}
	fmt.Println("Que voulez-vous acheter ? (potion/sort/autre)")
	var itemToBuy string
	if _, err := fmt.Scan(&itemToBuy); err != nil {
		fmt.Println("Erreur lors de la saisie.")
		return
	}

	switch itemToBuy {
	case "potion":
		fmt.Println("Quel type de potion voulez-vous acheter ? (soin/autre)")
		var potionType string
		if _, err := fmt.Scan(&potionType); err != nil {
			fmt.Println("Erreur lors de la saisie.")
			return
		}
		itemToBuy = "Potion de " + potionType
	case "sort":
		fmt.Println("Quel type de sort voulez-vous acheter ? (feu/autre)")
		var spellType string
		if _, err := fmt.Scan(&spellType); err != nil {
			fmt.Println("Erreur lors de la saisie.")
			return
		}
		itemToBuy = "Livre de sort : Boule de " + spellType
	}
	switch quantity := m.inventory[itemToBuy]; {
	case quantity > 0:
		if personnage.LimiteInventory() {
			m.inventory[itemToBuy]--
			price, exists := m.prices[itemToBuy]
			if exists {
				personnage.Gold -= price // Deduct the price from character's gold
				personnage.Inventory[itemToBuy]++
				fmt.Println("Vous avez acheté une", itemToBuy)
			} else {
				fmt.Println("Désolé, je ne connais pas le prix de cet objet.")
			}
		} else {
			fmt.Println("L'inventaire est plein, vous ne pouvez pas acheter cet objet.")
		}
	default:
		fmt.Println("Désolé, je n'ai pas cet objet en stock.")
	}
}
func (m *Marchand) Sell(Personnage *game.Personnage, item string, price int) {
	fmt.Println("Les objets disponibles à l'achat sont :")
	for key, value := range m.inventory {
		fmt.Println(key, ": ", value)
	}

	fmt.Println("Que voulez-vous vendre ? (Entrez le nom de l'objet)")
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

package utility

import (
	"ProjectRed/game"
	"fmt"
)

type Marchand struct {
	inventory map[string]int
	prices    map[string]int
	Gold      int
}

func NewMarchand(personnage *game.Personnage) *Marchand {
	return &Marchand{

		Gold: 200,
		inventory: map[string]int{
			"Augmentation d'inventaire":    3,
			"Potion de soin":               3,
			"Potion de poison":             3,
			"Potion de mana":               3,
			"Livre de sort : Boule de feu": 1,
			"Fourrure de Loup":             4,
			"Peau de Troll":                1,
			"Cuir de Sanglier":             1,
			"Plume de Corbeau":             1,
		},
		prices: map[string]int{
			"Potion de soin":   3,
			"Potion de poison": 6,
			"Potion de mana":   8,
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
	fmt.Println("Que voulez-vous acheter ? (1.Potion/2.sort/3.Objet)")
	var itemToBuy string
	if _, err := fmt.Scan(&itemToBuy); err != nil {
		fmt.Println("Erreur lors de la saisie.")
		return
	}

	switch itemToBuy {
	case "potion":
		fmt.Println("Quel type de potion voulez-vous acheter ? (soin/poison/mana)")
		var potionType string
		if _, err := fmt.Scan(&potionType); err != nil {
			fmt.Println("Erreur lors de la saisie.")
			return
		}
		itemToBuy = "Potion de " + potionType
		// Handle item purchase here...

	case "sort":
		fmt.Println("Quel type de sort voulez-vous acheter ? (feu/autre)")
		var spellType string
		if _, err := fmt.Scan(&spellType); err != nil {
			fmt.Println("Erreur lors de la saisie.")
			return
		}
		itemToBuy = "Livre de sort : Boule de " + spellType
		// Handle item purchase here...

	case "loot":
		fmt.Println("Quel type de loot souhaitez-vous acheter ? (1 : Loup/ 2 : Troll/ 3 : Sanglier/ 4 : Corbeau)")
		var loot int
		if _, err := fmt.Scan(&loot); err != nil {
			fmt.Println("Erreur lors de la saisie.")
			return
		}

		switch loot {
		case 1:
			fmt.Println("Vous souhaitez acheter une Fourrure de Loup ?(Oui/Non)")
			var lootype string
			if _, err := fmt.Scan(&lootype); err != nil {
				fmt.Println("Erreur lors de la saisie.")
				return
			}
			if lootype == "Oui" {
				itemToBuy = "Fourrure de Loup"
			} else {
				fmt.Println("Achat annulé.")
				return
			}
		case 2:
			fmt.Println("Vous souhaitez acheter une Peau de Troll ?")
			var lootype string
			if _, err := fmt.Scan(&lootype); err != nil {
				fmt.Println("Erreur lors de la saisie.")
				return
			}
			if lootype == "Oui" {
				itemToBuy = "Peau de Troll"
			} else {
				fmt.Println("Achat annulé.")
				return
			}
		case 3:
			fmt.Println("Vous souhaitez acheter une Cuir de Sanglier ?")
			var lootype string
			if _, err := fmt.Scan(&lootype); err != nil {
				fmt.Println("Erreur lors de la saisie.")
				return
			}
			if lootype == "Oui" {
				itemToBuy = "Cuir de Sanglier"
			} else {
				fmt.Println("Achat annulé.")
				return
			}
		case 4:
			fmt.Println("Vous souhaitez acheter une Plume de Corbeau ?")
			var lootype string
			if _, err := fmt.Scan(&lootype); err != nil {
				fmt.Println("Erreur lors de la saisie.")
				return
			}
			if lootype == "Oui" {
				itemToBuy = "Plume de Corbeau"
			} else {
				fmt.Println("Achat annulé.")
				return
			}
		default:
			fmt.Println("Type de loot invalide.")
			return
		}
	default:
		fmt.Println("Type d'objet invalide.")
		return
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
func (m *Marchand) Sell(Personnage *game.Personnage, item string) {
	fmt.Println("Les objets disponibles à la vente sont :")
	for key, value := range Personnage.Inventory {
		fmt.Println(key, value)
	}

	fmt.Println("Quel type d'objet voulez-vous vendre ?(Potion / Sort / Loot)")
	var itemToSell string
	if _, err := fmt.Scan(&itemToSell); err != nil {
		fmt.Println("Erreur lors de la saisie.")
		return
	}
	var price int
	switch itemToSell {
	case "Potion":
		fmt.Println("Quel type de potion voulez-vous vendre ? (soin/poison)")
		var potionType string
		if _, err := fmt.Scan(&potionType); err != nil {
			fmt.Println("Erreur lors de la saisie.")
			return
		}
		itemToSell = "Potion de " + potionType

	case "Sort":
		fmt.Println("Quel type de sort voulez-vous vendre ? (feu/autre)")
		var spellType string
		if _, err := fmt.Scan(&spellType); err != nil {
			fmt.Println("Erreur lors de la saisie.")
			return
		}
		itemToSell = "Livre de sort : Boule de " + spellType

	case "Loot":
		fmt.Println("Quel type de loot souhaitez-vous vendre ? (1 : Loup/ 2 : Troll/ 3 : Sanglier/ 4 : Corbeau)")
		var loot int
		if _, err := fmt.Scan(&loot); err != nil {
			fmt.Println("Erreur lors de la saisie.")
			return
		}

		switch loot {
		case 1:
			fmt.Println("Vous souhaitez vendre une Fourrure de Loup ?(Oui/Non)")
			var lootype string
			if _, err := fmt.Scan(&lootype); err != nil {
				fmt.Println("Erreur lors de la saisie.")
				return
			}
			if lootype == "Oui" {
				itemToSell = "Fourrure de Loup"
			} else {
				fmt.Println("Achat annulé.")
				return
			}
		case 2:
			fmt.Println("Vous souhaitez vendre une Peau de Troll ?")
			var lootype string
			if _, err := fmt.Scan(&lootype); err != nil {
				fmt.Println("Erreur lors de la saisie.")
				return
			}
			if lootype == "Oui" {
				itemToSell = "Peau de Troll"
			} else {
				fmt.Println("Achat annulé.")
				return
			}
		case 3:
			fmt.Println("Vous souhaitez vendre un Cuir de Sanglier ?")
			var lootype string
			if _, err := fmt.Scan(&lootype); err != nil {
				fmt.Println("Erreur lors de la saisie.")
				return
			}
			if lootype == "Oui" {
				itemToSell = "Cuir de Sanglier"
			} else {
				fmt.Println("Achat annulé.")
				return
			}
		case 4:
			fmt.Println("Vous souhaitez vendre une Plume de Corbeau ?")
			var lootype string
			if _, err := fmt.Scan(&lootype); err != nil {
				fmt.Println("Erreur lors de la saisie.")
				return
			}
			if lootype == "Oui" {
				itemToSell = "Plume de Corbeau"
			} else {
				fmt.Println("Achat annulé.")
				return
			}
		default:
			fmt.Println("Type de loot invalide.")
			return
		}

	}
	price, exists := m.prices[itemToSell]
	if !exists {
		fmt.Println("Désolé, le prix de vente de cet objet n'est pas défini.")
		return
	}
	if quantity, ok := Personnage.Inventory[itemToSell]; ok && quantity > 0 {
		// Vérifier si le marchand a assez d'argent pour acheter l'objet
		if m.Gold >= price {
			m.inventory[itemToSell]++
			m.Gold -= price
			Personnage.Gold += price
			Personnage.Inventory[itemToSell]--
			fmt.Println("Vous avez vendu un(e)", itemToSell, "pour", price, "pièces d'or.")
		} else {
			fmt.Println("Le marchand n'a pas assez d'argent pour acheter cet objet.")
		}
	} else {
		fmt.Println("Désolé, vous n'avez pas cet objet en stock.")
	}
}

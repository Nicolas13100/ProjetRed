package game

import (
	"fmt"
)

var (
	Marchand1 = Marchand{
		Name:       "John",
		Race:       "Humain",
		Level:      1,
		Xp:         0,
		XpMax:      100,
		HpMax:      100,
		Hp:         100,
		Gold:       1000,
		Mana:       40,
		ManaMax:    100,
		Atk:        5,
		Defense:    0,
		Initiative: 10,
		Inventory: map[string]int{
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
		Prices: map[string]int{
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
)

func Buy(personnage *Personnage, m1 *Marchand) {
	fmt.Println("Les objets disponibles à l'achat sont :")
	var itemToBuy string
	for key, value := range m1.Inventory {
		fmt.Println(key, ": ", value)
	}
	selectitem, err := getInput("Que voulez-vous acheter ? (1.Potion/2.sort/3.Objet) ")
	if err != nil {
		fmt.Println("Erreur lors de la saisie. Veuillez entrer un nombre.")
		return
	}

	switch selectitem {
	case 1:
		fmt.Println("Quel type de potion voulez-vous acheter ? (1.Soin/2.Poison/3.Mana)")
		var potionType int
		switch potionType {
		case 1:
			itemToBuy = "Potion de soin"
		case 2:
			itemToBuy = "Potion de poison"
		case 3:
			itemToBuy = "Potion de mana"
		}

	case 2:
		fmt.Println("Quel sort voulez-vous acheter ? (1. Boule de Feu /)")
		var spellType int
		if _, err := fmt.Scan(&spellType); err != nil {
			fmt.Println("Erreur lors de la saisie.")
			return
		}
		switch spellType {
		case 1:
			itemToBuy = "Livre de sort : Boule de feu"
		}

	case 3:
		fmt.Println("Quel type de loot souhaitez-vous acheter ? (1 : Loup/ 2 : Troll/ 3 : Sanglier/ 4 : Corbeau)")
		var loot int
		if _, err := fmt.Scan(&loot); err != nil {
			fmt.Println("Erreur lors de la saisie.")
			return
		}

		switch loot {
		case 1:
			if askYesNo("Vous souhaitez acheter une Fourrure de Loup ?") {
				itemToBuy = "Fourrure de Loup"
			} else {
				fmt.Println("Achat annulé.")
				return
			}

		case 2:
			if askYesNo("Vous souhaitez acheter une Peau de Troll ?") {
				itemToBuy = "Peau de Troll"
			} else {
				fmt.Println("Achat annulé.")
				return
			}

		case 3:
			if askYesNo("Vous souhaitez acheter une Cuir de Sanglier ?") {
				itemToBuy = "Cuir de Sanglier"
			} else {
				fmt.Println("Achat annulé.")
				return
			}

		case 4:
			if askYesNo("Vous souhaitez acheter une Plume de Corbeau ?") {
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

	switch quantity := m1.Inventory[itemToBuy]; {
	case quantity > 0:
		if personnage.LimiteInventory() {
			m1.Inventory[itemToBuy]--
			price, exists := m1.Prices[itemToBuy]
			if exists {
				personnage.Gold -= price // Deduct the price from character's gold
				personnage.Inventory[itemToBuy]++
				fmt.Println("Vous avez acheté ", itemToBuy)
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
func Sell(personnage *Personnage, m1 *Marchand) {
	fmt.Println("Les objets disponibles à la vente sont :")

	for key, value := range personnage.Inventory {
		fmt.Println(key, value)
	}

	selectItemType, err := getInput("Quel type d'objet voulez-vous vendre ? (1.Potion /2.Loot): ")
	if err != nil {
		fmt.Println("Erreur lors de la saisie. Veuillez entrer un nombre.")
		return
	}

	var itemToSell string
	var price int

	switch selectItemType {
	case 1:
		potionType, err := getInput("Quel type de potion voulez-vous vendre ? (1.soin/2.poison/3.mana): ")
		if err != nil {
			fmt.Println("Erreur lors de la saisie. Veuillez entrer un nombre.")
			return
		}

		switch potionType {
		case 1:
			itemToSell = "Potion de soin"
		case 2:
			itemToSell = "Potion de poison"
		case 3:
			itemToSell = "Potion de mana"
		default:
			fmt.Println("Type de potion invalide.")
			return
		}

	case 2:
		lootType, err := getInput("Quel type de loot souhaitez-vous vendre ? (1 : Loup/ 2 : Troll/ 3 : Sanglier/ 4 : Corbeau): ")
		if err != nil {
			fmt.Println("Erreur lors de la saisie. Veuillez entrer un nombre.")
			return
		}

		switch lootType {
		case 1:
			itemToSell = "Fourrure de Loup"
		case 2:
			itemToSell = "Peau de Troll"
		case 3:
			itemToSell = "Cuir de Sanglier"
		case 4:
			itemToSell = "Plume de Corbeau"
		default:
			fmt.Println("Type de loot invalide.")
			return
		}

	default:
		fmt.Println("Type d'objet invalide.")
		return
	}

	price, exists := m1.Prices[itemToSell]
	if !exists {
		fmt.Println("Désolé, le prix de vente de cet objet n'est pas défini.")
		return
	}

	quantity, ok := personnage.Inventory[itemToSell]
	if !ok || quantity <= 0 {
		fmt.Println("Désolé, vous n'avez pas cet objet en stock.")
		return
	}

	if m1.Gold >= price {
		m1.Inventory[itemToSell]++
		m1.Gold -= price
		personnage.Gold += price
		personnage.Inventory[itemToSell]--
		fmt.Printf("Vous avez vendu un(e) %s pour %d pièces d'or.\n", itemToSell, price)
	} else {
		fmt.Println("Le marchand n'a pas assez d'argent pour acheter cet objet.")
	}
}
func getInput(prompt string) (int, error) {
	var input int
	fmt.Print(prompt)
	_, err := fmt.Scan(&input)
	return input, err
}

func askYesNo(question string) bool {
	var choice int
	fmt.Print(question + " (1.Oui/2.Non): ")
	fmt.Scan(&choice)
	return choice == 1
}

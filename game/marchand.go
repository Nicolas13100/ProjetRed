package game

import (
	"fmt"
)

func NewMarchand() *Marchand {
	m1 := Marchand{
		Name:       "John",
		Race:       "Humain",
		Level:      1,
		Xp:         0,
		XpMax:      100,
		HpMax:      100,
		Hp:         100,
		Gold:       200,
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
	return &m1
}

func Buy(personnage *Personnage, m1 *Marchand) {
	fmt.Println("Les objets disponibles à l'achat sont :")
	var itemToBuy string
	for key, value := range m1.Inventory {
		fmt.Println(key, ": ", value)
	}
	fmt.Println("Que voulez-vous acheter ? (1.Potion/2.sort/3.Objet)")
	var selectitem int
	if _, err := fmt.Scan(&selectitem); err != nil {
		fmt.Println("Erreur lors de la saisie.")
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
			fmt.Println("Vous souhaitez acheter une Fourrure de Loup ?(1.Oui/2.Non)")
			var lootype int
			if _, err := fmt.Scan(&lootype); err != nil {
				fmt.Println("Erreur lors de la saisie.")
				return
			}
			if lootype == 1 {
				itemToBuy = "Fourrure de Loup"
			} else {
				fmt.Println("Achat annulé.")
				return
			}
		case 2:
			fmt.Println("Vous souhaitez acheter une Peau de Troll ?(1.Oui/2.Non)")
			var lootype int
			if _, err := fmt.Scan(&lootype); err != nil {
				fmt.Println("Erreur lors de la saisie.")
				return
			}
			if lootype == 1 {
				itemToBuy = "Peau de Troll"
			} else {
				fmt.Println("Achat annulé.")
				return
			}
		case 3:
			fmt.Println("Vous souhaitez acheter une Cuir de Sanglier ?(1.Oui/2.Non)")
			var lootype int
			if _, err := fmt.Scan(&lootype); err != nil {
				fmt.Println("Erreur lors de la saisie.")
				return
			}
			if lootype == 1 {
				itemToBuy = "Cuir de Sanglier"
			} else {
				fmt.Println("Achat annulé.")
				return
			}
		case 4:
			fmt.Println("Vous souhaitez acheter une Plume de Corbeau ?(1.Oui/2.Non)")
			var lootype int
			if _, err := fmt.Scan(&lootype); err != nil {
				fmt.Println("Erreur lors de la saisie.")
				return
			}
			if lootype == 1 {
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
func Sell(Personnage *Personnage, m1 *Marchand) {
	fmt.Println("Les objets disponibles à la vente sont :")
	for key, value := range Personnage.Inventory {
		fmt.Println(key, value)
	}

	fmt.Println("Quel type d'objet voulez-vous vendre ?(1.Potion /2.Loot)")
	var itemToSell string
	var truc int
	if _, err := fmt.Scan(&truc); err != nil {
		fmt.Println("Erreur lors de la saisie.")
		return
	}
	var price int
	switch truc {
	case 1:
		fmt.Println("Quel type de potion voulez-vous vendre ? (1.soin/2.poison/3.mana)")
		var potionType int
		if _, err := fmt.Scan(&potionType); err != nil {
			fmt.Println("Erreur lors de la saisie.")
			return
		}
		switch potionType {
		case 1:
			itemToSell = "Potion de soin"
		case 2:
			itemToSell = "potion de poison"
		case 3:
			itemToSell = "Potion de mana"
		}

	case 2:
		fmt.Println("Quel type d'objet souhaitez-vous vendre ? (1 : Loup/ 2 : Troll/ 3 : Sanglier/ 4 : Corbeau)")
		var loot int
		if _, err := fmt.Scan(&loot); err != nil {
			fmt.Println("Erreur lors de la saisie.")
			return
		}

		switch loot {
		case 1:
			fmt.Println("Vous souhaitez vendre une Fourrure de Loup ? (1.Oui 2.Non)")
			var lootype int
			if _, err := fmt.Scan(&lootype); err != nil {
				fmt.Println("Erreur lors de la saisie.")
				return
			}
			if lootype == 1 {
				itemToSell = "Fourrure de Loup"
			} else {
				fmt.Println("Achat annulé.")
				return
			}
		case 2:
			fmt.Println("Vous souhaitez vendre une Peau de Troll ? (1.Oui 2.Non)")
			var lootype int
			if _, err := fmt.Scan(&lootype); err != nil {
				fmt.Println("Erreur lors de la saisie.")
				return
			}
			if lootype == 1 {
				itemToSell = "Peau de Troll"
			} else {
				fmt.Println("Achat annulé.")
				return
			}
		case 3:
			fmt.Println("Vous souhaitez vendre un Cuir de Sanglier ? (1.Oui 2.Non)")
			var lootype int
			if _, err := fmt.Scan(&lootype); err != nil {
				fmt.Println("Erreur lors de la saisie.")
				return
			}
			if lootype == 1 {
				itemToSell = "Cuir de Sanglier"
			} else {
				fmt.Println("Achat annulé.")
				return
			}
		case 4:
			fmt.Println("Vous souhaitez vendre une Plume de Corbeau ? (1.Oui 2.Non)")
			var lootype int
			if _, err := fmt.Scan(&lootype); err != nil {
				fmt.Println("Erreur lors de la saisie.")
				return
			}
			if lootype == 1 {
				itemToSell = "Plume de Corbeau"
			} else {
				fmt.Println("Achat annulé.")
				return
			}
		default:
			fmt.Println("Type d'objet invalide.")
			return
		}

	}
	price, exists := m1.Prices[itemToSell]
	if !exists {
		fmt.Println("Désolé, le prix de vente de cet objet n'est pas défini.")
		return
	}
	if quantity, ok := Personnage.Inventory[itemToSell]; ok && quantity > 0 {
		// Vérifier si le marchand a assez d'argent pour acheter l'objet
		if m1.Gold >= price {
			m1.Inventory[itemToSell]++
			m1.Gold -= price
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

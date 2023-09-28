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
	for {
		selectitem, err := getInput("Que voulez-vous acheter ? (1.Potion/2.sort/3.Objet/0.Rien, je me suis trompé) ")
		if err != nil {
			fmt.Println("Erreur lors de la saisie. Veuillez entrer un nombre.")
			return
		}

		switch selectitem {
		case 1:
			potionType, err := getInput("Quel type de potion voulez-vous acheter ? (1.Soin/2.Poison/3.Mana) ")
			if err != nil {
				fmt.Println("Erreur lors de la saisie. Veuillez entrer un nombre.")
				return
			}
			switch potionType {
			case 1:
				itemToBuy = "Potion de soin"
			case 2:
				itemToBuy = "Potion de poison"
			case 3:
				itemToBuy = "Potion de mana"
			}

		case 2:
			spellType, err := getInput("Quel sort voulez-vous acheter ? (1. Boule de Feu /)")
			if err != nil {
				fmt.Println("Erreur lors de la saisie. Veuillez entrer un nombre.")
				return
			}
			switch spellType {
			case 1:
				itemToBuy = "Livre de sort : Boule de feu"
			}

		case 3:
			loot, err := getInput("Quel type de loot souhaitez-vous acheter ? (1 : Loup/ 2 : Troll/ 3 : Sanglier/ 4 : Corbeau)")
			if err != nil {
				fmt.Println("Erreur lors de la saisie. Veuillez entrer un nombre.")
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
			case 0:
				break
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
				personnage.RemoveZeroValueItems()
				price, exists := m1.Prices[itemToBuy]
				if exists {
					personnage.Gold -= price // Deduct the price from character's gold
					personnage.Inventory[itemToBuy]++
					fmt.Println("Vous avez acheté ", itemToBuy)
					if askYesNo("Voulez-vous acheter autrechose ?") {
						break
					} else {
						return
					}
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
}

func Sell(P1 *Personnage, m1 *Marchand) {

	for {

		var price int
		fmt.Println("Les objets disponibles à la vente sont :")
		SelleableItem := inventorySlice(P1.Inventory)
		for i, itemName := range SelleableItem {
			fmt.Printf("%d. %s\n", i+1, itemName)
		}
		var choix int
		fmt.Println("Quel objet souhaitez vous vendre ?")
		fmt.Scan(&choix)
		itemToSell := SelleableItem[choix-1]

		// Vérifiez si le personnage possède cet objet
		quantity, ok := P1.Inventory[itemToSell]
		if !ok || quantity <= 0 {
			fmt.Println("Désolé, vous n'avez pas cet objet en stock.")
			return
		}
		price, existePrix := Marchand1.Prices[itemToSell]
		if !existePrix {
			fmt.Printf("Désolé, nous n'achetons pas %s.\n", itemToSell)
			continue
		}
		// Vendre l'objet et mettre à jour l'inventaire et l'or du personnage
		P1.Inventory[itemToSell]--
		m1.Inventory[itemToSell]++
		Marchand1.Gold -= price
		P1.Gold += price
		P1.RemoveZeroValueItems()
		fmt.Printf("Vous avez vendu un(e) %s pour %d pièces d'or.\n", itemToSell, price)
		if askYesNo("Voulez-vous vendre autrechose ?") {
			continue
		} else {
			fmt.Println("Le marchand n'a pas assez d'argent pour acheter cet objet.")
			return

		}

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

func CalculerMontantTotal(items []string, quantite int, prices map[string]int) int {
	montantTotal := 0

	for _, itemName := range items {
		price, existe := prices[itemName]
		if !existe {
			fmt.Printf("Le prix de l'objet %s n'est pas défini.\n", itemName)
			continue
		}

		montantTotal += price * quantite
	}

	return montantTotal
}
func inventorySlice(inventory map[string]int) []string {
	// Créez une tranche (slice) des objets de l'inventaire
	var items []string
	for itemName := range inventory {
		items = append(items, itemName)
	}
	return items
}

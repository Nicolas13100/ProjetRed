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
			"Arc de l'aventurier":       1,
			"Epée de l'aventurier":      2,
			"Chapeau de l'aventurier":   1,
			"Tunique de l'aventurier":   1,
			"Jambiere de l'aventurier":  1,
			"Bottes de l'aventurier":    1,
			"Heaume en cuir":            1,
			"Plastron en cuir":          1,
			"Pantalon en cuir":          1,
			"Bottes en cuir":            1,
			"Augmentation d'inventaire": 3,
			"Potion de soin":            3,
			"Potion de poison":          3,
			"Potion de mana":            3,
			"Boule de feu":              1,
			"Stalactites glacés":        1,
			"Avalanche":                 1,
			"Tsunami":                   1,
			"Vase Aquatique":            1,
			"Extension de territoire : Sphère de l'Espace Infini": 1,
			"Picots de glace":     1,
			"Eruption Volcanique": 1,
			"Fourrure de Loup":    10,
			"Peau de Troll":       10,
			"Cuir de Sanglier":    10,
			"Plume de Corbeau":    10,
			"Heaume en métal":     1,
			"Plastron en métal":   1,
			"Pantalon en métal":   1,
			"Bottes en métal":     1,
			"Katana":              1,
			"Epée renforcé":       1,
			"Katana supérieur":    1,
			"Arc supérieur":       1,
		},
		Prices: map[string]int{
			"Potion de soin":           3,
			"Potion de poison":         6,
			"Potion de mana":           8,
			"Arc de l'aventurier":      15,
			"Epée de l'aventurier":     20,
			"Chapeau de l'aventurier":  25,
			"Tunique de l'aventurier":  35,
			"Jambiere de l'aventurier": 30,
			"Bottes de l'aventurier":   18,
			"Heaume en cuir":           50,
			"Plastron en cuir":         75,
			"Pantalon en cuir":         65,
			"Boule de feu":             25,
			"Stalactites glacés":       50,
			"Avalanche":                70,
			"Tsunami":                  75,
			"Vase Aquatique":           25,
			"Extension de territoire : Sphère de l'Espace Infini": 6999,
			"Picots de glace":           20,
			"Eruption Volcanique":       75,
			"Augmentation d'inventaire": 30,
			"Fourrure de Loup":          4,
			"Peau de Troll":             7,
			"Cuir de Sanglier":          3,
			"Plume de Corbeau":          1,
			"Bottes en cuir":            40,
			"Heaume en métal":           100,
			"Plastron en métal":         130,
			"Pantalon en métal":         115,
			"Bottes en métal":           80,
			"Katana":                    50,
			"Epée renforcé":             70,
			"Katana supérieur":          90,
			"Arc supérieur":             60,
		},
	}
)

func Buy(P1 *Personnage, m1 *Marchand) {

	for {
		ClearConsole()
		fmt.Println("Les objets disponibles à l'achat sont :")
		buyableItem := inventorySlice(m1.Inventory)
		for i, itemName := range buyableItem {
			price, existePrix := m1.Prices[itemName]
			if existePrix {
				fmt.Printf("%d. %s - Prix : %d pièces d'or\n", i+1, itemName, price)
			} else {
				fmt.Printf("Désolé, nous ne vendons pas %s.\n", itemName)
				continue
			}
		}
		var choix int
		fmt.Println("Quel objet souhaitez vous acheter ?")
		fmt.Scan(&choix)
		if choix == 0 {
			fmt.Println("Vous quittez la boutique.")
			return
		}
		if choix < 1 || choix > len(buyableItem) {
			fmt.Println("Choix invalide.")
			continue
		}
		itemToBuy := buyableItem[choix-1]
		// Vérifiez si le personnage possède cet objet
		quantity, ok := m1.Inventory[itemToBuy]
		if !ok || quantity <= 0 {
			fmt.Println("Désolé, nous n'avons pas cet objet en stock.")
			return
		}

		// Vendre l'objet et mettre à jour l'inventaire et l'or du personnage
		price, existePrix := m1.Prices[itemToBuy]
		if existePrix {
			P1.Inventory[itemToBuy]++
			m1.Inventory[itemToBuy]--
			m1.Gold += price
			P1.Gold -= price
			P1.RemoveZeroValueItems()
			fmt.Printf("Vous avez acheté un(e) %s pour %d pièces d'or.\n", itemToBuy, price)
		} else {
			fmt.Println("Vous n'avez pas assez d'argent pour acheter cet objet.")
			return
		}
		if askYesNo("Voulez-vous acheter autre-chose ?") {
			continue
		} else {
			break
		}
	}
}

func Sell(P1 *Personnage, m1 *Marchand) {

	for {
		ClearConsole()
		fmt.Println("Les objets disponibles à la vente sont :")

		SelleableItem := inventorySlice(P1.Inventory)
		for i, itemName := range SelleableItem {
			price, existePrix := m1.Prices[itemName]
			if existePrix {
				fmt.Printf("%d. %s - Prix : %d pièces d'or\n", i+1, itemName, price)

			}
		}
		var choix int
		fmt.Println("Quel objet souhaitez vous vendre ?")
		fmt.Scan(&choix)
		if choix == 0 {
			fmt.Println("Vous quittez la boutique.")
			return
		}
		itemToSell := SelleableItem[choix-1]

		// Vérifiez si le personnage possède cet objet
		quantity, ok := P1.Inventory[itemToSell]
		if !ok || quantity <= 0 {
			fmt.Println("Désolé, vous n'avez pas cet objet en stock.")
			return
		}
		price, existePrix := m1.Prices[itemToSell]
		if existePrix {
			P1.Inventory[itemToSell]--
			m1.Inventory[itemToSell]++
			m1.Gold -= price
			P1.Gold += price
			P1.RemoveZeroValueItems()

			fmt.Printf("Vous avez vendu un(e) %s pour %d pièces d'or.\n", itemToSell, price)
		} else {
			fmt.Println("Le marchand n'a pas assez d'argent pour acheter cet objet.")
			return
		}
		if askYesNo("Voulez-vous acheté autre-chose ?") {
			continue
		} else {
			break
		}
	}
}

func askYesNo(question string) bool {
	var choice int
	for {

		fmt.Print(question + " (1.Oui/2.Non): ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			return choice == 1
		case 2:
			return choice == 0
		default:
			fmt.Println("Je n'ai pas compris")
			continue
		}
	}
}

func inventorySlice(inventory map[string]int) []string {
	// Créez une tranche (slice) des objets de l'inventaire
	var items []string
	for itemName := range inventory {
		items = append(items, itemName)
	}
	return items
}

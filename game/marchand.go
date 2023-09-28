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
			"Arc de l'aventurier":          15,
			"Epée de l'aventurier":         20,
			"Heaume en cuir":               25,
			"Plastron en cuir":             35,
			"Pantalon en cuir":             30,
			"Rangers":                      18,
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
			"Potion de soin":               3,
			"Potion de poison":             6,
			"Potion de mana":               8,
			"Arc de l'aventurier":          15,
			"Epée de l'aventurier":         20,
			"Heaume en cuir":               25,
			"Plastron en cuir":             35,
			"Pantalon en cuir":             30,
			"Rangers":                      18,
			"Livre de sort : Boule de feu": 25,
			"Augmentation d'inventaire":    30,
			"Fourrure de Loup":             4,
			"Peau de Troll":                7,
			"Cuir de Sanglier":             3,
			"Plume de Corbeau":             1,
		},
	}
)

func Buy(P1 *Personnage, m1 *Marchand) {

	for {
		fmt.Println("Les objets disponibles à la vente sont :")
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
			if askYesNo("Voulez-vous acheté autre-chose ?") {
				continue
			} else {
				fmt.Println("Vous n'avez pas assez d'argent pour acheter cet objet.")
				return
			}
		}
	}
}

func Sell(P1 *Personnage, m1 *Marchand) {

	for {
		fmt.Println("Les objets disponibles à la vente sont :")
		fmt.Println(P1.Inventory)
		SelleableItem := inventorySlice(P1.Inventory)
		for _, itemName := range SelleableItem {
			price, existePrix := m1.Prices[itemName]
			if existePrix {
				fmt.Printf("%s - Prix : %d pièces d'or\n", itemName, price)

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
		if equipment, ok := P1.EquipementMap[itemToSell]; ok {
			// Améliorez les statistiques de l'équipement existant
			reconditionedEquipment := improveEquipmentStats(equipment)

			reconditionedEquipment.Name = equipment.Name + " reconditionné(e)"

			price, existePrix := m1.Prices[itemToSell]
			if existePrix {
				P1.Inventory[itemToSell]--
				m1.Inventory[itemToSell]++
				m1.Gold -= price
				P1.Gold += price

				P1.RemoveZeroValueItems()
				fmt.Printf("Vous avez vendu un(e) %s pour %d pièces d'or.\n", itemToSell, price)
				if askYesNo("Voulez-vous vendre autrechose ?") {
					continue
				}
				// Vendre l'objet et mettre à jour l'inventaire et l'or du personnage

			} else {
				fmt.Println("Le marchand n'a pas assez d'argent pour acheter cet objet.")
				return
			}
		}
	}
}

func askYesNo(question string) bool {
	var choice int
	fmt.Print(question + " (1.Oui/2.Non): ")
	fmt.Scan(&choice)
	return choice == 1
}

func inventorySlice(inventory map[string]int) []string {
	// Créez une tranche (slice) des objets de l'inventaire
	var items []string
	for itemName := range inventory {
		items = append(items, itemName)
	}
	return items
}
func improveEquipmentStats(originalEquipment Equipment) Equipment {
	improvedEquipment := originalEquipment
	// Augmentez les statistiques de l'équipement d'origine
	improvedEquipment.HPBonus = int(float64(originalEquipment.HPBonus) * 1.1)
	improvedEquipment.DefBonus = int(float64(originalEquipment.DefBonus) * 1.1)
	improvedEquipment.InitiativeBonus = int(float64(originalEquipment.InitiativeBonus) * 1.1)
	// Vous pouvez ajouter d'autres statistiques à améliorer si nécessaire
	return improvedEquipment
}

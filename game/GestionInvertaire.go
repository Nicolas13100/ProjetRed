package game

import "fmt"

func (P1 *Personnage) FightInventory() {
	for {
		ClearConsole()
		fmt.Println("\nInventaire:")
		fmt.Println("Cash :", P1.Gold)
		for key, value := range P1.Inventory {
			fmt.Printf("%s: %d \n", key, value)
		}
		fmt.Println("\nQue voulez-vous faire ?")
		fmt.Println("1. Utiliser une potion de soin")
		fmt.Println("2. Utiliser une potion de poison")
		fmt.Println("3. Utiliser une potion de mana")
		fmt.Println("0. Retourner en arrière")

		var input int
		fmt.Scan(&input)

		switch input {
		case 1:
			if count, ok := P1.Inventory["Potion de soin"]; ok && count > 0 {
				TakePot(P1)
				P1.InventoryUsed = true
				P1.AtkUsed = false
			} else {
				fmt.Println("Vous n'avez pas de Potion de soin dans votre inventaire.")
			}
		case 2:
			if count, ok := P1.Inventory["Potion de poison"]; ok && count > 0 {
				poisonPot(P1)
				P1.InventoryUsed = true
				P1.AtkUsed = false
			} else {
				fmt.Println("Vous n'avez pas de Potion de poison dans votre inventaire.")
			}
		case 3:
			if count, ok := P1.Inventory["Potion de mana"]; ok && count > 0 {
				ManaPot(P1)
				P1.InventoryUsed = true
				P1.AtkUsed = false
			} else {
				fmt.Println("Vous n'avez pas de Potion de mana dans votre inventaire.")
			}

		case 0:
			P1.InventoryUsed = false
			P1.AtkUsed = false
			return

		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func (P1 *Personnage) BaseInventory() {
	for {
		ClearConsole()
		fmt.Println("\nInventaire:")
		fmt.Println("Cash :", P1.Gold)
		for key, value := range P1.Inventory {
			fmt.Printf("%s: %d \n", key, value)

		}
		fmt.Println("\nQue voulez-vous faire ?")
		fmt.Println("1. Sélectionner une potion de soin")
		fmt.Println("2. Sélectionner une potion de poison")
		fmt.Println("3. Sélectionner une potion de mana")
		fmt.Println("4. Consulter les sorts acquis")
		fmt.Println("5. Equipements")
		fmt.Println("0. Retourner en arrière")

		var input int
		fmt.Scan(&input)

		switch input {
		case 1:
			if count, ok := P1.Inventory["Potion de soin"]; ok && count > 0 {
				TakePot(P1)
			} else {
				fmt.Println("Vous n'avez pas de Potion de soin dans votre inventaire.")
			}
		case 2:
			if count, ok := P1.Inventory["Potion de poison"]; ok && count > 0 {
				poisonPot(P1)
			} else {
				fmt.Println("Vous n'avez pas de Potion de poison dans votre inventaire.")
			}
		case 3:
			if count, ok := P1.Inventory["Potion de mana"]; ok && count > 0 {
				ManaPot(P1)
			} else {
				fmt.Println("Vous n'avez pas de Potion de mana dans votre inventaire.")
			}
		case 4:
		case 5:
			fmt.Println("Que voulez-vous faire ?")
			fmt.Println("1. Equiper de l'équipement")
			fmt.Println("2. Désequiper de l'équipement")
			fmt.Println("3. Afficher les bonus d'équipements")
			fmt.Println("4.Quitter")
			var input int
			fmt.Scan(&input)
			switch input {
			case 1:
				fmt.Println("Que souhaitez-vous équiper ?")
				fmt.Println("1. Casque")
				fmt.Println("2. Armure")
				fmt.Println("3. Pied")
				fmt.Println("4. Retour ")
				var input int
				fmt.Scan(&input)
				switch input {
				case 1:
					P1.EquiperHead()
				case 2:
					P1.EquiperBody()
				case 3:
					P1.EquiperLeg()
				case 4:
					continue

				}
			case 2:
				fmt.Println("Que souhaitez-vous désequiper ?")
				fmt.Println("1. Casque")
				fmt.Println("2. Armure")
				fmt.Println("3. Pied")
				fmt.Println("4. Retour ")
				var input int
				fmt.Scan(&input)
				switch input {
				case 1:
					P1.DesequiperHead()
				case 2:
					P1.DesequiperBody()
				case 3:
					P1.DesequiperLeg()
				case 4:
				}

			case 3:
				fmt.Println("Quel équipement souhaitez-vous afficher ?")
				fmt.Println("1. Casque")
				fmt.Println("2. Armure")
				fmt.Println("3. Pied")
				fmt.Println("4. Retour ")
				var input int
				fmt.Scan(&input)
				switch input {
				case 1:
					P1.ShowEquipBonus(input)
				case 2:
					P1.ShowEquipBonus(input)
				case 3:
					P1.ShowEquipBonus(input)

				}
			case 4:
				continue
			}
		case 0:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func (P1 *Personnage) LimiteInventory() bool {
	totalQuantity := 0
	for _, count := range P1.Inventory {
		totalQuantity += count
	}
	return totalQuantity < P1.InventoryCap
}

func (P1 *Personnage) UpgradeInventorySlot() {
	if P1.InventoryCap < 30 {
		P1.InventoryCap += 10
		return
	}
}

func (P1 *Personnage) RemoveZeroValueItems() {
	for key, value := range P1.Inventory {
		if value == 0 {
			delete(P1.Inventory, key)
		}
	}
}
func DropsToInventory(P1 *Personnage, itemDrop map[string]int) {
	for itemName, quantityDropped := range itemDrop {

		if existingQuantity, exists := P1.Inventory[itemName]; exists {

			P1.Inventory[itemName] = existingQuantity + quantityDropped
		} else {

			P1.Inventory[itemName] = quantityDropped
		}
	}
}

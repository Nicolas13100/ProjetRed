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

		case 0:
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
			fmt.Println("Voulez-vous 1.équiper ou 2.désequiper un equipement ? (3. retour)")
			var input int
			fmt.Scan(&input)
			switch input {
			case 1:
				fmt.Println("Que souhaitez-vous équiper ? 1. Casque / 2. Armure / 3. Pied / 4. retour ")
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
				fmt.Println("Que souhaitez-vous déquiper ? 1. Casque / 2. Armure / 3. Pied / 4. retour ")
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

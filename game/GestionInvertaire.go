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
				PlayerTurnTaken = true
			} else {
				fmt.Println("Vous n'avez pas de Potion de soin dans votre inventaire.")
			}
		case 2:
			if count, ok := P1.Inventory["Potion de poison"]; ok && count > 0 {
				poisonPot(P1)
				PlayerTurnTaken = true
			} else {
				fmt.Println("Vous n'avez pas de Potion de poison dans votre inventaire.")
			}
		case 3:
			if count, ok := P1.Inventory["Potion de mana"]; ok && count > 0 {
				ManaPot(P1)
				PlayerTurnTaken = true
			} else {
				fmt.Println("Vous n'avez pas de Potion de mana dans votre inventaire.")
			}

		case 0:
			PlayerTurnTaken = false
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func (P1 *Personnage) BaseInventory() {
	for {

		fmt.Println("\nInventaire:")
		fmt.Println("Cash :", P1.Gold)
		for key, value := range P1.Inventory {
			fmt.Printf("%s: %d \n", key, value)
		}
		fmt.Println("\nQue voulez-vous faire ?")
		fmt.Println("1. Sélectionner une potion de soin")
		fmt.Println("2. Sélectionner une potion de mana")
		fmt.Println("3. Consulter les sorts acquis")
		fmt.Println("4. Equipements")
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
			if count, ok := P1.Inventory["Potion de mana"]; ok && count > 0 {
				ManaPot(P1)
			} else {
				fmt.Println("Vous n'avez pas de Potion de mana dans votre inventaire.")
			}
		case 3:
			fmt.Println("Sort appris :")
			for i, spell := range P1.Spells {
				fmt.Printf("%d. %s (Type: %s, Dommages: %d, Coût en Mana: %d)\n", i+1, spell.Name, spell.Type, spell.Damage, spell.ManaCost)
			}
			waitForUserInput("Entrer 0 pour continuer")
		case 4:
			fmt.Println("Que voulez-vous faire ?")
			fmt.Println("1. Equiper de l'équipement")
			fmt.Println("2. Désequiper de l'équipement")
			fmt.Println("3. Afficher les équipements")
			fmt.Println("4.Quitter")
			var input int
			fmt.Scan(&input)
			switch input {
			case 1:
				equipableItems := ListEquipableItems(P1.Inventory)
				fmt.Println("Equipable Items:")
				for i, item := range equipableItems {
					fmt.Printf("%d. %s\n", i+1, item.Name)
				}
				fmt.Println("Que souhaitez-vous équiper ?")
				var input int
				fmt.Scan(&input)
				if input >= 1 && input <= len(equipableItems) {
					P1.Equip(equipableItems[input-1])
				} else {
					fmt.Println("Vous n'avez pas d'équipement sur cet emplacement.")
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

				case 2:

				case 3:

				case 4:
				}

			case 3:
				for i, equipements := range P1.Equipements {
					fmt.Printf("\n%d. %s :\n \nType: %s\nPV Bonus: %d\nAttaque Bonus: %d\nDéfense Bonus : %d\nInitiative Bonus : %d\n", i+1, equipements.Name, equipements.Type, equipements.HPBonus, equipements.AtkBonus, equipements.DefBonus, equipements.InitiativeBonus)
					waitForUserInput("Entrer 0  pour revenir au menu principal")
					continue
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

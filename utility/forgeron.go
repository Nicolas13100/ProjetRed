package utility

import (
	"ProjectRed/game"
	"fmt"
)

func Forgeron(personnage *game.Personnage) {
	var choice int
	fmt.Println("Bienvenue dans la forge, en quoi puis-je vous etres utile ? (1. Forge / 2. Autre)")
	fmt.Scan(&choice)
	switch choice {

	case 1:
		fmt.Println("Nous avons 3 objet forgeable, 1.Chapeau de l’aventurier, 2. Tunique de l’aventurier,3. Bottes de l’aventurier")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			var choice int
			fmt.Println("Attention, la fabrication des équipements nécessite des ressources, il faudra pour : Chapeau de l’aventurier,1 Plume de Corbeau et 1 Cuir de Sanglier")
			fmt.Println("Continuer ? 1.Oui 2. Non")
			fmt.Scan(&choice)
			switch choice {
			case 1:
				if count, ok := personnage.Inventory["Plume de Corbeau"]; ok && count > 0 {
					if count, ok := personnage.Inventory["Cuir de Sanglier"]; ok && count > 0 {
						if personnage.Gold >= 5 {
							if personnage.LimiteInventory() {
								personnage.Inventory["Chapeau de l'aventurier"]++
								personnage.Inventory["Plume de Corbeau"]--
								personnage.Inventory["Cuir de Sanglier"]--
								personnage.Gold = -5
							}
						} else {
							fmt.Println("Vous n'avez pas asser d'argent")
						}

					} else {
						fmt.Println("Vous n'avez pas asser de Cuir de Sanglier")

					}
				} else {
					fmt.Println("Vous n'avez pas asser de Plume de Corbeau")
				}
			case 2:
				return
			}
		case 2:
			var choice int
			fmt.Println("Attention, la fabrication des équipements nécessite des ressources, il faudra pour : Tunique de l’aventurier,2 Fourrure de loup et 1 Peau de Troll")
			fmt.Println("Continuer ? 1.Oui 2. Non")
			fmt.Scan(&choice)
			switch choice {
			case 1:
				if count, ok := personnage.Inventory["Fourrure de loup"]; ok && count > 1 {
					if count, ok := personnage.Inventory["Peau de Troll"]; ok && count > 0 {
						if personnage.Gold >= 5 {
							if personnage.LimiteInventory() {
								personnage.Inventory["Tunique de l'aventurier"]++
								personnage.Inventory["Fourrure de loup"] = -2
								personnage.Inventory["Peau de Troll"]--
								personnage.Gold = -5
							}
						} else {
							fmt.Println("Vous n'avez pas asser d'argent")
						}

					} else {
						fmt.Println("Vous n'avez pas asser de Peau de Troll")

					}
				} else {
					fmt.Println("Vous n'avez pas asser de Fourrure de loup")
				}
			case 2:
				return
			}
		case 3:
			var choice int
			fmt.Println("Attention, la fabrication des équipements nécessite des ressources, il faudra pour : Bouttes de l'aventurier,1 Fourrure de loup et 1 Cuir de Sanglier")
			fmt.Println("Continuer ? 1.Oui 2. Non")
			fmt.Scan(&choice)
			switch choice {
			case 1:
				if count, ok := personnage.Inventory["Fourrure de loup"]; ok && count > 0 {
					if count, ok := personnage.Inventory["Cuir de Sanglier"]; ok && count > 0 {
						if personnage.Gold >= 5 {
							if personnage.LimiteInventory() {
								personnage.Inventory["Bottes de l'aventurier"]++
								personnage.Inventory["Fourrure de loup"]--
								personnage.Inventory["Cuir de Sanglier"]--
								personnage.Gold = -5
							}
						} else {
							fmt.Println("Vous n'avez pas asser d'argent")
						}

					} else {
						fmt.Println("Vous n'avez pas asser de Cuir de Sanglier")

					}
				} else {
					fmt.Println("Vous n'avez pas asser de Plume de Corbeau")
				}
			case 2:
				return
			}
		}
	}
}

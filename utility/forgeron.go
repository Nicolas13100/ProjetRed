package utility

import (
	"ProjectRed/game"
	"fmt"
)

func Forgeron() {
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
				if count, ok := game.Inventory["Potion de soin"]; ok && count > 0 {

				}
			}
		}
	}
}

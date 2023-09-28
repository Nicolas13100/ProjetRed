package game

import (
	"fmt"
)

var (
	Forgeron1 = Forgeron{
		Name:       "Michel",
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
			"Failure":                  1,
			"Chapeau de l'aventurier":  1,
			"Arc de l'aventurier":      1,
			"Wadô Ichimonji":           1,
			"Epée de l'aventurier":     1,
			"Bottes de l'aventurier":   1,
			"Jambiere de l'aventurier": 1,
			"Tunique de l'aventurier":  1,
			"Heaume en cuir":           1,
			"Plastron en cuir":         1,
			"Pantalon en cuir":         1,
			"Bottes en cuir":           1,
			"Heaume en métal":          1,
			"Plastron en métal":        1,
			"Pantalon en métal":        1,
			"Bottes en métal":          1,
			"Katana":                   1,
			"Epée renforcé":            1,
			"Katana supérieur":         1,
			"Arc supérieur":            1,
		},
	}
)

func Forge(P1 *Personnage, f1 *Forgeron) {
	var choice int
	fmt.Println("Bienvenue dans la forge, en quoi puis-je vous être utile ? (1. Forge / 2. Recycler un item / 3. Quitter)")
	fmt.Scan(&choice)
	switch choice {

	case 1:
		ForgeableItems := P1.ListEquipableItems()
		for i, item := range ForgeableItems {
			fmt.Printf("%d. %s\n", i+1, item)
		}
		var choice int
		fmt.Print("Quel équipement souhaitez-vous forger ? (entrez le numéro) : ")
		fmt.Scan(&choice)
		// Récupérez l'objet choisi par l'utilisateur
		chosenItemName := ForgeableItems[choice-1]

		chosenItem := P1.GetItemByName(chosenItemName)

		materials := GetMaterialsFromEquipment(chosenItem)
		for material, quantity := range materials {
			if P1.Inventory[material] < 1 {
				fmt.Println("Vous n'avez pas les matériaux nécessaires pour forger cet item")
				return
			} else {
				P1.Inventory[material] -= quantity
			}
		}
		P1.AddItem(chosenItem.Name)

		fmt.Printf("Vous avez forgé %s et utilisé :\n", chosenItem.Name)
		for material, quantity := range materials {
			fmt.Printf("%s : %d\n", material, quantity)
		}

	case 2:
		recyclableItems := P1.ListEquipableItems()

		if len(recyclableItems) == 0 {
			fmt.Println("Vous n'avez pas d'objets recyclables.")
			return
		}

		fmt.Println("Objets recyclables :")
		for i, item := range recyclableItems {
			fmt.Printf("%d. %s\n", i+1, item)
		}

		// Demandez à l'utilisateur de choisir un objet à recycler
		var choice int
		fmt.Print("Quel équipement souhaitez-vous recycler ? (entrez le numéro) : ")
		fmt.Scan(&choice)
		// Récupérez l'objet choisi par l'utilisateur
		chosenItemName := recyclableItems[choice-1]

		chosenItem := P1.GetItemByName(chosenItemName)

		// Récupérez les matériaux utilisés pour créer l'arme et ajoutez-les à l'inventaire du joueur
		materials := GetMaterialsFromEquipment(chosenItem)
		for material, quantity := range materials {
			P1.Inventory[material] += (quantity / 2)
		}

		// Retirez l'objet recyclé de l'inventaire du joueur
		P1.RemoveItem(chosenItem.Name)

		fmt.Printf("Vous avez recyclé %s \n Vous avez récupéré les matériaux suivants :\n", chosenItem.Name)
		for material, quantity := range materials {
			fmt.Printf("%s : %d\n", material, quantity)
		}
	}

}
func (f *Forgeron) FListEquipableItems() []string {
	equipableItems := []string{}

	for item := range f.Inventory {
		switch item {
		case "Test", "Head Test2", "Head Test", "Failure", "Wadô Ichimonji", "Chapeau de l'aventurier", "Tunique de l'aventurier", "Jambiere de l'aventurier", "Bottes de l'aventurier", "Epée de l'aventurier", "Arc de l'aventurier":
			equipableItems = append(equipableItems, item)
		}
	}

	return equipableItems
}

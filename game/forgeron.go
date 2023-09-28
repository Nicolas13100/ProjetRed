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
		},
	}
)

func Forge(P1 *Personnage, f1 *Forgeron) {
	var choice int
	fmt.Println("Bienvenue dans la forge, en quoi puis-je vous être utile ? (1. Forge / 2. Recycler un item / 3. Quitter)")
	fmt.Scan(&choice)
	switch choice {

	case 1:

		var choice int
		fmt.Print("Quel équipement souhaitez-vous forger ? (entrez le numéro) : ")
		fmt.Scan(&choice)
		// Récupérez l'objet choisi par l'utilisateur

	case 2:

		// Demandez à l'utilisateur de choisir un objet à recycler
		var choice int
		fmt.Print("Quel équipement souhaitez-vous recycler ? (entrez le numéro) : ")
		fmt.Scan(&choice)
		// Récupérez l'objet choisi par l'utilisateur

		// Récupérez les matériaux utilisés pour créer l'arme et ajoutez-les à l'inventaire du joueur

		// Retirez l'objet recyclé de l'inventaire du joueur

	}

}

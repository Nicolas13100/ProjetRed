package game

import "fmt"

type Reducible interface {
	ReduceDurability(amount int)
}

func ApplyDamage(r Reducible, amount int) {
	r.ReduceDurability(amount)
}

func ReduceDurability(durability *int, amount int) {
	if *durability > 0 {
		*durability -= amount
		if *durability < 0 {
			*durability = 0
			fmt.Println("L'arme s'est brisée") // Assuming this is French for "The weapon has broken"
		}
	}
}

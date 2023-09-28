package game

import "fmt"

type Reducible interface {
	ReduceDurability(amount int)
}

func ApplyDamage(r Reducible, amount int) {
	r.ReduceDurability(amount)
}

func (w *Weapon) ReduceDurability(amount int) {
	if w.Durability > 0 {
		w.Durability -= amount
		if w.Durability < 0 {
			w.Durability = 0
			fmt.Println("L'arme c'est brisée")
		}
	}
}

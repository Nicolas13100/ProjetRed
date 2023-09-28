package game

import (
	"fmt"
	"math/rand"
	"time"
)

func GetRandomEquippedItem(p *Personnage) (string, error) {
	if len(p.EquippedItems) == 0 {
		return "", fmt.Errorf("aucun Objet Equipé")
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(p.EquippedItems))
	return p.EquippedItems[randomIndex], nil
}

func PickUpItem(itemKey string, p *Personnage) {
	newItem, ok := p.EquipementMap[itemKey]
	if !ok {
		fmt.Println("Item not found")
		return
	}
	newItem.Durability -= 2
	p.EquipementMap[itemKey] = newItem
}

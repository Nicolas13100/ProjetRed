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
		fmt.Println("No comprendo")
		return
	}
	fmt.Println(newItem)
	newItem.Durability -= 2
	fmt.Println(newItem)
	p.EquipementMap[itemKey] = newItem
	fmt.Println(p.EquipementMap)
}

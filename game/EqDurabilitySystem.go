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
		return
	}
	newItem.Durability -= 2
	p.EquipementMap[itemKey] = newItem
	if newItem.Durability <= 0 {
		item := p.GetItemByName(itemKey)
		fmt.Printf("Votre %s c'est brisé\n", itemKey)
		switch item.Type {
		case "Head":
			if p.Equipement.Head {
				p.UnequipItem(p.Head)
			}
		case "Armor":
			if p.Equipement.Armor {
				p.UnequipItem(p.Armor)
			}
		case "Legs":
			if p.Equipement.Legs {
				p.UnequipItem(p.Legs)
			}
		case "Boots":
			if p.Equipement.Boots {
				p.UnequipItem(p.Feets)
			}
		}
	}
}

func (p *Personnage) RepairAll() {
	for key, item := range p.BrokenEquipementMap {
		item.Durability = item.DurabilityMax // Set the durability to a desired value (e.g., 100 for full durability)
		p.EquipementMap[key] = item
	}

	for key, item := range p.EquipementMap {
		item.Durability = item.DurabilityMax // Set the durability to a desired value (e.g., 100 for full durability)
		p.EquipementMap[key] = item
	}
}
func (p *Personnage) RepairSelected() {

}

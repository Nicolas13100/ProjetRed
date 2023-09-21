package game

func (p1 *Personnage) LimiteInventory() bool {
	totalQuantity := 0
	for _, count := range p1.Inventory {
		totalQuantity += count
	}
	return totalQuantity < p1.InventoryCap
}

func (p1 *Personnage) UpgradeInventorySlot() {
	if p1.InventoryCap < 30 {
		p1.InventoryCap += 10
		return
	}
}

func (p1 *Personnage) RemoveZeroValueItems() {
	for key, value := range p1.Inventory {
		if value == 0 {
			delete(p1.Inventory, key)
		}
	}
}

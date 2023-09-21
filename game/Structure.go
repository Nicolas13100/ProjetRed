package game

type Personnage struct {
	Name         string
	Race         string
	Equipement   Equipment
	Level        int
	Xp           int
	XpMax        int
	HpMax        int
	Hp           int
	Spells       []string
	Inventory    map[string]int
	Skills       []string
	Gold         int
	InventoryCap int
	Mana         int
	ManaMax      int
	Atk          int
	Defense      int
	Initiative   int
}

type Equipment struct {
	Head            bool
	Body            bool
	Leg             bool
	HPBonus         int
	AtkBonus        int
	DefBonus        int
	InitiativeBonus int
}

type Marchand struct {
	Name         string
	Race         string
	Level        int
	Xp           int
	XpMax        int
	HpMax        int
	Hp           int
	Spells       []string
	Inventory    map[string]int
	Skills       []string
	Gold         int
	InventoryCap int
	Mana         int
	ManaMax      int
	Atk          int
	Defense      int
	Initiative   int
	Prices       map[string]int
}

type Monstre struct {
	Name       string
	HpMax      int
	Hp         int
	Atk        int
	Defense    int
	Initiative int
	XpDrop     int
}

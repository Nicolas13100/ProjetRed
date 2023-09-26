package game

type Personnage struct {
	Name          string
	Race          string
	Equipement    Equipment
	Equipements   []Equipment
	Level         int
	Xp            int
	XpMax         int
	HpMax         int
	Hp            int
	Inventory     map[string]int
	InventoryUsed bool
	AtkUsed       bool
	Spells        []Spell
	Gold          int
	InventoryCap  int
	Mana          int
	ManaMax       int
	Atk           int
	Defense       int
	Initiative    int
	x, y          int
}

type Equipment struct {
	Name            string
	Type            string
	Head            bool
	Body            bool
	Belt            bool
	Leg             bool
	Boots           bool
	Rings1          bool
	Ring2           bool
	Necklace        bool
	RHWeapon        bool
	LHWeapon        bool
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
	Niveau     int
	Name       string
	HpMax      int
	Hp         int
	ManaMax    int
	Mana       int
	Atk        int
	Defense    int
	Initiative int
	XpDrop     int
	Spells     []Spell
	x, y       int
	Type       string
	ItemDrop   map[string]int
}

type Spell struct {
	Name     string
	Type     string
	Damage   int
	ManaCost int
}

type Carte struct {
	x, y       int
	hasExit    bool
	hasReturn  bool
	floor      int
	HasMonster bool
	HasJar     bool
	Jars       []Jar
	Monsters   []Monstre
}

type Jar struct {
	Contenu     []Item
	ContentType ContentType
	x, y        int
	HasLoot     bool
	HasMonster  bool
	HasEvent    bool
}
type ContentType int

const (
	Loot ContentType = iota
	Monster
	Event
)

type Item struct {
	Name      string
	Quantity  int
	IsLoot    bool
	IsMonster bool
	IsEvent   bool
}

type SaveData struct {
	Level     int
	Stats     *Personnage
	Equipment Equipment
}

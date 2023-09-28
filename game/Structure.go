package game

type Personnage struct {
	Name                   string
	Race                   string
	Equipement             Equipment
	EquipementMap          map[string]Equipment
	EquipementDurabilities map[string]int
	Level                  int
	Xp                     int
	XpMax                  int
	HpMax                  int
	Hp                     int
	Inventory              map[string]int
	InventoryUsed          bool
	AtkUsed                bool
	Spells                 []Spell
	Gold                   int
	InventoryCap           int
	Mana                   int
	ManaMax                int
	Atk                    int
	Defense                int
	Initiative             int
	x, y                   int
	Weapon                 map[string]Equipment
	Head                   map[string]Equipment
	Armor                  map[string]Equipment
	Hands                  map[string]Equipment
	Legs                   map[string]Equipment
	Feets                  map[string]Equipment
	Rings1                 map[string]Equipment
	Rings2                 map[string]Equipment
	Necklace               map[string]Equipment
}

type Equipment struct {
	Name            string
	Type            string
	Materials       map[string]int
	Head            bool
	Armor           bool
	Hands           bool
	Belt            bool
	Legs            bool
	Boots           bool
	Rings1          bool
	Ring2           bool
	Necklace        bool
	Weapon          bool
	HPBonus         int
	AtkBonus        int
	DefBonus        int
	InitiativeBonus int
	Durability      int
	DurabilityMax   int
}

type Weapon struct {
	Name            string
	Type            string
	DefBonus        int
	Durability      int
	DurabilityMax   int
	HPBonus         int
	InitiativeBonus int
	AtkBonus        int
}

type Head struct {
	Name            string
	Type            string
	DefBonus        int
	Durability      int
	DurabilityMax   int
	HPBonus         int
	InitiativeBonus int
}
type Armor struct {
	Name            string
	Type            string
	DefBonus        int
	Durability      int
	DurabilityMax   int
	HPBonus         int
	InitiativeBonus int
}
type Hands struct {
	Name            string
	Type            string
	DefBonus        int
	Durability      int
	DurabilityMax   int
	HPBonus         int
	InitiativeBonus int
}
type Legs struct {
	Name            string
	Type            string
	DefBonus        int
	Durability      int
	DurabilityMax   int
	HPBonus         int
	InitiativeBonus int
}
type Feets struct {
	Name            string
	Type            string
	DefBonus        int
	Durability      int
	DurabilityMax   int
	HPBonus         int
	InitiativeBonus int
}
type Rings1 struct {
	Name            string
	Type            string
	DefBonus        int
	Durability      int
	DurabilityMax   int
	HPBonus         int
	InitiativeBonus int
}
type Rings2 struct {
	Name            string
	Type            string
	DefBonus        int
	Durability      int
	DurabilityMax   int
	HPBonus         int
	InitiativeBonus int
}
type Necklace struct {
	Name            string
	Type            string
	DefBonus        int
	Durability      int
	DurabilityMax   int
	HPBonus         int
	InitiativeBonus int
}
type Shield struct {
	Name          string
	Type          string
	DefBonus      int
	Durability    int
	DurabilityMax int
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

type Forgeron struct {
	Name       string
	Race       string
	Level      int
	Xp         int
	XpMax      int
	HpMax      int
	Hp         int
	Inventory  map[string]int
	Gold       int
	Mana       int
	ManaMax    int
	Atk        int
	Defense    int
	Initiative int
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

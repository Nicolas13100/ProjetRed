package game

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Personnage struct {
	Name         string
	Race         string
	Equipement   Equipment
	Level        int
	Xp           int
	XpMax        int
	HpMax        int
	Hp           int
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

func ClearConsole() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func isValidName(name string, maxLetters int) bool {
	if len(name) > maxLetters {
		return false
	}

	for _, char := range name {
		if !unicode.IsLetter(char) {
			return false
		}
	}

	return true
}
func CharCreation() *Personnage {
	equipement := Equipment{
		Head:            false,
		Body:            false,
		Leg:             false,
		HPBonus:         0,
		AtkBonus:        0,
		DefBonus:        0,
		InitiativeBonus: 0,
	}
	var name string
	var race string
	for {
		ClearConsole()
		fmt.Println("Création de votre personnage :")
		// Demander à l'utilisateur de choisir son nom
		fmt.Print("Entrez votre nom : ")
		fmt.Scan(&name)
		maxLetters := 10
		if !isValidName(name, maxLetters) {
			fmt.Printf("Le nom doit contenir uniquement des lettres et avoir au maximum %d caractères.\n", maxLetters)
			continue
		}
		// Vérifier si le nom contient uniquement des lettres
		if !OnlyLetters(name) {
			fmt.Println("Le nom doit contenir uniquement des lettres.")
			continue
		}

		// Confirm the name
		fmt.Printf("Votre nom est %s, est-ce correct ? (1.Oui/2.Non) ", name)
		var confirmation int
		fmt.Scan(&confirmation)

		if confirmation == 1 {
			// Create a Title case converter
			tc := cases.Title(language.English)

			// Convert the name to Title case
			name = tc.String(strings.ToLower(name))
			break
		} else if confirmation == 2 {
			// If user chooses "no", restart the loop
			continue
		}
	}

	for {
		ClearConsole()
		// Demander à l'utilisateur de choisir sa race (vous pouvez adapter cette partie selon vos besoins)
		fmt.Print("Choisissez votre race : \n")
		fmt.Print("Humain : Vous commencez avec 100 PV Max\n")
		fmt.Print("Elfe : Vous commencez avec 80 PV Max\n")
		fmt.Println("Nain : Vous commencez avec 120 PV Max")
		fmt.Scan(&race)
		switch race {
		case "Humain":
			p1 := &Personnage{
				Name:         name,
				Race:         race,
				Equipement:   equipement,
				Level:        1,
				Xp:           0,
				XpMax:        100,
				HpMax:        100,
				Hp:           50,
				Gold:         100,
				InventoryCap: 10,
				Skills:       []string{"coup de poing"},
				Inventory: map[string]int{
					"Potion de soin":   3,
					"Potion de poison": 3,
					"Potion de mana":   3,
				},
				Mana:       40,
				ManaMax:    100,
				Atk:        5,
				Defense:    0,
				Initiative: 10,
			}
			return p1

		case "Elfe":
			p1 := &Personnage{
				Name:         name,
				Race:         race,
				Equipement:   equipement,
				Level:        1,
				Xp:           0,
				XpMax:        100,
				HpMax:        80,
				Hp:           40,
				Gold:         100,
				InventoryCap: 10,
				Skills:       []string{"coup de poing"},
				Inventory: map[string]int{
					"Potion de soin":   3,
					"Potion de poison": 3,
					"Potion de mana":   3,
				},
				Mana:       50,
				ManaMax:    100,
				Atk:        5,
				Defense:    0,
				Initiative: 20,
			}
			return p1

		case "Nain":
			p1 := &Personnage{
				Name:         name,
				Race:         race,
				Equipement:   equipement,
				Level:        1,
				Xp:           0,
				XpMax:        100,
				HpMax:        120,
				Hp:           60,
				Gold:         100,
				InventoryCap: 10,
				Skills:       []string{"coup de poing"},
				Inventory: map[string]int{
					"Potion de soin":   3,
					"Potion de poison": 3,
					"Potion de mana":   3,
				},
				Mana:       40,
				ManaMax:    90,
				Atk:        5,
				Defense:    0,
				Initiative: 5,
			}
			fmt.Printf("Bienvenue à toi : %s\n", name)
			return p1

		case "Dieu":
			p1 := &Personnage{
				Name:         name,
				Race:         race,
				Equipement:   equipement,
				Level:        99,
				HpMax:        999999,
				Hp:           999999,
				Gold:         99999999,
				InventoryCap: 99999999,
				Skills:       []string{"coup de poing"},
				Inventory: map[string]int{
					"Potion de soin":   99,
					"Potion de poison": 99,
					"Potion de mana":   99,
				},
				Mana:       999,
				ManaMax:    999,
				Atk:        999,
				Defense:    999,
				Initiative: 999,
			}
			fmt.Printf("Bienvenue à toi : %s\n", name)
			return p1
		default:
			fmt.Printf("Race inconnue : %s\n", race)
			continue
		}
	}
}

func OnlyLetters(s string) bool {
	for _, char := range s {
		if !('a' <= char && char <= 'z' || 'A' <= char && char <= 'Z') {
			return false
		}
	}
	return true
}

func (p1 *Personnage) DisplayInfo() {
	// Affichage des informations du Personnage p1

	for {
		ClearConsole()
		fmt.Println("Nom:", p1.Name)
		fmt.Println("Classe:", p1.Race)
		fmt.Println("Niveau:", p1.Level)
		fmt.Println("Points de vie maximum:", p1.HpMax+p1.Equipement.HPBonus)
		fmt.Println("Points de vie actuels:", p1.Hp+p1.Equipement.HPBonus)
		fmt.Println("Cash:", p1.Gold)
		fmt.Println("Sorts", p1.Skills)
		// Affichage de l'inventaire
		fmt.Println("\nType 'return' to come back to main menu")

		// Read user input
		var input string
		fmt.Scan(&input)

		// Check if the user wants to return to the main menu
		if input == "return" {
			break
		}
	}
}

func (p1 *Personnage) AccessInventory(m1 *Monstre, spells *Spell) {
	for {
		ClearConsole()
		fmt.Println("\nInventaire:")
		fmt.Println(p1.Gold)
		for key, value := range p1.Inventory {
			fmt.Printf("%s: %d \n", key, value)

		}
		fmt.Println("\nQue voulez-vous faire ?")
		fmt.Println("1. Sélectionner une potion de soin")
		fmt.Println("2. Sélectionner une potion de poison")
		fmt.Println("3. Sélectionner une potion de mana")
		fmt.Println("4. Sélectionner un livre de sort")
		fmt.Println("5. Equipements")
		fmt.Println("6. Retourner en arrière")

		var input int
		fmt.Scan(&input)

		switch input {
		case 1:
			if count, ok := p1.Inventory["Potion de soin"]; ok && count > 0 {
				TakePot(p1)
			} else {
				fmt.Println("Vous n'avez pas de Potion de soin dans votre inventaire.")
			}
		case 2:
			if count, ok := p1.Inventory["Potion de poison"]; ok && count > 0 {
				poisonPot(p1, m1)
			} else {
				fmt.Println("Vous n'avez pas de Potion de poison dans votre inventaire.")
			}
		case 3:
<<<<<<< HEAD
			if count, ok := p1.Inventory["Potion de mana"]; ok && count > 0 {
				ManaPot(p1, m1)
			} else {
				fmt.Println("Vous n'avez pas de Potion de mana dans votre inventaire.")
=======
			if count, ok := p1.Inventory["Livre de sort : Boule de feu"]; ok && count > 0 {
				SpellBook(p1, spells)
>>>>>>> 804a4a289e68669a8b79439eb72448d1884a6812
			}
		case 4:
			if count, ok := p1.Inventory["Livre de sort : Boule de feu"]; ok && count > 0 {
				SpellBook(p1, spells[])
			}
		case 5:
			fmt.Println("Voulez-vous 1.équiper ou 2.désequiper un equipement ? (3. retour)")
			var input int
			fmt.Scan(&input)
			switch input {
			case 1:
				fmt.Println("Que souhaitez-vous équiper ? 1. Casque / 2. Armure / 3. Pied / 4. retour ")
				var input int
				fmt.Scan(&input)
				switch input {
				case 1:
					p1.EquiperHead()
				case 2:
					p1.EquiperBody()
				case 3:
					p1.EquiperLeg()
				case 4:
					continue

				}
			case 2:
				fmt.Println("Que souhaitez-vous déquiper ? 1. Casque / 2. Armure / 3. Pied / 4. retour ")
				var input int
				fmt.Scan(&input)
				switch input {
				case 1:
					p1.DesequiperHead()
				case 2:
					p1.DesequiperBody()
				case 3:
					p1.DesequiperLeg()
				case 4:
					continue
				}
			case 3:
				continue
			}
		case 6:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func (p1 *Personnage) Dead() {
	if p1.Hp <= 0 {
		fmt.Printf("Vous êtes mort ! \n")
		p1.Hp = p1.HpMax / 2
		fmt.Printf("Vous avez été ressuscité avec %d PV\n", p1.Hp)
	}

}
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

func NewEquipement(personnage *Personnage) *Equipment {
	return &Equipment{
		Head: false,
		Body: false,
		Leg:  false,
	}
}

func (p1 *Personnage) EquiperHead() {
	if p1.Inventory["Chapeau de l'aventurier"] > 0 {
		if p1.Equipement.Head {
			// Si le personnage a déjà un chapeau équipé, le remettre dans l'inventaire
			p1.Inventory["Chapeau de l'aventurier"]++
		}
		p1.Equipement.Head = true
		p1.Equipement.HPBonus += 10
		delete(p1.Inventory, "Chapeau de l'aventurier")
	}
}

func (p1 *Personnage) EquiperBody() {
	if p1.Inventory["Tunique de l'aventurier"] > 0 {
		if p1.Equipement.Body {
			p1.Inventory["Tunique de l'aventurier"]++
		}
		p1.Equipement.Body = true
		p1.Equipement.HPBonus += 25
		delete(p1.Inventory, "Tunique de l'aventurier")
	}
}

func (p1 *Personnage) EquiperLeg() {
	if p1.Inventory["Bottes de l'aventurier"] > 0 {
		if p1.Equipement.Leg {
			p1.Inventory["Bottes de l'aventurier"]++
		}
		p1.Equipement.Leg = true
		p1.Equipement.HPBonus += 15
		delete(p1.Inventory, "Bottes de l'aventurier")
	}
}

func (p1 *Personnage) DesequiperHead() {
	if p1.Equipement.Head {
		p1.Inventory["Chapeau de l'aventurier"]++
		p1.Equipement.HPBonus -= 10
		p1.Equipement.Head = false
	}
}

func (p1 *Personnage) DesequiperBody() {
	if p1.Equipement.Body {
		p1.Inventory["Tunique de l'aventurier"]++
		p1.Equipement.HPBonus -= 25
		p1.Equipement.Body = false
	}
}

func (p1 *Personnage) DesequiperLeg() {
	if p1.Equipement.Leg {
		p1.Inventory["Bottes de l'aventurier"]++
		p1.Equipement.HPBonus -= 15
		p1.Equipement.Leg = false
	}
}

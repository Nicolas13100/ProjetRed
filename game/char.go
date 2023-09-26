package game

import (
	"fmt"
	"math/rand"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

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
	var Name string
	var Race string
	for {
		ClearConsole()
		text := "Création de votre personnage :\n Seulement des lettres de l'alphabet latin sont accepté, 10 lettres maximum\nEntrez votre nom : \n"
		centeredText := CenterText(text)
		fmt.Println(centeredText)
		fmt.Scan(&Name)
		maxLetters := 10
		if !IsValidName(Name, maxLetters) {
			text := "Le nom doit contenir uniquement des lettres et avoir au maximum %d caractères.\n"
			centeredText := CenterText(text)
			fmt.Printf(centeredText, maxLetters)
			continue
		}
		// Vérifier si le nom contient uniquement des lettres
		if !OnlyLetters(Name) {
			text := "Le nom doit contenir uniquement des lettres.\n"
			centeredText := CenterText(text)
			fmt.Println(centeredText)
			continue
		}

		// Confirm the name
		tc := cases.Title(language.English)
		Name = tc.String(strings.ToLower(Name))
		text1 := "Votre nom sera %s, est-ce correct ? (1.Oui/2.Non) \n"
		centeredText1 := CenterText(text1)
		fmt.Printf(centeredText1, Name)
		var confirmation int
		fmt.Scan(&confirmation)

		if confirmation == 1 {
			// Create a Title case converter
			tc := cases.Title(language.English)

			// Convert the name to Title case
			Name = tc.String(strings.ToLower(Name))
			break
		} else if confirmation == 2 {
			// If user chooses "no", restart the loop
			continue
		}
	}

	for {
		ClearConsole()
		// Demander à l'utilisateur de choisir sa race (vous pouvez adapter cette partie selon vos besoins)
		text := "Choisissez votre race : \nHumain : Vous commencez avec 100 PV Max, initiative moyenne\n Elfe : Vous commencez avec 80 PV Max, initiative elevé\nNain : Vous commencez avec 120 PV Max, initiative basse\n"
		centeredText := CenterText(text)
		fmt.Println(centeredText)
		fmt.Scan(&Race)
		switch Race {
		case "Humain":
			P1 := &Personnage{
				Name:         Name,
				Race:         Race,
				Equipement:   equipement,
				Equipements:  []Equipment{Chapeau, Tunique},
				Level:        1,
				Xp:           0,
				XpMax:        100,
				HpMax:        100,
				Hp:           50,
				Gold:         100,
				InventoryCap: 10,
				Spells:       []Spell{CoupdePoing},
				Inventory: map[string]int{
					"Potion de soin":   3,
					"Potion de poison": 3,
					"Potion de mana":   3,
				},
				InventoryUsed: false,
				AtkUsed:       false,
				Mana:          40,
				ManaMax:       100,
				Atk:           5,
				Defense:       0,
				Initiative:    10,
				x:             rand.Intn(gridSize),
				y:             rand.Intn(gridSize),
			}
			return P1

		case "Elfe":
			P1 := &Personnage{
				Name:         Name,
				Race:         Race,
				Equipement:   equipement,
				Level:        1,
				Xp:           0,
				XpMax:        100,
				HpMax:        80,
				Hp:           40,
				Gold:         100,
				InventoryCap: 10,
				Spells:       []Spell{CoupdePoing},
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
				x:          rand.Intn(gridSize),
				y:          rand.Intn(gridSize),
			}
			return P1

		case "Nain":
			P1 := &Personnage{
				Name:         Name,
				Race:         Race,
				Equipement:   equipement,
				Level:        1,
				Xp:           0,
				XpMax:        100,
				HpMax:        120,
				Hp:           60,
				Gold:         100,
				InventoryCap: 10,
				Spells:       []Spell{CoupdePoing},
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
				x:          rand.Intn(gridSize),
				y:          rand.Intn(gridSize),
			}
			return P1

		case "Dieu":
			P1 := &Personnage{
				Name:         Name,
				Race:         Race,
				Equipement:   equipement,
				Level:        99,
				HpMax:        999999,
				Hp:           999999,
				Gold:         9999999,
				InventoryCap: 99999999,
				Spells:       []Spell{CoupdePoing, fireball, iceBlast},
				Inventory: map[string]int{
					"Potion de soin":   99,
					"Potion de poison": 99,
					"Potion de mana":   99,
					"Fourrure de Loup": 99,
					"Peau de Troll":    99,
					"Cuir de Sanglier": 99,
					"Plume de Corbeau": 99,
				},
				Mana:       999,
				ManaMax:    999,
				Atk:        999,
				Defense:    999,
				Initiative: 999,
				x:          rand.Intn(gridSize),
				y:          rand.Intn(gridSize),
			}
			return P1
		default:
			fmt.Printf("Race inconnue : %s\n", Race)
			continue
		}
	}
}

func (P1 Personnage) DisplayInfo() {
	// Affichage des informations du Personnage p1

	for {
		ClearConsole()
		text1 := "\nNom : %s\nRace :%s\n\nNiveau : %d\n\nPoints de d'XP actuels : %d\nPoints de d'XP avant le prochain niveau : %d\n\nPoints de vie actuels : %d\nPoints de vie maximum : %d\nCash : %d\n \n"
		centeredText1 := CenterText(text1)
		fmt.Printf(centeredText1, P1.Name, P1.Race, P1.Level, P1.Hp+P1.Equipement.HPBonus, P1.HpMax+P1.Equipement.HPBonus, P1.Xp, P1.XpMax, P1.Gold)

		text11 := "\nSorts:"
		centeredText11 := CenterText(text11)
		fmt.Println(centeredText11)
		text12 := "\n%d.Spell Name : %s\nType :   %s\nDamage :  %d\nMana Cost : %d\n"
		centeredText12 := CenterText(text12)
		for i, spell := range P1.Spells {
			fmt.Printf(centeredText12, i+1, spell.Name, spell.Type, spell.Damage, spell.ManaCost)

		}
		text21 := "\nEquipements:"
		centeredText21 := CenterText(text21)
		fmt.Println(centeredText21)
		text22 := "\n%d. %s :\n Type: %s\n"
		centeredText22 := CenterText(text22)
		for i, equipements := range P1.Equipements {
			fmt.Printf(centeredText22, i+1, equipements.Name, equipements.Type)
		}

		text := "\nType 0 to come back to main menu"
		centeredText := CenterText(text)
		fmt.Println(centeredText)

		var input int
		fmt.Scan(&input)
		if input == 0 {
			break
		}
	}
}

func (p1 *Personnage) Dead() {
	if p1.Hp <= 0 {
		text := "Vous avez perdu connaissance !\nVous vous retrouve face a une ombre, impossible de savoir de quoi il s'agit\nL'ombre vous propose un marché, 5% de vos pièce d'Or en enchange de votre resurection en pleine santé.\n"
		centeredText := CenterText(text)
		fmt.Println(centeredText)
		if askYesNo("Acceptez-vous ?") && p1.Gold > 0 {
			p1.Hp = p1.HpMax
			p1.Mana = p1.ManaMax
			p1.Gold -= p1.Gold * 5 / 100
			fmt.Printf("Vous avez été ressuscité avec %d PV et %d Mana\n", p1.Hp, p1.Mana)
			fmt.Printf("Vous avez été ressuscité avec %d pièce d'Or\n", p1.Gold)
			waitForUserInput("Entrer 0 pour continuer ...")
		} else if askYesNo("Acceptez-vous ?") && p1.Gold < 0 {
			p1.Hp = p1.HpMax / 2
			p1.Mana = p1.ManaMax / 2
			fmt.Printf("Vous n'aviez pas de pièce d'Or")
			fmt.Printf("Vous avez été ressuscité avec %d PV et %d Mana\n", p1.Hp, p1.Mana)
			waitForUserInput("Entrer 0 pour continuer ...")
		} else {
			p1.Hp = p1.HpMax / 2
			p1.Mana = p1.ManaMax / 2
			fmt.Printf("Vous avez été ressuscité avec %d PV et %d Mana\n", p1.Hp, p1.Mana)
			waitForUserInput("Entrer 0 pour continuer ...")
		}
	}

}

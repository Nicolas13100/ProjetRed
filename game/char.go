package game

import (
	"fmt"
	"math/rand"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func CharCreation() Personnage {
	equipement := Equipment{
		Weapon:          false,
		Head:            false,
		Armor:           false,
		Hands:           false,
		Legs:            false,
		Boots:           false,
		Rings1:          false,
		Ring2:           false,
		Necklace:        false,
		HPBonus:         0,
		AtkBonus:        0,
		DefBonus:        0,
		InitiativeBonus: 0,
	}
	var Name string
	var Race string
	for {
		ClearConsole()
		text := "Création de votre personnage :\n Seuls les lettres de l'alphabet latin sont acceptées, 10 lettres maximum\nEntrez votre nom : \n"
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

RestartLoop:
	for {
		ClearConsole()
		// Demander à l'utilisateur de choisir sa race (vous pouvez adapter cette partie selon vos besoins)
		text := "Choisissez votre race : \nPourfendeur : Vous commencez avec 100 PV Max, initiative moyenne\n Exorciste : Vous commencez avec 80 PV Max, initiative elevé\nShinigami : Vous commencez avec 120 PV Max, initiative basse\n"
		centeredText := CenterText(text)
		fmt.Println(centeredText)
		fmt.Scan(&Race)

		tc := cases.Title(language.English)
		Race = tc.String(strings.ToLower(Race))
		text1 := "Votre race sera %s, est-ce correct ? (1.Oui/2.Non) \n"
		centeredText1 := CenterText(text1)
		fmt.Printf(centeredText1, Race)
		var confirmation int
		fmt.Scan(&confirmation)

		if confirmation == 1 {
			// Create a Title case converter
			tc := cases.Title(language.English)

			// Convert the name to Title case
			Race = tc.String(strings.ToLower(Race))
			break
		} else if confirmation == 2 {
			// If user chooses "no", restart the loop
			continue
		}

	}
	switch Race {
	case "Pourfendeur":
		P1 := Personnage{
			Name:                   Name,
			Race:                   Race,
			Equipement:             equipement,
			EquipementMap:          map[string]Equipment{},
			BrokenEquipementMap:    map[string]Equipment{},
			EquipementDurabilities: map[string]int{},
			EquippedItems:          []string{},
			Level:                  1,
			Xp:                     0,
			XpMax:                  100,
			HpMax:                  100,
			Hp:                     50,
			Gold:                   100,
			InventoryCap:           12,
			Spells:                 []Spell{CoupdePoing},
			Inventory: map[string]int{
				"Potion de soin":   3,
				"Potion de poison": 3,
				"Potion de mana":   3,
			},
			InventoryUsed:    false,
			AtkUsed:          false,
			Mana:             40,
			ManaMax:          100,
			Atk:              8,
			Defense:          2,
			Initiative:       12,
			x:                rand.Intn(gridSize),
			y:                rand.Intn(gridSize),
			Weapon:           map[string]Equipment{},
			EquippedWeapon:   "",
			Head:             map[string]Equipment{},
			EquippedHead:     "",
			Armor:            map[string]Equipment{},
			EquippedArmor:    "",
			Hands:            map[string]Equipment{},
			EquippedHands:    "",
			Legs:             map[string]Equipment{},
			Equippedlegs:     "",
			Feets:            map[string]Equipment{},
			EquippedFeets:    "",
			Rings1:           map[string]Equipment{},
			EquippedWRings1:  "",
			Rings2:           map[string]Equipment{},
			EquippedRings2:   "",
			Necklace:         map[string]Equipment{},
			EquippedNecklace: "",
		}
		return P1

	case "Exorciste":
		P1 := Personnage{
			Name:                   Name,
			Race:                   Race,
			Equipement:             equipement,
			EquipementMap:          map[string]Equipment{},
			BrokenEquipementMap:    map[string]Equipment{},
			EquipementDurabilities: map[string]int{},
			EquippedItems:          []string{},
			Level:                  1,
			Xp:                     0,
			XpMax:                  100,
			HpMax:                  80,
			Hp:                     40,
			Gold:                   100,
			InventoryCap:           10,
			Spells:                 []Spell{CoupdePoing},
			Inventory: map[string]int{
				"Potion de soin":   3,
				"Potion de poison": 3,
				"Potion de mana":   3,
			},
			Mana:             60,
			ManaMax:          100,
			Atk:              6,
			Defense:          0,
			Initiative:       20,
			x:                rand.Intn(gridSize),
			y:                rand.Intn(gridSize),
			Weapon:           map[string]Equipment{},
			EquippedWeapon:   "",
			Head:             map[string]Equipment{},
			EquippedHead:     "",
			Armor:            map[string]Equipment{},
			EquippedArmor:    "",
			Hands:            map[string]Equipment{},
			EquippedHands:    "",
			Legs:             map[string]Equipment{},
			Equippedlegs:     "",
			Feets:            map[string]Equipment{},
			EquippedFeets:    "",
			Rings1:           map[string]Equipment{},
			EquippedWRings1:  "",
			Rings2:           map[string]Equipment{},
			EquippedRings2:   "",
			Necklace:         map[string]Equipment{},
			EquippedNecklace: "",
		}
		return P1

	case "Shinigami":
		P1 := Personnage{
			Name:                   Name,
			Race:                   Race,
			Equipement:             equipement,
			EquipementMap:          map[string]Equipment{},
			BrokenEquipementMap:    map[string]Equipment{},
			EquipementDurabilities: map[string]int{},
			EquippedItems:          []string{},
			Level:                  1,
			Xp:                     0,
			XpMax:                  100,
			HpMax:                  120,
			Hp:                     60,
			Gold:                   100,
			InventoryCap:           10,
			Spells:                 []Spell{Shikai},
			Inventory: map[string]int{
				"Potion de soin":   3,
				"Potion de poison": 3,
				"Potion de mana":   3,
			},
			Mana:             50,
			ManaMax:          120,
			Atk:              7,
			Defense:          0,
			Initiative:       5,
			x:                rand.Intn(gridSize),
			y:                rand.Intn(gridSize),
			Weapon:           map[string]Equipment{},
			EquippedWeapon:   "",
			Head:             map[string]Equipment{},
			EquippedHead:     "",
			Armor:            map[string]Equipment{},
			EquippedArmor:    "",
			Hands:            map[string]Equipment{},
			EquippedHands:    "",
			Legs:             map[string]Equipment{},
			Equippedlegs:     "",
			Feets:            map[string]Equipment{},
			EquippedFeets:    "",
			Rings1:           map[string]Equipment{},
			EquippedWRings1:  "",
			Rings2:           map[string]Equipment{},
			EquippedRings2:   "",
			Necklace:         map[string]Equipment{},
			EquippedNecklace: "",
		}
		return P1

	case "Mentor":
		P1 := Personnage{
			Name:                   Name,
			Race:                   Race,
			Equipement:             equipement,
			EquipementMap:          map[string]Equipment{},
			BrokenEquipementMap:    map[string]Equipment{},
			EquipementDurabilities: map[string]int{},
			EquippedItems:          []string{},
			Level:                  99,
			HpMax:                  999999,
			Hp:                     999999,
			Gold:                   9999999,
			InventoryCap:           99999999,
			Spells:                 []Spell{CoupdePoing, Pistol, Shikai, BouleFeu, ProjectGlace, Infini, Eruption, Stalactites, Avalanche, VaseEau, Tsunami},
			Inventory: map[string]int{
				"Potion de soin":      99,
				"Potion de poison":    99,
				"Potion de mana":      99,
				"Fourrure de Loup":    99,
				"Peau de Troll":       99,
				"Cuir de Sanglier":    99,
				"Plume de Corbeau":    99,
				"Arc de l'aventurier": 1,
				"Failure":             1,
				"Wadô Ichimonji":      1,
			},
			Mana:             999,
			ManaMax:          999,
			Atk:              20,
			Defense:          10,
			Initiative:       999,
			x:                rand.Intn(gridSize),
			y:                rand.Intn(gridSize),
			Weapon:           map[string]Equipment{},
			EquippedWeapon:   "",
			Head:             map[string]Equipment{},
			EquippedHead:     "",
			Armor:            map[string]Equipment{},
			EquippedArmor:    "",
			Hands:            map[string]Equipment{},
			EquippedHands:    "",
			Legs:             map[string]Equipment{},
			Equippedlegs:     "",
			Feets:            map[string]Equipment{},
			EquippedFeets:    "",
			Rings1:           map[string]Equipment{},
			EquippedWRings1:  "",
			Rings2:           map[string]Equipment{},
			EquippedRings2:   "",
			Necklace:         map[string]Equipment{},
			EquippedNecklace: "",
		}
		return P1

	default:
		fmt.Printf("Race inconnue : %s\n", Race)
		waitForUserInput("Entrer 0 pour retourner au choix de race")
		goto RestartLoop
	}
}

func (P1 Personnage) DisplayInfo() {
	// Affichage des informations du Personnage p1

	for {
		ClearConsole()
		text1 := "\nNom : %s\nRace : %s\nNiveau : %d\nPoints d'XP actuels : %d\nPoints d'XP avant le prochain niveau : %d\n\nPoints de vie actuels : %d\nPoints de vie maximum : %d\n Attaque : %d\n Défense : %d\nInitiative : %d\nCash : %d\n"
		centeredText1 := CenterText(text1)
		fmt.Printf(centeredText1, P1.Name, P1.Race, P1.Level, P1.Xp, P1.XpMax, P1.Hp+P1.Equipement.HPBonus, P1.HpMax+P1.Equipement.HPBonus, P1.Atk+P1.Equipement.AtkBonus, P1.Defense+P1.Equipement.DefBonus, P1.Initiative+P1.Equipement.InitiativeBonus, P1.Gold)

		text11 := "\nSorts:"
		centeredText11 := CenterText(text11)
		fmt.Println(centeredText11)
		text12 := "\n%d.Nom du sort : %s\nType :   %s\n Dégâts :  %d\n Coût en mana : %d\n"
		centeredText12 := CenterText(text12)
		for i, spell := range P1.Spells {
			fmt.Printf(centeredText12, i+1, spell.Name, spell.Type, spell.Damage, spell.ManaCost)

		}
		text21 := "\nEquipements:"
		centeredText21 := CenterText(text21)
		fmt.Println(centeredText21)
		text22 := "\nArme :\n%s:\n Type: %s\n Attaque: %d\n Défense: %d\n HP: %d\n Initiative: %d\n Durabilité %d / %d \n"
		centeredText22 := CenterText(text22)
		for _, value := range P1.EquipementMap {
			fmt.Printf(centeredText22, value.Name, value.Type, value.AtkBonus, value.DefBonus, value.HPBonus, value.InitiativeBonus, value.Durability, value.DurabilityMax)
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
		text := "Vous avez perdu connaissance !\nVous vous retrouve face à une ombre, impossible de savoir de quoi il s'agit\nL'ombre vous propose un marché, 5% de vos pièces d'or en échange de votre résurrection à pleine santé.\n"
		centeredText := CenterText(text)
		fmt.Println(centeredText)
		if askYesNo("Acceptez-vous ?") && p1.Gold > 0 {
			p1.Hp = p1.HpMax
			p1.Mana = p1.ManaMax
			p1.Gold -= p1.Gold * 5 / 100
			fmt.Printf("Vous avez été ressuscité avec %d PV et %d Mana\n", p1.Hp, p1.Mana)
			fmt.Printf("Vous avez été ressuscité avec %d pièces d'or\n", p1.Gold)
			waitForUserInput("Entrer 0 pour continuer ...")
		} else if askYesNo("Acceptez-vous ?") && p1.Gold < 0 {
			p1.Hp = p1.HpMax / 2
			p1.Mana = p1.ManaMax / 2
			fmt.Printf("Vous n'aviez pas de pièces d'or")
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

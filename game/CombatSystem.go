package game

import (
	"fmt"
)

var Tours int
var PlayerTurnTaken = false

func Tutorial(p *Personnage) {
	ClearConsole()
	fmt.Println("Bienvenue dans le tutoriel")
	fmt.Println("Nous allons vous apprendre les bases du combat en tour par tour.")
	fmt.Println("Vous allez vous battre contre un goblin d'entrainement.")
	fmt.Println("Bonne chance")
	fmt.Println("-----------------------------------------------------------------------------")

	goblin := Monstre{
		Name:       "Goblin d'entrainement",
		HpMax:      30,
		Hp:         20,
		ManaMax:    10,
		Mana:       10,
		Atk:        8,
		Defense:    2,
		Initiative: 12,
	}
	if P1.Hp >= 10 {
		fmt.Println("Vous n'avez plus beaucoup de vie, prenez une potion de soin dans votre inventaire.")
	}
	Fight(p, &goblin)
}

func Entrainement(p *Personnage) {
	ClearConsole()
	text := "Bienvenue dans la zone d'entrainement\nAmusez-vous a tapé le mannequin d'entrainement\n-----------------------------------------------------------------------------\n"
	centeredText := CenterText(text)
	fmt.Println(centeredText)

	mannequin := Monstre{
		Name:       "Mannequin d'entrainement",
		Niveau:     1,
		HpMax:      30,
		Hp:         30,
		Defense:    5,
		Initiative: 0,
		XpDrop:     20,
	}
	Fight(p, &mannequin)

}

func Fight(p *Personnage, Monstre1 *Monstre) {
	Tours = 1

	fmt.Println("Début du combat")
	for p.Hp > 0 && Monstre1.Hp > 0 {
		fmt.Printf("Vous êtes au tour %d\n", Tours)
		if p.Initiative >= Monstre1.Initiative {
			if Tours == 1 {
				fmt.Println("Vous avez l'initiative et attaquez en premier !")
			}
			for { // Player's turn
				charTurn(p, Monstre1)
				if PlayerTurnTaken {
					PlayerTurnTaken = false // Reset playerTurnTaken for the next turn
					break
				}
			}
			// Check if the Monster is defeated
			if Monstre1.Hp <= 0 {
				Monstre1.DeadMonstre(p)
				break
			}
			MonsterPattern(Monstre1)
			attaqueMonstre := Monstre1.Atk - p.Defense
			if attaqueMonstre < 0 {
				attaqueMonstre = 0
			}
			p.Hp -= attaqueMonstre

			fmt.Printf("%s attaque %s et lui inflige %d points de dégâts.\n", Monstre1.Name, p.Name, attaqueMonstre)

			// Check if the Player is defeated
			if p.Hp <= 0 {
				p.Dead()
				break
			}
		}
		if p.Initiative < Monstre1.Initiative {
			if Tours == 1 {
				fmt.Println("Le monstre est plus rapide que vous et attaque en premier !")
			}
			MonsterPattern(Monstre1)
			// Monster's turn
			attaqueMonstre := Monstre1.Atk - (p.Defense + p.Equipement.DefBonus)
			if attaqueMonstre < 0 {
				attaqueMonstre = 0
			}
			p.Hp -= attaqueMonstre

			fmt.Printf("%s attaque %s et lui inflige %d points de dégâts.\n", Monstre1.Name, p.Name, attaqueMonstre)

			// Check if the Player is defeated
			if p.Hp <= 0 {
				fmt.Printf("%s est vaincu!\n", p.Name)
				p.Dead()
				break
			}
			for { // Player's turn
				charTurn(p, Monstre1)
				if PlayerTurnTaken {
					PlayerTurnTaken = false // Reset playerTurnTaken for the next turn
					break
				}
			}

			// Check if the Monster is defeated
			if Monstre1.Hp <= 0 {
				Monstre1.DeadMonstre(p)
				break
			}

			Tours++
			fmt.Println("-----------------------------------------------------------------------------")
		}
	}
}

func GetTours() int {
	return Tours + 1
}

func charTurn(p *Personnage, Monstre1 *Monstre) {
	fmt.Printf("Vous avez %d Pv / %d Pv \n", p.Hp, p.HpMax)
	fmt.Printf("Vous avez %d de mana restant / %d mana \n", p.Mana, p.ManaMax)
	fmt.Printf("Il reste %d Pv / %d Pv à l'ennemi \n", Monstre1.Hp, Monstre1.HpMax)
	fmt.Printf("Il reste %d de mana restant / %d mana à l'ennemi \n", Monstre1.Mana, Monstre1.ManaMax)

	fmt.Println("C'est votre tour, que voulez-vous faire ?")
	fmt.Println("1 : Attaquer")
	fmt.Println("2 : Ouvrir l'inventaire")
	fmt.Println("3 : Utiliser un sort")
	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("Attaque basique")
		attaqueJoueur := (p.Atk + p.Equipement.AtkBonus) - Monstre1.Defense
		if attaqueJoueur < 0 {
			attaqueJoueur = 0
		}

		fmt.Printf("Vous avez infligé %d points de dégats\n", attaqueJoueur)
		Monstre1.Hp -= attaqueJoueur
		PlayerTurnTaken = true
	case 2:
		p.FightInventory(Monstre1)

	case 3:
		fmt.Println("Sélectionnez un sort :")
		for i, spell := range p.Spells {
			fmt.Printf("%d. %s (Type: %s, Dégâts: %d, Coût en Mana: %d)\n", i+1, spell.Name, spell.Type, spell.Damage, spell.ManaCost)

		}
		var selectedSpellIndex int
		fmt.Print("Entrez le numéro du sort que vous souhaitez lancer : ")
		fmt.Scan(&selectedSpellIndex)
		if selectedSpellIndex >= 1 && selectedSpellIndex <= len(p.Spells) {
			selectedSpell := p.Spells[selectedSpellIndex-1]
			p.CastSpell(selectedSpell, Monstre1)
			PlayerTurnTaken = true

		} else {
			fmt.Println("Sélection de sort invalide.")
			return
		}
	}
}

func GagnerXp(p *Personnage, Monstre1 *Monstre) {
	p.Xp += Monstre1.XpDrop
	fmt.Printf("Vous avez gagné %d points d'XP \n", Monstre1.XpDrop)
	fmt.Printf("Vous avez %d points d'XP / %d points d'XP avant le prochain niveau \n", p.Xp, p.XpMax)
	if p.Xp >= p.XpMax {
		p.Level += 1
		surplusxp := p.Xp - p.XpMax
		p.Xp = 0
		p.XpMax += 10
		p.Xp += surplusxp
		fmt.Printf("Vous avez gagné un niveau et êtes au niveau : %d \n", p.Level)
		fmt.Printf("Vous avez %d points d'XP / %d points d'XP avant le prochain niveau \n", p.Xp, p.XpMax)
	}
}

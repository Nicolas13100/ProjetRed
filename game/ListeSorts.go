package game

import "fmt"

var (
	CoupdePoing = Spell{
		Name:     "Coup de Poing",
		Type:     "Physic",
		Damage:   6,
		ManaCost: 0,
	}
	Shikai = Spell{
		Name:     "Coup d'épée",
		Type:     "Physic",
		Damage:   10,
		ManaCost: 2,
	}
	Pistol = Spell{
		Name:     "Coup de Poing",
		Type:     "Physic",
		Damage:   6,
		ManaCost: 0,
	}

	Fireball = Spell{
		Name:     "Fireball",
		Type:     "Fire",
		Damage:   20,
		ManaCost: 10,
	}

	iceBlast = Spell{
		Name:     "Ice Blast",
		Type:     "Ice",
		Damage:   15,
		ManaCost: 8,
	}
)

func (P1 *Personnage) CastSpell(spell Spell, target *Monstre) {
	if P1.Mana >= spell.ManaCost {
		P1.Mana -= spell.ManaCost
		target.Hp -= spell.Damage
		fmt.Printf("%s casts %s on %s for %d damage!\n", P1.Name, spell.Name, target.Name, spell.Damage)
	} else {
		fmt.Printf("%s n'a pas assez de mana pour lancer %s !\n", P1.Name, spell.Name)
	}
}

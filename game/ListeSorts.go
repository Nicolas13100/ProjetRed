package game

import "fmt"

var (
	CoupdePoing = Spell{
		Name:     "Coup de Poing",
		Type:     "Physique",
		Damage:   6,
		ManaCost: 0,
	}
	Shikai = Spell{
		Name:     "Shikai",
		Type:     "Physique",
		Damage:   10,
		ManaCost: 2,
	}
	Pistol = Spell{
		Name:     "Gomu Gomu no Pistol",
		Type:     "Physique",
		Damage:   6,
		ManaCost: 0,
	}

	BouleFeu = Spell{
		Name:     "Boule de feu",
		Type:     "Feu",
		Damage:   20,
		ManaCost: 7,
	}
	Eruption = Spell{
		Name:     "Eruption Volcanique",
		Type:     "Feu",
		Damage:   80,
		ManaCost: 35,
	}
	ProjectGlace = Spell{
		Name:     "Picots de glace",
		Type:     "Glace",
		Damage:   15,
		ManaCost: 8,
	}
	Stalactites = Spell{
		Name:     "Stalactites glacés",
		Type:     "Glace",
		Damage:   35,
		ManaCost: 15,
	}
	Avalanche = Spell{
		Name:     "Avalanche",
		Type:     "Glace",
		Damage:   70,
		ManaCost: 10,
	}

	VaseEau = Spell{
		Name:     "Vase Aquatique",
		Type:     "Eau",
		Damage:   20,
		ManaCost: 7,
	}

	Tsunami = Spell{
		Name:     "Tsunami",
		Type:     "Eau",
		Damage:   90,
		ManaCost: 50,
	}
	Infini = Spell{
		Name:     "Extension de territoire : Sphère de l'Espace Infini",
		Type:     "Infini",
		Damage:   999999,
		ManaCost: 999,
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

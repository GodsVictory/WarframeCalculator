package lib

import "errors"

type Weapon struct {
    Name string
    Damage []WeaponDamage
    AttackSpeed float64
    CritChance float64
    CritMulti float64
    StatusChance float64
	Mod []Mod
}

type WeaponDamage struct {
	Type string
	Value float64
}

func setupWeapons() (weapons []Weapon) {
	weapons = append(weapons, Weapon{
		Name: "Orthos Prime",
		AttackSpeed: 1.17,
		Damage: []WeaponDamage{
			{ Type: "impact", Value: 35.1 },
			{ Type: "puncture", Value: 35.1 },
			{ Type: "slash", Value: 163.8 },
		},
		CritChance: 0.24,
		CritMulti: 2.2,
		StatusChance: 0.36,
	})
	weapons = append(weapons, Weapon{
		Name: "Paracesis",
		AttackSpeed: 0.917 + 0.15,
		Damage: []WeaponDamage{
			{ Type: "impact", Value: 48.8 },
			{ Type: "puncture", Value: 17.8 },
			{ Type: "slash", Value: 155.4 },
		},
		CritChance: 0.31,
		CritMulti: 2.6,
		StatusChance: 0.22,
	})
	weapons = append(weapons, Weapon{
		Name: "Broken War",
		AttackSpeed: 1.0 + 0.25,
		Damage: []WeaponDamage{
			{ Type: "impact", Value: 18.7 },
			{ Type: "puncture", Value: 18.7 },
			{ Type: "slash", Value: 149.6 },
		},
		CritChance: 0.35,
		CritMulti: 2.2,
		StatusChance: 0.2,
	})
	weapons = append(weapons, Weapon{
		Name: "Rakta Dark Dagger",
		AttackSpeed: 1.0 + 0.5,
		Damage: []WeaponDamage{
			{ Type: "puncture", Value: 88 },
			{ Type: "slash", Value: 62 },
			{ Type: "radiation", Value: 96 },
		},
		CritChance: 0.12,
		CritMulti: 1.8,
		StatusChance: 0.3,
	})
	weapons = append(weapons, Weapon{
		Name: "Lesion",
		AttackSpeed: 1.0,
		Damage: []WeaponDamage{
			{ Type: "impact", Value: 47.4 },
			{ Type: "puncture", Value: 11.9 },
			{ Type: "slash", Value: 177.8 },
		},
		Mod: []Mod{
			{
				Name: "Lesion",
				Modifiers: []Modifier{
					{ Type: "toxin", Value: 1.0 },
					{ Type: "attackSpeed", Value: 0.15 },
				},
			},
		},
		CritChance: 0.15,
		CritMulti: 2.0,
		StatusChance: 0.37,
	})
	return
}

var Weapons []Weapon = setupWeapons()

func GetWeaponByName(name string) (weapon Weapon, err error) {
	for _, w := range Weapons {
		if w.Name == name {
			weapon = w
			return
		}
	}
	err = errors.New("Weapon not found")
	return
}
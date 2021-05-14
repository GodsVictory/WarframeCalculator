package lib

type Weapon struct {
    Name string
    Damage []WeaponDamage
    AttackSpeed float64
    CritChance float64
    CritMulti float64
    StatusChance float64
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
		AttackSpeed: 0.917,
		Damage: []WeaponDamage{
			{ Type: "impact", Value: 48.8 },
			{ Type: "puncture", Value: 17.8 },
			{ Type: "slash", Value: 155.4 },
		},
		CritChance: 0.31,
		CritMulti: 2.6,
		StatusChance: 0.22,
	})
	return
}

var Weapons []Weapon = setupWeapons()
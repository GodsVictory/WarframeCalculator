package lib

type Mod struct {
    Name string
    Modifiers []Modifier
}

type Modifier struct {
	Type string
	Value float64
}

var comboCount int = 12
// var Mods []Mod = setupMods()
var Mods []Mod = myMods()

func setupMods() (mods []Mod) {
	mods = append(mods, Mod{
		Name: "Primed Fever Strike",
		Modifiers: []Modifier{
			{ Type: "toxin", Value: 1.65 },
		},
	})
	mods = append(mods, Mod{
		Name: "Northern Wind",
		Modifiers: []Modifier{
			{ Type: "cold", Value: 0.9 },
		},
	})
	mods = append(mods, Mod{
		Name: "Molten Impact",
		Modifiers: []Modifier{
			{ Type: "heat", Value: 0.9 },
		},
	})
	mods = append(mods, Mod{
		Name: "Shocking Touch",
		Modifiers: []Modifier{
			{ Type: "electricity", Value: 0.9 },
		},
	})
	mods = append(mods, Mod{
		Name: "Sacrificial Steel",
		Modifiers: []Modifier{
			{ Type: "critChance", Value: 2.2 },
		},
	})
	mods = append(mods, Mod{
		Name: "Organ Shatter",
		Modifiers: []Modifier{
			{ Type: "critMulti", Value: 0.9 },
		},
	})
	mods = append(mods, Mod{
		Name: "Carnis Mandible",
		Modifiers: []Modifier{
			{ Type: "slash", Value: 0.9 },
			{ Type: "statusChance", Value: 0.6 },
		},
	})
	mods = append(mods, Mod{
		Name: "Primed Pressure Point",
		Modifiers: []Modifier{
			{ Type: "base", Value: 1.65 },
		},
	})
	mods = append(mods, Mod{
		Name: "Berserker",
		Modifiers: []Modifier{
			{ Type: "attackSpeedMulti", Value: 0.75 },
		},
	})
	mods = append(mods, Mod{
		Name: "Weeping Wounds",
		Modifiers: []Modifier{
			{ Type: "statusChance", Value: 0.4 * float64(comboCount - 1)},
		},
	})
	mods = append(mods, Mod{
		Name: "Blood Rush",
		Modifiers: []Modifier{
			{ Type: "critChance", Value: 0.6 * float64(comboCount - 1)},
		},
	})
	mods = append(mods, Mod{
		Name: "Condition Overload",
		Modifiers: []Modifier{
			{ Type: "base", Value: 1.2 },
		},
	})
	mods = append(mods, Mod{
		Name: "Vicious Frost",
		Modifiers: []Modifier{
			{ Type: "cold", Value: 0.6 },
			{ Type: "statusChance", Value: 0.6 },
		},
	})
	mods = append(mods, Mod{
		Name: "Virulent Scourge",
		Modifiers: []Modifier{
			{ Type: "toxin", Value: 0.6 },
			{ Type: "statusChance", Value: 0.6 },
		},
	})
	mods = append(mods, Mod{
		Name: "Volcanic Edge",
		Modifiers: []Modifier{
			{ Type: "heat", Value: 0.6 },
			{ Type: "statusChance", Value: 0.6 },
		},
	})
	mods = append(mods, Mod{
		Name: "Voltaic Strike",
		Modifiers: []Modifier{
			{ Type: "electricity", Value: 0.6 },
			{ Type: "statusChance", Value: 0.6 },
		},
	})
	// mods = append(mods, Mod{
	// 	Name: "Primed Fury",
	// 	Modifiers: []Modifier{
	// 		{ Type: "attackSpeed", Value: 0.55 },
	// 	},
	// })
	return
}
func myMods() (mods []Mod) {
	mods = append(mods, Mod{
		Name: "Fever Strike",
		Modifiers: []Modifier{
			{ Type: "toxin", Value: 0.9 },
		},
	})
	mods = append(mods, Mod{
		Name: "Northern Wind",
		Modifiers: []Modifier{
			{ Type: "cold", Value: 0.9 },
		},
	})
	mods = append(mods, Mod{
		Name: "Molten Impact",
		Modifiers: []Modifier{
			{ Type: "heat", Value: 0.9 },
		},
	})
	mods = append(mods, Mod{
		Name: "Shocking Touch",
		Modifiers: []Modifier{
			{ Type: "electricity", Value: 0.9 },
		},
	})
	mods = append(mods, Mod{
		Name: "True Steel",
		Modifiers: []Modifier{
			{ Type: "critChance", Value: 1.2 },
		},
	})
	mods = append(mods, Mod{
		Name: "Organ Shatter",
		Modifiers: []Modifier{
			{ Type: "critMulti", Value: 0.9 },
		},
	})
	mods = append(mods, Mod{
		Name: "Carnis Mandible",
		Modifiers: []Modifier{
			{ Type: "slash", Value: 0.9 },
			{ Type: "statusChance", Value: 0.6 },
		},
	})
	mods = append(mods, Mod{
		Name: "Pressure Point",
		Modifiers: []Modifier{
			{ Type: "base", Value: 1.2 },
		},
	})
	mods = append(mods, Mod{
		Name: "Berserker",
		Modifiers: []Modifier{
			{ Type: "attackSpeedMulti", Value: 0.75 },
		},
	})
	mods = append(mods, Mod{
		Name: "Weeping Wounds",
		Modifiers: []Modifier{
			{ Type: "statusChance", Value: 0.4 * float64(comboCount - 1)},
		},
	})
	mods = append(mods, Mod{
		Name: "Blood Rush",
		Modifiers: []Modifier{
			{ Type: "critChance", Value: 0.6 * float64(comboCount - 1)},
		},
	})
	mods = append(mods, Mod{
		Name: "Condition Overload",
		Modifiers: []Modifier{
			{ Type: "base", Value: 1.2 },
		},
	})
	mods = append(mods, Mod{
		Name: "Vicious Frost",
		Modifiers: []Modifier{
			{ Type: "cold", Value: 0.6 },
			{ Type: "statusChance", Value: 0.6 },
		},
	})
	mods = append(mods, Mod{
		Name: "Virulent Scourge",
		Modifiers: []Modifier{
			{ Type: "toxin", Value: 0.6 },
			{ Type: "statusChance", Value: 0.6 },
		},
	})
	return
}

func GetModArrays() (elementMods []Mod, otherMods []Mod) {
	for _, mod := range Mods {
		var isElement bool = false
		for _, modifier := range mod.Modifiers {
			for _, element := range []string{"cold", "toxin", "heat", "electricity"} {
				if modifier.Type == element {
					isElement = true
				}
			}
		}
		if isElement {
			elementMods = append(elementMods, mod)
		} else {
			otherMods = append(otherMods, mod)
		}
	}
	return
}
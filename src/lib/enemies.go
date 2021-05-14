package lib

import (
	"math"
)

type Enemy struct {
	Name string
	Faction string
	Health Health
	Armor Armor
	Shield Shield
	Level int
	Count int
	Status []Status
	Viral []Proc
	Corrosive []Proc
	Heat []Proc
}

type Health struct {
	Type string
	Value float64
}

type Armor struct {
	Type string
	Value float64
}

type Shield struct {
	Type string
	Value float64
}

type Status struct {
	Type string
	Value float64
	TicksLeft int
	Delay int
}

type Proc struct {
	Value float64
	TicksLeft int
}

func (e Enemy) GetHealthAtLevel(level int) (health float64) {
	if level < e.Level {
		level = e.Level
	}
	if level < 70 {
		health = e.Health.Value * (1 + 0.015 * math.Pow(float64(level - e.Level), 2))
	} else {
		health = e.Health.Value * (1 + ((24 * math.Sqrt(5)) / 5) * math.Pow(float64(level - e.Level), 0.5))
	}
	return
}

func (e Enemy) GetArmorAtLevel(level int) (armor float64) {
	if level < e.Level {
		level = e.Level
	}
	if level < 70 {
		armor = e.Armor.Value * (1 + 0.005 * math.Pow(float64(level - e.Level), 1.75))
	} else {
		armor = e.Armor.Value * (1 + 0.4 * math.Pow(float64(level - e.Level), 0.75))
	}
	return
}

func (e Enemy) GetShieldAtLevel(level int) (shield float64) {
	if level < e.Level {
		level = e.Level
	}
	if level < 70 {
		shield = e.Shield.Value * (1 + 0.02 * math.Pow(float64(level - e.Level), 1.75))
	} else {
		shield = e.Shield.Value * (1 + 1.6 * math.Pow(float64(level - e.Level), 0.75))
	}
	return
}

func (e Enemy) GetDamageModifier(damage Damage, corrosiveEffect float64, heatEffect float64) (damageModifier float64) {
	var armorBonus float64 = damage.GetBonus(e.Armor.Type)
	var armorPenalty float64 = -damage.GetPenalty(e.Armor.Type)
	var healthBonus float64 = damage.GetBonus(e.Health.Type)
	var healthPenalty float64 = -damage.GetPenalty(e.Health.Type)
	var shieldBonus float64 = damage.GetBonus(e.Shield.Type)
	var shieldPenalty float64 = -damage.GetPenalty(e.Shield.Type)
	var currentArmor float64 = e.Armor.Value * (1 - corrosiveEffect) * (1 - heatEffect)

    if e.Shield.Value > 0 && damage.Type != "toxin" {
		damageModifier = 1 + shieldBonus + shieldPenalty
	} else if currentArmor > 0 {
		damageModifier = (300 / (300 + currentArmor * (1 - armorBonus))) *
        	(1 + armorBonus + armorPenalty) *
        	(1 + healthBonus + shieldPenalty)
    } else {
		damageModifier = 1 + healthBonus + healthPenalty
	}
	return
}

func (e *Enemy) setStats(lvl int) {
	e.Health.Value = e.GetHealthAtLevel(lvl)
	e.Armor.Value = e.GetArmorAtLevel(lvl)
	e.Shield.Value = e.GetShieldAtLevel(lvl)
	return
}

func SpawnEnemies(lvl int) (enemies []Enemy) {
	enemies = append(enemies, Enemy{
		Name: "Charger",
		Faction: "infested",
		Health: Health{ Type: "infested", Value: 80 },
		Level: 1,
		Count: 40,
	})
	enemies = append(enemies, Enemy{
		Name: "Leaper",
		Faction: "infested",
		Health: Health{ Type: "infested", Value: 100 },
		Level: 1,
		Count: 40,
	})
	enemies = append(enemies, Enemy{
		Name: "Crawler",
		Faction: "infested",
		Health: Health{ Type: "infestedFlesh", Value: 50 },
		Level: 1,
		Count: 30,
	})
	enemies = append(enemies, Enemy{
		Name: "Mutalist Osprey",
		Faction: "infested",
		Health: Health{ Type: "infestedFlesh", Value: 200 },
		Level: 10,
		Count: 15,
	})
	enemies = append(enemies, Enemy{
		Name: "Ancient Healer",
		Faction: "infested",
		Health: Health{ Type: "fossil", Value: 400 },
		Level: 1,
		Count: 25,
	})
	enemies = append(enemies, Enemy{
		Name: "Brood Mother",
		Faction: "infested",
		Health: Health{ Type: "fossil", Value: 700 },
		Level: 12,
		Count: 10,
	})
	enemies = append(enemies, Enemy{
		Name: "Juggernaut",
		Faction: "infested",
		Health: Health{ Type: "infested", Value: 3500 },
		Armor: Armor{ Type: "ferrite", Value: 200 },
		Level: 15,
		Count: 1,
	})
	enemies = append(enemies, Enemy{
		Name: "Butcher",
		Faction: "grineer",
		Health: Health{ Type: "cloned", Value: 50 },
		Armor: Armor{ Type: "ferrite", Value: 5 },
		Level: 1,
		Count: 10,
	})
	enemies = append(enemies, Enemy{
		Name: "Lancer",
		Faction: "grineer",
		Health: Health{ Type: "cloned", Value: 100 },
		Armor: Armor{ Type: "ferrite", Value: 100 },
		Level: 1,
		Count: 50,
	})
	enemies = append(enemies, Enemy{
		Name: "Elite Lancer",
		Faction: "grineer",
		Health: Health{ Type: "cloned", Value: 150 },
		Armor: Armor{ Type: "alloy", Value: 200 },
		Level: 1,
		Count: 40,
	})
	enemies = append(enemies, Enemy{
		Name: "Heavy Gunner",
		Faction: "grineer",
		Health: Health{ Type: "cloned", Value: 300 },
		Armor: Armor{ Type: "alloy", Value: 500 },
		Level: 8,
		Count: 40,
	})
	enemies = append(enemies, Enemy{
		Name: "Hyekka Master",
		Faction: "grineer",
		Health: Health{ Type: "cloned", Value: 650 },
		Armor: Armor{ Type: "ferrite", Value: 50 },
		Level: 1,
		Count: 15,
	})
	enemies = append(enemies, Enemy{
		Name: "Nox",
		Faction: "grineer",
		Health: Health{ Type: "cloned", Value: 350 },
		Armor: Armor{ Type: "alloy", Value: 250 },
		Level: 1,
		Count: 10,
	})
	enemies = append(enemies, Enemy{
		Name: "Napalm",
		Faction: "grineer",
		Health: Health{ Type: "cloned", Value: 600 },
		Armor: Armor{ Type: "alloy", Value: 500 },
		Level: 6,
		Count: 20,
	})
	enemies = append(enemies, Enemy{
		Name: "Crewman",
		Faction: "corpus",
		Health: Health{ Type: "flesh", Value: 60 },
		Armor: Armor{ Type: "shield", Value: 150 },
		Level: 1,
		Count: 50,
	})
	enemies = append(enemies, Enemy{
		Name: "Corpus Tech",
		Faction: "corpus",
		Health: Health{ Type: "flesh", Value: 700 },
		Armor: Armor{ Type: "proto", Value: 250 },
		Level: 15,
		Count: 15,
	})
	enemies = append(enemies, Enemy{
		Name: "MOA",
		Faction: "corpus",
		Health: Health{ Type: "robotic", Value: 60 },
		Shield: Shield{ Type: "shield", Value: 150 },
		Level: 1,
		Count: 25,
	})
	enemies = append(enemies, Enemy{
		Name: "Oxium Osprey",
		Faction: "corpus",
		Health: Health{ Type: "robotic", Value: 750 },
		Armor: Armor{ Type: "ferrite", Value: 40 },
		Shield: Shield{ Type: "shield", Value: 150 },
		Level: 5,
		Count: 10,
	})
	enemies = append(enemies, Enemy{
		Name: "Isolator Bursa",
		Faction: "corpus",
		Health: Health{ Type: "robotic", Value: 1200 },
		Armor: Armor{ Type: "alloy", Value: 200 },
		Shield: Shield{ Type: "shield", Value: 700 },
		Level: 1,
		Count: 1,
	})
	for i := range enemies {
		enemies[i].setStats(lvl)
	}
	return
}

func tickProcs(procs []Proc) {
	for _, proc := range procs {
		proc.TicksLeft -= 1
	}
}

func getViralProcs(procs []Proc) (value float64) {
	for _, proc := range procs {
		if proc.TicksLeft > 0 {
			value += proc.Value
		}
	}
	if (value > 3.25) {
		value = 3.25
	}
	return
}

func addViralProc(procs []Proc, statusChance float64, distribution float64, attackSpeed float64) (proc Proc) {
	if getViralProcs(procs) < 1 {
		proc.Value = (statusChance * distribution) / attackSpeed
	} else {
		proc.Value = (0.25 * statusChance * distribution) / attackSpeed
	}
	proc.TicksLeft = 6
	return
}

func getCorrosiveProcs(procs []Proc) (value float64) {
	for _, proc := range procs {
		if proc.TicksLeft > 0 {
			value += proc.Value
		}
	}
	if (value > 0.8) {
		value = 0.8
	}
	return
}

func addCorrosiveProc(procs []Proc, statusChance float64, distribution float64, attackSpeed float64) (proc Proc) {
	if getCorrosiveProcs(procs) < 0.26 {
		proc.Value = (0.26 * statusChance * distribution) / attackSpeed
	} else {
		proc.Value = (0.06 * statusChance * distribution) / attackSpeed
	}
	proc.TicksLeft = 8
	return
}

func getHeatProcs(procs []Proc) (value float64) {
	for _, proc := range procs {
		if proc.TicksLeft > 0 {
			value += proc.Value
		}
	}
	if (value > 0.5) {
		value = 0.5
	}
	return
}

func addHeatProc(procs []Proc, statusChance float64, distribution float64, attackSpeed float64) (proc Proc) {
	proc.Value = (0.1 * statusChance * distribution) / attackSpeed
	proc.TicksLeft = 8
	return
}

func (e *Enemy) Hit(baseDamage float64, baseModifier float64, distribution float64, damage Damage, statusChance float64, statusDuration float64, avgDamageMulti float64, attackSpeed float64, factionBonus float64) (damageInflicted float64, avgHit float64, dot float64) {
	var viralEffect float64 = getViralProcs(e.Viral)
	var corrosiveEffect float64 = getCorrosiveProcs(e.Corrosive)
	var heatEffect float64 = getHeatProcs(e.Heat)

	var currentDot float64
	for i := range e.Status {
		if e.Status[i].TicksLeft >= 0 && e.Status[i].Type == damage.Type {
			if e.Status[i].Delay == 0 {
				currentDot += e.Status[i].Value
			} else {
				e.Status[i].Delay -= 1
			}
			e.Status[i].TicksLeft -= 1
		}
	}

	var damageValue float64 = damage.Value * avgDamageMulti * (1 + factionBonus)
	var damageModifier float64 = e.GetDamageModifier(damage, corrosiveEffect, heatEffect)
	if e.Shield.Value > 0 && damage.Type != "toxin" {
		e.Shield.Value -= ((damageValue + currentDot) * damageModifier) / attackSpeed
		damageInflicted = ((damageValue + currentDot) * damageModifier) / attackSpeed
		avgHit = damageValue * damageModifier
		dot = currentDot * damageModifier
	} else {
		e.Health.Value -= ((damageValue + currentDot) * damageModifier * (1 + viralEffect)) / attackSpeed
		damageInflicted = ((damageValue + currentDot) * damageModifier * (1 + viralEffect)) / attackSpeed
		avgHit = damageValue * damageModifier * (1 + viralEffect)
		dot = currentDot * damageModifier * (1 + viralEffect)
	}

	// add new procs
	if damage.Dot.Damage > 0 {
		var status Status
		var moddedDamage float64 = baseDamage * (1 + baseModifier) * (1 + factionBonus)
		if damage.Type == "slash" {
			damageModifier = 1
		}
		var modifier float64 = 1
		if damage.Type != "slash" && damage.Type != "gas" {
			modifier = 1 + damage.Modifier
		}

		if damage.Type == "heat" {
			var heatIndex int = -1
			for i := range e.Status {
				if e.Status[i].Type == "heat" {
					heatIndex = i
				}
			}
			if heatIndex > -1 {
				e.Status[heatIndex].Value += (damage.Dot.Damage * moddedDamage * modifier * (1 + factionBonus) * distribution * statusChance * avgDamageMulti) / attackSpeed
				e.Status[heatIndex].TicksLeft = int(math.Floor(1 * (6 * (1 + statusDuration) - float64(damage.Dot.Delay))) + 1)
			} else {
				status.Value = (damage.Dot.Damage * moddedDamage * modifier * (1 + factionBonus) * distribution * statusChance * avgDamageMulti) / attackSpeed
				status.Type = damage.Type
				status.TicksLeft = int(math.Floor(1 * (6 * (1 + statusDuration) - float64(damage.Dot.Delay))) + 1)
				status.Delay = damage.Dot.Delay
				e.Status = append(e.Status, status)
			}
		} else {
			status.Value = (damage.Dot.Damage * moddedDamage * modifier * (1 + factionBonus) * distribution * statusChance * avgDamageMulti) / attackSpeed
			status.Type = damage.Type
			status.TicksLeft = int(math.Floor(1 * (6 * (1 + statusDuration) - float64(damage.Dot.Delay))) + 1)
			status.Delay = damage.Dot.Delay
			e.Status = append(e.Status, status)
		}
	}
	return
}

func (e *Enemy) Kill(damages []Damage, totalDamage float64, baseDamage float64, baseModifier float64, statusChance float64, statusDuration float64, avgDamageMulti float64, attackSpeed float64, factionBonus float64) (ttk int, avgDps float64, avgAvgHit float64, avgDot float64) {
	ttk = -1
	var totalDps float64
	var totalAvgHit float64
	var totalDot float64
	for e.Health.Value > 0 {
		ttk += 1
		var procViral Proc
		var procCorrosive Proc
		var procHeat Proc
		for _, damage := range damages {
			var distribution float64 = damage.Value / totalDamage
			dps, avgHit, dot := e.Hit(baseDamage, baseModifier, distribution, damage, statusChance, statusDuration, avgDamageMulti, attackSpeed, factionBonus)
			totalDps += dps
			totalAvgHit += avgHit
			totalDot += dot
			if damage.Type == "viral" {
				procViral = addViralProc(e.Viral, statusChance, distribution, attackSpeed)
			} else if damage.Type == "corrosive" {
				procCorrosive = addCorrosiveProc(e.Corrosive, statusChance, distribution, attackSpeed)
			} else if damage.Type == "heat" {
				procHeat = addHeatProc(e.Heat, statusChance, distribution, attackSpeed)
			}
		}
		if procViral.Value > 0 {
			e.Viral = append(e.Viral, procViral)
		}
		if procCorrosive.Value > 0 {
			e.Corrosive = append(e.Corrosive, procCorrosive)
		}
		if procHeat.Value > 0 {
			e.Heat = append(e.Heat, procHeat)
		}
	}
	avgDps = totalDps / float64(ttk+1)
	avgAvgHit = totalAvgHit / float64(ttk+1)
	avgDot = totalDot / float64(ttk+1)
	ttk *= e.Count
	return
}
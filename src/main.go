package main

import (
	"WFCalc/src/lib"
	"flag"
	"fmt"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	var allowedMods int
	flag.IntVar(&allowedMods, "mods", 8, "Mod Slots")
	var weaponName string
	flag.StringVar(&weaponName, "weapon", "Orthos Prime", "Weapon to test")
	var enemyLevelInput string
	flag.StringVar(&enemyLevelInput, "level", "10,25,50,75,100,150", "Enemy levels")
	flag.Parse()

	var enemyLevels []int
	for _, i := range strings.Split(enemyLevelInput, ",") {
		t, _ := strconv.Atoi(i)
		enemyLevels = append(enemyLevels, t)
	}

	var NGoRoutines = MaxParallelism()

	weapon, err := lib.GetWeaponByName(weaponName)
	if err != nil {
		fmt.Println(err)
		return
	}

	elementMods, otherMods := lib.GetModArrays()
	// var elementModsLen int = len(elementMods)

	var allModSets [][]lib.Mod
	var eModSet [][]lib.Mod = getCombinations(elementMods, 2)
	for i := 0; i <= 4; i++ {
		// var totalBuilds int = int(math.Pow(float64(elementModsLen), float64(i)))
		var otherModSets [][]lib.Mod = getCombinations(otherMods, allowedMods-i)
		for _, otherModSet := range otherModSets {
			if hasDup(otherModSet) {
				continue
			}

			if i == 1 {
				for _, emod := range elementMods {
					var modSet []lib.Mod
					modSet = append(otherModSet, emod)
					allModSets = append(allModSets, modSet)
				}
			} else if i == 2 {
				for _, ec := range eModSet {
					var modSet []lib.Mod
					for _, om := range otherModSet {
						modSet = append(modSet, om)
					}
					for _, ecm := range ec {
						modSet = append(modSet, ecm)
					}
					if !hasDup(modSet) {
						allModSets = append(allModSets, modSet)
					}
				}
			} else if i == 3 {
				for _, ec := range eModSet {
					for _, emod := range elementMods {
						var modSet []lib.Mod
						for _, om := range otherModSet {
							modSet = append(modSet, om)
						}
						for _, ecm := range ec {
							modSet = append(modSet, ecm)
						}
						modSet = append(modSet, emod)
						if !hasDup(modSet) {
							allModSets = append(allModSets, modSet)
						}
					}
				}
			} else if i == 4 {
				for _, ec1 := range eModSet {
					for _, ec2 := range eModSet {
						var modSet []lib.Mod
						for _, om := range otherModSet {
							modSet = append(modSet, om)
						}
						for _, ecm := range ec1 {
							modSet = append(modSet, ecm)
						}
						for _, ecm := range ec2 {
							modSet = append(modSet, ecm)
						}
						if !hasDup(modSet) {
							allModSets = append(allModSets, modSet)
						}
					}
				}
			}
			// for j := 0; j < totalBuilds; j++ {
			// 	var elementModSet []lib.Mod = ConvertToLengthBase(j, elementMods, elementModsLen, i)
			// 	if len(elementModSet) != i {continue}
			// 	var modSet []lib.Mod
			// 	for _, om := range otherModSet {
			// 		modSet = append(modSet, om)
			// 	}
			// 	for _, em := range elementModSet {
			// 		modSet = append(modSet, em)
			// 	}
			// 	allModSets = append(allModSets, modSet)
			// }
		}
	}
	// sort.Slice(allModSets, func(i, j int) bool {
	// 	for k := range allModSets[i] {
	// 		if allModSets[i][k].Name < allModSets[j][k].Name {
	// 			return true
	// 		}
	// 		if allModSets[i][k].Name > allModSets[j][k].Name {
	// 			return false
	// 		}
	// 	}
	// 	return false
	// })

	runtime.GOMAXPROCS(NGoRoutines)
	var ranks [][]Rank = make([][]Rank, NGoRoutines)
	var status []bool = make([]bool, NGoRoutines)
	var totalBuilds int = len(allModSets)

	for i := 0; i < NGoRoutines; i++ {
		go func(i int) {
			for j := i; j < totalBuilds; j += NGoRoutines {
				var rank Rank = simulate(weapon, allModSets[j], enemyLevels)
				ranks[i] = append(ranks[i], rank)
			}
			status[i] = true
		}(i)
	}

	var startTime int64 = time.Now().UnixNano() / int64(time.Second)
	for i := range status {
		for !status[i] {
			var buildsCompleted int = 0
			for _, v := range ranks {
				buildsCompleted += len(v)
			}
			var percentCompleted float64 = float64(buildsCompleted) / float64(totalBuilds) * 100
			var currentTime int64 = time.Now().UnixNano() / int64(time.Second)

			var bps int
			var eta int
			if currentTime-startTime > 0 {
				bps = buildsCompleted / int(currentTime-startTime)
				eta = (totalBuilds - buildsCompleted) / bps
			}
			fmt.Printf("\r[%s%s] %5.2f%% %14s %5dbps %3ds eta",
				strings.Repeat("#", int(percentCompleted)),
				strings.Repeat("-", 100-int(percentCompleted)),
				percentCompleted,
				fmt.Sprintf("(%d/%d)", buildsCompleted, totalBuilds),
				bps,
				eta,
			)
			time.Sleep(50 * time.Millisecond)
		}
	}
	var parsedRanks []Rank
	for _, rankArray := range ranks {
		for _, rank := range rankArray {
			parsedRanks = append(parsedRanks, rank)
		}
	}

	sort.SliceStable(parsedRanks, func(i, j int) bool {
		if parsedRanks[i].TTK == parsedRanks[j].TTK {
			return parsedRanks[i].DPS < parsedRanks[j].DPS
		} else {
			return parsedRanks[i].TTK < parsedRanks[j].TTK
		}
	})
	printRank(parsedRanks[0])
	return
}

func printRank(rank Rank) {
	fmt.Printf("\rBest Build%s\n", strings.Repeat(" ", 150))
	fmt.Printf("  Weapon: %s\n", rank.Weapon)
	fmt.Printf("  TTK: %d\n", rank.TTK)
	fmt.Printf("  DPS: %.2f\n", rank.DPS)
	fmt.Printf("  AvgHit: %.2f\n", rank.AvgHit)
	fmt.Printf("  DoT: %.2f\n", rank.DoT)
	fmt.Printf("  Attack Speed: %.2f\n", rank.attackSpeed)
	fmt.Printf("  Crit Chance: %.2f\n", rank.CritChance)
	fmt.Printf("  Crit Multi: %.2f\n", rank.CritMulti)
	fmt.Printf("  Status Chance: %.2f\n", rank.StatusChance)
	fmt.Println("  Build")
	for _, m := range rank.Set {
		fmt.Printf("    %s\n", m.Name)
	}
	fmt.Println("  Damage")
	for _, d := range rank.Damages {
		fmt.Printf("    %s: %.2f\n", d.Type, d.Value)
	}
	return
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func getModifierForType(Type string, set []lib.Mod) (value float64) {
	value = 0
	for _, mod := range set {
		for _, modifier := range mod.Modifiers {
			if modifier.Type == Type {
				value += modifier.Value
			}
		}
	}
	return
}

func getProcCount(weapon lib.Weapon, modSet []lib.Mod) (procCount int) {
	var damages []lib.Damage
	for _, d := range weapon.Damage {
		var nd lib.Damage = lib.GetDamageInfo(d.Type)
		damages = append(damages, nd)
	}

	var activeElements []string
	for _, mod := range modSet {
		for _, modifier := range mod.Modifiers {
			for _, element := range lib.Elements {
				if modifier.Type == element {
					activeElements = append(activeElements, element)
				}
			}
		}
	}

	for len(activeElements) >= 2 {
		var firstElement string = activeElements[0]
		activeElements = activeElements[1:]
		var secondElement string = activeElements[0]
		activeElements = activeElements[1:]
		for _, d := range lib.Damages {
			if contains(d.Mix, firstElement) && contains(d.Mix, secondElement) {
				var nd lib.Damage = lib.GetDamageInfo(d.Type)
				damages = append(damages, nd)
				break
			}
		}
	}
	if len(activeElements) > 0 {
		var nd lib.Damage = lib.GetDamageInfo(activeElements[0])
		nd.Modifier = getModifierForType(nd.Type, modSet)
		damages = append(damages, nd)
	}
	procCount = len(damages)
	return
}

func simulate(weapon lib.Weapon, inModSet []lib.Mod, enemyLevels []int) (stats Rank) {
	var modSet []lib.Mod
	for _, m := range inModSet {
		modSet = append(modSet, m)
	}
	for _, wm := range weapon.Mod {
		modSet = append(modSet, wm)
	}
	var damages []lib.Damage
	var moddedCritChance = weapon.CritChance * (1 + getModifierForType("critChance", modSet))
	var moddedCritMulti = weapon.CritMulti * (1 + getModifierForType("critMulti", modSet))
	var moddedStatusChance = weapon.StatusChance * (1 + getModifierForType("statusChance", modSet))
	var moddedStatusDuration = getModifierForType("statusDuration", modSet)
	var moddedAttackSpeed = 1 / (weapon.AttackSpeed * (1 + getModifierForType("attackSpeed", modSet)) * (1 + getModifierForType("attackSpeedMulti", modSet)))
	var avgDamageMulti = 1 + moddedCritChance*(moddedCritMulti-1)
	var baseModifier float64
	var procCount int = getProcCount(weapon, modSet)

	for _, mod := range modSet {
		for _, modifier := range mod.Modifiers {
			if mod.Name == "Condition Overload" {
				baseModifier += modifier.Value * float64(procCount)
			} else if modifier.Type == "base" {
				baseModifier += modifier.Value
			}
		}
	}

	var baseDamage float64
	for _, d := range weapon.Damage {
		var nd lib.Damage = lib.GetDamageInfo(d.Type)
		nd.Base = d.Value
		nd.Modifier = getModifierForType(nd.Type, modSet)
		nd.Value = d.Value * (1 + baseModifier) * (1 + nd.Modifier)
		damages = append(damages, nd)
		baseDamage += nd.Base
	}

	var activeElements []string
	for _, mod := range modSet {
		for _, modifier := range mod.Modifiers {
			for _, element := range lib.Elements {
				if modifier.Type == element {
					var exists bool = false
					for _, e := range activeElements {
						if element == e {
							exists = true
						}
					}
					if !exists {
						activeElements = append(activeElements, element)
					}
				}
			}
		}
	}

	for len(activeElements) >= 2 {
		var firstElement string = activeElements[0]
		activeElements = activeElements[1:]
		var secondElement string = activeElements[0]
		activeElements = activeElements[1:]
		for _, d := range lib.Damages {
			if contains(d.Mix, firstElement) && contains(d.Mix, secondElement) {
				var nd lib.Damage = lib.GetDamageInfo(d.Type)
				nd.Modifier = getModifierForType(firstElement, modSet) + getModifierForType(secondElement, modSet)
				nd.Base = baseDamage
				nd.Value = baseDamage * (1 + baseModifier) * (nd.Modifier)
				damages = append(damages, nd)
				break
			}
		}
	}
	if len(activeElements) > 0 {
		var nd lib.Damage = lib.GetDamageInfo(activeElements[0])
		nd.Modifier = getModifierForType(nd.Type, modSet)
		nd.Base = baseDamage
		nd.Value = baseDamage * (1 + baseModifier) * (nd.Modifier)
		damages = append(damages, nd)
	}

	var totalDamage float64
	for _, d := range damages {
		totalDamage += d.Value
	}

	var totalDps []float64
	var totalAvgHit []float64
	var totalDot []float64
	var totalTtk int
	for _, lvl := range enemyLevels {
		var enemies []lib.Enemy
		enemies = lib.SpawnEnemies(lvl)
		for _, e := range enemies {
			ttk, dps, avgHit, dot := e.Kill(damages, totalDamage, baseDamage, baseModifier, moddedStatusChance, moddedStatusDuration, avgDamageMulti, moddedAttackSpeed, getModifierForType(e.Faction, modSet))
			totalDps = append(totalDps, dps)
			totalAvgHit = append(totalAvgHit, avgHit)
			totalDot = append(totalDot, dot)
			totalTtk += ttk
		}
	}

	var avgDps float64
	for _, v := range totalDps {
		avgDps += v
	}
	avgDps = avgDps / float64(len(totalDps))

	var avgHit float64
	for _, v := range totalAvgHit {
		avgHit += v
	}
	avgHit = avgHit / float64(len(totalAvgHit))

	var avgDot float64
	var countDot int
	for _, v := range totalDot {
		if v > 0 {
			avgDot += v
			countDot++
		}
	}
	avgDot = avgDot / float64(countDot)

	return Rank{
		Weapon:       weapon.Name,
		DPS:          avgDps,
		TTK:          totalTtk,
		AvgHit:       avgHit,
		DoT:          avgDot,
		Set:          inModSet,
		Damages:      damages,
		StatusChance: moddedStatusChance,
		CritChance:   moddedCritChance,
		CritMulti:    moddedCritMulti,
		attackSpeed:  weapon.AttackSpeed * (1 + getModifierForType("attackSpeed", modSet)) * (1 + getModifierForType("attackSpeedMulti", modSet)),
	}
}

func ConvertToLengthBase(n int, arr []lib.Mod, len int, L int) (set []lib.Mod) {
	for i := 0; i < L; i++ {
		for _, m := range set {
			if m.Name == arr[n%len].Name {
				set = []lib.Mod{}
				return
			}
		}
		set = append(set, arr[n%len])
		n = int(n / len)
	}
	return
}

func hasDup(mods []lib.Mod) (result bool) {
	for i, m1 := range mods {
		for j, m2 := range mods {
			if i != j && m1.Name == m2.Name {
				return true
			}
		}
	}
	return false
}

func combinationUtil(arr []lib.Mod, data []lib.Mod, start int, end int, index int, r int, sets *[][]lib.Mod) (set []lib.Mod) {
	if index == r {
		for j := 0; j < r; j++ {
			set = append(set, data[j])
		}
		*sets = append(*sets, set)
		return
	}
	for i := start; i <= end && end-i+1 >= r-index; i++ {
		data[index] = arr[i]
		combinationUtil(arr, data, i+1, end, index+1, r, sets)
	}
	return
}

func getCombinations(arr []lib.Mod, r int) (out [][]lib.Mod) {
	var data []lib.Mod = make([]lib.Mod, len(arr))
	copy(data, arr)
	var n int = len(arr)
	combinationUtil(arr, data, 0, n-1, 0, r, &out)
	return
}

func MaxParallelism() int {
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	if maxProcs < numCPU {
		return maxProcs
	}
	return numCPU
}

type Rank struct {
	Weapon       string
	DPS          float64
	TTK          int
	AvgHit       float64
	DoT          float64
	Set          []lib.Mod
	Damages      []lib.Damage
	StatusChance float64
	CritChance   float64
	CritMulti    float64
	attackSpeed  float64
}

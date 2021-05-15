package lib

type Damage struct {
	Type     string
	Bonus    []Bonus
	Penalty  []Penalty
	Dot      Dot
	Mix      []string
	Base     float64
	Value    float64
	Modifier float64
}

type Bonus struct {
	Type  string
	Value float64
}

type Penalty struct {
	Type  string
	Value float64
}

type Dot struct {
	Damage float64
	Delay  int
}

func (d Damage) GetBonus(test string) (bonus float64) {
	bonus = 0
	for _, b := range d.Bonus {
		if b.Type == test {
			bonus += b.Value
		}
	}
	return
}

func (d Damage) GetPenalty(test string) (penalty float64) {
	penalty = 0
	for _, b := range d.Penalty {
		if b.Type == test {
			penalty -= b.Value
		}
	}
	return
}

func setupDamages() (damages []Damage) {
	damages = append(damages, Damage{
		Type: "slash",
		Bonus: []Bonus{
			{Type: "flesh", Value: 0.25},
			{Type: "cloned", Value: 0.25},
			{Type: "fossil", Value: 0.15},
			{Type: "infested", Value: 0.25},
			{Type: "infestedFlesh", Value: 0.25},
		},
		Penalty: []Penalty{
			{Type: "robotic", Value: 0.25},
			{Type: "ferrite", Value: 0.15},
			{Type: "alloy", Value: 0.5},
		},
		Dot: Dot{
			Damage: 0.35,
			Delay:  1,
		},
	})
	damages = append(damages, Damage{
		Type: "impact",
		Bonus: []Bonus{
			{Type: "machine", Value: 0.25},
			{Type: "shield", Value: 0.5},
			{Type: "proto", Value: 0.15},
		},
		Penalty: []Penalty{
			{Type: "flesh", Value: 0.25},
			{Type: "cloned", Value: 0.25},
		},
	})
	damages = append(damages, Damage{
		Type: "puncture",
		Bonus: []Bonus{
			{Type: "sinew", Value: 0.25},
			{Type: "robotic", Value: 0.25},
			{Type: "ferrite", Value: 0.5},
			{Type: "alloy", Value: 0.15},
		},
		Penalty: []Penalty{
			{Type: "shield", Value: 0.2},
			{Type: "proto", Value: 0.5},
		},
	})
	damages = append(damages, Damage{
		Type: "cold",
		Bonus: []Bonus{
			{Type: "sinew", Value: 0.25},
			{Type: "shield", Value: 0.5},
			{Type: "alloy", Value: 0.25},
		},
		Penalty: []Penalty{
			{Type: "fossil", Value: 0.25},
			{Type: "infestedFlesh", Value: 0.5},
		},
	})
	damages = append(damages, Damage{
		Type: "electricity",
		Bonus: []Bonus{
			{Type: "machine", Value: 0.5},
			{Type: "robotic", Value: 0.5},
		},
		Penalty: []Penalty{
			{Type: "alloy", Value: 0.5},
		},
		Dot: Dot{
			Damage: 0.5,
			Delay:  0,
		},
	})
	damages = append(damages, Damage{
		Type: "heat",
		Bonus: []Bonus{
			{Type: "cloned", Value: 0.25},
			{Type: "infested", Value: 0.25},
			{Type: "infestedFlesh", Value: 0.5},
		},
		Penalty: []Penalty{
			{Type: "proto", Value: 0.5},
		},
		Dot: Dot{
			Damage: 0.5,
			Delay:  1,
		},
	})
	damages = append(damages, Damage{
		Type: "toxin",
		Bonus: []Bonus{
			{Type: "flesh", Value: 0.5},
		},
		Penalty: []Penalty{
			{Type: "fossil", Value: 0.5},
			{Type: "machine", Value: 0.25},
			{Type: "robotic", Value: 0.25},
		},
		Dot: Dot{
			Damage: 0.5,
			Delay:  1,
		},
	})
	damages = append(damages, Damage{
		Type: "blast",
		Bonus: []Bonus{
			{Type: "fossil", Value: 0.5},
			{Type: "machine", Value: 0.75},
		},
		Penalty: []Penalty{
			{Type: "sinew", Value: 0.5},
			{Type: "ferrite", Value: 0.25},
		},
		Mix: []string{"heat", "cold"},
	})
	damages = append(damages, Damage{
		Type: "corrosive",
		Bonus: []Bonus{
			{Type: "fossile", Value: 0.75},
			{Type: "ferrite", Value: 0.75},
		},
		Penalty: []Penalty{
			{Type: "proto", Value: 0.5},
		},
		Mix: []string{"toxin", "electricity"},
	})
	damages = append(damages, Damage{
		Type: "gas",
		Bonus: []Bonus{
			{Type: "infested", Value: 0.75},
			{Type: "infestedFlesh", Value: 0.5},
		},
		Penalty: []Penalty{
			{Type: "flesh", Value: 0.25},
			{Type: "cloned", Value: 0.5},
		},
		Dot: Dot{
			Damage: 0.5,
			Delay:  0,
		},
		Mix: []string{"heat", "toxin"},
	})
	damages = append(damages, Damage{
		Type: "magnetic",
		Bonus: []Bonus{
			{Type: "shield", Value: 0.75},
			{Type: "proto", Value: 0.75},
		},
		Penalty: []Penalty{
			{Type: "alloy", Value: 0.5},
		},
		Mix: []string{"cold", "electricity"},
	})
	damages = append(damages, Damage{
		Type: "radiation",
		Bonus: []Bonus{
			{Type: "sinew", Value: 0.5},
			{Type: "robotic", Value: 0.25},
		},
		Penalty: []Penalty{
			{Type: "fossil", Value: 0.75},
			{Type: "infested", Value: 0.5},
			{Type: "shield", Value: 0.25},
		},
		Mix: []string{"heat", "electricity"},
	})
	damages = append(damages, Damage{
		Type: "viral",
		Bonus: []Bonus{
			{Type: "flesh", Value: 0.5},
			{Type: "cloned", Value: 0.75},
		},
		Penalty: []Penalty{
			{Type: "infested", Value: 0.5},
			{Type: "machine", Value: 0.25},
		},
		Mix: []string{"cold", "toxin"},
	})
	return
}

func GetDamageInfo(Type string) (damage Damage) {
	for _, d := range Damages {
		if d.Type == Type {
			damage = d
			break
		}
	}
	return
}

// var MixMap = map[string]string{
// 	"coldtoxin": "viral",
// }

var Damages []Damage = setupDamages()

var Elements = []string{"cold", "toxin", "heat", "electricity"}

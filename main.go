package main

import (
	"fmt"
	"slices"
)

type Item struct {
	s1 Skill
	s2 Skill
	s3 Skill
}

type Skill int

const (
	Warrior Skill = iota
	Magician
	Thief
	Bowman
	Pirate
	Attack
	Damage
	Crit
	Registry
	White
	Dimensional
	Other
)

var ItemMap = make(map[Item]int)
var ItemArr = make([]Item, 10)
var ItemCombs = make([]Combination, 10)
var CombsCount = make(map[*Combination]int)

type Combination struct {
	SkillCounts map[Skill]int // initialized to the type of combination
	CombArr     []Item
}

func main() {
}

func add(item Item) {
	ItemArr = append(ItemArr, item)
	ItemMap[item]++
	added := true

	if ItemMap[item] > 1 {
		updateOptions(item, added)
	}

}

func remove(item Item) {
	for i := range ItemArr {
		if ItemArr[i] == item {
			temp := ItemArr[:i]
			temp = append(temp, ItemArr[i+1:]...)
			ItemArr = temp
			ItemMap[item]--
			break
		}
	}

	if ItemMap[item] == 0 {
		updateOptions(item, false)
	}

}

func clear() {
	ItemArr = make([]Item, 10)
	ItemMap = make(map[Item]int)
}

func updateOptions(item Item, added bool) {
	if added {
		for _, comb := range ItemCombs {
			if comb.SkillCounts[item.s1] > 0 &&
				comb.SkillCounts[item.s2] > 0 &&
				comb.SkillCounts[item.s3] > 0 {
				comb.SkillCounts[item.s1]++
				comb.SkillCounts[item.s2]++
				comb.SkillCounts[item.s3]++
				comb.CombArr = append(comb.CombArr, item)
				CombsCount[&comb]++
			}
		}
	} else {
		for _, comb := range ItemCombs {
			if slices.Contains(comb.CombArr, item) {
				comb.SkillCounts[item.s1]--
				comb.SkillCounts[item.s2]--
				comb.SkillCounts[item.s3]--
				var i int
				for i = range comb.CombArr {
					if comb.CombArr[i] == item {
						break
					}
				}
				temp := comb.CombArr[:i]
				temp = append(temp, comb.CombArr[i+1:]...)
				CombsCount[&comb]--
			}
		}
	}
}

func createCombination(combType string) Combination {
	var combination Combination

	switch combType {

	case "eu":
		combination.SkillCounts[Warrior] = 1
		combination.SkillCounts[Registry] = 3
		combination.SkillCounts[Damage] = 3
		combination.SkillCounts[Crit] = 3
		combination.SkillCounts[Attack] = 3

	case "chinese":
		combination.SkillCounts[Warrior] = 1
		combination.SkillCounts[Thief] = 1
		combination.SkillCounts[Bowman] = 1
		combination.SkillCounts[Registry] = 3
		combination.SkillCounts[Dimensional] = 3
		combination.SkillCounts[White] = 3
		combination.SkillCounts[Damage] = 3
		combination.SkillCounts[Crit] = 3
		combination.SkillCounts[Attack] = 3

	default:
		combination.SkillCounts[Warrior] = 1
		combination.SkillCounts[Thief] = 1
		combination.SkillCounts[Bowman] = 1
		combination.SkillCounts[Magician] = 1
		combination.SkillCounts[Pirate] = 1
		combination.SkillCounts[Registry] = 1
		combination.SkillCounts[Damage] = 3
		combination.SkillCounts[Crit] = 3
		combination.SkillCounts[Attack] = 3
	}

	combination.CombArr = []Item{}
	return combination

}

func displayItems() {
	for i := range ItemArr {
		fmt.Println(ItemArr[i])
	}
}

func displayCombinations() {

}

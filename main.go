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

func add(item Item, combType string) {
	ItemArr = append(ItemArr, item)
	ItemMap[item]++

	if ItemMap[item] > 1 {
		addOption(item, combType)
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
		removeOption(item)
	}

}

func clear() {
	ItemArr = make([]Item, 10)
	ItemMap = make(map[Item]int)
}

func addOption(item Item, combType string) {
	for _, comb := range ItemCombs {
		if comb.SkillCounts[item.s1] > 0 &&
			comb.SkillCounts[item.s2] > 0 &&
			comb.SkillCounts[item.s3] > 0 {
			var temp Combination
			for i := range comb.CombArr {
				temp.CombArr = append(temp.CombArr, comb.CombArr[i])
			}
			for i := range comb.SkillCounts {
				temp.SkillCounts[i] = comb.SkillCounts[i]
			}
			temp.SkillCounts[item.s1]++
			temp.SkillCounts[item.s2]++
			temp.SkillCounts[item.s3]++
			temp.CombArr = append(temp.CombArr, item)
			ItemCombs = append(ItemCombs, temp)
			CombsCount[&temp] = CombsCount[&comb] + 1 
		}
	}

	newComb := createCombination(combType)
	newComb.SkillCounts[item.s1]--
	newComb.SkillCounts[item.s2]--
	newComb.SkillCounts[item.s3]--

	ItemCombs = append(ItemCombs, newComb)
}

func removeOption(item Item) {
	removed := []int{}
	for i, comb := range ItemCombs {
		if slices.Contains(comb.CombArr, item) {
			removed = append(removed, i)	
		}
	}

	newCombs := []Combination{}
	for i := range ItemCombs {
		for _, idx := range removed {
			if i == idx {
				continue
			}
			newCombs = append(newCombs, ItemCombs[i])
		}
	}
	ItemCombs = newCombs
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
	slices.SortFunc(ItemCombs, func(a int, b int) int{
		 if CombCounts[ItemCombs[a]] > CombCounts[ItemCombs[b]] {
			 return 1
		 }
		 return 0
	})
}

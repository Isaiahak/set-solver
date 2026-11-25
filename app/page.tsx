'use client'

import Image from "next/image";
import { useState } from "react";

const enum Skill {
	Warrior = 1,
	Magician = 2,
	Thief = 3,
	Bowman = 4, 
	Pirate = 5, 
	Attack = 6, 
	Damage = 7, 
	Crit = 8, 
	Registry = 9, 
	White = 10, 
	Dimensional = 11,
	Other = 12,
}

const enum EquipmentSlot {
	Hat = 1,
	Top = 2,
	Bottom = 3,
	Cape = 4,
	Gloves = 5,
	Boots = 6,
}


interface Item {
	type: EquipmentSlot | null
	s1: Skill | null 
	s2: Skill | null
	S3: Skill | null
}

interface Combination {
	SkillCounts: Map<Skill,int>
	EquipmentCounts: Map<Equipment,int>
	CombArr: Item[]
}

export default function Home() {
	
const [itemMap, setItemMap] = useState<Map<Item,int>>([])
const [itemList, setItemList] = useState<Item[]>([])
const [combList, setCombList] = useState<Combination[]>([])
const [combsCount, setCombsCount] = useState<Map<Combination,int> | nil >([])

function add(item: Item, combType: string) {

}

function remove(item: Item) {
	
}

function clear() {

}

function addOption(item: Item, combType: string) {
	for (const comb of ItemCombs) {
		if (comb.SkillCounts[item.s1] > 0 &&
			comb.SkillCounts[item.s2] > 0 &&
			comb.SkillCounts[item.s3] > 0) {
			var temp : Combination
			for (const i in comb.CombArr) {
				temp.CombArr = append(temp.CombArr, comb.CombArr[i])
			}
			for (const key of comb.SkillCounts.keys()) {
				temp.SkillCounts[key] = comb.SkillCounts[key]
			}

			temp.SkillCounts[item.s1]++
			temp.SkillCounts[item.s2]++
			temp.SkillCounts[item.s3]++
			temp.CombArr = append(temp.CombArr, item)
			ItemCombs = append(ItemCombs, temp)
			CombsCount[temp] = CombsCount[&comb] + 1 
		}
	}
}

function removeOption(item: Item) {
	var remove : int[]

	for (i in ItemCombs) {
		if (ItemCombs[i].contains(item)) {
			remove.append(i)
		}
	}

	var newComb : Combination[]
	for (i in ItemCombs) {
		if (removed.length == 0) {
			break 
		}
		if (i == removed[0]) {
			removed = removed.slice(1)
			continue
		}

		newComb.append(ItemCombs[i])
	}

	ItemCombs = newComb
}

function createCombination(combType: string) : Combination {
	var combination : Combination
	switch (combType) {
		case combType == "eu":

			combination.SkillCounts[Registry] = 3
			combination.SkillCounts[Dimensional] = 3
			combination.SkillCounts[White] = 3
			combination.SkillCounts[Damage] = 3
			combination.SkillCounts[Crit] = 3
			combination.SkillCounts[Attack] = 3

		case combType == "chinese":

			combination.SkillCounts[Registry] = 3
			combination.SkillCounts[Dimensional] = 3
			combination.SkillCounts[White] = 3
			combination.SkillCounts[Registry] = 3
			combination.SkillCounts[Crit] = 3
			combination.SkillCounts[Attack] = 3
			combination.SkillCounts[Warrior] = 1
			combination.SkillCounts[Thief] = 1
			combination.SkillCounts[Bowman] = 1

		default:

			combination.SkillCounts[Registry] = 3
			combination.SkillCounts[Dimensional] = 3
			combination.SkillCounts[White] = 3
			combination.SkillCounts[Registry] = 3
			combination.SkillCounts[Attack] = 3
			combination.SkillCounts[Crit] = 1
			combination.SkillCounts[Warrior] = 1
			combination.SkillCounts[Thief] = 1
			combination.SkillCounts[Pirate] = 1
			combination.SkillCounts[Magician] = 1
			combination.SkillCounts[Bowman] = 1

	}
	
	combbination.CombArr = []
	return combination
}

	  return (
		<div className="flex flex-col flex-1 items-center justify-center bg-zinc-50 font-sans dark:bg-black">
			<main className="flex flex-1 w-full max-w-3xl flex-col items-center justify-between py-32 px-16 bg-white dark:bg-black sm:items-start">
			</main>
		</div>
	  );
}

const EquipmentList = (e: Equipment) => {
	
	const [list, setList] = useState<item[]>([])

	return (
		<div>
		</div>
	)
}

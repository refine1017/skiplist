package skiplist

import (
	"fmt"
	"math/rand"
)

const PLAYER_LENGTH = 3000

var hero_list = []uint32{2011, 2031, 3011, 3021, 4011, 4021}

var data_map = make(map[uint64]*LadderHero)

type LadderHero struct {
	owner   uint32
	entry   uint32
	ability uint32
}

func heroRealLadderLess(d1 interface{}, d2 interface{}) bool {
	h1 := d1.(*LadderHero)
	h2 := d2.(*LadderHero)

	if h1.ability != h2.ability {
		return h1.ability > h2.ability
	}

	if h1.owner != h2.owner {
		return h1.owner > h2.owner
	}

	return h1.entry > h2.entry
}

func randomAdd(heroRealLadder *SkipList, owner, entry uint32) uint64 {

	ability := uint32(rand.Int31n(100000))

	ladderData := &LadderHero{
		owner:   owner,
		entry:   entry,
		ability: ability,
	}

	key := uint64(owner)<<32 | uint64(entry)

	if heroRealLadder.Set(key, ladderData) == false {
		fmt.Printf("【Delete Fail】 %10v - %10v - %10v - %10v\n", key, ladderData.owner, ladderData.entry, ladderData.ability)
		old_data := data_map[key]
		if old_data != nil {
			fmt.Printf("【OLD DATA】 %10v - %10v - %10v - %10v\n", key, old_data.owner, old_data.entry, old_data.ability)
		} else {
			fmt.Printf("【OLD DATA】 NIL\n")
		}
	} else {
		data_map[key] = ladderData
	}

	return key
}

func checkLadder(heroRealLadder *SkipList) error {
	var ladder_length = PLAYER_LENGTH * len(hero_list)

	if heroRealLadder.Length() != ladder_length {
		return checkLengthError(heroRealLadder)
	}

	var last_ability uint32 = 100000
	var last_index int = 0
	elements := heroRealLadder.Top(ladder_length)
	for index, ele := range elements {
		if ele == nil {
			return fmt.Errorf("checkLadder: ele is nil, index = %v", index)
		}

		hero_data := ele.(*LadderHero)
		if hero_data == nil {
			return fmt.Errorf("checkLadder: hero_data is nil, index = %v", index)
		}

		if hero_data.owner < 1 || hero_data.owner > PLAYER_LENGTH {
			return fmt.Errorf("checkLadder: owner is error, owner = %v", hero_data.owner)
		}

		if hero_data.entry == 0 {
			return fmt.Errorf("checkLadder: entry is zero, owner = %v", hero_data.owner)
		}

		if hero_data.ability > last_ability {
			return fmt.Errorf("checkLadder: index[%v][%v]>index[%v][%v]", index, hero_data.ability, last_index, last_ability)
		}
	}

	return nil
}

func checkLengthError(heroRealLadder *SkipList) error {
	var key_map = make(map[uint64]*LadderHero)

	elements := heroRealLadder.Top(heroRealLadder.Length())
	for index, ele := range elements {
		if ele == nil {
			return fmt.Errorf("checkLengthError: ele is nil, index = %v", index)
		}

		hero_data := ele.(*LadderHero)
		if hero_data == nil {
			return fmt.Errorf("checkLengthError: hero_data is nil, index = %v", index)
		}

		if hero_data.owner < 1 || hero_data.owner > PLAYER_LENGTH {
			return fmt.Errorf("checkLengthError: owner is error, owner = %v", hero_data.owner)
		}

		if hero_data.entry == 0 {
			return fmt.Errorf("checkLengthError: entry is zero, owner = %v", hero_data.owner)
		}

		key := uint64(hero_data.owner)<<32 | uint64(hero_data.entry)

		if key_map[key] != nil {
			fmt.Printf("【Duplicate】 %10v - %10v - %10v - %10v\n", index, hero_data.owner, hero_data.entry, hero_data.ability)
			fmt.Printf("%10v - %10v - %10v - %10v\n", key, key_map[key].owner, key_map[key].entry, key_map[key].ability)
		}
	}

	return fmt.Errorf("checkLengthError, ladder length = %v", heroRealLadder.Length())
}

func randomChange(heroRealLadder *SkipList) {
	for i := 0; i < 10000; i++ {
		owner := uint32(rand.Int31n(PLAYER_LENGTH)) + 1
		hero_index := uint32(rand.Int31n(int32(len(hero_list)) - 1))

		randomAdd(heroRealLadder, owner, hero_list[hero_index])
	}
}

func printTop(heroRealLadder *SkipList, num int) error {
	fmt.Printf("printTop ###########[%v]##########\n", num)

	elements := heroRealLadder.Top(num)
	for index, ele := range elements {
		if ele == nil {
			return fmt.Errorf("printTop: ele is nil, index = %v", index)
		}

		hero_data := ele.(*LadderHero)
		if hero_data == nil {
			return fmt.Errorf("printTop: hero_data is nil, index = %v", index)
		}

		fmt.Printf("%10v - %10v - %10v\n", hero_data.owner, hero_data.entry, hero_data.ability)
	}

	return nil
}

func main() {
	var heroRealLadder = NewSkipList()
	heroRealLadder.Less = heroRealLadderLess

	for i := 1; i <= PLAYER_LENGTH; i++ {
		for _, hero := range hero_list {
			randomAdd(heroRealLadder, uint32(i), hero)
		}

	}

	if e := checkLadder(heroRealLadder); e != nil {
		panic(e)
	}

	randomChange(heroRealLadder)

	if e := checkLadder(heroRealLadder); e != nil {
		panic(e)
	}
}

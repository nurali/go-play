package maparray

import (
	"fmt"
	"testing"
)

func testMap(t *testing.T) {
	slotMap := NewSlotMap()

	// for slotNo := 1; slotNo <= 100; slotNo++ {
	// 	slotMap.Add(&Slot{number: slotNo, free: true})
	// }
	// for slotNo := 1; slotNo <= 100; slotNo++ {
	// 	if slotNo%2 == 0 {
	// 		slotMap.Remove(slotNo)
	// 	}
	// }

	slotNo := 1
	slotMap.Add(&Slot{number: slotNo, free: true})
	slot := slotMap.Get(slotNo)
	fmt.Println(slot)
	slotMap.Remove(slotNo)
	slot = slotMap.Get(slotNo)
	fmt.Println(slot)
}

func testArr(t *testing.T) {
	slotArr := NewSlotArr()

	slotNo := 1
	slotArr.Add(&Slot{number: slotNo, free: true})
	slot := slotArr.Get(slotNo)
	fmt.Println(slot)
	slotArr.Remove(slotNo)
	slot = slotArr.Get(slotNo)
	fmt.Println(slot)
}

var SLOTS = 1000

func BenchmarkMapAdd(b *testing.B) {
	slots := NewSlotMap()
	AddMany(slots)
	slotNo := SLOTS + 1
	for i := 0; i < b.N; i++ {
		slots.Add(&Slot{number: slotNo, free: true})
	}
}

func BenchmarkArrAdd(b *testing.B) {
	slots := NewSlotArr()
	AddMany(slots)
	slotNo := SLOTS + 1
	for i := 0; i < b.N; i++ {
		slots.Add(&Slot{number: slotNo, free: true})
	}
}

func BenchmarkMapGet(b *testing.B) {
	slots := NewSlotMap()
	AddMany(slots)
	slotNo := SLOTS + 1
	slots.Add(&Slot{number: slotNo, free: true})
	for i := 0; i < b.N; i++ {
		slots.Get(slotNo)
	}
}

func BenchmarkArrGet(b *testing.B) {
	slots := NewSlotArr()
	AddMany(slots)
	slotNo := SLOTS + 1
	slots.Add(&Slot{number: slotNo, free: true})
	for i := 0; i < b.N; i++ {
		slots.Get(slotNo)
	}
}

func BenchmarkMapRemove(b *testing.B) {
	slots := NewSlotMap()
	AddMany(slots)
	slotNo := SLOTS + 1
	slots.Add(&Slot{number: slotNo, free: true})
	for i := 0; i < b.N; i++ {
		slots.Remove(slotNo)
	}
}

func BenchmarkArrRemove(b *testing.B) {
	slots := NewSlotArr()
	AddMany(slots)
	slotNo := SLOTS + 1
	slots.Add(&Slot{number: slotNo, free: true})
	for i := 0; i < b.N; i++ {
		slots.Remove(slotNo)
	}
}

func BenchmarkMapNextSlot1(b *testing.B) {
	slots := NewSlotMap()
	AddMany(slots)
	RemoveLast(slots)
	for i := 0; i < b.N; i++ {
		slots.NextFreeSlot()
	}
}

func BenchmarkArrNextSlot1(b *testing.B) {
	slots := NewSlotArr()
	AddMany(slots)
	RemoveLast(slots)
	for i := 0; i < b.N; i++ {
		slots.NextFreeSlot()
	}
}

func BenchmarkMapNextSlot2(b *testing.B) {
	slots := NewSlotMap()
	AddMany(slots)
	RemoveMany(slots)
	for i := 0; i < b.N; i++ {
		slots.NextFreeSlot()
	}
}

func BenchmarkArrNextSlot2(b *testing.B) {
	slots := NewSlotArr()
	AddMany(slots)
	RemoveMany(slots)
	for i := 0; i < b.N; i++ {
		slots.NextFreeSlot()
	}
}

func AddMany(slots SlotData) {
	for i := 1; i <= SLOTS; i++ {
		slots.Add(&Slot{number: i, free: true})
	}
}

func RemoveMany(slots SlotData) {
	for i := 1; i <= SLOTS; i++ {
		if i%2 == 0 {
			slots.Remove(i)
		}
	}
}

func RemoveLast(slots SlotData) {
	slots.Remove(SLOTS)
}

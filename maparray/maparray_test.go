package maparray

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	slots := NewSlotMap()
	for i := 1; i <= 10; i++ {
		slots.Add(&Slot{number: i, free: true})
	}

	// slot no 11
	slotNo := 11
	slots.Add(&Slot{number: slotNo, free: true})
	slot := slots.Get(slotNo)
	fmt.Println(slot)
	slots.Remove(slotNo)
	slot = slots.Get(slotNo)
	fmt.Println(slot)
	slotNo = slots.NextFreeSlot()
	fmt.Println(slotNo)

	// slot no 5
	slotNo = 5
	slots.Remove(slotNo)
	slot = slots.Get(slotNo)
	fmt.Println(slot)
	slotNo = slots.NextFreeSlot()
	fmt.Println(slotNo)
}

func TestArr(t *testing.T) {
	slots := NewSlotArr()
	for i := 1; i <= 10; i++ {
		slots.Add(&Slot{number: i, free: true})
	}

	// slot no 11
	slotNo := 11
	slots.Add(&Slot{number: slotNo, free: true})
	slot := slots.Get(slotNo)
	fmt.Println(slot)
	slots.Remove(slotNo)
	slot = slots.Get(slotNo)
	fmt.Println(slot)
	slotNo = slots.NextFreeSlot()
	fmt.Println(slotNo)

	// slot no 5
	slotNo = 5
	slots.Remove(slotNo)
	slot = slots.Get(slotNo)
	fmt.Println(slot)
	slotNo = slots.NextFreeSlot()
	fmt.Println(slotNo)
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

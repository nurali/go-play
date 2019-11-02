package maparray

type SlotData interface {
	Add(slot *Slot)
	Get(slotNo int) *Slot
	Remove(slotNo int)
	NextFreeSlot() int
}

type Slot struct {
	number int
	free   bool
}

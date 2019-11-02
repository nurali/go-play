package maparray

type SlotArr struct {
	slots []*Slot
}

func NewSlotArr() *SlotArr {
	return &SlotArr{
		slots: make([]*Slot, 0),
	}
}

func (a *SlotArr) Add(slot *Slot) {
	slot.free = false
	a.slots = append(a.slots, slot)
}

func (a *SlotArr) Get(slotNo int) *Slot {
	for _, slot := range a.slots {
		if slotNo == slot.number {
			return slot
		}
	}
	return nil
}

func (a *SlotArr) Remove(slotNo int) {
	for _, slot := range a.slots {
		if slotNo == slot.number {
			slot.free = true
		}
	}
}

func (a *SlotArr) NextFreeSlot() int {
	slotNo := 10000
	for _, slot := range a.slots {
		if slot.free == true {
			return slot.number
		}
	}
	return slotNo
}

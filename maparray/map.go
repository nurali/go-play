package maparray

type SlotMap struct {
	slots map[int]*Slot
}

func NewSlotMap() *SlotMap {
	return &SlotMap{
		slots: make(map[int]*Slot),
	}
}

func (m *SlotMap) Add(slot *Slot) {
	slot.free = false
	m.slots[slot.number] = slot
}

func (m *SlotMap) Get(slotNo int) *Slot {
	return m.slots[slotNo]
}

func (m *SlotMap) Remove(slotNo int) {
	m.slots[slotNo].free = true
	// delete(m.slots, slotNo)
}

func (m *SlotMap) NextFreeSlot() int {
	slotNo := 10000
	for key, val := range m.slots {
		if val.free == true && key < slotNo {
			slotNo = key
		}
	}
	return slotNo
}

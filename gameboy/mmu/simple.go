package mmu

import (
	"strings"
)

type MemoryManagementUnitSimple struct {
	memory_arr [0xFFFF]byte
	timer      *Timer
}

func NewMemoryManagementUnitSimple() *MemoryManagementUnitSimple {
	timer := TimerNew()
	return &MemoryManagementUnitSimple{
		timer: timer,
	}
}

func (m *MemoryManagementUnitSimple) Dump() string {
	var str strings.Builder
	str.WriteString("\n")
	// for i := 0; i < len(m.memory_arr); i += 2 {
	// 	if i%16 == 0 && i != 0 {
	// 		str.WriteString("\n")
	// 	}
	// 	s := fmt.Sprintf("%02x%02x ", m.memory_arr[i], m.memory_arr[i+1])
	// 	str.WriteString(s)
	//
	// }
	str.WriteString("\n")
	return strings.ToUpper(str.String())
}

func (m *MemoryManagementUnitSimple) Init(rom []byte) {
	m.timer.Init()

	for i, v := range BOOTROM {
		m.memory_arr[i] = v
	}

	for i, v := range rom {
		m.memory_arr[ROM_START+i] = v
	}
}

func (m *MemoryManagementUnitSimple) RB(address uint16) byte {
	return m.memory_arr[address]
}

func (m *MemoryManagementUnitSimple) WB(address uint16, value byte) {
	m.memory_arr[address] = value
}

func (m *MemoryManagementUnitSimple) RW(address uint16) uint16 {
	var b1 = m.memory_arr[address]
	var b2 = m.memory_arr[address+1]

	return uint16(b1)<<8 | uint16(b2)
}

func (m *MemoryManagementUnitSimple) DoCycle(ticks uint32) {
	m.timer.DoCycle(ticks)
}

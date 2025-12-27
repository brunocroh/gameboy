package mmu

import (
	"fmt"
	"strings"
)

type MemoryManagementUnitSimple struct {
	memory_arr [0xFFFFF]byte
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
	region := ROM_START - 16
	for i := 0; i < 32; i += 2 {
		if i%16 == 0 && i != 0 {
			str.WriteString("\n")
		}
		s := fmt.Sprintf("%02x%02x ", m.memory_arr[region+i], m.memory_arr[region+i+1])
		str.WriteString(s)

	}
	str.WriteString("\n")
	return strings.ToUpper(str.String())
}

func (m *MemoryManagementUnitSimple) Init(rom []byte) {
	m.timer.Init()

	for i, v := range BOOTROM {
		m.memory_arr[HRAM_START+i] = v
	}

	copy(m.memory_arr[:], rom)
}

func (m *MemoryManagementUnitSimple) RB(address uint16) byte {
	return m.memory_arr[address]
}

func (m *MemoryManagementUnitSimple) WB(address uint16, value byte) {
	if address == 0xFF02 || address == 0xFF01 {
		fmt.Print("WRITE AT SERIAL PORT")
	}
	m.memory_arr[address] = value
}

func (m *MemoryManagementUnitSimple) RW(address uint16) uint16 {
	var b1 = m.memory_arr[address]
	var b2 = m.memory_arr[address+1]

	return uint16(b1) | uint16(b2)<<8
}

func (m *MemoryManagementUnitSimple) DoCycle(ticks uint32) {
	m.timer.DoCycle(ticks)
}

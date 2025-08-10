package cpu

import (
	"fmt"

	"github.com/brunocroh/gameboy/gameboy/mmu"
)

type CPU struct {
	mmu *mmu.MemoryManagementUnit

	pc uint16
	sp uint16

	a uint8
	b uint8
	c uint8
	d uint8
	e uint8
	h uint8
	l uint8

	f uint8
}

func New(mmu *mmu.MemoryManagementUnit) *CPU {

	return &CPU{
		mmu: mmu,
	}
}

func (m *CPU) Init(mmu *mmu.MemoryManagementUnit) {
	m.pc = 0x0100
	m.sp = 0xFFFE
	m.a = 0x01
	m.f = 0xB0 //check later the initial value and how to handle flags
	m.b = 0x00
	m.c = 0x13
	m.d = 0x00
	m.e = 0xD8
	m.h = 0x01
	m.l = 0x4D
}

func Cycle() {

	fetch()
	decode()
	instruction()
}

func fetch() {
	fmt.Println("fetch")

}

func decode() {
	fmt.Println("fetch")
}

func instruction() {
	fmt.Println("instruction")
}

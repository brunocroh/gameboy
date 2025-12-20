package mmu

import (
	"github.com/brunocroh/gameboy/gameboy/mbc"
)

type SimpleMemoryManagementUnit struct {
	hram [0x100]byte
	wram [0x8000]byte
	vram [0x4000]byte
	mbc  *mbc.MemoryBankController

	timer *Timer
}

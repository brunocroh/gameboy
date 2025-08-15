package cpu

type instructions struct {
}

func instructionsNew() *instructions {
	return &instructions{}
}

func (m *instructions) jpAddr(cpu *CPU, addr uint16) uint32 {
	cpu.pc = uint16(addr)
	return 4
}

package cpu

type instructions struct {
}

func instructionsNew() *instructions {
	return &instructions{}
}

// ---- 8-Bit Load Instructions ----

/*
0x41 - LD r, r’: Load register (register)

Load to the 8-bit register r, data from the 8-bit register r'.

Machine Cycles: 1
*/
func (m *instructions) ldBC(cpu *CPU) uint32 {
	cpu.register.b = cpu.register.c
	return 0
}

/*
0x06 - LD r, n: Load register (immediate)

Load to the 8-bit register r, the immediate data n.

Machine Cycles: 2
*/
func (m *instructions) ldBAddr(cpu *CPU, addr uint16) uint32 {
	cpu.register.b = cpu.rb(addr)
	return 1
}

// ---- 16-Bit Load Instructions ----

// ---- 8-Bit Arithmetic and logical ----

// ---- 16-Bit Arithmetic and logical ----

// ---- Rotate, shift and bit ----

// ---- FLOW ----

/*
0xC3 - JP nn: Jump

Unconditional jump to the absolute address specified by the 16-bit immediate operand nn.

Machine Cycles: 4
*/
func (m *instructions) jpAddr(cpu *CPU, addr uint16) uint32 {
	cpu.pc = cpu.rw(addr)
	return 3
}

// ---- MISC ----

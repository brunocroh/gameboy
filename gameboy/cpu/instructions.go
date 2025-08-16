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
func (m *instructions) ld_rr(cpu *CPU) uint32 {
	cpu.register.b = cpu.register.c
	return 0
}

/*
0x06 - LD r, n: Load register (immediate)

Load to the 8-bit register r, the immediate data n.

Machine Cycles: 2
*/
func (m *instructions) ld_r_n(cpu *CPU) uint32 {
	cpu.register.b = cpu.mmu.RB(cpu.popPC())
	return 1
}

/*
0x46 - LD r, (HL): Load register (indirect HL)

Load to the 8-bit register r, data from the absolute address specified by the 16-bit register HL.

Machine Cycles: 2
*/
func (m *instructions) ld_r_HL(cpu *CPU) uint32 {
	cpu.register.b = cpu.mmu.RB(cpu.register.HL())
	return 1
}

/*
0x70 - LD (HL), r: Load from register (indirect HL)

Load to the absolute address specified by the 16-bit register HL, data from the 8-bit register r.

Machine Cycles: 2
*/
func (m *instructions) ld_HL_r(cpu *CPU) uint32 {
	cpu.mmu.WB(cpu.register.HL(), cpu.register.b)
	return 1
}

/*
0x36 - LD (HL), n: Load from immediate data (indirect HL)

Load to the absolute address specified by the 16-bit register HL, the immediate data n.

Machine Cycles: 3
*/
func (m *instructions) ld_HL_n(cpu *CPU) uint32 {
	n := cpu.mmu.RB(cpu.popPC())
	cpu.mmu.WB(cpu.register.HL(), n)
	return 2
}

/*
0x0A - LD A, (BC): Load accumulator (indirect BC)

Load to the 8-bit A register, data from the absolute address specified by the 16-bit register BC.

Machine Cycles: 2
*/
func (m *instructions) ld_A_BC(cpu *CPU) uint32 {
	cpu.register.a = cpu.mmu.RB(cpu.register.BC())
	return 1
}

/*
0x1A - LD A, (DE): Load accumulator (indirect DE)

Load to the 8-bit A register, data from the absolute address specified by the 16-bit register DE.

Machine Cycles: 2
*/
func (m *instructions) ld_A_DE(cpu *CPU) uint32 {
	cpu.register.a = cpu.mmu.RB(cpu.register.DE())
	return 1
}

/*
0x02 - A: Load from accumulator (indirect BC)

Load to the absolute address specified by the 16-bit register BC, data from the 8-bit A register.

Machine Cycles: 2
*/
func (m *instructions) ld_BC_A(cpu *CPU) uint32 {
	cpu.mmu.WB(cpu.register.BC(), cpu.register.a)
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
func (m *instructions) jp_nn(cpu *CPU) uint32 {
	cpu.PC = cpu.rw(cpu.PC)
	return 3
}

// ---- MISC ----

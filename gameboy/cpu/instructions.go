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
	return 1
}

/*
0x06 - LD r, n: Load register (immediate)

Load to the 8-bit register r, the immediate data n.

Machine Cycles: 2
*/
func (m *instructions) ld_r_n(cpu *CPU) uint32 {
	cpu.register.b = cpu.mmu.RB(cpu.popPC())
	return 2
}

/*
0x46 - LD r, (HL): Load register (indirect HL)

Load to the 8-bit register r, data from the absolute address specified by the 16-bit register HL.

Machine Cycles: 2
*/
func (m *instructions) ld_r_HL(cpu *CPU) uint32 {
	hl := uint16(cpu.register.h)<<8 | uint16(cpu.register.l)
	cpu.register.b = cpu.mmu.RB(hl)
	return 2
}

/*
0x70 - LD (HL), r: Load from register (indirect HL)

Load to the absolute address specified by the 16-bit register HL, data from the 8-bit register r.

Machine Cycles: 2
*/
func (m *instructions) ld_HL_r(cpu *CPU) uint32 {
	hl := uint16(cpu.register.h)<<8 | uint16(cpu.register.l)
	cpu.mmu.WB(hl, cpu.register.b)
	return 2
}

/*
0x36 - LD (HL), n: Load from immediate data (indirect HL)

Load to the absolute address specified by the 16-bit register HL, the immediate data n.

Machine Cycles: 3
*/
func (m *instructions) ld_HL_n(cpu *CPU) uint32 {
	hl := uint16(cpu.register.h)<<8 | uint16(cpu.register.l)
	n := cpu.mmu.RB(cpu.popPC())
	cpu.mmu.WB(hl, n)
	return 3
}

/*
0x0A - LD A, (BC): Load accumulator (indirect BC)

Load to the 8-bit A register, data from the absolute address specified by the 16-bit register BC.

Machine Cycles: 2
*/
func (m *instructions) ld_A_BC(cpu *CPU) uint32 {
	bc := uint16(cpu.register.b)<<8 | uint16(cpu.register.c)
	cpu.register.a = cpu.mmu.RB(bc)
	return 2
}

/*
0x1A - LD A, (DE): Load accumulator (indirect DE)

Load to the 8-bit A register, data from the absolute address specified by the 16-bit register DE.

Machine Cycles: 2
*/
func (m *instructions) ld_A_DE(cpu *CPU) uint32 {
	de := uint16(cpu.register.d)<<8 | uint16(cpu.register.e)
	cpu.register.a = cpu.mmu.RB(de)
	return 2
}

/*
0x02 - A: Load from accumulator (indirect BC)

Load to the absolute address specified by the 16-bit register BC, data from the 8-bit A register.

Machine Cycles: 2
*/
func (m *instructions) ld_BC_A(cpu *CPU) uint32 {
	bc := uint16(cpu.register.b)<<8 | uint16(cpu.register.c)
	cpu.mmu.WB(bc, cpu.register.a)
	return 2
}

/*
0x12 - LD (DE), A: Load from accumulator (indirect DE)

Load to the absolute address specified by the 16-bit register DE, data from the 8-bit A register.

Machine Cycles: 2
*/
func (m *instructions) ld_DE_A(cpu *CPU) uint32 {
	de := uint16(cpu.register.d)<<8 | uint16(cpu.register.e)
	cpu.mmu.WB(de, cpu.register.a)
	return 2
}

/*
0xFA - LD A, (nn): Load accumulator (direct)

Load to the 8-bit A register, data from the absolute address specified by the 16-bit operand nn.

Machine Cycles: 4
*/
func (m *instructions) ld_A_nn(cpu *CPU) uint32 {
	addr := uint16(cpu.mmu.RB(cpu.popPC()))<<8 | uint16(cpu.mmu.RB(cpu.popPC()))
	cpu.register.a = cpu.mmu.RB(addr)
	return 4
}

/*
0xEA - LD (nn), A: Load from accumulator (direct)

Load to the absolute address specified by the 16-bit operand nn, data from the 8-bit A register.

Machine Cycles: 4
*/
func (m *instructions) ld_nn_A(cpu *CPU) uint32 {
	addr := uint16(cpu.mmu.RB(cpu.popPC()))<<8 | uint16(cpu.mmu.RB(cpu.popPC()))
	cpu.mmu.WB(addr, cpu.register.a)
	return 4
}

/*
0xF2 - LDH A, (C): Load accumulator (indirect 0xFF00+C)

Load to the 8-bit A register, data from the address specified by the 8-bit C register. The full
16-bit absolute address is obtained by setting the most significant byte to 0xFF and the least
significant byte to the value of C, so the possible range is 0xFF00-0xFFFF.

Machine Cycles: 2
*/
func (m *instructions) ldh_A_C(cpu *CPU) uint32 {
	cpu.register.a = cpu.mmu.RB(0xFF00 | uint16(cpu.register.c))
	return 2
}

/*
0xE2 - LDH (C), A: Load from accumulator (indirect 0xFF00+C)

Load to the address specified by the 8-bit C register, data from the 8-bit A register. The full
16-bit absolute address is obtained by setting the most significant byte to 0xFF and the least
significant byte to the value of C, so the possible range is 0xFF00-0xFFFF.

Machine Cycles: 2
*/
func (m *instructions) ldh_C_A(cpu *CPU) uint32 {
	cpu.mmu.WB(0xFF00|uint16(cpu.register.c), cpu.register.a)
	return 2
}

/*
0xF0 - LDH A, (n): Load accumulator (direct 0xFF00+n)

Load to the 8-bit A register, data from the address specified by the 8-bit immediate data n. The
full 16-bit absolute address is obtained by setting the most significant byte to 0xFF and the
least significant byte to the value of n, so the possible range is 0xFF00-0xFFFF.

Machine Cycles: 3
*/
func (m *instructions) ldh_A_n(cpu *CPU) uint32 {
	n := cpu.mmu.RB(cpu.popPC())
	cpu.register.a = cpu.mmu.RB(0xFF00 | uint16(n))
	return 3
}

/*
0xE0 - LDH (n), A: Load from accumulator (direct 0xFF00+n)

Load to the address specified by the 8-bit immediate data n, data from the 8-bit A register. The
full 16-bit absolute address is obtained by setting the most significant byte to 0xFF and the
least significant byte to the value of n, so the possible range is 0xFF00-0xFFFF.

Machine Cycles: 3
*/
func (m *instructions) ldh_n_A(cpu *CPU) uint32 {
	n := cpu.mmu.RB(cpu.popPC())
	cpu.mmu.WB(0xFF00|uint16(n), cpu.register.a)
	return 3
}

/*
0x3A - LD A, (HL-): Load accumulator (indirect HL, decrement)

Load to the 8-bit A register, data from the absolute address specified by the 16-bit register HL.
The value of HL is decremented after the memory read.

Machine Cycles: 2
*/
func (m *instructions) ld_A_HLd(cpu *CPU) uint32 {
	hl := uint16(cpu.register.h)<<8 | uint16(cpu.register.l)
	value := hl - 1
	cpu.register.h = uint8(value >> 8)
	cpu.register.l = uint8(value & 0x00FF)
	cpu.register.a = cpu.mmu.RB(hl)
	return 2
}

/*
0x32 - LD (HL-), A: Load from accumulator (indirect HL, decrement)

Load to the absolute address specified by the 16-bit register HL, data from the 8-bit A register.
The value of HL is decremented after the memory write.

Machine Cycles: 2
*/
func (m *instructions) ld_HLd_A(cpu *CPU) uint32 {
	hl := uint16(cpu.register.h)<<8 | uint16(cpu.register.l)
	value := hl - 1
	cpu.register.h = uint8(value >> 8)
	cpu.register.l = uint8(value & 0x00FF)
	cpu.mmu.WB(hl, cpu.register.a)
	return 2
}

/*
0x2A - LD A, (HL+): Load accumulator (indirect HL, increment)

Load to the 8-bit A register, data from the absolute address specified by the 16-bit register HL.
The value of HL is incremented after the memory read

Machine Cycles: 2
*/
func (m *instructions) ld_A_HLi(cpu *CPU) uint32 {
	hl := uint16(cpu.register.h)<<8 | uint16(cpu.register.l)
	value := hl + 1
	cpu.register.h = uint8(value >> 8)
	cpu.register.l = uint8(value & 0x00FF)
	cpu.register.a = cpu.mmu.RB(hl)
	return 2
}

/*
0x22 - LD (HL+), A: Load from accumulator (indirect HL, increment)

Load to the absolute address specified by the 16-bit register HL, data from the 8-bit A register.
The value of HL is incremented after the memory write.

Machine Cycles: 2
*/
func (m *instructions) ld_HLi_A(cpu *CPU) uint32 {
	hl := uint16(cpu.register.h)<<8 | uint16(cpu.register.l)
	value := hl + 1
	cpu.register.h = uint8(value >> 8)
	cpu.register.l = uint8(value & 0x00FF)
	cpu.mmu.WB(hl, cpu.register.a)
	return 2
}

// ---- 16-Bit Load Instructions ----

/*
0x01 - LD rr, nn: Load 16-bit register / register pair

Load to the 16-bit register rr, the immediate 16-bit data nn.

Machine Cycles: 3
*/
func (m *instructions) ld_rr_nn(cpu *CPU) uint32 {
	value := cpu.rw(cpu.PC)

	cpu.register.b = uint8(value >> 8)
	cpu.register.c = uint8(value & 0x00FF)

	return 3
}

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
	return 4
}

// ---- MISC ----

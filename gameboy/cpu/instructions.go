package cpu

import "fmt"

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
	word := cpu.rw(cpu.PC)

	cpu.register.b = uint8(word >> 8)
	cpu.register.c = uint8(word & 0x00FF)
	return 3
}

/*
0x08 - LD (nn), SP: Load from stack pointer (direct)

Load to the absolute address specified by the 16-bit operand nn, data from the 16-bit SP register.

Machine Cycles: 5
*/
func (m *instructions) ld_nn_sp(cpu *CPU) uint32 {
	word := cpu.rw(cpu.PC)

	sp_msb := uint8(cpu.SP >> 8)
	sp_lsb := uint8(cpu.SP | 0x00FF)
	cpu.mmu.WB(word, sp_lsb)
	word += 1
	cpu.mmu.WB(word, sp_msb)
	return 5
}

/*
0xF9 - Load stack pointer from HL

Load to the 16-bit SP register, data from the 16-bit HL register.

Machine Cycles: 2
*/
func (m *instructions) ld_sp_HL(cpu *CPU) uint32 {
	cpu.SP = uint16(cpu.register.h)<<8 | uint16(cpu.register.l)
	return 2
}

/*
0xC5 - PUSH rr: Push to stack

Push to the stack memory, data from the 16-bit register rr.

Machine Cycles: 4
*/
func (m *instructions) ld_push_rr(cpu *CPU) uint32 {
	cpu.SP -= 1
	cpu.mmu.WB(cpu.SP, cpu.register.b)
	cpu.SP -= 1
	cpu.mmu.WB(cpu.SP, cpu.register.c)
	return 4
}

/*
0xC1 - POP rr: Pop from stack

Pops to the 16-bit register rr, data from the stack memory.
This instruction does not do calculations that affect flags, but POP AF completely replaces the
F register value, so all flags are changed based on the 8-bit data that is read from memory.

Machine Cycles: 3
*/
func (m *instructions) ld_pop_rr(cpu *CPU) uint32 {
	word := cpu.mmu.RW(cpu.SP)
	cpu.SP += 2

	cpu.register.b = uint8(word >> 8)
	cpu.register.c = uint8(word & 0x00FF)
	return 3
}

/*
0xF8 - LD HL, SP+e: Load HL from adjusted stack pointer

Load to the HL register, 16-bit data calculated by adding the signed 8-bit operand e to the 16-
bit value of the SP register.

Machine Cycles: 3
*/
func (m *instructions) ld_HL_spe(cpu *CPU) uint32 {
	e := cpu.mmu.RB(cpu.popPC())
	result := cpu.SP + uint16(e)

	tmp := cpu.SP ^ uint16(e) ^ result

	cpu.register.h = uint8(result >> 8)
	cpu.register.l = uint8(result & 0x00FF)

	cpu.register.setFlag("Z", false)
	cpu.register.setFlag("N", false)
	cpu.register.setFlag("H", (tmp&0x10) == 0x10)
	cpu.register.setFlag("C", (tmp&0x100) == 0x100)
	return 3
}

// ---- 8-Bit Arithmetic and logical ----

/*
0x80 - ADD r: Add (register)

adds to the 8-bit A register, the 8-bit register r, and stores the result back into the A register

Machine Cycles: 1
*/
func (m *instructions) add_r(cpu *CPU) uint32 {
	a := cpu.register.a
	b := cpu.register.b

	sum := a + b
	cpu.register.a = sum

	cpu.register.setFlag("Z", sum == 0x0)
	cpu.register.setFlag("N", false)
	cpu.register.setFlag("H", (a&0xF)+(b&0xF) > 0xF)
	cpu.register.setFlag("C", uint16(a)+uint16(b) > 0xFF)

	return 1
}

/*
0x86 - ADD (HL): Add (indirect HL)

Adds to the 8-bit A register, data from the absolute address specified by the 16-bit register HL,
and stores the result back into the A register.

Machine Cycles: 2
*/
func (m *instructions) add_HL(cpu *CPU) uint32 {
	hl := uint16(cpu.register.h)<<8 | uint16(cpu.register.l)
	a := cpu.register.a

	n := cpu.mmu.RB(hl)

	sum := a + n

	cpu.register.setFlag("Z", sum == 0x0)
	cpu.register.setFlag("N", false)
	cpu.register.setFlag("H", (a&0xF)+(n&0xF) > 0xF)
	cpu.register.setFlag("C", uint16(a)+uint16(n) > 0xFF)

	return 2
}

/*
0xC6 - ADD n: Add (immediate)

Adds to the 8-bit A register, the immediate data n, and stores the result back into the A register.

Machine Cycles: 2
*/
func (m *instructions) add_n(cpu *CPU) uint32 {
	n := cpu.mmu.RB(cpu.popPC())
	a := cpu.register.a

	sum := a + n

	cpu.register.a = sum

	cpu.register.setFlag("Z", sum == 0x0)
	cpu.register.setFlag("N", false)
	cpu.register.setFlag("H", (a&0xF)+(n&0xF) > 0xF)
	cpu.register.setFlag("C", uint16(a)+uint16(n) > 0xFF)
	return 2
}

/*
0x88 - ADC r: Add with carry (register)

Adds to the 8-bit A register, the carry flag and the 8-bit register r, and stores the result back
into the A register.

Machine Cycles: 1
*/
func (m *instructions) adc_r(cpu *CPU) uint32 {
	c := uint8(0)
	if cpu.register.getFlag("C") {
		c = 1
	}

	a := cpu.register.a
	b := cpu.register.b

	sum := a + b + c

	cpu.register.a = sum

	cpu.register.setFlag("Z", sum == 0x0)
	cpu.register.setFlag("N", false)
	cpu.register.setFlag("H", (a&0xF)+(b&0xF)+c > 0xF)
	cpu.register.setFlag("C", uint16(a)+uint16(b)+uint16(c) > 0xFF)

	return 1
}

/*
0x8E - ADC (HL): Add with carry (indirect HL)

Adds to the 8-bit A register, the carry flag and data from the absolute address specified by the
16-bit register HL, and stores the result back into the A register.

Machine Cycles: 2
*/
func (m *instructions) adc_HL(cpu *CPU) uint32 {
	hl := uint16(cpu.register.h)<<8 | uint16(cpu.register.l)
	c := uint8(0)
	if cpu.register.getFlag("C") {
		c = 1
	}

	n := cpu.mmu.RB(hl)
	a := cpu.register.a

	sum := a + n + c

	cpu.register.a = sum

	cpu.register.setFlag("Z", sum == 0x0)
	cpu.register.setFlag("N", false)
	cpu.register.setFlag("H", (a&0xF)+(n&0xF)+c > 0xF)
	cpu.register.setFlag("C", uint16(a)+uint16(n)+uint16(c) > 0xFF)

	return 2
}

/*
0xCE - ADC n: Add with carry (immediate)

Adds to the 8-bit A register, the carry flag and the immediate data n, and stores the result back
into the A register

Machine Cycles: 2
*/
func (m *instructions) adc_n(cpu *CPU) uint32 {
	c := uint8(0)
	if cpu.register.getFlag("C") {
		c = 1
	}

	n := cpu.mmu.RB(cpu.popPC())
	a := cpu.register.a

	sum := a + n + c

	cpu.register.a = sum

	cpu.register.setFlag("Z", sum == 0x0)
	cpu.register.setFlag("N", false)
	cpu.register.setFlag("H", (a&0xF)+(n&0xF)+c > 0xF)
	cpu.register.setFlag("C", uint16(a)+uint16(n)+uint16(c) > 0xFF)

	return 2
}

/*
0x90 - SUB r: Subtract (register)

Subtracts from the 8-bit A register, the 8-bit register r, and stores the result back into the A
register.

Machine Cycles: 1
*/
func (m *instructions) sub_r(cpu *CPU) uint32 {
	a := cpu.register.a
	b := cpu.register.b

	r := a - b

	cpu.register.a = r

	cpu.register.setFlag("Z", r == 0x0)
	cpu.register.setFlag("N", true)
	cpu.register.setFlag("H", (a&0x0F) < (b&0x0F))
	cpu.register.setFlag("C", uint16(a) < uint16(b))

	return 1
}

/*
0x96 - SUB (HL): Subtract (indirect HL)

Subtracts from the 8-bit A register, data from the absolute address specified by the 16-bit
register HL, and stores the result back into the A register.

Machine Cycles: 2
*/
func (m *instructions) sub_HL(cpu *CPU) uint32 {
	a := cpu.register.a
	h := cpu.register.h
	l := cpu.register.l
	hl := uint16(h)<<8 | uint16(l)

	n := cpu.mmu.RB(hl)

	r := a - n

	cpu.register.a = r

	cpu.register.setFlag("Z", r == 0x0)
	cpu.register.setFlag("N", true)
	cpu.register.setFlag("H", (a&0x0F) < (n&0x0F))
	cpu.register.setFlag("C", uint16(a) < uint16(n))

	return 2
}

/*
0xD6 - SUB n: Subtract (immediate)

Subtracts from the 8-bit A register, the immediate data n, and stores the result back into the A
register.

Machine Cycles: 2
*/
func (m *instructions) sub_n(cpu *CPU) uint32 {
	a := cpu.register.a
	n := cpu.mmu.RB(cpu.popPC())

	r := a - n

	cpu.register.a = r

	cpu.register.setFlag("Z", r == 0x0)
	cpu.register.setFlag("N", true)
	cpu.register.setFlag("H", (a&0x0F) < (n&0x0F))
	cpu.register.setFlag("C", uint16(a) < uint16(n))

	return 2
}

/*
0x98 - SBC r: Subtract with carry (register)

Subtracts from the 8-bit A register, the carry flag and the 8-bit register r, and stores the result
back into the A register.

Machine Cycles: 1
*/
func (m *instructions) sbc_r(cpu *CPU) uint32 {
	a := cpu.register.a
	b := cpu.register.b
	c := uint8(0)
	if cpu.register.getFlag("C") {
		c = 1
	}

	r := a - b - c

	cpu.register.a = r

	cpu.register.setFlag("Z", r == 0x0)
	cpu.register.setFlag("N", true)
	cpu.register.setFlag("H", (a&0x0F) < ((b&0x0F)+c))
	cpu.register.setFlag("C", uint16(a) < (uint16(b)+uint16(c)))

	return 1
}

/*
0x9E - SBC (HL): Subtract with carry (indirect HL)

Subtracts from the 8-bit A register, the carry flag and the 8-bit register r, and stores the result
back into the A register.

Machine Cycles: 2
*/
func (m *instructions) sbc_HL(cpu *CPU) uint32 {
	a := cpu.register.a
	h := cpu.register.h
	l := cpu.register.l
	hl := uint16(h)<<8 | uint16(l)

	c := uint8(0)
	if cpu.register.getFlag("C") {
		c = 1
	}

	n := cpu.mmu.RB(hl)

	r := a - n - c

	cpu.register.a = r

	cpu.register.setFlag("Z", r == 0x0)
	cpu.register.setFlag("N", true)
	cpu.register.setFlag("H", (a&0x0F) < ((n&0x0F)+c))
	cpu.register.setFlag("C", uint16(a) < (uint16(n)+uint16(c)))

	return 2
}

/*
0xDE - SBC n: Subtract with carry (immediate)

Subtracts from the 8-bit A register, the carry flag and the immediate data n, and stores the
result back into the A register.

Machine Cycles: 2
*/
func (m *instructions) sbc_n(cpu *CPU) uint32 {
	a := cpu.register.a

	c := uint8(0)
	if cpu.register.getFlag("C") {
		c = 1
	}

	n := cpu.mmu.RB(cpu.popPC())

	r := a - n - c

	cpu.register.a = r

	cpu.register.setFlag("Z", r == 0x0)
	cpu.register.setFlag("N", true)
	cpu.register.setFlag("H", (a&0x0F) < ((n&0x0F)+c))
	cpu.register.setFlag("C", uint16(a) < (uint16(n)+uint16(c)))

	return 2
}

/*
0xB8 - CP r: Compare (register)

Subtracts from the 8-bit A register, the 8-bit register r, and updates flags based on the result.
This instruction is basically identical to SUB r, but does not update the A register.

Machine Cycles: 1
*/
func (m *instructions) cp_r(cpu *CPU) uint32 {
	a := cpu.register.a
	b := cpu.register.b

	r := a - b

	cpu.register.setFlag("Z", r == 0x0)
	cpu.register.setFlag("N", true)
	cpu.register.setFlag("H", (a&0x0F) < (b&0x0F))
	cpu.register.setFlag("C", uint16(a) < uint16(b))

	return 1
}

/*
0xBE - CP (HL): Compare (indirect HL)

Subtracts from the 8-bit A register, data from the absolute address specified by the 16-bit
register HL, and updates flags based on the result. This instruction is basically identical to SUB
(HL), but does not update the A register

Machine Cycles: 2
*/
func (m *instructions) cp_HL(cpu *CPU) uint32 {
	a := cpu.register.a
	h := cpu.register.h
	l := cpu.register.l
	hl := uint16(h)<<8 | uint16(l)

	n := cpu.mmu.RB(hl)

	r := a - n

	cpu.register.setFlag("Z", r == 0x0)
	cpu.register.setFlag("N", true)
	cpu.register.setFlag("H", (a&0x0F) < (n&0x0F))
	cpu.register.setFlag("C", uint16(a) < uint16(n))

	return 2
}

/*
0xFE - CP n: Compare (immediate)

Subtracts from the 8-bit A register, the immediate data n, and updates flags based on the result.
This instruction is basically identical to SUB n, but does not update the A register.

Machine Cycles: 2
*/
func (m *instructions) cp_n(cpu *CPU) uint32 {
	a := cpu.register.a
	n := cpu.mmu.RB(cpu.popPC())

	r := a - n

	cpu.register.setFlag("Z", r == 0x0)
	cpu.register.setFlag("N", true)
	cpu.register.setFlag("H", (a&0x0F) < (n&0x0F))
	cpu.register.setFlag("C", uint16(a) < uint16(n))

	return 2
}

/*
0x04 - INC r: Increment (register)

# Increments data in the 8-bit register r

Machine Cycles: 1
*/
func (m *instructions) inc_r(cpu *CPU) uint32 {
	b := cpu.register.b

	r := b + 1
	cpu.register.b = r

	cpu.register.setFlag("Z", r == 0x0)
	cpu.register.setFlag("N", false)
	cpu.register.setFlag("H", (b&0x0F)+1 > 0x0F)

	return 1
}

/*
0x34 - INC (HL): Increment (indirect HL)

# Increments data at the absolute address specified by the 16-bit register HL

Machine Cycles: 3
*/
func (m *instructions) inc_HL(cpu *CPU) uint32 {
	h := cpu.register.h
	l := cpu.register.l
	hl := uint16(h)<<8 | uint16(l)

	data := cpu.mmu.RB(hl)

	r := data + 1

	cpu.mmu.WB(hl, r)

	cpu.register.setFlag("Z", r == 0x0)
	cpu.register.setFlag("N", false)
	cpu.register.setFlag("H", (data&0x0F)+1 > 0x0F)

	return 3
}

/*
0x05 - DEC r: Decrement (register)

# Decrements data in the 8-bit register r

Machine Cycles: 1
*/
func (m *instructions) dec_r(cpu *CPU) uint32 {
	b := cpu.register.b

	r := b - 1
	cpu.register.b = r

	cpu.register.setFlag("Z", r == 0x0)
	cpu.register.setFlag("N", false)
	cpu.register.setFlag("H", (b&0x0F)-1 > 0x0F)

	return 1
}

/*
0x35 - DEC (HL): Decrement (indirect HL)

Decrements data at the absolute address specified by the 16-bit register HL.

Machine Cycles: 3
*/
func (m *instructions) dec_HL(cpu *CPU) uint32 {
	h := cpu.register.h
	l := cpu.register.l
	hl := uint16(h)<<8 | uint16(l)

	data := cpu.mmu.RB(hl)

	r := data - 1

	cpu.mmu.WB(hl, r)

	cpu.register.setFlag("Z", r == 0x0)
	cpu.register.setFlag("N", false)
	cpu.register.setFlag("H", (data&0x0F)-1 > 0x0F)

	return 3
}

/*
0xA0 - AND r: Bitwise AND (register)

Performs a bitwise AND operation between the 8-bit A register and the 8-bit register r, and
stores the result back into the A register.

Machine Cycles: 1
*/
func (m *instructions) and_r(cpu *CPU) uint32 {
	a := cpu.register.a
	b := cpu.register.b

	r := a & b

	cpu.register.a = r

	cpu.register.setFlag("Z", r == 0x0)
	cpu.register.setFlag("N", false)
	cpu.register.setFlag("H", true)
	cpu.register.setFlag("C", false)

	return 1
}

/*
0xA6 - AND (HL): Bitwise AND (indirect HL)

Performs a bitwise AND operation between the 8-bit A register and data from the absolute
address specified by the 16-bit register HL, and stores the result back into the A register.

Machine Cycles: 2
*/
func (m *instructions) and_HL(cpu *CPU) uint32 {
	a := cpu.register.a
	h := cpu.register.h
	l := cpu.register.l
	hl := uint16(h)<<8 | uint16(l)

	data := cpu.mmu.RB(hl)

	r := a & data

	cpu.register.a = r

	cpu.register.setFlag("Z", r == 0x0)
	cpu.register.setFlag("N", false)
	cpu.register.setFlag("H", true)
	cpu.register.setFlag("C", false)

	return 2
}

/*
0xE6 - AND n: Bitwise AND (immediate)

Performs a bitwise AND operation between the 8-bit A register and immediate data n, and
stores the result back into the A register.

Machine Cycles: 2
*/
func (m *instructions) and_n(cpu *CPU) uint32 {
	a := cpu.register.a
	n := cpu.mmu.RB(cpu.popPC())

	r := a & n

	cpu.register.a = r

	cpu.register.setFlag("Z", r == 0x0)
	cpu.register.setFlag("N", false)
	cpu.register.setFlag("H", true)
	cpu.register.setFlag("C", false)

	return 2
}

/*
0xB0 - OR r: Bitwise OR (register)

Performs a bitwise OR operation between the 8-bit A register and the 8-bit register r, and stores
the result back into the A register.

Machine Cycles: 1
*/
func (m *instructions) or_r(cpu *CPU) uint32 {
	a := cpu.register.a
	b := cpu.register.b

	r := a | b

	cpu.register.a = r

	cpu.register.setFlag("Z", r == 0x0)
	cpu.register.setFlag("N", false)
	cpu.register.setFlag("H", false)
	cpu.register.setFlag("C", false)

	return 1
}

/*
0xB6 - OR (HL): Bitwise OR (indirect HL)

Performs a bitwise OR operation between the 8-bit A register and data from the absolute
address specified by the 16-bit register HL, and stores the result back into the A register.

Machine Cycles: 2
*/
func (m *instructions) or_HL(cpu *CPU) uint32 {
	a := cpu.register.a
	h := cpu.register.h
	l := cpu.register.l

	hl := uint16(h)<<8 | uint16(l)

	data := cpu.mmu.RB(hl)

	r := a | data

	cpu.register.a = r

	cpu.register.setFlag("Z", r == 0x0)
	cpu.register.setFlag("N", false)
	cpu.register.setFlag("H", false)
	cpu.register.setFlag("C", false)

	return 2
}

/*
0xF6 - OR n: Bitwise OR (immediate)

# Performs a bitwise OR operation between the 8-bit A register and immediate data n, and stores
the result back into the A register.

Machine Cycles: 2
*/
func (m *instructions) or_n(cpu *CPU) uint32 {
	a := cpu.register.a
	n := cpu.mmu.RB(cpu.popPC())
	r := a | n

	cpu.register.a = r

	cpu.register.setFlag("Z", r == 0x0)
	cpu.register.setFlag("N", false)
	cpu.register.setFlag("H", false)
	cpu.register.setFlag("C", false)

	return 2
}

/*
0xAE - XOR (HL): Bitwise XOR (indirect HL)

Performs a bitwise XOR operation between the 8-bit A register and data from the absolute
address specified by the 16-bit register HL, and stores the result back into the A register.

Machine Cycles: 2
*/
func (m *instructions) xor_HL(cpu *CPU) uint32 {
	a := cpu.register.a
	h := cpu.register.h
	l := cpu.register.l

	hl := uint16(h)<<8 | uint16(l)

	data := cpu.mmu.RB(hl)

	r := a ^ data

	cpu.register.a = r

	cpu.register.setFlag("Z", r == 0x0)
	cpu.register.setFlag("N", false)
	cpu.register.setFlag("H", false)
	cpu.register.setFlag("C", false)

	return 2
}

/*
0xA8 - XOR r: Bitwise XOR (register)

Performs a bitwise XOR operation between the 8-bit A register and the 8-bit register r, and
stores the result back into the A register.

Machine Cycles: 1
*/
func (m *instructions) xor_r(cpu *CPU) uint32 {
	a := cpu.register.a
	b := cpu.register.b

	r := a ^ b

	cpu.register.a = r

	cpu.register.setFlag("Z", r == 0x0)
	cpu.register.setFlag("N", false)
	cpu.register.setFlag("H", false)
	cpu.register.setFlag("C", false)

	return 1
}

/*
0xEE - XOR n: Bitwise XOR (immediate)

Performs a bitwise XOR operation between the 8-bit A register and immediate data n, and
stores the result back into the A register.

Machine Cycles: 2
*/
func (m *instructions) xor_n(cpu *CPU) uint32 {
	a := cpu.register.a
	n := cpu.mmu.RB(cpu.popPC())

	r := a ^ n

	cpu.register.a = r

	cpu.register.setFlag("Z", r == 0x0)
	cpu.register.setFlag("N", false)
	cpu.register.setFlag("H", false)
	cpu.register.setFlag("C", false)

	return 2
}

/*
0x3F - CCF: Complement carry flag

Flips the carry flag, and clears the N and H flags.

Machine Cycles: 1
*/
func (m *instructions) ccf(cpu *CPU) uint32 {
	cpu.register.setFlag("N", false)
	cpu.register.setFlag("H", false)
	cpu.register.setFlag("C", !cpu.register.getFlag("C"))

	return 1
}

/*
0x37 - SCF: Set carry flag

Sets the carry flag, and clears the N and H flags.

Machine Cycles: 1
*/
func (m *instructions) scf(cpu *CPU) uint32 {
	cpu.register.setFlag("N", false)
	cpu.register.setFlag("H", false)
	cpu.register.setFlag("C", true)

	return 1
}

/*
0x27 - DAA: Decimal adjust accumulator

# TODO

Machine Cycles: 1
*/
func (m *instructions) daa(_ *CPU) uint32 {
	fmt.Println("NOT IMPLEMENTED")
	return 1
}

/*
0x2F - CPL: Complement accumulator

Flips all the bits in the 8-bit A register, and sets the N and H flags.

Machine Cycles: 1
*/
func (m *instructions) cpl(cpu *CPU) uint32 {

	cpu.register.a = ^cpu.register.a
	cpu.register.setFlag("N", true)
	cpu.register.setFlag("H", true)
	return 1
}

// ---- 16-Bit Arithmetic and logical ----

/*
0x03 - INC rr: Increment 16-bit register

Increments data in the 16-bit register rr.

Machine Cycles: 2
*/
func (m *instructions) inc_rr(cpu *CPU) uint32 {
	b := cpu.register.b
	c := cpu.register.c

	bc := uint16(b)<<8 | uint16(c)

	bc += 1

	cpu.register.b = uint8(bc >> 8)
	cpu.register.c = uint8(b & 0x00FF)

	return 2
}

/*
0x0B - DEC rr: Decrement 16-bit register

Decrements data in the 16-bit register rr.

Machine Cycles: 2
*/
func (m *instructions) dec_rr(cpu *CPU) uint32 {
	b := cpu.register.b
	c := cpu.register.c

	bc := uint16(b)<<8 | uint16(c)

	bc -= 1

	cpu.register.b = uint8(bc >> 8)
	cpu.register.c = uint8(b & 0x00FF)

	return 2
}

/*
0x09 - ADD HL, rr: Add (16-bit register)

Adds to the 16-bit HL register pair, the 16-bit register rr, and stores the result back into the HL
register pair.

Machine Cycles: 2
*/
func (m *instructions) add_HL_rr(cpu *CPU) uint32 {
	b := cpu.register.b
	c := cpu.register.c
	bc := uint16(b)<<8 | uint16(c)

	h := cpu.register.h
	l := cpu.register.l
	hl := uint16(h)<<8 | uint16(l)

	r := hl + bc

	cpu.register.setFlag("N", false)
	cpu.register.setFlag("H", (bc&0x07FF)+(hl&0x07FF) > 0x07FF)
	cpu.register.setFlag("C", bc > 0xFFFF-hl)

	cpu.register.h = uint8(r >> 8)
	cpu.register.l = uint8(r & 0x00FF)

	return 2
}

/*
0xE8 - ADD SP, e: Add to stack pointer (relative)

Loads to the 16-bit SP register, 16-bit data calculated by adding the signed 8-bit operand e to
the 16-bit value of the SP register.

Machine Cycles: 4
*/
func (m *instructions) add_sp_e(cpu *CPU) uint32 {
	e := uint16(cpu.mmu.RB(cpu.popPC()))

	r := cpu.SP + e

	cpu.register.setFlag("Z", false)
	cpu.register.setFlag("N", false)
	cpu.register.setFlag("H", (cpu.SP&0x000F)+(e&0x000F) > 0x000F)
	cpu.register.setFlag("C", (cpu.SP&0x00FF)+(e&0x00FF) > 0x00FF)

	cpu.SP = r

	return 4
}

// ---- Rotate, shift and bit ----

/*
0x07 - RLCA: Rotate left circular (accumulator)

Rotates the 8-bit A register value left in a circular manner (carry flag is updated but not used).
Every bit is shifted to the left (e.g. bit 1 value is copied from bit 0). Bit 7 is copied both to bit
0 and the carry flag. Note that unlike the related RLC r  instruction, RLCA always sets the zero
flag to 0 without looking at the resulting value of the calculation.

Machine Cycles: 1
*/
func (m *instructions) rlca(cpu *CPU) uint32 {
	a := cpu.register.a

	b7 := (a & (1 << 7)) >> 7

	cpu.register.setFlag("Z", false)
	cpu.register.setFlag("N", false)
	cpu.register.setFlag("H", false)
	cpu.register.setFlag("C", b7 != 0)

	cpu.register.a = a<<1 | b7

	return 1
}

/*
0x0F - RRCA: Rotate right circular (accumulator)

Rotates the 8-bit A register value right in a circular manner (carry flag is updated but not used).
Every bit is shifted to the right (e.g. bit 1 value is copied to bit 0). Bit 0 is copied both to bit 7
and the carry flag. Note that unlike the related RRC r  instruction, RRCA always sets the zero
flag to 0 without looking at the resulting value of the calculation.

Machine Cycles: 1
*/
func (m *instructions) rrca(cpu *CPU) uint32 {
	a := cpu.register.a

	b0 := (a & (1 << 0))

	cpu.register.setFlag("Z", false)
	cpu.register.setFlag("N", false)
	cpu.register.setFlag("H", false)
	cpu.register.setFlag("C", b0 != 0)

	cpu.register.a = a>>1 | b0

	return 1
}

/*
0x17 - RLA: Rotate left (accumulator)

Rotates the 8-bit A register value left through the carry flag.
Every bit is shifted to the left (e.g. bit 1 value is copied from bit 0). The carry flag is copied to bit
0, and bit 7 is copied to the carry flag. Note that unlike the related RL r  instruction, RLA always
sets the zero flag to 0 without looking at the resulting value of the calculation.

Machine Cycles: 1
*/
func (m *instructions) rla(cpu *CPU) uint32 {
	c := uint8(0)
	if cpu.register.getFlag("C") {
		c = 1
	}

	a := cpu.register.a

	b7 := (a & (1 << 7)) >> 7

	cpu.register.setFlag("Z", false)
	cpu.register.setFlag("N", false)
	cpu.register.setFlag("H", false)
	cpu.register.setFlag("C", b7 != 0)

	cpu.register.a = a<<1 | c

	return 1
}

/*
0x1F - RRA: Rotate right (accumulator)

Rotates the 8-bit A register value right through the carry flag.
Every bit is shifted to the right (e.g. bit 1 value is copied to bit 0). The carry flag is copied to bit
7, and bit 0 is copied to the carry flag. Note that unlike the related RR r  instruction, RRA always
sets the zero flag to 0 without looking at the resulting value of the calculation.

Machine Cycles: 1
*/
func (m *instructions) rra(cpu *CPU) uint32 {
	c := uint8(0)
	if cpu.register.getFlag("C") {
		c = 1
	}
	a := cpu.register.a

	b0 := (a & (1 << 0))

	cpu.register.setFlag("Z", false)
	cpu.register.setFlag("N", false)
	cpu.register.setFlag("H", false)
	cpu.register.setFlag("C", b0 != 0)

	cpu.register.a = a>>1 | c

	return 1
}

/*
0xCB + 0x00 - RLC r: Rotate left circular (register)

Rotates the 8-bit register r value left in a circular manner (carry flag is updated but not used).
Every bit is shifted to the left (e.g. bit 1 value is copied from bit 0). Bit 7 is copied both to bit 0

Machine Cycles: 2
*/
func (m *instructions) rlc_r(cpu *CPU) uint32 {
	b := cpu.register.b

	b7 := (b & (1 << 7)) >> 7

	cpu.register.setFlag("Z", false)
	cpu.register.setFlag("N", false)
	cpu.register.setFlag("H", false)
	cpu.register.setFlag("C", b7 != 0)

	cpu.register.b = b<<1 | b7

	return 2
}

/*
0xCB + 0x06 - RLC (HL): Rotate left circular (indirect HL)

Rotates, the 8-bit data at the absolute address specified by the 16-bit register HL, left in a
circular manner (carry flag is updated but not used).
Every bit is shifted to the left (e.g. bit 1 value is copied from bit 0). Bit 7 is copied both to bit 0
and the carry flag.

Machine Cycles: 4
*/
func (m *instructions) rlc_HL(cpu *CPU) uint32 {
	h := cpu.register.h
	l := cpu.register.l
	hl := uint16(h)<<8 | uint16(l)

	data := cpu.mmu.RB(hl)

	b7 := (data & (1 << 7)) >> 7

	cpu.register.setFlag("Z", b7 == 0)
	cpu.register.setFlag("N", false)
	cpu.register.setFlag("H", false)
	cpu.register.setFlag("C", b7 != 0)

	cpu.register.b = data<<1 | b7

	return 4
}

/*
0xCB + 0x08 - RRC r: Rotate right circular (register)

Rotates the 8-bit register r value right in a circular manner (carry flag is updated but not used).
Every bit is shifted to the right (e.g. bit 1 value is copied to bit 0). Bit 0 is copied both to bit 7
and the carry flag.

Machine Cycles: 2
*/
func (m *instructions) rrc_r(cpu *CPU) uint32 {
	b := cpu.register.b

	b0 := (b & (1 << 0))

	cpu.register.setFlag("Z", false)
	cpu.register.setFlag("N", false)
	cpu.register.setFlag("H", false)
	cpu.register.setFlag("C", b0 != 0)

	cpu.register.b = b>>1 | b0

	return 2
}

/*
0xCB + 0x0E - RLC (HL): Rotate left circular (indirect HL)

Rotates, the 8-bit data at the absolute address specified by the 16-bit register HL, left in a
circular manner (carry flag is updated but not used).
Every bit is shifted to the left (e.g. bit 1 value is copied from bit 0). Bit 7 is copied both to bit 0
and the carry flag.

Machine Cycles: 4
*/
func (m *instructions) rrc_HL(cpu *CPU) uint32 {
	h := cpu.register.h
	l := cpu.register.l
	hl := uint16(h)<<8 | uint16(l)

	data := cpu.mmu.RB(hl)

	b7 := (data & (1 << 7)) >> 7

	cpu.register.setFlag("Z", b7 == 0)
	cpu.register.setFlag("N", false)
	cpu.register.setFlag("H", false)
	cpu.register.setFlag("C", b7 != 0)

	cpu.register.b = data<<1 | b7

	return 4
}

/*
0xCB + 0x10 - RL r:  Rotate left (register)

Rotates the 8-bit register r value left through the carry flag.
Every bit is shifted to the left (e.g. bit 1 value is copied from bit 0). The carry flag is copied to bit
0, and bit 7 is copied to the carry flag.

Machine Cycles: 2
*/
func (m *instructions) rl_r(cpu *CPU) uint32 {
	b := cpu.register.b

	b0 := (b & (1 << 0))

	cpu.register.setFlag("Z", false)
	cpu.register.setFlag("N", false)
	cpu.register.setFlag("H", false)
	cpu.register.setFlag("C", b0 != 0)

	cpu.register.b = b>>1 | b0

	return 2
}

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

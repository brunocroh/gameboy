package cpu

type instructions struct {
}

func instructionsNew() *instructions {
	return &instructions{}
}

// ---- 8-Bit Load Instructions ----

// ldBB - Load register(register) loads a 16-bit immediate value into the BC register pair.
// Load to the 8-bit register r, data from the 8-bit register r'.
//
// OPCODE: 0x41
// Cycles: 1
func (m *instructions) ldBC(cpu *CPU) uint32 {
	cpu.register.b = cpu.register.c
	return 0
}

// ---- 16-Bit Load Instructions ----

// ---- 8-Bit Arithmetic and logical ----

// ---- 16-Bit Arithmetic and logical ----

// ---- Rotate, shift and bit ----

// ---- FLOW ----

// jpAddr - Load register(register) loads a 16-bit immediate value into the BC register pair.
// Load to the 8-bit register r, data from the 8-bit register r'.
//
// OPCODE: 0xC3
// Cycles: 4
func (m *instructions) jpAddr(cpu *CPU, addr uint16) uint32 {
	cpu.pc = uint16(addr)
	return 3
}

// ---- MISC ----

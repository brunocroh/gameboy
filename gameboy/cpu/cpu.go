package cpu

import (
	"fmt"

	"github.com/brunocroh/gameboy/gameboy/mmu"
)

type CPU struct {
	mmu mmu.MemoryManagementUnit

	ins       *instructions
	register  *register
	interrupt *interrupt

	PC uint16
	SP uint16
}

func New(mmu mmu.MemoryManagementUnit) *CPU {
	i := instructionsNew()
	r := registerNew()
	interrupt := interruptNew(mmu)
	return &CPU{
		mmu:       mmu,
		ins:       i,
		register:  r,
		interrupt: interrupt,
	}
}

func (m *CPU) Init() {
	m.PC = 255
	m.SP = 0xFFFE
	m.register.Init()
	m.interrupt.Init()
}

func (m *CPU) Cycle() {
	opcode := m.fetchOpcode()

	interruptOutput := m.interrupt.handleInterrupt(m.PC)

	m.doCycle(1)
	if interruptOutput != 0 {
		return
	}

	m.execInstruction(opcode)
}

func (m *CPU) popPC() uint16 {
	pc := m.PC
	m.PC += 1
	return pc
}

func (m *CPU) fetchOpcode() byte {
	return m.mmu.RB(m.popPC())
}

func (m *CPU) execInstruction(opcode byte) {
	var ticks uint32 = 1
	switch opcode {
	case 0x00:
		ticks = m.ins.nop()
	case 0x01:
		ticks = m.ins.ld_rr_nn(m, &m.register.b, &m.register.c)
	case 0x02:
		ticks = m.ins.ld_BC_A(m)
	case 0x03:
		ticks = m.ins.inc_rr(m, &m.register.b, &m.register.c)
	case 0x04:
		ticks = m.ins.inc_r(m, &m.register.b)
	case 0x05:
		ticks = m.ins.dec_r(m, &m.register.b)
	case 0x06:
		ticks = m.ins.ld_r_n(m, &m.register.b)
	case 0x07:
		ticks = m.ins.rlca(m)
	case 0x08:
		ticks = m.ins.ld_nn_sp(m)
	case 0x09:
		ticks = m.ins.add_HL_rr(m, &m.register.b, &m.register.c)
	case 0x0A:
		ticks = m.ins.ld_A_BC(m)
	case 0x0B:
		ticks = m.ins.dec_rr(m, &m.register.b, &m.register.c)
	case 0x0C:
		ticks = m.ins.inc_r(m, &m.register.c)
	case 0x0D:
		ticks = m.ins.dec_r(m, &m.register.c)
	case 0x0E:
		ticks = m.ins.ld_r_n(m, &m.register.c)
	case 0x0F:
		ticks = m.ins.rrca(m)
	case 0x10:
		// fmt.Println("0x10 - STOP Instruction")
	case 0x11:
		ticks = m.ins.ld_rr_nn(m, &m.register.d, &m.register.e)
	case 0x12:
		ticks = m.ins.ld_DE_A(m)
	case 0x13:
		ticks = m.ins.inc_rr(m, &m.register.d, &m.register.e)
	case 0x14:
		ticks = m.ins.inc_r(m, &m.register.d)
	case 0x15:
		ticks = m.ins.dec_r(m, &m.register.d)
	case 0x16:
		ticks = m.ins.ld_r_n(m, &m.register.d)
	case 0x17:
		ticks = m.ins.rla(m)
	case 0x18:
		ticks = m.ins.jr_e(m)
	case 0x19:
		ticks = m.ins.add_HL_rr(m, &m.register.d, &m.register.e)
	case 0x1A:
		ticks = m.ins.ld_A_DE(m)
	case 0x1B:
		ticks = m.ins.dec_rr(m, &m.register.d, &m.register.e)
	case 0x1C:
		ticks = m.ins.inc_r(m, &m.register.e)
	case 0x1D:
		ticks = m.ins.dec_r(m, &m.register.e)
	case 0x1E:
		ticks = m.ins.ld_r_n(m, &m.register.e)
	case 0x1F:
		ticks = m.ins.rra(m)
	case 0x20:
		ticks = m.ins.jr_nz(m)
	case 0x21:
		ticks = m.ins.ld_rr_nn(m, &m.register.h, &m.register.l)
	case 0x22:
		ticks = m.ins.ld_HLi_A(m)
	case 0x23:
		ticks = m.ins.inc_rr(m, &m.register.h, &m.register.l)
	case 0x24:
		ticks = m.ins.inc_r(m, &m.register.h)
	case 0x25:
		ticks = m.ins.dec_r(m, &m.register.h)
	case 0x26:
		ticks = m.ins.ld_r_n(m, &m.register.h)
	case 0x27:
		ticks = m.ins.daa(m)
	case 0x28:
		ticks = m.ins.jr_z(m)
	case 0x29:
		ticks = m.ins.add_HL_rr(m, &m.register.h, &m.register.l)
	case 0x2A:
		ticks = m.ins.ld_A_HLi(m)
	case 0x2B:
		ticks = m.ins.dec_rr(m, &m.register.h, &m.register.l)
	case 0x2C:
		ticks = m.ins.inc_r(m, &m.register.l)
	case 0x2D:
		ticks = m.ins.dec_r(m, &m.register.l)
	case 0x2E:
		ticks = m.ins.ld_r_n(m, &m.register.l)
	case 0x2F:
		ticks = m.ins.cpl(m)
	case 0x30:
		ticks = m.ins.jr_nc(m)
	case 0x31:
		ticks = m.ins.ld_SP_nn(m)
	case 0x32:
		ticks = m.ins.ld_HLd_A(m)
	case 0x33:
		ticks = m.ins.inc_sp(m)
	case 0x34:
		ticks = m.ins.inc_HL(m)
	case 0x35:
		ticks = m.ins.dec_HL(m)
	case 0x36:
		ticks = m.ins.ld_HL_n(m)
	case 0x37:
		ticks = m.ins.scf(m)
	case 0x38:
		ticks = m.ins.jr_c(m)
	case 0x39:
		msb := uint8(m.SP >> 8)
		lsb := uint8(m.SP)
		ticks = m.ins.add_HL_rr(m, &msb, &lsb)
	case 0x3A:
		ticks = m.ins.ld_A_HLd(m)
	case 0x3B:
		ticks = m.ins.dec_sp(m)
	case 0x3C:
		ticks = m.ins.inc_r(m, &m.register.a)
	case 0x3D:
		ticks = m.ins.dec_r(m, &m.register.a)
	case 0x3E:
		ticks = m.ins.ld_r_n(m, &m.register.a)
	case 0x3F:
		ticks = m.ins.ccf(m)
	case 0x40, 0x41, 0x42, 0x43, 0x44, 0x45, 0x47:
		ticks = m.ins.ld_rr(&m.register.b, getRegister(m, opcode))
	case 0x46:
		ticks = m.ins.ld_r_HL(m, &m.register.b)
	case 0x48, 0x49, 0x4A, 0x4B, 0x4C, 0x4D, 0x4F:
		ticks = m.ins.ld_rr(&m.register.c, getRegister(m, opcode))
	case 0x4E:
		ticks = m.ins.ld_r_HL(m, &m.register.c)
	case 0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x57:
		ticks = m.ins.ld_rr(&m.register.d, getRegister(m, opcode))
	case 0x56:
		ticks = m.ins.ld_r_HL(m, &m.register.d)
	case 0x58, 0x59, 0x5A, 0x5B, 0x5C, 0x5D, 0x5F:
		ticks = m.ins.ld_rr(&m.register.e, getRegister(m, opcode))
	case 0x5E:
		ticks = m.ins.ld_r_HL(m, &m.register.e)
	case 0x60, 0x61, 0x62, 0x63, 0x64, 0x65, 0x67:
		ticks = m.ins.ld_rr(&m.register.h, getRegister(m, opcode))
	case 0x66:
		ticks = m.ins.ld_r_HL(m, &m.register.h)
	case 0x68, 0x69, 0x6A, 0x6B, 0x6C, 0x6D, 0x6F:
		ticks = m.ins.ld_rr(&m.register.l, getRegister(m, opcode))
	case 0x6E:
		ticks = m.ins.ld_r_HL(m, &m.register.l)
	case 0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x77:
		ticks = m.ins.ld_HL_r(m, getRegister(m, opcode))
	case 0x78, 0x79, 0x7A, 0x7B, 0x7C, 0x7D, 0x7F:
		ticks = m.ins.ld_rr(&m.register.a, getRegister(m, opcode))
	case 0x7E:
		ticks = m.ins.ld_r_HL(m, &m.register.a)
	case 0x80:
		ticks = m.ins.add_r(m)
	case 0x86:
		ticks = m.ins.add_HL(m)
	case 0x88:
		ticks = m.ins.adc_r(m)
	case 0x8E:
		ticks = m.ins.adc_HL(m)
	case 0x90:
		ticks = m.ins.sub_r(m)
	case 0x96:
		ticks = m.ins.sub_HL(m)
	case 0x98:
		ticks = m.ins.sbc_r(m)
	case 0x9E:
		ticks = m.ins.sbc_HL(m)
	case 0xA0:
		ticks = m.ins.and_r(m)
	case 0xA6:
		ticks = m.ins.and_HL(m)
	case 0xA8, 0xA9, 0xAA, 0xAB, 0xAC, 0xAD, 0xAF:
		ticks = m.ins.xor_r(m, getRegister(m, opcode))
	case 0xAE:
		ticks = m.ins.xor_HL(m)
	case 0xB0, 0xB1, 0xB2, 0xB3, 0xB4, 0xB5, 0xB7:
		ticks = m.ins.or_a_r(m, getRegister(m, opcode))
	case 0xB6:
		ticks = m.ins.or_a_HL(m)
	case 0xB8, 0xB9, 0xBA, 0xBB, 0xBC, 0xBD, 0xBF:
		ticks = m.ins.cp_A_r(m, getRegister(m, opcode))
	case 0xBE:
		ticks = m.ins.cp_HL(m)
	case 0xC0:
		ticks = m.ins.ret_cc(m, !m.register.getFlag("Z")) // NZ
	case 0xC1:
		ticks = m.ins.ld_pop_rr(m, &m.register.b, &m.register.c)
	case 0xC2:
		ticks = m.ins.jp_cc_nn(m, !m.register.getFlag("Z")) // NZ
	case 0xC3:
		ticks = m.ins.jp_nn(m)
	case 0xC4:
		ticks = m.ins.call_cc_nn(m)
	case 0xC5:
		ticks = m.ins.ld_push_rr(m, &m.register.b, &m.register.c)
	case 0xC6:
		ticks = m.ins.add_n(m)
	case 0xC8:
		ticks = m.ins.ret_cc(m, m.register.getFlag("Z")) // Z
	case 0xC9:
		ticks = m.ins.ret(m)
	case 0xCA:
		ticks = m.ins.jp_cc_nn(m, m.register.getFlag("Z")) // Z
	case 0xCB:
		op := m.fetchOpcode()
		switch op & 0xF0 {
		case 0x00:
			if op < 0x08 {
				if op == 0x06 {
					ticks = m.ins.rlc_HL(m)
				} else {
					ticks = m.ins.rlc_r(m, getRegister(m, op))
				}
			} else {
				if op == 0x0E {
					ticks = m.ins.rrc_HL(m)
				} else {
					ticks = m.ins.rrc_r(m, getRegister(m, op))
				}
			}
		case 0x10:
			if op < 0x18 {
				if op == 0x16 {
					ticks = m.ins.rl_HL(m)
				} else {
					ticks = m.ins.rl_r(m, getRegister(m, op))
				}
			} else {
				if op == 0x1E {

					ticks = m.ins.rr_HL(m)
				} else {

					ticks = m.ins.rr_r(m, getRegister(m, op))
				}
			}
		case 0x20:
			if op < 0x28 {
				if op == 0x26 {
					ticks = m.ins.sla_HL(m)
				} else {
					ticks = m.ins.sla_r(m, getRegister(m, op))
				}
			} else {
				if op == 0x2E {
					ticks = m.ins.sra_HL(m)
				} else {
					ticks = m.ins.sra_r(m, getRegister(m, op))
				}
			}
		case 0x30:
			if op < 0x38 {
				if op == 0x36 {
					ticks = m.ins.swap_HL(m)
				} else {
					ticks = m.ins.swap_r(m, getRegister(m, op))
				}
			} else {
				if op == 0x3E {
					ticks = m.ins.srl_HL(m)
				} else {
					ticks = m.ins.srl_r(m, getRegister(m, op))
				}
			}
		case 0x40:
			if op < 0x48 {
				if op == 0x46 {
					ticks = m.ins.bit_u3_HL(m, 0)
				} else {
					ticks = m.ins.bit_u3_r(m, 0, getRegister(m, op))
				}
			} else {
				if op == 0x4E {
					ticks = m.ins.bit_u3_HL(m, 1)
				} else {
					ticks = m.ins.bit_u3_r(m, 1, getRegister(m, op))
				}
			}
		case 0x50:
			if op < 0x58 {
				if op == 0x56 {
					ticks = m.ins.bit_u3_HL(m, 2)
				} else {
					ticks = m.ins.bit_u3_r(m, 2, getRegister(m, op))
				}
			} else {
				if op == 0x5E {
					ticks = m.ins.bit_u3_HL(m, 3)
				} else {
					ticks = m.ins.bit_u3_r(m, 3, getRegister(m, op))
				}
			}
		case 0x60:
			if op < 0x68 {
				if op == 0x66 {
					ticks = m.ins.bit_u3_HL(m, 4)
				} else {
					ticks = m.ins.bit_u3_r(m, 4, getRegister(m, op))
				}
			} else {
				if op == 0x6E {
					ticks = m.ins.bit_u3_HL(m, 5)
				} else {
					ticks = m.ins.bit_u3_r(m, 5, getRegister(m, op))
				}
			}
		case 0x70:
			if op < 0x78 {
				if op == 0x76 {
					ticks = m.ins.bit_u3_HL(m, 6)
				} else {
					ticks = m.ins.bit_u3_r(m, 6, getRegister(m, op))
				}
			} else {
				if op == 0x7E {
					ticks = m.ins.bit_u3_HL(m, 7)
				} else {
					ticks = m.ins.bit_u3_r(m, 7, getRegister(m, op))
				}

			}
		case 0x80:
			if op == 0x86 {
				ticks = m.ins.res_b_HL(m)
			} else {
				ticks = m.ins.res_b_r(m, getRegister(m, op))
			}
		case 0xC0:
			if op == 0xC6 {
				ticks = m.ins.set_b_HL(m)
			} else {
				ticks = m.ins.set_b_r(m, getRegister(m, op))
			}
		default:
			fmt.Printf("CB opcode (0x%x) not implemented\n", op)
		}
	case 0xCD:
		ticks = m.ins.call_nn(m)
	case 0xCE:
		ticks = m.ins.adc_n(m)
	case 0xCF:
		ticks = m.ins.rst_n(m, 0x08)
	case 0xD0:
		ticks = m.ins.ret_cc(m, !m.register.getFlag("C")) // NC
	case 0xD1:
		ticks = m.ins.ld_pop_rr(m, &m.register.d, &m.register.e)
	case 0xD2:
		ticks = m.ins.jp_cc_nn(m, !m.register.getFlag("C")) // NC
	case 0xD5:
		ticks = m.ins.ld_push_rr(m, &m.register.d, &m.register.e)
	case 0xD6:
		ticks = m.ins.sub_n(m)
	case 0xD8:
		ticks = m.ins.ret_cc(m, m.register.getFlag("C")) // C
	case 0xDA:
		ticks = m.ins.jp_cc_nn(m, m.register.getFlag("C")) // C
	case 0xDE:
		ticks = m.ins.sbc_n(m)
	case 0xDF:
		ticks = m.ins.rst_n(m, 0x18)
	case 0xE0:
		ticks = m.ins.ldh_n_A(m)
	case 0xE1:
		ticks = m.ins.ld_pop_rr(m, &m.register.h, &m.register.l)
	case 0xE2:
		ticks = m.ins.ldh_C_A(m)
	case 0xE5:
		ticks = m.ins.ld_push_rr(m, &m.register.h, &m.register.l)
	case 0xE6:
		ticks = m.ins.and_n(m)
	case 0xE8:
		ticks = m.ins.add_sp_e(m)
	case 0xEA:
		ticks = m.ins.ld_nn_A(m)
	case 0xEE:
		ticks = m.ins.xor_n(m)
	case 0xEF:
		ticks = m.ins.rst_n(m, 0x28)
	case 0xE9:
		ticks = m.ins.jp_HL(m)
	case 0xF0:
		ticks = m.ins.ldh_A_n(m)
	case 0xF1:
		ticks = m.ins.ld_pop_rr(m, &m.register.a, &m.register.f)
	case 0xF2:
		ticks = m.ins.ldh_A_C(m)
	case 0xF3:
		ticks = m.ins.di(m)
	case 0xF5:
		ticks = m.ins.ld_push_rr(m, &m.register.a, &m.register.f)
	case 0xF6:
		ticks = m.ins.or_n(m)
	case 0xF8:
		ticks = m.ins.ld_HL_spe(m)
	case 0xF9:
		ticks = m.ins.ld_sp_HL(m)
	case 0xFA:
		ticks = m.ins.ld_A_nn(m)
	case 0xFE:
		ticks = m.ins.cp_n(m)
	case 0xFF:
		ticks = m.ins.rst_n(m, 0x38)
	default:
		fmt.Printf("opcode (0x%x) not implemented\n", opcode)
	}

	if ticks != 0 {
		m.doCycle(ticks)
	}

	fmt.Printf("A:%02x F:%02x B:%02x C:%02x D:%02x E:%02x H:%02x L:%02x SP:%04x PC:%04x PCMEM:%02x,%02x,%02x,%02x\n",
		m.register.a,
		m.register.f,
		m.register.b,
		m.register.c,
		m.register.d,
		m.register.e,
		m.register.h,
		m.register.l,
		m.SP,
		m.PC,
		m.mmu.RB(m.PC),
		m.mmu.RB(m.PC+1),
		m.mmu.RB(m.PC+2),
		m.mmu.RB(m.PC+3))

}

func getRegister(m *CPU, opcode byte) *uint8 {
	switch opcode & 0x0F {
	case 0x00, 0x08:
		return &m.register.b
	case 0x01, 0x09:
		return &m.register.c
	case 0x02, 0x0A:
		return &m.register.d
	case 0x03, 0x0B:
		return &m.register.e
	case 0x04, 0x0C:
		return &m.register.h
	case 0x05, 0x0D:
		return &m.register.l
	case 0x07, 0x0F:
		return &m.register.a
	}

	fmt.Printf("no register BRO %02x\n", opcode&0x0F)

	return nil
}

func (m *CPU) rw(addr uint16) uint16 {
	value := m.mmu.RW(addr)
	m.PC += 2
	return value
}

func (m *CPU) doCycle(ticks uint32) {
	m.mmu.DoCycle(ticks * 4)
}

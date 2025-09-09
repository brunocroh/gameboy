package cpu

import (
	"fmt"

	"github.com/brunocroh/gameboy/gameboy/mmu"
)

type CPU struct {
	mmu *mmu.MemoryManagementUnit

	ins      *instructions
	register *register

	PC uint16
	SP uint16
}

func New(mmu *mmu.MemoryManagementUnit) *CPU {
	i := instructionsNew()
	r := registerNew()
	return &CPU{
		mmu:      mmu,
		ins:      i,
		register: r,
	}
}

func (m *CPU) Init() {
	m.PC = 0x0100
	m.SP = 0xFFFE
	m.register.Init()
}

func (m *CPU) Cycle() {
	opcode := m.fetchOpcode()
	m.doCycle(1)
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
	case 0x01:
		ticks = m.ins.ld_rr_nn(m)
	case 0x02:
		ticks = m.ins.ld_BC_A(m)
	case 0x03:
		ticks = m.ins.inc_rr(m)
	case 0x04:
		ticks = m.ins.inc_r(m)
	case 0x05:
		ticks = m.ins.dec_r(m)
	case 0x06:
		ticks = m.ins.ld_r_n(m)
	case 0x07:
		ticks = m.ins.rlca(m)
	case 0x08:
		ticks = m.ins.ld_nn_sp(m)
	case 0x09:
		ticks = m.ins.add_HL_rr(m)
	case 0x0A:
		ticks = m.ins.ld_A_BC(m)
	case 0x0B:
		ticks = m.ins.dec_rr(m)
	case 0x0F:
		ticks = m.ins.rrca(m)
	case 0x12:
		ticks = m.ins.ld_DE_A(m)
	case 0x17:
		ticks = m.ins.rla(m)
	case 0x18:
		ticks = m.ins.jr_e(m)
	case 0x1A:
		ticks = m.ins.ld_A_DE(m)
	case 0x1F:
		ticks = m.ins.rra(m)
	case 0x20:
		ticks = m.ins.jr_cc(m)
	case 0x22:
		ticks = m.ins.ld_HLi_A(m)
	case 0x27:
		ticks = m.ins.daa(m)
	case 0x2A:
		ticks = m.ins.ld_A_HLi(m)
	case 0x2F:
		ticks = m.ins.cpl(m)
	case 0x32:
		ticks = m.ins.ld_HLd_A(m)
	case 0x34:
		ticks = m.ins.inc_HL(m)
	case 0x35:
		ticks = m.ins.dec_HL(m)
	case 0x36:
		ticks = m.ins.ld_HL_n(m)
	case 0x37:
		ticks = m.ins.scf(m)
	case 0x3A:
		ticks = m.ins.ld_A_HLd(m)
	case 0x3F:
		ticks = m.ins.ccf(m)
	case 0x41:
		ticks = m.ins.ld_rr(m)
	case 0x46:
		ticks = m.ins.ld_r_HL(m)
	case 0x70:
		ticks = m.ins.ld_HL_r(m)
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
	case 0xA8:
		ticks = m.ins.xor_r(m)
	case 0xAE:
		ticks = m.ins.xor_HL(m)
	case 0xB0:
		ticks = m.ins.or_r(m)
	case 0xB6:
		ticks = m.ins.or_HL(m)
	case 0xB8:
		ticks = m.ins.cp_r(m)
	case 0xBE:
		ticks = m.ins.cp_HL(m)
	case 0xC0:
		ticks = m.ins.ret_cc(m)
	case 0xC1:
		ticks = m.ins.ld_pop_rr(m)
	case 0xC2:
		ticks = m.ins.jp_cc_nn(m)
	case 0xC3:
		ticks = m.ins.jp_nn(m)
	case 0xC4:
		ticks = m.ins.call_cc_nn(m)
	case 0xC5:
		ticks = m.ins.ld_push_rr(m)
	case 0xC6:
		ticks = m.ins.add_n(m)
	case 0xC9:
		ticks = m.ins.ret(m)
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
			if op == 0x46 {
				ticks = m.ins.bit_b_HL(m)
			} else {
				ticks = m.ins.bit_b_r(m, getRegister(m, op))
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
		}
	case 0xCD:
		ticks = m.ins.call_nn(m)
	case 0xCE:
		ticks = m.ins.adc_n(m)
	case 0xD6:
		ticks = m.ins.sub_n(m)
	case 0xDE:
		ticks = m.ins.sbc_n(m)
	case 0xDF:
		ticks = m.ins.rst_n(m)
	case 0xE0:
		ticks = m.ins.ldh_n_A(m)
	case 0xE2:
		ticks = m.ins.ldh_C_A(m)
	case 0xE6:
		ticks = m.ins.and_n(m)
	case 0xE8:
		ticks = m.ins.add_sp_e(m)
	case 0xEA:
		ticks = m.ins.ld_nn_A(m)
	case 0xEE:
		ticks = m.ins.xor_n(m)
	case 0xE9:
		ticks = m.ins.jp_HL(m)
	case 0xF0:
		ticks = m.ins.ldh_A_n(m)
	case 0xF2:
		ticks = m.ins.ldh_A_C(m)
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
	default:
		fmt.Printf("opcode (0x%x) not implemented\n", opcode)
	}

	if ticks != 0 {
		m.doCycle(ticks)
	}
}

func getRegister(m *CPU, opcode byte) *uint8 {
	switch opcode & 0x0F {
	case 0x00:
	case 0x08:
		return &m.register.b
	case 0x01:
	case 0x09:
		return &m.register.c
	case 0x02:
	case 0x0A:
		return &m.register.d
	case 0x03:
	case 0x0B:
		return &m.register.e
	case 0x04:
	case 0x0C:
		return &m.register.h
	case 0x05:
	case 0x0D:
		return &m.register.l
	case 0x07:
	case 0x0F:
		return &m.register.a
	}

	return nil
}

func (m *CPU) rw(addr uint16) uint16 {
	value := m.mmu.RW(addr)
	m.PC += 2
	return value
}

func (m *CPU) doCycle(ticks uint32) uint32 {
	return m.mmu.DoCycle(ticks * 4)
}

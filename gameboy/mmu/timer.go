package mmu

const DIVIDER = 0xFF04
const COUNTER = 0xFF05
const MODULO = 0xFF06
const CONTROL = 0xFF07

type Timer struct {
	divider uint8
	counter uint8
	modulo  uint8
	control uint32
	enabled bool
}

func TimerNew() *Timer {
	return &Timer{}
}

func (m *Timer) Init() {
	m.enabled = true
	m.divider = 0
	m.counter = 0
	m.modulo = 0
	m.control = 0
}

func IsTimerAddress(address uint16) bool {
	if address >= DIVIDER || address <= CONTROL {
		return true
	}
	return false
}

func (m *Timer) read(address uint16) uint8 {
	switch address {
	case DIVIDER:
		return m.divider
	case COUNTER:
		return m.counter
	case MODULO:
		return m.modulo
	default:
		return m.modulo //TODO update to control
	}
}

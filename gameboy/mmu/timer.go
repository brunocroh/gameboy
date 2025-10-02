package mmu

const (
	DIV  = 0xFF04 // Divider
	TIMA = 0xFF05 // Counter
	TMA  = 0xFF06 // Modulo
	TAC  = 0xFF07 // Control
)

type Timer struct {
	divider byte
	counter byte
	modulo  byte
	control byte
	step    byte
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
	if address >= DIV || address <= TAC {
		return true
	}
	return false
}

func (m *Timer) Read(address uint16) byte {
	switch address {
	case DIV:
		return m.divider
	case TIMA:
		return m.counter
	case TMA:
		return m.modulo
	default:
		return m.modulo //TODO update to control
	}
}

func (m *Timer) Write(address uint16, v byte) {
	switch address {
	case DIV:
		m.divider = 0
		break
	case TIMA:
		m.counter = v
		break
	case TMA:
		m.modulo = v
	default:
		m.enabled = (v & 0x4) != 0

	}
}

func (m *Timer) DoCycle(ticks uint32) {
	switch address {
	case DIV:
		m.divider = 0
		break
	case TIMA:
		m.counter = v
		break
	case TMA:
		m.modulo = v
	default:
		return m.modulo //TODO update to control
	}
}

package mmu

const (
	DIV  = 0xFF04 // Divider
	TIMA = 0xFF05 // Counter
	TMA  = 0xFF06 // Modulo
	TAC  = 0xFF07 // Control
)

type Timer struct {
	divider         byte
	internalDivider uint32
	counter         byte
	internalCounter uint32
	modulo          byte
	step            uint32
	enabled         bool
	Interrupt       byte
}

func TimerNew() *Timer {
	return &Timer{}
}

func (m *Timer) Init() {
	m.enabled = true
	m.divider = 0
	m.internalDivider = 0
	m.counter = 0
	m.internalCounter = 0
	m.modulo = 0
	m.step = 256
	m.Interrupt = 0
}

func IsTimerAddress(address uint16) bool {
	if address >= DIV || address <= TAC {
		return true
	}
	return false
}

func (m *Timer) read(address uint16) byte {
	switch address {
	case DIV:
		return m.divider
	case TIMA:
		return m.counter
	case TMA:
		return m.modulo
	default:
		var stepCount = byte(0)

		switch m.step {
		case 16:
			stepCount = 1
		case 64:
			stepCount = 2
		case 256:
			stepCount = 3
		default:
			stepCount = byte(0)
		}

		if m.enabled {
			return 0xF8 | 0x4 | stepCount
		} else {
			return 0xF8 | stepCount
		}
	}
}

func (m *Timer) write(address uint16, v byte) {
	switch address {
	case DIV:
		m.divider = 0
	case TIMA:
		m.counter = v
	case TMA:
		m.modulo = v
	case TAC:
		m.enabled = (v & 0x4) != 0
		switch v & 0x3 {
		case 1:
			m.step = 16
		case 2:
			m.step = 64
		case 3:
			m.step = 256
		default:
			m.step = 1024
		}
	}
}

func (m *Timer) DoCycle(ticks uint32) {
	m.internalDivider += ticks

	for m.internalDivider >= 256 {
		m.divider = m.divider + 1
		m.internalDivider -= 256
	}

	if m.enabled {
		m.internalCounter += ticks

		for m.internalCounter >= m.step {
			m.counter = m.counter + 1
			if m.counter == 0 {
				m.counter = m.modulo
				m.Interrupt |= 0x04
			}
			m.internalCounter -= m.step
		}
	}
}

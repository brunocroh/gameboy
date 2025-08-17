package cpu

type register struct {
	a uint8
	b uint8
	c uint8
	d uint8
	e uint8
	h uint8
	l uint8

	f uint8
}

func registerNew() *register {
	return &register{}
}

func (m *register) Init() {
	m.a = 0x01
	m.f = 0xB0 //check later the initial value and how to handle flags
	m.b = 0x00
	m.c = 0x13
	m.d = 0x00
	m.e = 0xD8
	m.h = 0x01
	m.l = 0x4D
}

func (m *register) getFlag(flag string) bool {
	switch flag {
	case "Z":
		return m.f&(1<<7) != 0
	case "N":
		return m.f&(1<<6) != 0
	case "H":
		return m.f&(1<<5) != 0
	case "C":
		return m.f&(1<<4) != 0
	default:
		return false
	}
}

func (m *register) setFlag(flag string, value bool) {
	var newBit uint8
	if value == true {
		newBit = 1
	} else {
		newBit = 0
	}

	switch flag {
	case "Z":
		m.f = (m.f & ^uint8(1<<7)) | (newBit << 7)
	case "N":
		m.f = (m.f & ^uint8(1<<6)) | (newBit << 6)
	case "H":
		m.f = (m.f & ^uint8(1<<5)) | (newBit << 5)
	case "C":
		m.f = (m.f & ^uint8(1<<4)) | (newBit << 4)
	}
}

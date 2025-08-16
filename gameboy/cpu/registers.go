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

func (m *register) HL() uint16 {
	return uint16(m.h)<<8 | uint16(m.l)
}

func (m *register) BC() uint16 {
	return uint16(m.b)<<8 | uint16(m.c)
}

func (m *register) DE() uint16 {
	return uint16(m.d)<<8 | uint16(m.e)
}

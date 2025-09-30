package cpu

type interupt struct {
	IME bool
	IE  uint16
	IF  uint16
}

func interuptNew() *interupt {
	return &interupt{}
}

func (m *interupt) Init() {
	m.IME = false
	m.IE = uint16(0x0000)
	m.IF = uint16(0x0000)
}

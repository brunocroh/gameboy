package mbc

type MemoryBankController struct {
}

func New() *MemoryBankController {

	return &MemoryBankController{}
}

func (m *MemoryBankController) RB(address uint16) byte {
	return 0xFF
}

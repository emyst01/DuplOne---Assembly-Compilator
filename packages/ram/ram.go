package ram

import (
	"fmt"
)

type RAM struct {
	memory map[uint16]uint8
}

func NewRAM() *RAM {
	return &RAM{
		memory: make(map[uint16]uint8),
	}
}

func (ram *RAM) Write(address uint16, value uint8) {
	ram.memory[address] = value
}

func (ram *RAM) Read(address uint16) uint8 {
	if value, exists := ram.memory[address]; exists {
		return value
	}
	return 0
}

func (ram *RAM) Dump() {
	for address, value := range ram.memory {
		fmt.Printf("Address: %d, Value: %d\n", address, value)
	}
}

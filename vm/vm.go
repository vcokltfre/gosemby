package vm

import (
	"encoding/binary"
	"fmt"
	"os"
)

func ExecBytecode(bytecode []byte, stacksize int, memsize int) {
	stack := Stack{}
	stack.Init(stacksize)

	registers := make(map[int]int)
	memory := make([]int, memsize)

	ptr := 0
	length := len(bytecode)

	for ptr < length {
		rawInstr := bytecode[ptr]
		instr := binary.LittleEndian.Uint16([]byte{rawInstr, 0})

		switch instr {
		case 0x00:
			// HALT
			os.Exit(stack.Pop())
		case 0x01:
			// LD
			register := binary.LittleEndian.Uint16([]byte{bytecode[ptr + 1], 0})
			stack.Push(registers[int(register)])
			ptr += 2
		case 0x02:
			// ST
			register := binary.LittleEndian.Uint16([]byte{bytecode[ptr + 1], 0})
			registers[int(register)] = stack.Pop()
			ptr += 2
		case 0x03:
			// STC
			register := binary.LittleEndian.Uint16([]byte{bytecode[ptr + 1], 0})
			data := binary.LittleEndian.Uint64(bytecode[ptr + 2:ptr + 10])
			registers[int(register)] = int(data)
			ptr += 10
		case 0x04:
			// DUP
			stack.Dup()
			ptr += 1
		case 0x05:
			// PUSH
			data := binary.LittleEndian.Uint64(bytecode[ptr + 1:ptr + 9])
			stack.Push(int(data))
			ptr += 9
		case 0x10:
			// JMP
			ptr = int(binary.LittleEndian.Uint32(bytecode[ptr + 1:ptr + 5]))
		case 0x11:
			// JMPZ
			if stack.Pop() == 0 {
				ptr = int(binary.LittleEndian.Uint32(bytecode[ptr + 1:ptr + 5]))
			} else {
				ptr += 5
			}
		case 0x12:
			// JMPNZ
			if stack.Pop() != 0 {
				ptr = int(binary.LittleEndian.Uint32(bytecode[ptr + 1:ptr + 5]))
			} else {
				ptr += 5
			}
		case 0x13:
			// JMPP
			if stack.Pop() > 0 {
				ptr = int(binary.LittleEndian.Uint32(bytecode[ptr + 1:ptr + 5]))
			} else {
				ptr += 5
			}
		case 0x20:
			// ADD
			stack.Push(stack.Pop() + stack.Pop())
			ptr += 1
		case 0x21:
			// SUB
			stack.Push(stack.Pop() - stack.Pop())
			ptr += 1
		case 0x22:
			// MUL
			stack.Push(stack.Pop() * stack.Pop())
			ptr += 1
		case 0x23:
			// DIV
			stack.Push(stack.Pop() / stack.Pop())
			ptr += 1
		case 0x24:
			// MOD
			stack.Push(stack.Pop() % stack.Pop())
			ptr += 1
		case 0x30:
			// OUT
			fmt.Printf("%d", stack.Pop())
			ptr += 1
		case 0x31:
			// OUTC
			char := string(stack.Pop())
			fmt.Print(char)
			ptr += 1
		case 0x40:
			// MDP
			loc := stack.Pop()
			val := stack.Pop()
			memory[loc] = val
			ptr += 1
		case 0x41:
			// MLD
			loc := stack.Pop()
			stack.Push(memory[loc])
			ptr += 1
		case 0x42:
			// CPY
			loc_from := stack.Pop()
			loc_to := stack.Pop()
			memory[loc_to] = memory[loc_from]
			ptr += 1
		case 0xF0:
			// TRC - Not implemented
			ptr += 1
		case 0xF1:
			ptr += 5
			datalen := int(binary.LittleEndian.Uint32(bytecode[ptr:ptr+4]))
			ptr += 4 + datalen
		default:
			panic(fmt.Sprintf("invalid instruction: %d", instr))
		}
	}
}

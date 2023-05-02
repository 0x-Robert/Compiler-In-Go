package code

import (
	"encoding/binary"
	"fmt"
)

//바이트 코드는 명령어로 되어있다.
//명령어는 바이트 열이고,  명령어 하나는 명령코드 하나와 피연산자를 0개 이상 가진다.
//명령코드는 1ㅏ이트 정도 크기를 가진다.

type Instructions []byte //바이트 슬라이스
type Opcode byte         //바이트

const (
	//상수개념이 필요하다. 여기서 상수란 상수표현식을 말한다.
	OpConstant Opcode = iota //피연산자를 하나 가진다.
)

// op 코드의 정의는 필드 Name과 OperandWidths를 갖는다.
type Definition struct {
	Name          string //Opcode를 사람이 읽을 수 있는 이름으로 담는다.
	OperandWidths []int  //OperandWidth는 각각의 피연산자가 차지하는 바이트의 크기를 담는다.
}

var definitions = map[Opcode]*Definition{
	OpConstant: {"OpConstant", []int{2}},
}

func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}
	return def, nil
}

func Make(op Opcode, operands ...int) []byte {
	def, ok := definitions[op]
	if !ok {
		return []byte{}
	}

	instructionLen := 1
	for _, w := range def.OperandWidths {
		instructionLen += w
	}

	instruction := make([]byte, instructionLen)
	instruction[0] = byte(op)

	offset := 1
	for i, o := range operands {
		width := def.OperandWidths[i]
		switch width {
		case 2:
			binary.BigEndian.PutUint16(instruction[offset:], uint16(o))

		}
		offset += width

	}
	return instruction
}

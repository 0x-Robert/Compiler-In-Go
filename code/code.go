package code

import "fmt"

//바이트 코드는 명령어로 되어있다.
//명령어는 바이트 열이고,  명령어 하나는 명령코드 하나와 피연산자를 0개 이상 가진다.
//명령코드는 1ㅏ이트 정도 크기를 가진다.

type Instructions []byte //바이트 슬라이스
type Opcode byte         //바이트

const (
	//상수개념이 필요하다. 여기서 상수란 상수표현식을 말한다.
	OpConstant Opcode = iota //피연산자를 하나 가진다.
)

type Definition struct {
	Name          string
	OperandWidths []int
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

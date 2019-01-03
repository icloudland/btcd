package txscript

import (
	"bytes"
	"fmt"
	"strings"
)

func DisasmStringHex(buf []byte) (string, error) {
	var disbuf bytes.Buffer
	opcodes, err := parseScript(buf)
	for _, pop := range opcodes {
		str:=pop.print1()
		if len(str) > 2 {
			str = "14" + str
		}
		disbuf.WriteString(str)
	}
	if err != nil {
		disbuf.WriteString("[error]")
	}
	return disbuf.String(), err
}

func (pop *parsedOpcode) print1() string {
	opcodeName := pop.opcode.name
	if replName, ok := opcodeOnelineRepls[opcodeName]; ok {
		opcodeName = replName
	} else {
		opcodeName = fmt.Sprintf("%x", pop.opcode.value)
	}

	// Nothing more to do for non-data push opcodes.
	if pop.opcode.length == 1 {
		return opcodeName
	}

	return fmt.Sprintf("%x", pop.data)

}

func DisasmStringHexGen(asm string) string  {
	asms := strings.Split(asm, " ")
	var disbuf bytes.Buffer
	for _, item := range asms {
		isFind := false
		for _, c := range opcodeArray {
			if c.name == item {
				disbuf.WriteString(fmt.Sprintf("%x", c.value))
				isFind = true
				break
			}
		}

		if !isFind {
			disbuf.WriteString(item)
		}
	}

	return disbuf.String()

}
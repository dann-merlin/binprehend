package hex

import (
	"fmt"
	"strings"
	"unicode"
)

const separator = " "
const separator8Byte = " "
const missingbyte = "  "

func ByteToAscii(data byte) rune {
	if unicode.IsPrint(rune(data)) {
		return rune(data)
	} else {
		return '.'
	}
}

func ByteToHex(data byte) string {
	return fmt.Sprintf("%02x", data)
}

func DumpToString(data []byte) string {
	var result strings.Builder
	for i := 0; i < len(data); i += 16 {
		result.WriteString(fmt.Sprintf("%08x: ", i))  // address
		for j := 0; j < 16; j++ {
			if j == 8 {
				result.WriteString(separator8Byte)
			}
			index := i + j
			if index < len(data) {
				result.WriteString(fmt.Sprintf("%02x%s", data[index], separator))
			} else {
				result.WriteString(fmt.Sprintf("%s%s", missingbyte, separator))
			}
		}
		result.WriteString("\n")
	}

	return result.String()
}

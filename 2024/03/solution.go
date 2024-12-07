package main

import (
	"bufio"
	"fmt"
	"os"
)

type Mul struct {
	operand1 int
	operand2 int
}

func (m Mul) result() int {
	return m.operand1 * m.operand2
}

func peekOperand(fileReader *bufio.Reader, offset int, delimiter rune) (int, int, bool) {
	register := int(0)
	for i := 0; i < 3; i++ {
		offset++
		peekedByte, err := fileReader.Peek(offset)
		if err != nil {
			return 0, 0, false
		}

		fmt.Printf("Peeked byte for operand: %c\n", peekedByte[len(peekedByte)-1])
		if peekedByte[len(peekedByte)-1] >= '0' && peekedByte[len(peekedByte)-1] <= '9' {
			register = register*10 + int(peekedByte[len(peekedByte)-1]-'0')
			continue
		}
		if peekedByte[len(peekedByte)-1] == byte(delimiter) {
			return register, offset - 1, true
		}
		return 0, 0, false
	}

	return register, offset, true
}

func searchMul(fileScanner *bufio.Reader) (Mul, int, bool) {
	prefixBytes := []rune{'m', 'u', 'l', '('}

	offset := len(prefixBytes)
	peekedBytes, err := fileScanner.Peek(offset)
	if err != nil {
		return Mul{}, 0, false
	}
	for i := 0; i < len(prefixBytes); i++ {
		if peekedBytes[i] != byte(prefixBytes[i]) {
			return Mul{}, 0, false
		}
	}

	fmt.Println("Found mul command")

	operand1, offset, ok := peekOperand(fileScanner, offset, ',')
	if !ok {
		return Mul{}, 0, false
	}

	fmt.Printf("Operand 1: %d\n", operand1)
	offset++
	peekedBytes, err = fileScanner.Peek(offset)
	if err != nil || peekedBytes[len(peekedBytes)-1] != ',' {
		return Mul{}, 0, false
	}

	operand2, offset, ok := peekOperand(fileScanner, offset, ')')
	if !ok {
		return Mul{}, 0, false
	}

	fmt.Printf("Operand 2: %d\n", operand2)

	offset++
	peekedBytes, err = fileScanner.Peek(offset)
	if err != nil || peekedBytes[len(peekedBytes)-1] != ')' {
		return Mul{}, 0, false
	}

	return Mul{operand1, operand2}, offset, true
}

func parseCommandString(fileScanner *bufio.Reader, command string) (bool, int) {
	expectedBytes := []rune(command)
	offset := len(expectedBytes)
	peekedBytes, err := fileScanner.Peek(offset)
	if err != nil {
		return false, 0
	}
	for i := 0; i < len(expectedBytes); i++ {
		if peekedBytes[i] != byte(expectedBytes[i]) {
			return false, 0
		}
	}

	return true, offset
}

func parseDo(fileScanner *bufio.Reader) bool {
	expectedBytes := []rune{'d', 'o', '(', ')'}
	offset := len(expectedBytes)
	peekedBytes, err := fileScanner.Peek(offset)
	if err != nil {
		return false
	}
	for i := 0; i < len(expectedBytes); i++ {
		if peekedBytes[i] != byte(expectedBytes[i]) {
			return false
		}
	}

	return true
}

func parseDont(fileScanner *bufio.Reader) (bool, int) {
	expectedBytes := []rune{'d', 'o', 'n', 't', '(', ')'}
	offset := len(expectedBytes)
	peekedBytes, err := fileScanner.Peek(offset)
	if err != nil {
		return false, 0
	}
	for i := 0; i < len(expectedBytes); i++ {
		if peekedBytes[i] != byte(expectedBytes[i]) {
			return false, 0
		}
	}

	return true, offset
}

func parseCommands(filename string) []Mul {
	commands := make([]Mul, 0)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file")
		fmt.Println(err)
		return nil
	}

	fileReader := bufio.NewReader(file)

	enabled := true
	for {
		if enabled {
			dont, offset := parseCommandString(fileReader, "don't()")
			if dont {
				fmt.Println("Found dont() command")
				enabled = false
				fmt.Printf("Discarding %d bytes\n", offset)
				fileReader.Discard(offset)
				continue
			}
		}

		if !enabled {
			do, offset := parseCommandString(fileReader, "do()")
			if do {
				fmt.Println("Found do() command")
				enabled = true
				fmt.Printf("Discarding %d bytes\n", offset)
				fileReader.Discard(offset)
				continue
			} else {
				fmt.Println("Skipping byte")
				fileReader.Discard(1)
				continue
			}
		}

		mul, offset, ok := searchMul(fileReader)
		if ok {
			fmt.Printf("Found mul(%d, %d) command\n", mul.operand1, mul.operand2)
			commands = append(commands, mul)
			fmt.Printf("Discarding %d bytes\n", offset)
			fileReader.Discard(offset)
		} else {
			peek, _, err := fileReader.ReadRune()
			if err != nil {
				fmt.Println("Error reading rune")
				fmt.Println(err)
				break
			}

			fmt.Printf("Found unknown command: %c\n", peek)
		}
	}

	return commands
}

func main() {
	commands := parseCommands("input.data")
	fmt.Println("Commands:")
	for _, command := range commands {
		fmt.Printf("mul(%d, %d)\n", command.operand1, command.operand2)
	}
	fmt.Println("Result:")
	result := 0
	for _, command := range commands {
		result += command.result()
	}
	fmt.Println(result)
}

package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed test.txt
var test string

//go:embed input.txt
var input string

type CPUInfo struct {
	registerA int
	registerB int
	registerC int
	ip        int // instruction pointer
	output    []int
	program   []int
}

func main() {
	input := ParseInput(input)
	out := input.Run()

	strOut := make([]string, len(out))
	for i, v := range out {
		strOut[i] = strconv.Itoa(v)
	}
	fmt.Println(strings.Join(strOut, ","))
}

func (c *CPUInfo) Run() []int {
	for c.ip < len(c.program)-1 {
		op := c.program[c.ip]
		opr := c.program[c.ip+1]
		if c.Execute(op, opr) {
			c.ip += 2
		}
	}

	return c.output
}

func (c *CPUInfo) Execute(op, opr int) bool {
	comboOpr := c.GetComboOperator(opr)
	fmt.Printf("%+v\n", *c)
	switch op {
	case 0:
		fmt.Println("adv", comboOpr)
		num := c.registerA
		val := float64(num) / (math.Pow(2, float64(comboOpr)))
		c.registerA = int(math.Trunc(val))
	case 1:
		fmt.Println("bxl", opr)
		c.registerB = c.registerB ^ opr
	case 2:
		fmt.Println("bst", comboOpr)
		c.registerB = comboOpr % 8
	case 3:
		fmt.Println("jnz", opr)
		if c.registerA == 0 {
			return true
		}
		c.ip = opr
		return false
	case 4:
		fmt.Println("bxc")
		c.registerB = c.registerB ^ c.registerC
	case 5:
		fmt.Println("out", comboOpr)
		c.output = append(c.output, comboOpr%8)
	case 6:
		fmt.Println("bdv", comboOpr)
		num := c.registerA
		val := float64(num) / (math.Pow(2, float64(comboOpr)))
		c.registerB = int(math.Trunc(val))
	case 7:
		fmt.Println("cdv", opr, comboOpr)
		num := c.registerA
		val := float64(num) / (math.Pow(2, float64(comboOpr)))
		c.registerC = int(math.Trunc(val))

	}

	return true
}

func (c *CPUInfo) GetComboOperator(opr int) int {
	switch opr {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 3
	case 4:
		return c.registerA
	case 5:
		return c.registerB
	case 6:
		return c.registerC
	default:
		return -1
	}
}

func ParseInput(in string) CPUInfo {
	fields := strings.Split(in, "\n")

	var registerA, registerB, registerC int
	fmt.Sscanf(fields[0], "Register A: %d", &registerA)
	fmt.Sscanf(fields[1], "Register B: %d", &registerB)
	fmt.Sscanf(fields[2], "Register C: %d", &registerC)

	var program []int
	programString := strings.TrimPrefix(fields[4], "Program: ")
	programValues := strings.Split(programString, ",")

	for _, v := range programValues {
		e, _ := strconv.Atoi(v)
		program = append(program, e)
	}

	return CPUInfo{registerA: registerA, registerB: registerB, registerC: registerC, program: program, ip: 0}

}

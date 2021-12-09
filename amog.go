package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

const (
	INST_IMPOSTER = iota
	INST_TEAMMATE
	INST_YELL
	INST_SUS
	INST_TRUST
	INST_EJECT
)

type Inst struct {
	action int
	target string
}

type Program struct {
	imposters map[string]int
	teammates map[string]int

	source []Inst
	ip int
}

func parseLine(program *Program, line string, lineNr int) {
	line = strings.Trim(strings.Split(line, "#")[0], " ")

	if len(line) == 0 {
		return
	}

	words := strings.Split(line, " ")

	if len(words) == 1 {
		fmt.Printf("[line %d] no target provided, kinda sus, ngl\n", lineNr)
		os.Exit(1)
	}

	inst := Inst{
		target: words[1],
	}

	switch words[0] {
	case "imposter":
		inst.action = INST_IMPOSTER
	case "teammate":
		inst.action = INST_TEAMMATE
	case "yell":
		inst.action = INST_YELL
	case "sus":
		inst.action = INST_SUS
	case "trust":
		inst.action = INST_TRUST
	case "eject":
		inst.action = INST_EJECT
	default:
		fmt.Printf("[line %d] unknown action `%s`. kinda sus, ngl\n", lineNr, inst.action)
		os.Exit(1)
	}

	program.source = append(program.source, inst)
}

func runProgram(program Program) {
	for program.ip < len(program.source) {
		inst := program.source[program.ip]

		switch (inst.action) {
		case INST_IMPOSTER:
			program.imposters[inst.target] = 0
		case INST_TEAMMATE:
			program.teammates[inst.target] = 0
		case INST_EJECT:
			if _, ok := program.imposters[inst.target]; ok {
				delete(program.teammates, inst.target);
				fmt.Printf("ejected imposter %s\n", inst.target)
			} else if _, ok := program.teammates[inst.target]; ok {
				delete(program.teammates, inst.target);
				fmt.Printf("ejected teammate %s\n", inst.target)
			} else {
				fmt.Printf("unknown player %s. kinda sus, ngl\n", inst.target)
				os.Exit(1)
			}

		case INST_TRUST:
			if _, ok := program.imposters[inst.target]; ok {
				program.imposters[inst.target]++
			} else {
				fmt.Printf("unknown imposter %s. kinda sus, ngl\n", inst.target)
				os.Exit(1)
			}
			fmt.Printf("%s is not sus\n", inst.target)

		case INST_SUS:
			if _, ok := program.imposters[inst.target]; ok {
				program.imposters[inst.target]--
			} else {
				fmt.Printf("unknown imposter %s. kinda sus, ngl\n", inst.target)
				os.Exit(1)
			}
			fmt.Printf("%s is kinda sus, ngl\n", inst.target)

		case INST_YELL:
			if _, ok := program.imposters[inst.target]; ok {
				fmt.Println(program.imposters[inst.target])
			} else if _, ok := program.teammates[inst.target]; ok {
				fmt.Println(program.teammates[inst.target])
			} else {
				fmt.Printf("unknown player %s. kinda sus, ngl\n", inst.target)
				os.Exit(1)
			}

		default:
			panic("not implemented")
		}

		program.ip++
	}
}

func main() {
	if len(os.Args) == 2 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Printf("could not load `%s`. kinda sus, ngl\n", os.Args[1])
			os.Exit(1)
		}

		program := Program{}
		program.imposters = make(map[string]int)
		program.teammates = make(map[string]int)

		scanner := bufio.NewScanner(file)
		for lineNr := 1; scanner.Scan(); lineNr++ {
			parseLine(&program, scanner.Text(), lineNr)
		}

		if err := scanner.Err(); err != nil {
			os.Exit(1)
		}

		runProgram(program)
	} else {
		fmt.Println("file not provided. kinda sus, ngl")
		os.Exit(1)
	}
}

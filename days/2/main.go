package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var extended = flag.Bool("x", false, "solve for problem extension")

// Instruction represents a submarine instruction
type Instruction struct {
	Dir string
	Val int
}

// Simulator is a submarine simulator
type Simulator struct {
	X, Y, Aim int
}

func (sim *Simulator) Execute(ins ...Instruction) {
	for _, instruction := range ins {
		switch strings.ToLower(instruction.Dir) {
		case "up":
			sim.Y -= instruction.Val
		case "down":
			sim.Y += instruction.Val
		case "forward":
			sim.X += instruction.Val
		}
	}
}

type SimulatorV2 struct {
	X, Y, Aim int
}

func (sim *SimulatorV2) Execute(ins ...Instruction) {
	for _, instruction := range ins {
		switch strings.ToLower(instruction.Dir) {
		case "up":
			sim.Aim -= instruction.Val
		case "down":
			sim.Aim += instruction.Val
		case "forward":
			sim.X += instruction.Val
			sim.Y += sim.Aim * instruction.Val
		}
	}
}

func parseInstructions(in io.Reader) (ins []Instruction, err error) {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		ln := scanner.Text()
		items := strings.Fields(ln)
		if len(items) != 2 {
			return ins, fmt.Errorf("syntax error: %q is invalid input", ln)
		}
		v, err := strconv.Atoi(items[1])
		if err != nil {
			return ins, fmt.Errorf("syntax error: %q: expected integer", ln)
		}
		ins = append(ins, Instruction{
			Dir: items[0],
			Val: v,
		})
	}
	return ins, scanner.Err()
}

func main() {
	flag.Parse()
	instructions, err := parseInstructions(os.Stdin)
	if err != nil {
		panic(err)
	}

	// todo
	if *extended {
		sim := new(SimulatorV2)
		sim.Execute(instructions...)
		fmt.Println(sim.X * sim.Y)
		return
	}
	sim := new(Simulator)
	sim.Execute(instructions...)
	fmt.Println(sim.X * sim.Y)
}

package main

import (
	"fmt"
	"github.com/chzyer/readline"
	"io"
	"os"
)

type I1 interface {
	X() int
	Y() int
}

type I2 interface {
	X() int
}

type T1 struct {
	I1
	T2
}

func (o T1) X() int {
	return 1
}

type T2 struct {
	I1
}

type T3 struct {
}

func (o T3) X() int {
	return 2
}

func (o T3) Y() int {
	return 3
}

type t4 struct {
	I1
}

func (o t4) X() int {
	return 4
}

func (o t4) Y() int {
	return 5
}

type t5 struct {
	I2
}

func (o t5) X() int {
	return 6
}

// if this is not omitted, t5 implements I1:
// func (o t5) Y() int {
// 	return 7
// }

func main() {
	rl, err := readline.New("")
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	defer rl.Close()

loop:
	for {
		line, err := rl.Readline()
		switch err {
		case readline.ErrInterrupt:
			if len(line) == 0 {
				break
			} else {
				continue
			}
		case io.EOF:
			break loop
		}

		var x interface{}
		switch line {
		case "T1":
			x = T1{}
		case "*T1":
			x = &T1{}
		case "T2":
			x = T2{}
		case "*T2":
			x = &T2{}
		case "T3":
			x = T3{}
		case "*T3":
			x = &T3{}
		case "t4":
			x = t4{}
		case "*t4":
			x = &t4{}
		case "t5":
			x = t5{}
		case "*t5":
			x = &t5{}
		case "int":
			x = 3
		case "*int":
			x = &make([]int, 1)[0]
		case "string":
			x = "hi there"
		case "*string":
			x = &make([]string, 1)[0]
		case "[5]string":
			x = make([]string, 5)
		case "[5][5]string":
			x = make([][]string, 5, 5)
		case "nil":
			x = nil
		default:
			fmt.Fprintf(os.Stderr, "unrecognized command %s\n", line)
			continue
		}
		fmt.Printf("%s (%v) is a %s\n", line, x, WhatIsThis(x))
	}
}

// switch on type is implemented as if each case is compared in the
// order given.  So, all concrete types must precede all interface
// (abstract) types.  Further, more-specific interface types must
// precede less-specific ones, where "specific" refers to the
// methods/receivers specified as being implemented by any concrete
// type implementing the interface type.
//
// So, the I2 case must follow the I1 case, below, or t4 will match I2.
//
// Note that if there's an actual implementation of Y() on t5 (func (o
// t5) Y int {}), the t5 case matches I1, even though t5 is defined as
// implementing only I2.
//
// That is, a type implements an interface either if it is declared to
// do so (even if there are missing implementations) or it actually
// does so (due to present implementations).
//
// Also, pointer/references to interface types are evidently ignored
// as cases, in that they never seem to match. Might as well omit them.
//
// (The t4 and t5 types are omitted below, as this example is designed
// to illustrate how to handle switching on types in the presence of
// private types implementing public interfaces.)
func WhatIsThis(x interface{}) (is string) {
	switch x.(type) {
	case T1:
		is = "T1"
	case *T1:
		is = "*T1"
	case T2:
		is = "T2"
	case *T2:
		is = "*T2"
	case T3:
		is = "T3"
	case *T3:
		is = "*T3"
	case *I1:
		is = "*I1"
	case I1:
		is = "I1"
	case *I2:
		is = "*I2"
	case I2:
		is = "I2"
	case [5][5]string:
		is = "[5][5]string"
	case [][5]string:
		is = "[][5]string"
	case [5][]string:
		is = "[5][]string"
	case [][]string:
		is = "[][]string"
	case [5]string:
		is = "[5]string"
	case []string:
		is = "[]string"
	case *interface{}:
		is = "*interface{}"
	case interface{}:
		is = "interface{}"
	default:
		is = "default"
	}
	return
}

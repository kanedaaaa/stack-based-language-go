package main 

import (
  "fmt"
  "github.com/kanedaaaa/gostuff/src/core"
  "github.com/kanedaaaa/gostuff/src/extra"
)

var STACK []int
var CAP = 69

func simulate(program [][]int) {
  extra.Require(core.COUNT == 6, "CORE:001", "")
  for _, value := range program { 
    switch value[0] {
    case core.PUSH:
      STACK = append(STACK, value[1])
    case core.PLUS:
      extra.Require(len(STACK) > 1, "CORE:002", " Help: core.PLUS")
      a, b := extra.GetLastTwo(&STACK)

      STACK = append(STACK, a + b)
    case core.MINUS:
      extra.Require(len(STACK) > 1, "CORE:002", " Help: core.MINUS")
      a, b := extra.GetLastTwo(&STACK)

      STACK = append(STACK, a - b)

    case core.MUL:
      extra.Require(len(STACK) > 1, "CORE:002", " Help: core.MUL")
      a, b := extra.GetLastTwo(&STACK)

      STACK = append(STACK, a * b)

    case core.DIV:
      extra.Require(len(STACK) > 1, "CORE:002", " Help: core.DIV")
      a, b := extra.GetLastTwo(&STACK)
      extra.Require(a != 0 && b != 0, "CORE:003", " Help core.DIV duh...")

      STACK = append(STACK, a / b)
    case core.DUMP:
      fmt.Println(STACK)
    }
  }
}

func main() {
 var program = [][]int{core.Push(20), core.Push(3), core.Mul(), core.Dump()}
 
 simulate(program)
}


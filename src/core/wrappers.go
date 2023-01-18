package core

func Push(x int) []int {
  arr := [2]int{PUSH, x}
  return arr[:]
}

func Plus() []int {
  arr := [2]int{PLUS, 0}
  return arr[:]
}

func Minus() []int {
  arr := [2]int{MINUS, 0}
  return arr[:]
}

func Mul() []int {
  arr := [2]int{MUL, 0}
  return arr[:]
}

func Div() []int {
  arr := [2]int{DIV, 0}
  return arr[:]
}

func Dump() []int {
  arr := [2]int{DUMP, 0}
  return arr[:]
}
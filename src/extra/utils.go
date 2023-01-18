package extra

func Require(condition bool, errCode string, extra string) {
	err := ToThrow(errCode)
	if (!condition) {
	  panic(err + extra)
	}
}

func Pop(arr *[]int) int {
	ret := (*arr)[len(*arr) - 1]
	*arr = (*arr)[:len(*arr) - 1]
	
	return ret
}

func GetLastTwo(arr *[]int) (int, int) {
	a := Pop(arr)
	b := Pop(arr)

	return a, b
}


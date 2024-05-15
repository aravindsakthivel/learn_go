package main

import "fmt"

func infVariadic(inp ...interface{}) {
	for i := 0; i < len(inp); i++ {
		fmt.Println(inp[i])
	}
}

func anyVar(inp ...any) {
	for i := 0; i < len(inp); i++ {
		fmt.Println(inp[i])
	}
}

func intVar(inp ...int) {
	for i := 0; i < len(inp); i++ {
		fmt.Println(inp[i])
	}

}

func main() {
	infVariadic(1, 2, 3, 4, 5, "Hello", "World")
	intVar(1, 2, 3, 4, 5)
	anyVar(1, 2, 3, 4, 5, "Hello", "World")
}

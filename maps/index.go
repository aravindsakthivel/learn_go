package main

func main() {
	var mapX map[string]int = map[string]int{}
	mapX["one"] = 1
	mapX["two"] = 2

	println(mapX["one"], mapX["two"])
}

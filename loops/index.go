package main

func forLoop() {
	for i := 0; i < 10; i++ {
		println(i)
	}
}

func whileLoop() {
	i := 10
	for i < 20 {
		i++
		println(i)
	}

	rm := 4 % 3
	const div float32 = 4 / 3
	println(rm, div)

	tl := 12 % 4
	println(tl)
}

func main() {
	forLoop()
	whileLoop()
}

package main

import "fmt"

func main() {
	/* integers */
	var zz int = 0
	x := 10
	var z int = x
	var y int8 = int8(x + 1)
	const n = 5001
	const c int = 5001
	var e float32 = 6
	var f float32 = e

	fmt.Print(zz, z, z, y, n, c, e, f)
}

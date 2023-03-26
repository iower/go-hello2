package main

import "fmt"

func main() {
	var a int8 = -1
	var b uint8 = 1
	var c byte = 11

	var d int16 = -2
	var e uint16 = 2

	var f int32 = -3
	var g uint32 = 3
	var h rune = -33

	var i int64 = -4
	var j uint64 = 4

	var k int = -5
	var l uint = 5

	var m float32 = 1.1e10
	var n float64 = 1.1e100

	var o complex64 = 1 + 2i  // both float32
	var p complex128 = 3 + 4i // both float64

	var q bool
	var r string = "str1\nstr2\rstr3\tstr4\"str5\\str6"

	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r)
}

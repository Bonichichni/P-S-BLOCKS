package main

import (
	"fmt"

	"github.com/Bonichichni/P-S-BLOCKS.git/transformation"
)

func main() {
	var test = [8]int{1, 0, 0, 1, 1, 1, 0, 1}
	res := transformation.TransformPBlock(test)
	fmt.Println(res)
	res = transformation.InverseTransformPBlock(res)
	fmt.Println(res)
	var testS = [8]int{1, 0, 0, 1, 1, 1, 0, 1}
	testSS := transformation.TransformSBlock(testS, transformation.SBox)
	invTable := transformation.InverseTable(transformation.SBox)
	fmt.Println(transformation.TransformSBlock(testSS, invTable))
}

package main

import "fmt"

func main() {
	mapA := make(map[string]string)
	mapA["chenjie"] = "super"
	mapB := mapA
	delete(mapB, "chenjie")
	mapA["chenjie1"] = "super1"
	fmt.Println(mapA, mapB)
}

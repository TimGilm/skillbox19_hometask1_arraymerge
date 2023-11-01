/*
	Задание 1. Слияние отсортированных массивов

Что нужно сделать Напишите функцию, которая производит слияние двух отсортированных массивов
длиной четыре и пять в один массив длиной девять.
package main
*/
package main

import "fmt"

const firstArrSize = 4
const secondArrSize = 5
const thirdArrSize = 9

var arrayOne = [firstArrSize]int{1, 3, 6, 9}
var arrayTwo = [secondArrSize]int{2, 4, 5, 7, 8}

func arrayMerge(A [firstArrSize]int, B [secondArrSize]int) [thirdArrSize]int {
	var result [thirdArrSize]int
	a := 0
	b := 0
	for i := 0; i < thirdArrSize; i++ {
		if a >= len(A) {
			result[i] = B[b]
			b++
			continue
		} else if b >= len(B) {
			result[i] = A[a]
			a++
			continue
		}
		if A[a] > B[b] {
			result[i] = B[b]
			b++
		} else {
			result[i] = A[a]
			a++
		}
	}
	return result
}

func main() {
	fmt.Print(arrayMerge(arrayOne, arrayTwo))
}

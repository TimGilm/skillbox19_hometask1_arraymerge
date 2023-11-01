/*
	Задание 1. Слияние отсортированных массивов

Что нужно сделать Напишите функцию, которая производит слияние двух отсортированных массивов длиной четыре и пять в один массив длиной девять.
package main
*/
package main

import "fmt"

var arrayOne = [4]int{1, 3, 6, 9}
var arrayTwo = [5]int{2, 4, 5, 7, 8}
var arrayUnited [9]int

func arrayMerge([4]int, [5]int) [9]int {
	n := 0
	for i := 0; i < len(arrayOne); i++ {
		if arrayOne[i] < arrayTwo[i] {
			arrayUnited[n] = arrayOne[i]
			n += 1
			arrayUnited[n] = arrayTwo[i]
			n += 1
		} else {
			arrayUnited[n] = arrayTwo[i]
			n += 1
			arrayUnited[n] = arrayOne[i]
			n += 1
		}
	}
	m := len(arrayTwo) - 1
	arrayUnited[n] = arrayTwo[m]
	if arrayUnited[n] < arrayUnited[n-1] {
		arrayUnited[n], arrayUnited[n-1] = arrayUnited[n-1], arrayUnited[n]
	}
	return arrayUnited
}

func main() {

	arrayMerge(arrayOne, arrayTwo)
	fmt.Print(arrayUnited)

}

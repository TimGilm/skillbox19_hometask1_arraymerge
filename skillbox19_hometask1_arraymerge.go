/*
	Задание 1. Слияние отсортированных массивов

Что нужно сделать Напишите функцию, которая производит слияние двух отсортированных массивов длиной четыре и пять в один массив длиной девять.
package main
*/
package main

import "fmt"

var arrayOne = [4]int{1, 2, 3, 4}
var arrayTwo = [5]int{5, 6, 7, 8, 9}
var arrayUnited [9]int

func main() {
	for i := 0; i < len(arrayUnited); i++ {
		if i < len(arrayOne) {
			arrayUnited[i] = arrayOne[i]
		} else {
			arrayUnited[i] = arrayTwo[i-len(arrayOne)]
		}
	}
	fmt.Print(arrayUnited)
}

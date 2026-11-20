package main

import "fmt"

func main() {
	marks := [5]int{5, 4, 3, 5, 4}
	// получите срез marks[:]
	//Выведите исходный массив, полученный срез, его длину и ёмкость.
	marksSlices := marks[1:4]
	marksSlices = append(marksSlices, 2, 5)

	fmt.Printf("Marks: %v\n", marksSlices)
}

/*
Мини-шпаргалка от дипсика
замучала я его вопросами...
s := []int{1, 2, 3}
s = append(s, 4)              // [1 2 3 4]
s = append(s, 5, 6)           // несколько за раз
fmt.Println(len(s), s[0], s[len(s)-1])
sub := s[1:3]                 // [2 3], но ссылается на тот же массив
// независимая копия
c := make([]int, len(s))
copy(c, s)
// удалить элемент i
s = append(s[:i], s[i+1:]...)
// вставить в позицию i
s = append(s[:i], append([]int{x}, s[i:]...)...)
*/

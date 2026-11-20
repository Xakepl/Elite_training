package main

import "fmt"

func main() {
	fmt.Print("============================================ \n")
	fmt.Print("TASK 1\n")
	fmt.Print("============================================\n")
	//Базовые вычисления
	var name string = "Yaroslava"
	var age int = 19
	avegare := 0.75
	active := true
	const courseName = "Бэкенд-разработка на языке Go (базовый уровень)"

	fmt.Printf("Name: %v type: %T \n", name, name)
	fmt.Printf("age: %v type: %T \n", age, age)
	fmt.Printf("average: %v type: %T \n", avegare, avegare)
	fmt.Printf("active: %v type: %T \n", active, active)
	fmt.Printf("courseName: %v type: %T \n", courseName, courseName)
	fmt.Print("============================================ \n")
	fmt.Print("TASK 2\n")
	fmt.Print("============================================ \n")
	//калькуряторв
	var a int = 6
	var b int = 7
	var operation string = "+"

	switch operation {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: деление на ноль")
		} else {
			fmt.Println(float64(a) / float64(b)) // явное преобразование
		}
	case "%":
		if b == 0 {
			fmt.Println("Ошибка: деление на ноль")
		} else {
			fmt.Println(a % b)
		}
	default:
		fmt.Println("Неизвестная операция:", operation)

	}
	fmt.Print("============================================ \n")
	fmt.Print("TASK 3\n")
	fmt.Print("============================================ \n")
	/*Задайте n типа int. Если n находится вне диапазона от 1 до 100 включительно, выведите сообщение и пропустите расчёт.
	Иначе циклом for найдите сумму чисел от 1 до n и количество чётных чисел в этом диапазоне.
	Проверку диапазона запишите с логической операцией.*/
	var n int = 25
	if n < 1 || n > 100 {
		fmt.Printf("ERROR, n must be э [1, 100]")
	} else {
		result := 0
		k := 0
		for i := 0; i <= n; i++ {
			result += i
			if i%2 == 0 {
				k++
			}
		}

		fmt.Printf("result: %v evens: %v\n", result, k)
	}
	fmt.Print("============================================ \n")
	fmt.Print("TASK 4\n")
	fmt.Print("============================================ \n")
	/*
		Управление циклом. Отдельным циклом выведите чётные числа от 1 до n, но не более первых трёх.
		Нечётные числа пропускайте через continue, а после третьего выведенного числа завершайте цикл через break.
		Для корректного n меньше 6 выведите столько чётных чисел, сколько есть.
	*/
	N := 10
	count := 0
	for i := 0; i <= N; i++ {
		if i%2 != 0 {
			continue
		}
		count++
		if count == 3 {
			break
		}
	}
	fmt.Printf("Count: %v \n", count)
}

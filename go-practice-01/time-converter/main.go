package main

import "fmt"

func main() {
	fmt.Print("============================================ \n")
	fmt.Print("TIME-CONVERTER\n")
	fmt.Print("============================================\n")
	//Базовые вычисления
	seconds := 0
	fmt.Scan(&seconds)
	fmt.Printf("%v:%v:%v", seconds/3600, (seconds%3600)/60, seconds%60)
}

package main

import "fmt"

const (
	tF1 float64 = 1.8
	tF2 float64 = 32
)

func celsiusToFah(celsius float64) float64 {
	return celsius*tF1 + tF2
}

func main() {
	var tC float64
	fmt.Print("Введите температуру в градусах Цельсия: ")
	fmt.Scan(&tC)
	fah := celsiusToFah(tC)
	fmt.Printf("Температура в градусах Фаренгейта: %.2f\n", fah)
	return
}

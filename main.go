package main

import (
	"fmt"
	"strconv"
)

func LuhnAlgorithm(cardNumber string) ([]int, error) {
	var digits []int
	var credit_card []int

	for _, digit := range cardNumber {
		digito, err := strconv.Atoi(string(digit))
		if err != nil {
			fmt.Println("Error al convertir el digito: ", err)
			return nil, err
		}
		digits = append(digits, digito)
	}

	for i := range digits {
		if (i+1)%2 == 0 {
			result1 := digits[i] * 1
			if result1 > 9 {
				result1 = (result1 % 10) + (result1 / 10)
			}
			credit_card = append(credit_card, result1)
		} else {
			result2 := digits[i] * 2
			if result2 > 9 {
				result2 = (result2 % 10) + (result2 / 10)
			}
			credit_card = append(credit_card, result2)
		}
	}
	return credit_card, nil
}

func main() {
	fmt.Println(LuhnAlgorithm("3379513561108795"))
}

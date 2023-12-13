package main

import (
	"fmt"
	"strconv"
)

func LuhnAlgorithm(cardNumber string) bool {
	var digits []int
	var credit_card []int
	var sum int = 0

	if isLounge(cardNumber) != 0 {
		for _, digit := range cardNumber {
			digito, err := strconv.Atoi(string(digit))
			if err != nil {
				fmt.Println("Error al convertir el digito: ", err)
				return false
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

		for i := 0; i < len(credit_card); i++ {
			sum += credit_card[i]
		}

		return sum%10 == 0
	} else {
		return false
	}
}

func isLounge(cardNumber string) int {

	for {
		if len(cardNumber) == 16 {
			return 1
		} else {
			fmt.Println("Numero de tarjeta invalido.")
			return 0
		}
	}
}

func main() {
	var cardNumber string

	fmt.Print("Ingresa el numero de tu tarjeta bancaria: ")
	fmt.Scanln(&cardNumber)

	if LuhnAlgorithm(cardNumber) {
		fmt.Println("La tarjeta de credito es válida.")
	} else {
		fmt.Println("Tarjeta de credito inválida.")
	}
}

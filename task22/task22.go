package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

/*
	Большие числа и операции

	Разработать программу, которая перемножает,
	делит, складывает, вычитает две числовых
	переменных a, b, значения которых > 2^20 (больше 1 миллион).
*/

func main() {
	scanner := bufio.NewReader(os.Stdin)

	fmt.Print("Enter (sep. by a space) a, b:\n")
	var aStr, bStr string
	_, err := fmt.Fscanf(scanner, "%s %s\n", &aStr, &bStr)
	if err != nil {
		fmt.Printf("Fscan error: %v", err)
	}

	a := new(big.Int)
	b := new(big.Int)
	if _, ok := a.SetString(aStr, 10); !ok {
		fmt.Println("неверный формат числа а")
		return
	}
	if _, ok := b.SetString(bStr, 10); !ok {
		fmt.Println("неверный формат числа а")
		return
	}

	sum := new(big.Int).Add(a, b)
	diff := new(big.Int).Sub(a, b)
	prod := new(big.Int).Mul(a, b)
	div := new(big.Int)
	if b.Sign() == 0 {
		fmt.Println("делить на 0 нежелательно, здесь не вышмат")
	} else {
		div.Div(a, b)
	}
	fmt.Println("\nРезультаты:")
	fmt.Printf("a = %s\nb = %s\n", aStr, bStr)
	fmt.Printf("a + b = %s\n", sum.String())
	fmt.Printf("a - b = %s\n", diff.String())
	fmt.Printf("a * b = %s\n", prod.String())
	if b.Sign() != 0 {
		fmt.Printf("a/b=%s\n", div.String())
	}

}

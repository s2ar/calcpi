package main

import (
	"fmt"
)

func main() {
	/*
		demolist := [7]int{1, 6, 9, 9, 10, 1, 10}
		pilist := make([]int, 0, 6)

		fmt.Println("pilist", pilist)

		helddigits := 0

		for i, q := range demolist {
			fmt.Println(q)

			if q == 9 {
				helddigits++
			} else if q == 10 {
				q = 0
				for k := 1; k <= helddigits; k++ {
					replaced := pilist[i-k]
					fmt.Println("replaced", replaced)
					if replaced == 9 {
						replaced = 0
					} else {
						replaced++
					}
					pilist[i-k] = replaced
					fmt.Println("pilist i k", pilist)
				}
				helddigits = 1
			} else {
				helddigits = 1
			}
			pilist = append(pilist, q)

		}

		fmt.Println(pilist)
		os.Exit(1)*/

	n := 15
	count := n * 10 / 3
	helddigits := 0
	pilist := make([]int, 0, count)
	volume := make([]int, count, count)
	// заполняем двойками
	for i := 0; i < len(volume); i++ {
		volume[i] = 2
	}

	//unconfirmeFigure := 0

	for i := 0; i < n; i++ {
		carryover := 0
		sum := 0

		for j := count - 1; j >= 0; j-- {
			volume[j] *= 10
			sum = volume[j] + carryover
			quotient := int(sum / (j*2 + 1)) // результат деления суммы на знаменатель
			volume[j] = sum % (j*2 + 1)      // остаток от деления суммы на знаменатель
			carryover = quotient * j         // j - числитель
		}
		volume[0] = sum % 10
		q := int(sum / 10) // новая цифра числа Пи
		fmt.Println("q ", q)

		// регулировка недействительных цифр
		if q == 9 {
			helddigits++
		} else if q == 10 {
			q = 0
			for k := 1; k <= helddigits; k++ {
				replaced := pilist[i-k]
				fmt.Println("replaced", replaced)
				if replaced == 9 {
					replaced = 0
				} else {
					replaced++
				}
				pilist[i-k] = replaced
				fmt.Println("pilist i k", pilist)
			}
			helddigits = 1
		} else {
			helddigits = 1
		}
		pilist = append(pilist, q)

	}
	fmt.Println(volume)
	fmt.Println("pilist", pilist)
}

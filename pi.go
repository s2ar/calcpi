package calcpi

import (
	"fmt"
	"strconv"
	"strings"
)

// To находит число Pi до заданой точности
func To(n int) (string, error) {

	if n < 2 {
		return "", fmt.Errorf("number n must be greater than one")
	}

	count := n * 10 / 3
	helddigits := 0
	pilist := make([]string, 0, count)
	volume := make([]int, count, count)
	// заполняем двойками
	for i := 0; i < len(volume); i++ {
		volume[i] = 2
	}

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

		// регулировка недействительных цифр
		if q == 9 {
			helddigits++
		} else if q == 10 {
			q = 0
			for k := 1; k <= helddigits; k++ {
				replaced := pilist[i-k]

				if replaced == "9" {
					replaced = strconv.Itoa(0)
				} else {
					replacedint, err := strconv.Atoi(replaced)
					if err != nil {
						return "", err
					}
					replacedint++
					replaced = strconv.Itoa(replacedint)
				}
				pilist[i-k] = replaced
			}
			helddigits = 1
		} else {
			helddigits = 1
		}
		pilist = append(pilist, strconv.Itoa(q))
	}

	result := pilist[0] + "," + strings.Join(pilist[1:], "")
	return result, nil

}

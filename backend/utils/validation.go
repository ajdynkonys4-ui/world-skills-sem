package utils

import (
	"strings"
)

func Validation(iin string) string {
	//Кусок кода для проверки валидности входящих данных
	strValidate := strings.TrimSpace(iin)
	for _, v := range iin {
		if v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z' {
			return "Ошибка"
		}
	}
	if len(iin) > 12 || strValidate == "" {
		return "Ошибка"
	}
	if len(iin) < 12 {
		return "Ошибка"
	}
	//Функция проверки данных
	savedNumbers := []int{}
	for _, v := range iin {
		savedNumbers = append(savedNumbers, int(v-'0'))
	}
	firstCount := []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	secondCount := []int{2, 1, 11, 10, 9, 8, 7, 6, 5, 4, 3}
	temp := 0
	for i, v := range firstCount {
		temp += v * savedNumbers[i]
	}
	if temp%11 <= 9 && temp%11 >= 0 {
		return "Валиден"
	}
	if temp%11 == 10 {
		temp = 0
		for i, v := range secondCount {
			temp += savedNumbers[i] * v
		}

		if temp%11 <= 9 && temp%11 >= 0 {
			return "Валиден"
		} else {
			return "Не валиден"
		}
	}
	return "Не Валиден"
}

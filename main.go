package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func romanToArabic(roman string) (int, error) {
	risks := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100}
	sum := 0
	prevValue := 0

	for i := len(roman) - 1; i >= 0; i-- {
		value := risks[rune(roman[i])]
		if value < prevValue {
			sum -= value
		} else {
			sum += value
		}
		prevValue = value
	}

	if sum > 100 {
		return 0, fmt.Errorf("число больше 100")
	}

	return sum, nil
}

func arabicToRoman(num int) (string, error) {
	if num < 1 {
		return "", fmt.Errorf("результат меньше единицы")
	}
	if num > 100 {
		return "", fmt.Errorf("число больше 100")
	}

	var valuePairs = []struct {
		Value int
		Digit string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	result := ""
	for _, pair := range valuePairs {
		for num >= pair.Value {
			num -= pair.Value
			result += pair.Digit
		}
	}

	return result, nil
}

func isRomanNumeral(s string) bool {
	for _, char := range s {
		if char != 'I' && char != 'V' && char != 'X' && char != 'L' && char != 'C' {
			return false
		}
	}
	return true
}

func calculate(op1, op2 int, operator string) (int, error) {
	switch operator {
	case "+":
		return op1 + op2, nil
	case "-":
		return op1 - op2, nil
	case "*":
		return op1 * op2, nil
	case "/":
		if op2 == 0 {
			return 0, fmt.Errorf("деление на ноль невозможно")
		}
		return op1 / op2, nil
	default:
		return 0, fmt.Errorf("неизвестная операция")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите выражение: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	input = strings.ToUpper(strings.TrimSpace(input))
	parts := strings.Fields(input)
	if len(parts) != 3 {
		panic("Неверный формат ввода")
	}

	isNum1Roman := isRomanNumeral(parts[0])
	isNum2Roman := isRomanNumeral(parts[2])
	if isNum1Roman != isNum2Roman {
		panic("Оба числа должны быть в одной системе счисления")
	}

	var op1, op2 int

	if isNum1Roman {
		op1, err = romanToArabic(parts[0])
		if err != nil {
			panic(err.Error())
		}
		op2, err = romanToArabic(parts[2])
		if err != nil {
			panic(err.Error())
		}
	} else {
		op1, err = strconv.Atoi(parts[0])
		if err != nil {
			panic("Ошибка при преобразовании первого числа")
		}
		op2, err = strconv.Atoi(parts[2])
		if err != nil {
			panic("Ошибка при преобразовании второго числа")
		}
	}

	result, err := calculate(op1, op2, parts[1])
	if err != nil {
		panic(err.Error())
	}

	if isNum1Roman && result < 1 {
		panic("Результатом работы калькулятора с римскими числами могут быть только положительные числа")
	}

	if isNum1Roman {
		romanResult, err := arabicToRoman(result)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Результат:", romanResult)
	} else {
		fmt.Println("Результат:", result)
	}
}

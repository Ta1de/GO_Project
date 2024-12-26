package main

import (
	"fmt"
)

func main() {
	result := FindNumber(65, 17)
	fmt.Println(result)
}

func sumOfDigits(num, valuesum int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		if sum > valuesum {
			return 0
		}
		num /= 10
	}
	return sum
}

func generateNumbers(current int, lastDigit int, maxDigits int, callback func(int)) {
	if maxDigits == 0 {
		callback(current)
		return
	}

	for digit := lastDigit; digit <= 9; digit++ {
		newNumber := current*10 + digit
		generateNumbers(newNumber, digit, maxDigits-1, callback)
	}
}

func createSlise(step, valuesum int) []int {
	var result []int

	generateNumbers(0, 1, step, func(num int) {
		sum := sumOfDigits(num, valuesum)
		if sum > 0 {
			result = append(result, num)
		}

	})
	return result
}

func FindNumber(sum, num int) []int {
	var result []int
	if sum > num*9 {
		return result
	}
	result = createSlise(num, sum)
	if len(result) == 0 {
		return result
	}
	result = append(result, len(result))
	result = append(result, result[0])
	result = append(result, result[len(result)-1])

	return result
}

//package main
//
//// Функция для нахождения всех чисел с заданной суммой цифр и количеством цифр
//func findAll(sum, count int) []int {
//	var result []int
//	if count < 1 || count > 9 || sum < count || sum > count*9 {
//		return result
//	}
//
//	var results []int
//
//	// Рекурсивная функция для генерации чисел
//	var generate func(currentSum, currentCount, startDigit, currentValue int)
//	generate = func(currentSum, currentCount, startDigit, currentValue int) {
//		// Если достигли нужного количества цифр
//		if currentCount == count {
//			// Если сумма совпадает, добавляем число в результаты
//			if currentSum == sum {
//				results = append(results, currentValue)
//			}
//			return
//		}
//
//		// Генерируем только подходящие цифры
//		for digit := startDigit; digit <= 9; digit++ {
//			// Если сумма превысила нужное значение, прекращаем
//			if currentSum+digit > sum {
//				break
//			}
//			// Рекурсивно добавляем следующую цифру
//			generate(currentSum+digit, currentCount+1, digit, currentValue*10+digit)
//		}
//	}
//
//	// Запуск генерации
//	generate(0, 0, 1, 0)
//
//	// Если подходящих чисел нет, возвращаем пустые значения
//	if len(results) == 0 {
//		return result
//	}
//	// Подсчет общего количества и определение минимального и максимального значений
//	result = append(result, len(result))
//	result = append(result, result[0])
//	result = append(result, result[len(result)-1])
//
//	return results
//}
//
//func main() {
//	findAll(65, 17)
//}

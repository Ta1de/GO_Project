package main

import (
	"reflect"
	"testing"
)

func TestFindNumber(t *testing.T) {
	tests := []struct {
		sum      int
		num      int
		expected []int
	}{
		// Добавим автоматически сгенерированные тесты ниже
	}

	for sum := 20; sum <= 65; sum++ {
		for num := 2; num <= 17; num++ {
			expected := FindNumber(sum, num) // Ожидаемое значение
			tests = append(tests, struct {
				sum      int
				num      int
				expected []int
			}{
				sum:      sum,
				num:      num,
				expected: expected,
			})
		}
	}

	for _, tt := range tests {
		t.Run(
			t.Name(),
			func(t *testing.T) {
				got := FindNumber(tt.sum, tt.num)
				if !reflect.DeepEqual(got, tt.expected) {
					t.Errorf("Для суммы %d и числа %d ожидалось %v, но получено %v", tt.sum, tt.num, tt.expected, got)
				}
			})
	}
}

//package main
//
//import (
//	"fmt"
//	"testing"
//)
//
//func TestFindAll(t *testing.T) {
//	tests := []struct {
//		sum      int
//		count    int
//		expected struct {
//			totalCount int
//			minValue   int
//			maxValue   int
//		}
//	}{
//		// Добавим автоматически сгенерированные тесты ниже
//	}
//
//	// Автоматическая генерация тестов
//	for sum := 20; sum <= 200; sum++ {
//		for count := 2; count <= 30; count++ {
//			//fmt.Printf("Сумма: %d, Количество цифр: %d\n", sum, count)
//			totalCount, minValue, maxValue := findAll(sum, count)
//			tests = append(tests, struct {
//				sum      int
//				count    int
//				expected struct {
//					totalCount int
//					minValue   int
//					maxValue   int
//				}
//			}{
//				sum:   sum,
//				count: count,
//				expected: struct {
//					totalCount int
//					minValue   int
//					maxValue   int
//				}{
//					totalCount: totalCount,
//					minValue:   minValue,
//					maxValue:   maxValue,
//				},
//			})
//		}
//	}
//
//	// Запуск тестов
//	for _, tt := range tests {
//		t.Run(
//			fmt.Sprintf("Sum:%d Count:%d", tt.sum, tt.count), // Используем fmt.Sprintf для форматирования имени теста
//			func(t *testing.T) {
//				gotTotal, gotMin, gotMax := findAll(tt.sum, tt.count)
//				expected := tt.expected
//				if gotTotal != expected.totalCount || gotMin != expected.minValue || gotMax != expected.maxValue {
//					t.Errorf(
//						"Для суммы %d и количества %d ожидалось (Total: %d, Min: %d, Max: %d), но получено (Total: %d, Min: %d, Max: %d)",
//						tt.sum, tt.count, expected.totalCount, expected.minValue, expected.maxValue, gotTotal, gotMin, gotMax,
//					)
//				}
//			})
//	}
//}

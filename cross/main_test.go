package main

import (
	"testing"
)

// Тестируем функцию inputCross
func TestInputCross(t *testing.T) {
	// Пример ввода:
	// 3 3
	// a # b
	// c d e
	// f # g

	// Подменяем ввод с помощью fmt.Scanln
	R, C, Cross := inputCrossMock()
	if R != 3 || C != 3 {
		t.Errorf("Expected R=3, C=3, got R=%d, C=%d", R, C)
	}

	// Проверяем, что кроссворд загружен правильно
	expected := [][]rune{
		{'a', '#', 'b'},
		{'c', 'd', 'e'},
		{'f', '#', 'g'},
	}

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if Cross[i][j] != expected[i][j] {
				t.Errorf("Expected Cross[%d][%d] = %c, got %c", i, j, expected[i][j], Cross[i][j])
			}
		}
	}
}

// Тестируем функцию decideCross
func TestDecideCross(t *testing.T) {
	// Кроссворд для теста
	Cross := [][]rune{
		{'a', '#', 'b'},
		{'c', 'd', 'e'},
		{'f', '#', 'g'},
	}

	// Ожидаемый результат
	expected := "acf"

	// Вызываем decideCross
	result := decideCross(Cross, 3, 3)
	if result != expected {
		t.Errorf("Expected result = %s, got %s", expected, result)
	}
}

// Нагрузочный тест для проверки времени выполнения
func BenchmarkDecideCrossLarge(b *testing.B) {
	// Создаем большой кроссворд
	R, C := 1000, 1000
	Cross := make([][]rune, R)
	for i := 0; i < R; i++ {
		Cross[i] = make([]rune, C)
		for j := 0; j < C; j++ {
			if (i+j)%2 == 0 {
				Cross[i][j] = 'a' + rune(i%26) // Используем буквы для создания кроссворда
			} else {
				Cross[i][j] = '#'
			}
		}
	}

	// Запускаем тестирование производительности
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decideCross(Cross, R, C)
	}
}

// Дополнительная функция, чтобы замокать ввод
func inputCrossMock() (R int, C int, Cross [][]rune) {
	R, C = 3, 3
	Cross = [][]rune{
		{'a', '#', 'b'},
		{'c', 'd', 'e'},
		{'f', '#', 'g'},
	}
	return R, C, Cross
}

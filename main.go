// Package add предоставляет простые арифметические операции.
//
// Основная цель пакета — демонстрация минимального Go-модуля
// с публичным API и корректной документацией GoDoc.
package add

import "golang.org/x/exp/constraints"

// Number описывает числовые типы, поддерживающие оператор сложения.
type Number interface {
	constraints.Integer | constraints.Float
}

// Add складывает два числовых значения и возвращает их сумму.
//
// Подробное описание операции сложения см. в документации:
// https://example.com/addition.html
func Add[T Number](a, b T) T {
	return a + b
}

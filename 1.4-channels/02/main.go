package main

import (
	"fmt"
	"time"
)

// gather выполняет переданные функции одновременно
// и возвращает срез с результатами, когда они готовы
func gather(funcs []func() any) []any {
	// начало решения
	type res struct {
		idx    int
		result any
	}
	resChan := make(chan res)

	for idx := range funcs {
		go func() {
			resChan <- res{
				idx:    idx,
				result: funcs[idx](),
			}
		}()
	}

	result := make([]any, len(funcs))
	for i := 0; i < len(funcs); i++ {
		res := <-resChan
		result[res.idx] = res.result
	}
	return result
	// выполните все переданные функции,
	// соберите результаты в срез
	// и верните его

	// конец решения
}

// squared возвращает функцию,
// которая считает квадрат n
func squared(n int) func() any {
	return func() any {
		time.Sleep(time.Duration(n) * 100 * time.Millisecond)
		return n * n
	}
}

func main() {
	funcs := []func() any{squared(2), squared(3), squared(4)}

	start := time.Now()
	nums := gather(funcs)
	elapsed := float64(time.Since(start)) / 1_000_000

	fmt.Println(nums)
	fmt.Printf("Took %.0f ms\n", elapsed)
}

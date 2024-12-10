package main

import "fmt"

func main() {
	// Создаем два множества как срезы
	set1 := []int{1, 2, 3, 4}
	set2 := []int{3, 4, 5, 6}

	// Получаем пересечение множеств
	intersection_set := intersect(set1, set2)

	// Выводим результат
	fmt.Println("Пересечение множеств:", intersection_set)
}

func intersect(set1, set2 []int) []int {

	intersection_set := []int{}

	// Создаем map для второго множества для быстрого поиска
	set2Map := make(map[int]struct{})
	for _, value := range set2 {
		set2Map[value] = struct{}{}
	}

	// Проходим по первому множеству и проверяем наличие элемента во втором
	for _, value := range set1 {
		if _, exists := set2Map[value]; exists {
			// Если элемент есть в обоих множествах, добавляем его в пересечение
			intersection_set = append(intersection_set, value)
		}
	}

	return intersection_set
}

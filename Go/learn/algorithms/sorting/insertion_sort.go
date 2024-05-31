package sorting

import "Go/time_completion"

func Insertion_sort(toSort []int) (sorted []int) {
	defer time_completion.FunctionTimer(Insertion_sort)()

	for i := 1; i < len(toSort); i++ {
		j := i - 1
		next := toSort[i]
		for j >= 0 && toSort[j] > next {
			toSort[j+1] = toSort[j]
			j = j - 1
		}
		toSort[j+1] = next
	}
	return toSort
}

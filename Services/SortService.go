package Services

import (
	"GoApp/Model"
)

type Sort interface {
	Sort(slice []int) (sorted []int)
}

type SortService struct{}

func (SortService) Sort(slice []int) (sorted []int) {
	return mergeSort(slice, Model.NewRepeatedList(), 0)
}

/**
this is basically a variant of merge-sort that order numbers and put into a shared list the repeated numbers
to put them at the end of the list.
Complexity: O(N*log(N))
*/
func mergeSort(items []int, repeatedList *Model.RepeatedList, round int) []int {
	var num = len(items)

	if num <= 1 {
		return items
	}

	middle := int(num / 2)
	var (
		left  = make([]int, middle)
		right = make([]int, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	result := merge(mergeSort(left, repeatedList, round+1), mergeSort(right, repeatedList, round+1), repeatedList)
	if round == 0 {
		for j := 0; j < len(repeatedList.List); j++ {
			result = append(result, repeatedList.List[j])
		}
	}
	return result
}

func merge(left, right []int, repeated *Model.RepeatedList) []int {
	result := make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			if left[0] > right[0] {
				result[i] = right[0]
				right = right[1:]
			} else {
				result[i] = right[0]
				right = right[1:]
				repeated.Add(left[0])
				left = left[1:]
			}
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}
	// since we move repeated numbers, the len of  list is not always (len(left) + len(right)) but it can be considered a ceiling
	return result[:i]
}

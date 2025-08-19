package main

import "sort"

func solution(numbers []int) int {
    sort.Ints(numbers)

    n := len(numbers)
    return numbers[n-1] * numbers[n-2]
}

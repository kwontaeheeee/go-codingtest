package main

func solution(numbers []int, direction string) []int {
    n := len(numbers)
    result := make([]int, n)

    if direction == "right" {
        for i := 0; i < n; i++ {
            result[(i+1)%n] = numbers[i]
        }
    } else {
        for i := 0; i < n; i++ {
            result[i] = numbers[(i+1)%n]
        }
    }
    return result
}

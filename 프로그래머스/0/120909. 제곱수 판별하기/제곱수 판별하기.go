package main

import "math"

func solution(n int) int {
    sqrtN := int(math.Sqrt(float64(n)))
    if sqrtN*sqrtN == n {
        return 1
    }
    return 2
}

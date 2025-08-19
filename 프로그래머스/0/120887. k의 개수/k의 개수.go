package main

func solution(i int, j int, k int) int {
    count := 0
    for n := i; n <= j; n++ {
        x := n
        for x > 0 {
            if x%10 == k {
                count++
            }
            x /= 10
        }
    }
    return count
}

package main

func solution(balls int, share int) int {
    return int(comb(balls, share))
}

func comb(n, r int) int64 {
    if r == 0 || n == r {
        return 1
    }
    if r == 1 {
        return int64(n)
    }
    return comb(n-1, r-1) + comb(n-1, r)
}

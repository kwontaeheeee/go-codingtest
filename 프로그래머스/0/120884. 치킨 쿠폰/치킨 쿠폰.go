func solution(chicken int) int {
    coupons := chicken   
    service := 0         

    for coupons >= 10 {
        newChicken := coupons / 10   
        service += newChicken
        coupons = coupons%10 + newChicken 
    }

    return service
}

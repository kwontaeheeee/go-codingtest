func solution(hp int) int {
    generals := hp / 5        
    hp = hp % 5

    soldiers := hp / 3         
    hp = hp % 3

    workers := hp              

    return generals + soldiers + workers
}

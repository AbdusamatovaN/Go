func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    s := x
    res := 0
    for {
        if x == 0 {
            break
        }
        res = res * 10 + x % 10
        x /= 10
    }
    return res == s
}

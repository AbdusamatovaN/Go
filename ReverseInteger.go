func reverse(x int) int {
    var f, res int
    if x < 0 {
        f = -1
    } else {
        f = 1
    }
    res = int(math.Abs(float64(x % 10)))
    x /= 10
    for int(math.Abs(float64(x))) > 0 {
        res = res*10 + int(math.Abs(float64(x % 10)))
        x /= 10
        fmt.Println(x, res)
    }
    if res < -2147483648 || res > 2147483647 {
        return 0
    }
    return res*f
}

func myAtoi(s string) int {
    var i, res, flag int = 0, 0, 1
    for i < len(s){
        if s[i] == ' ' {
            i++
        } else {
            break
        }
    }
    if i < len(s) && s[i] == '-' {
        flag = -1
        i++
    } else if i < len(s) && s[i] == '+'{
        i++
    }
    L:
        for i < len(s) {
            switch {
                case s[i] >= '0' && s[i] <= '9': 
                    res = res*10 + int(s[i] - '0')
                    if res > math.MaxInt32{
                        if flag == 1 {
                            return int(math.MaxInt32)
                        } else { 
                            return int(math.MinInt32)
                        }
                    }
                default : 
                    break L
            }
            i++
        }
    return res*flag
}

func longestPalindrome(s string) string {
    if len(s) < 1 {
        return ""
    }
    var r, l int
    for i := 0; i < len(s); i++ {
        l1 := expan(s, i, i)
        r1 := expan(s, i, i+1)
        leng := int(math.Max(float64(l1), float64(r1)))
        if leng > r - l {
            l = i - (leng-1)/2
            r = i + leng/2
        }
    }
    return string(s[l:r+1])
}

func expan(s string, l, r int) int {
    for {
        if l >= 0 && r < len(s) && s[l] == s[r] {
            l--
            r++
        } else {
            break
        } 
    }
    return r - l - 1
}

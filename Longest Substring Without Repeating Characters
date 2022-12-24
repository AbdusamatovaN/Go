func lengthOfLongestSubstring(s string) int {
    if s == "" {
        return 0
    }
    if len(s) == 1 {
        return 1
    }
    r := []rune(s)
    str := ""
    strMax := ""
    for i := 0; i < len(r); i++ {
        for j := 0; j < len(str); j++ {
            //fmt.Println(j, len(str), string(str[j]), string(r[i]))
            if string(str[j]) == string(r[i]) {
                if len(strMax) < len(str) {
                    strMax = str
                }
                fmt.Println(str)
                str = str[j+1:]
                fmt.Println(str)
                str += string(r[i])
                fmt.Println(str)
                break
            }
            if j == len(str) - 1 {
                str += string(r[i])
                break
            }
            //fmt.Println(j, strMax, str, string(r[i]))
        }
        if str == "" {
            str += string(r[i])
        }
    }
    if len(strMax) < len(str) {
        strMax = str
        str = ""
    }
    fmt.Println(strMax, str)
    return(len(strMax))
}

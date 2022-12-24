func rim(s string) int {
    switch s {
        case "I" : return 1
        case "V" : return 5
        case "X" : return 10
        case "L" : return 50
        case "C" : return 100
        case "D" : return 500
        case "M" : return 1000
    }
    return 0
}

func romanToInt(s string) int {
    r := []rune(s)
    var res, flag int
    for i := 0; i < len(r); i++ {
        if i != len(r) - 1 {
            if string(r[i]) == "I" {
                if string(r[i+1]) == "V" {
                    res += 4
                    i++
                    flag = 1
                } else if string(r[i+1]) == "X" {
                    res += 9
                    i++
                    flag = 1
                }
            } else if string(r[i]) == "X" {
                if string(r[i+1]) == "L" {
                    res += 40
                    i++
                    flag = 1
                } else if string(r[i+1]) == "C" {
                    res += 90
                    i++
                    flag = 1
                }
            } else if string(r[i]) == "C" {
                if string(r[i+1]) == "D" {
                    res += 400
                    i++
                    flag = 1
                } else if string(r[i+1]) == "M" {
                    res += 900
                    i++
                    flag = 1
                }
            }
        }
        if flag != 1 {
            res += rim(string(r[i]))
        }
        flag = 0
    }
    return res
}

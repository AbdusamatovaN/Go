func convert(s string, numRows int) string {
    if len(s) <= 1 || numRows == 1 { return s}
    mas := make([][]rune, numRows)
    var c, f int
    str := ""
    for _, runeValue := range s {
        mas[c] = append(mas[c], runeValue)
        if ((c + 1) / numRows == 1 || f == 1) && c != 0 {
            c--
            f = 1
        } else if c == 0 || f == 0 {
            c++
            f = 0
        } 
    }
    for i := 0; i < numRows; i++ {
        for j := 0; j < len(mas[i]); j++ {
            str += string(mas[i][j])
        }
    } 
    return str
}

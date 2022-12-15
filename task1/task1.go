package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("test.txt")
	rd := bufio.NewReader(f)

	nstr, _ := rd.ReadString('\n')
	nstr = strings.Trim(nstr, "\n")

	n, _ := strconv.Atoi(nstr)

	for i := 0; i < n; i++ {
		s, _ := rd.ReadString('\n')

		data := strings.Split(strings.Trim(s, "\n"), ",")

		//Различные символы в ФИО:
		fio := data[0] + data[1] + data[2]
		fiorunes := make(map[rune]int)
		for _, r := range fio {
			fiorunes[r]++
		}

		fioRuneCount := len(fiorunes)
		//Сумма цифр в дне и месяце рождения
		dateDigitsSum := 0
		dateDigits := data[3] + data[4]
		for _, d := range dateDigits {
			digit, _ := strconv.Atoi(string(d))
			dateDigitsSum += digit
		}
		//Номер в алфавите первой буквы фамилии
		var nf int
		if int(fio[0]) >= 65 && int(fio[0]) <= 90 {
			nf = int(fio[0]) - 64
		} else if int(fio[0]) >= 97 && int(fio[0]) <= 122 {
			nf = int(fio[0]) - 96
		}

		// Конвертация двоичного значения в шестнадцатеричное, оставить только 3 последних разряда
		val := int64(fioRuneCount + dateDigitsSum*64 + nf*256&0xfff)

		v := strconv.FormatInt(val, 16)
		fmt.Print(v, " ")
	}
}

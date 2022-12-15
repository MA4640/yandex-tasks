package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N int
	fmt.Scan(&N)
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < N; i++ {
		str, _ := reader.ReadString('\n')
		str = strings.Trim(str, " \n")

		//разбиение на подстроки
		strs := strings.Split(str, ",")
		var ints []int

		for _, s := range strs {
			num, err := strconv.Atoi(s)
			if err == nil {
				ints = append(ints, num)
			}
		}

		word := strings.Join(strs, "")
		counts := make(map[rune]int)
		for _, r := range word {
			if r >= '0' && r <= '9' {
				continue
			}
			counts[r]++
		}

		//fmt.Printf("%v", counts)

		//Подсчитывается количество различных символов в ФИО (регистр важен, А и а — разные символы).
		//упущена конкатенация (планировала скастить в одну строку в массив рун при помощи карты
		// (ключ руна - значение количество))
		//здесь упущен блок
		p1 := len(counts)

		//Берётся сумма цифр в дне и месяце рождения, умноженная на 64.
		p2 := 64 * (ints[0]/10 + ints[0]%10 + ints[1]/10 + ints[1]%10)

		//Для первой (по позиции в слове) буквы фамилии определяется её номер в алфавите (в 1-индексации), умноженный (далее!) на 256 (регистр буквы не важен).
		var p3 int
		if int(str[0]) >= 65 && int(str[0]) <= 90 {
			p3 = int(str[0]) - 64
		} else if int(str[0]) >= 97 && int(str[0]) <= 122 {
			p3 = int(str[0]) - 96
		}

		p4 := int64(p1+p2+p3*256) & 0xfff

		// Конвертация двоичного значения в шестнадцатеричное
		v := strconv.FormatInt(p4, 16)

		//У результата сохраняются только 3 младших разряда (если значимых разрядов меньше, то шифр дополняется до 3-х разрядов ведущими нулями)
		/*run := []rune(v)

		for i := 3; i > 0; i-- {
			fmt.Println(string(vrun[len(vrun)-i]))
		}*/
		fmt.Print(v, " ")
		//	fmt.Print(" ")
	}
}

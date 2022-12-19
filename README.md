Задача — помочь вычислить для каждого кандидата его шифр. 

Заботясь о персональных данных соискателей, компания придумала хитрый алгоритм шифрования: 
    • Подсчитывается количество различных символов в ФИО (регистр важен, А и а — разные символы). 
    • Берётся сумма цифр в дне и месяце рождения, умноженная на 64. 
    • Для первой (по позиции в слове) буквы фамилии определяется её номер в алфавите (в 1-индексации), умноженный на 256 (регистр буквы не важен). 
    • Полученные числа суммируются. 
    • Результат переводится в 16-чную систему счисления (в верхнем регистре). 
    • У результата сохраняются только 3 младших разряда (если значимых разрядов меньше, то шифр дополняется до 3-х разрядов ведущими нулями). 

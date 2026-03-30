package parser

import(
	"bufio"
	"fmt"
	"os"
	"strings"
	"productStorage/internal/models"
)

func ParseProductsFromFile(path string)([]*models.Product,error){ // создаём функцию, которая будет проходить все строчки в файле, возвращает массив с продуктами и ошибку
	file, err := os.Open(path) // открываем файл по пути
	if err != nil{
		return nil, fmt.Errorf("Error: %w", err) // если файл не открылся выдаём ошибку
	}

	defer file.Close() // закрываем файл после отработки функции

	var Products []*models.Product // создаём переменную, которая хранит массив продуктов
	lineNum := 1 // создаём переменную которая указывает номер строки
	scanner := bufio.NewScanner(file) // создаём сканнер, который будет обходить(сканировать) весь файл

	for scanner.Scan(){ // сканируем файл построчно
		line := scanner.Text() // создаём переменную, которая хранит в себе всю строку

		if strings.TrimSpace(line) == ""{ // проверка, что строка не пустая
			lineNum++ // если она пустая, увеличиваем номер строки
			continue // продолжаем чтение следующей строки
		}

		parts := strings.Split(line, ";") // переменная, которая хранит массив частей строки, части разделяются по ; 

		if len(parts) != 3{ // проверка что частей в строке должно быть 3
			fmt.Printf("[Строка: %d] Недостаточно данных для сохранения продукта", lineNum) // если их не 3, то выводим ошибку
			lineNum++ // увеличиваем номер строки
			continue //  продолжаем чтение следующей строки
		}

		name := strings.TrimSpace(parts[0]) // создаём переменную, которая хранит название продукта
		sbin := strings.TrimSpace(parts[1]) // создаём переменную, которая хранит sbin продукта
		date := strings.TrimSpace(parts[2]) // создаём переменную, которая хранит дату срока годности продукта

		product, err := models.NewProduct(name, sbin, date) // создаём новый продукт

		if err != nil{
			fmt.Printf("Товар %s на строке %d отклонён. Error: %w\n", name, lineNum, err) // если во время создания произошла ошибка(не прошёл валидацию), выводим эту ошибку
		} else{
			Products = append(Products, product) // иначе добавляем продукт в массив
		}
		lineNum++ // увеличиваем номер строки

	}

	if err:= scanner.Err(); err != nil{
		return nil, fmt.Errorf("Ошибка чтения файла %w",err) // если во время чтения файла произошла ошибка, выводим её, а массив не возвращаем
	}

	return Products, nil // если всё прошло гладко, возвращаем массив с продуктами и nil в качестве ошибки
}
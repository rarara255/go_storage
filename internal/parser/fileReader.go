package parser

import(
	"bufio"
	"fmt"
	"os"
	"strings"
	"productStorage/internal/models"
)

func ParseProductsFromFile(path string)([]*models.Product,error){
	file, err := os.Open(path)
	if err != nil{
		return nil, fmt.Errorf("Error: %w", err)
	}

	defer file.Close()

	var Products []*models.Product
	lineNum := 1
	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		line := scanner.Text()

		if strings.TrimSpace(line) == ""{
			lineNum++
			continue
		}

		parts := strings.Split(line, ";")

		if len(parts) != 3{
			fmt.Printf("[Строка: %d] Недостаточно данных для сохранения продукта", lineNum)
			lineNum++
			continue
		}

		name := strings.TrimSpace(parts[0])
		sbin := strings.TrimSpace(parts[1])
		date := strings.TrimSpace(parts[2])

		product, err := models.NewProduct(name, sbin, date)

		if err != nil{
			fmt.Printf("Товар %s на строке %d отклонён. Error: %w\n", name, lineNum, err)
		} else{
			Products = append(Products, product)
		}
		lineNum++

	}

	if err:= scanner.Err(); err != nil{
		return nil, fmt.Errorf("Ошибка чтения файла %w",err)
	}

	return Products, nil
}
package main

import (
	"flag"
	"fmt"
	"log"
	"go_storage/internal/parser"
	"go_storage/internal/postgres"
)

func main() {
	// rep := storage.NewRepository[*models.Product]()

	fileNamePTR := flag.String("file", "data.txt", "Передайте адрес data файла")
	pgHost := flag.String("pg-host", "localhost", "Postgre host")
	pgPort := flag.String("pg-port", "5432", "Postgre port")
	pgUser := flag.String("pg-user", "postgres", "Postgre user")
	pgPass := flag.String("pg-pass", "", "Postgre password")
	pgName := flag.String("pg-db", "productdb", "Postgre database name")

	flag.Parse()

	cfg := postgres.Config{
		Host:     *pgHost,
		Port:     *pgPort,
		User:     *pgUser,
		Password: *pgPass,
		DBName:   *pgName,
	}

	repo, err := postgres.NewPGStorage(cfg)
	if err != nil {
		log.Fatal("Don't connected to DB %w", err)
	}
	fmt.Println("Connected to Postgre. Table created successfully")
	// используется для того чтобы компилятор прочитал переданный флаг

	currentFile := *fileNamePTR

	fmt.Printf("Система настроена на чтение из файла: %s", currentFile)
	fmt.Printf("Чтение из файла (%s)", currentFile)

	Products, err := parser.ParseProductsFromFile(currentFile)
	if err != nil {
		log.Fatalf("Fatal Error: %v\n", err)
	}

	//var repos storage.productStorage = storage.

	for _, product := range Products {
		repo.Add(product)
	}

	items := repo.GetAll()

	fmt.Printf("На склад отгружено %d товаров \n", len(items))

	for _, item := range items {
		fmt.Printf("-Name: %10s | SBIN: %20s | Годен до %s \n", item.Name, item.SBIN, item.DateToString())
	}

	fmt.Scanln()
}

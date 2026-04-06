package main

import(
	"fmt"
	"flag"
	"log"
	"productStorage/internal/models"
	"productStorage/internal/parser"
	"productStorage/internal/storage"
)

func main(){
	// rep := storage.NewRepository[*models.Product]()

	fileNamePTR := flag.String("file", "data.txt", "Передайте адрес data файла")

	flag.Parse()
	// используется для того чтобы компилятор прочитал переданный флаг

	currentFile := *fileNamePTR

	fmt.Printf("Система настроена на чтение из файла: %s", currentFile)
	fmt.Printf("Чтение из файла (%s)", path)

	var repos storage.productStorage = storage.
	// if err != nil{
	// 	log.Fatalf("Fatal Error: %v\n", err)
	// }

	// for _, product := range Products{
	// 	rep.Add(product)
	// }

	// items := rep.GetAll()

	// fmt.Printf("На склад отгружено %d товаров \n",len(items))

	// for _, item := range items{
	// 	fmt.Printf("-Name: %10s | SBIN: %20s | Годен до %s \n", item.Name, item.SBIN, item.DateToString())
	// }

	fmt.Scanln()
}

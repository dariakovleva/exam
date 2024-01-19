package main

import (
	"fmt"
	"log"

	"example.com/m/ticket1"
	_ "example.com/m/ticket3"
	_ "example.com/m/ticket6"
)

// 1
func main() {
	person1, err := ticket1.NewPerson(
		15,
		"Константин",
	)
	if err != nil {
		log.Fatal(err)
	}

	person2, err := ticket1.NewPerson(
		98,
		"Макс",
	)
	if err != nil {
		log.Fatal(err)
	}

	person3, err := ticket1.NewPerson(
		27,
		"Дарья",
	)
	if err != nil {
		log.Fatal(err)
	}

	// Создайте срез, содержащий несколько экземпляров этой структуры.
	srez_of_persons := make([]ticket1.Person, 0)
	ticket1.TryAdd(&srez_of_persons, *person1)
	ticket1.TryAdd(&srez_of_persons, *person2)
	ticket1.TryAdd(&srez_of_persons, *person3)

	fmt.Println(person1)
	fmt.Println(srez_of_persons)
	fmt.Println(ticket1.AverageAgeValue(srez_of_persons))
}

//

// //3
// func main() {
// 	book1, err := ticket3.NewBook(1852, "Детство")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	book2, err := ticket3.NewBook(1872, "Бесы")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	book3, err := ticket3.NewBook(1841, "Шинель")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	srez_of_books := make([]ticket3.Book, 0)
// 	ticket3.TryAdd(&srez_of_books, *book1)
// 	ticket3.TryAdd(&srez_of_books, *book2)
// 	ticket3.TryAdd(&srez_of_books, *book3)

// 	fmt.Println(book1)
// 	fmt.Println(srez_of_books)
// 	fmt.Println(ticket3.AverageYearValue(srez_of_books))
// }

//

// // 6
// func main() {
// 	product1, err := ticket6.NewProduct("торт", 5)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	product2, err := ticket6.NewProduct("леденцы", 31)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	product3, err := ticket6.NewProduct("печенье", 45)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	srez_of_products := make([]ticket6.Product, 0)
// 	ticket6.TryAdd(&srez_of_products, *product1)
// 	ticket6.TryAdd(&srez_of_products, *product2)
// 	ticket6.TryAdd(&srez_of_products, *product3)

// 	fmt.Println(product1)
// 	fmt.Println(srez_of_products)
// 	fmt.Println("среднее значение:", ticket6.AveragePriceValue(srez_of_products))
// }

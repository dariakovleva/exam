package ticket3

import (
	"errors"
)

type Book struct {
	year   int
	title  string
	author string
}

func NewBook(year int, title string) (*Book, error) {
	book := &Book{
		title: title,
	}
	if err := book.setYear(year); err != nil {
		return nil, err
	}
	if err := book.setAuthor(book.title); err != nil {
		return nil, err
	}
	return book, nil
}

var errorYearIsIncorrect = errors.New("Год выпуска должен быть от 1800 до 2023")

func (b *Book) setYear(year int) error {
	if year < 1800 || year > 2023 {
		return errorYearIsIncorrect
	}
	b.year = year
	return nil
}

func (b *Book) setAuthor(author string) error {
	switch {
	case b.title == "Детство" || b.title == "Оторочество" || b.title == "Юность":
		b.author = "Л.Н.Толстой"
	case b.title == "Идиот" || b.title == "Бесы" || b.title == "Бедные люди":
		b.author = "Ф.М.Достоевский"
	case b.title == "Вий" || b.title == "Шинель" || b.title == "Мёртвые души":
		b.author = "Н.В.Гоголь"
	}
	return nil
}

func AverageYearValue(srez []Book) float64 {
	if len(srez) == 0 {
		return 0
	}
	sumOfYears := 0
	for _, b := range srez {
		sumOfYears += b.year
	}
	result := float64(sumOfYears) / float64(len(srez))
	return result
}

func TryAdd(srez *[]Book, newBooks Book) bool {
	for _, b := range *srez {
		if b == newBooks {
			return false
		}
	}
	*srez = append(*srez, newBooks)
	return true
}

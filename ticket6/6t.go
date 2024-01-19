package ticket6

import (
	"errors"
)

type Product struct {
	name   string
	price  float64
	amount int
}

func NewProduct(name string, amount int) (*Product, error) {
	product := &Product{
		name: name,
	}
	if err := product.setAmount(amount); err != nil {
		return nil, err
	}
	if err := product.setPrice(product.name); err != nil {
		return nil, err
	}
	return product, nil
}

var errorPrice = errors.New("Недопустимое  значение цены")
var errorAmount = errors.New("Недопустимое значение количества продукта")

func (p *Product) setPrice(name string) error {
	switch {
	case p.name == "мармелад":
		p.price = 620
	case p.name == "печенье":
		p.price = 230
	case p.name == "леденцы":
		p.price = 190
	case p.name == "вафли":
		p.price = 340
	case p.name == "торт":
		p.price = 1520
	}
	if p.price < 0 || p.price > 100000 {
		return errorPrice
	}
	return nil
}

func (p *Product) setAmount(amount int) error {
	if amount < 0 || amount > 1000 {
		return errorAmount
	}
	p.amount = amount
	return nil
}

func AveragePriceValue(srez []Product) float64 {
	if len(srez) == 0 {
		return 0
	}
	summOfPrices := 0
	for _, p := range srez {
		summOfPrices += int(p.price)
	}
	result := summOfPrices / len(srez)
	return float64(result)
}

func TryAdd(srez *[]Product, newProd Product) bool {
	for _, p := range *srez {
		if p == newProd {
			return false
		}
	}
	*srez = append(*srez, newProd)
	return true
}

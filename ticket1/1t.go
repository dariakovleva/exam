package ticket1

import (
	"errors"
)

// Создайте структуру для следующей предметной области "Человек" с полями "Возраст", "Имя" и "Статус".
type Person struct {
	age    int
	name   string
	status string
}

func NewPerson(age int, name string) (*Person, error) {
	person := &Person{
		name: name,
	}
	if err := person.SetAge(age); err != nil {
		return nil, err
	}
	return person, nil
}

func (p *Person) SetAge(age int) error {
	switch {
	// Установите ограничения для возраста от 0 до 150.
	// Статус должен автоматически устанавливаться в зависимости от возраста и не иметь собственного сеттера.
	case age >= 0 && age <= 150:
		p.age = age
		switch {
		case p.age < 2:
			p.status = "младенец"
		case p.age < 14:
			p.status = "ребенок"
		case p.age < 18:
			p.status = "подросток"
		case p.age < 25:
			p.status = "молодой"
		case p.age < 60:
			p.status = "взрослый"
		case p.age <= 150:
			p.status = "пожилой"
		}
		return nil
	// Запретите устанавливать значения вне этого диапазона, возвращая ошибку при попытке.
	default:
		var errorAgeIsIncorrect string
		switch {
		case age < 0:
			errorAgeIsIncorrect = "указано значение возраста менее допустимого"
		case age > 150:
			errorAgeIsIncorrect = "указано значение возраста более допустимого"
		}
		return errors.New(errorAgeIsIncorrect)
	}
}

// Напишите метод для вычисления среднего возраста, который принимает срез структур в качестве входных данных и возвращает сумму возрастов.
func AverageAgeValue(srez []Person) float64 {
	if len(srez) == 0 {
		return 0
	}
	sumOfAges := 0
	for _, p := range srez {
		sumOfAges += p.age
	}
	result := float64(sumOfAges) / float64(len(srez))
	return result
}

// Также реализуйте метод с именем tryAdd, который принимает указатель на срез структур и экземпляр структуры.
func TryAdd(srez *[]Person, newPers Person) bool {
	// Этот метод должен добавить структуру в срез и вернуть true, если структуры еще нет в срезе, иначе вернуть false.
	for _, p := range *srez {
		if p == newPers {
			return false
		}
	}
	*srez = append(*srez, newPers)
	return true
}

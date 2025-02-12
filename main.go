Ось базовий приклад обробки даних в Go:

```go
package main

import (
	"fmt"
	"sort"
	"strings"
)

// Структура для студентів
type Student struct {
	Name string
	Age  int
	GPA  float64
}

// Функція для виводу інформації про студента
func (s Student) String() string {
	return fmt.Sprintf("Name: %s, Age: %d, GPA: %.2f", s.Name, s.Age, s.GPA)
}

// Функція для сортування студентів за іменем
type ByName []Student

func (n ByName) Len() int           { return len(n) }
func (n ByName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n ByName) Less(i, j int) bool { return n[i].Name < n[j].Name }

func main() {
	// Вхідні дані
	students := []Student{
		{"John", 20, 3.7},
		{"Amy", 22, 3.9},
		{"Peter", 21, 3.8},
	}

	// Виведемо студентів
	fmt.Println("Students:")
	for _, student := range students {
		fmt.Println(student)
	}

	// Пошук студентів за GPA
	minGPA := 3.8
	fmt.Printf("\nStudents with GPA greater than %.2f:\n", minGPA)
	for _, student := range students {
		if student.GPA > minGPA {
			fmt.Println(student)
		}
	}

	// Сортування студентів за іменем
	sort.Sort(ByName(students))
	fmt.Println("\nStudents sorted by name:")
	for _, student := range students {
		fmt.Println(student)
	}

	// Перетворення імені студентів на верхній регістр
	fmt.Println("\nStudent names in uppercase:")
	for _, student := range students {
		fmt.Println(strings.ToUpper(student.Name))
	}
}
```

Цей скрипт має наступні функції:

1. Виведення інформації про студентів.
2. Пошук студентів з великим середнім балом (GPA).
3. Сортування студентів за іменем.
4. Перетворення імені студентів на верхній регістр.

Пам'ятайте, що даний код потребує додаткового тестування та корекції для використання в реальних умовах.
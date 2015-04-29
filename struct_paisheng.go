// package main

// import (
// 	"fmt"
// )

// type Animal struct {
// 	Name string
// }

// func (animal *Animal) Say() {
// 	fmt.Println("WangWang~")
// }

// func (animal Animal) MyName() {
// 	fmt.Println("MyName is " + animal.Name)
// }

// type Dog struct {
// 	Animal
// }

// func main() {
// 	dog := &Dog{Name: "10"}
// 	fmt.Println(dog)
// 	// fmt.Println(dog.Name)
// 	// dog.Say()
// 	// dog.Animal.Say()
// }

// TEST_1
// package main

// import "fmt"

// type Person struct {
// 	name string
// 	age  int
// 	tel  string
// }

// type Student struct {
// 	person Person // 有另一个字段
// 	school string
// }

// func main() {
// 	anna := new(Student)
// 	anna.person.name = "Anna"
// 	anna.person.tel = "12345678"
// 	//anna.school = "mit"

// 	fmt.Printf("My name is %s, and my tel number is %s\n", anna.person.name, anna.person.tel)

// }

// TEST_2
// package main

// import "fmt"

// type Person struct {
// 	name string
// 	age  int
// 	tel  string
// }

// type Student struct {
// 	person Person // 有另一个字段
// 	school string
// }

// //在person上面定义了一个传值的method
// func (p Person) Hello() {
// 	fmt.Printf("My name is %s, and my tel number is %s\n", p.name, p.tel)
// }

// func main() {
// 	anna := new(Student)
// 	anna.person.name = "Anna"
// 	anna.person.tel = "2345678"
// 	//anna.school = "mit"

// 	anna.person.Hello()
// }

// TEST_3
// package main

// import "fmt"

// type Person struct {
// 	name string
// 	age  int
// 	tel  string
// }

// type Student struct {
// 	person Person // 有另一个字段
// 	school string
// }

// //在person上面定义了一个传指针的method
// func (p *Person) Hello() {
// 	fmt.Printf("My name is %s, and my tel number is %s\n", p.name, p.tel)
// }
// func main() {
// 	anna := new(Student)
// 	anna.person.name = "Anna"
// 	anna.person.tel = "345678"
// 	//anna.school = "mit"

// 	anna.person.Hello()
// }

// TEST_4
// package main

// import "fmt"

// type Person struct {
// 	name string
// 	age  int
// 	tel  string
// }

// type Student struct {
// 	Person // 有一个匿名字段
// 	school string
// }

// func main() {
// 	anna := new(Student)
// 	anna.Person.name = "Anna" // 用类型代替字段名
// 	anna.tel = "123456789"    // 或者直接引用
// 	//anna.school = "mit"

// 	fmt.Printf("My name is %s, and my tel number is %s\n", anna.Person.name, anna.tel)

// }

// TEST_5
// package main

// import "fmt"

// type Person struct {
// 	name string
// 	age  int
// 	tel  string
// }

// type Student struct {
// 	Person // 有另一个字段
// 	school string
// }

// //在person上面定义了一个传值的method
// func (p Person) Hello() {
// 	fmt.Printf("My name is %s, and my tel number is %s\n", p.name, p.tel)
// }

// func main() {
// 	anna := new(Student)
// 	anna.Person.name = "Anna"
// 	anna.tel = "2345"
// 	//anna.school = "mit"

// 	anna.Hello() // 匿名字段，相当于匿名结构直接嵌入，在同一层面

// }

// TEST_6
// package main

// import "fmt"

// type Person struct {
// 	name string
// 	age  int
// 	tel  string
// }

// type Student struct {
// 	Person // 有另一个字段
// 	school string
// }

// //在person上面定义了一个传指针的method
// func (p *Person) Hello() {
// 	fmt.Printf("My name is %s, and my tel number is %s\n", p.name, p.tel)
// }
// func main() {
// 	anna := new(Student)
// 	anna.Person.name = "Anna"
// 	anna.tel = "345"
// 	//anna.school = "mit"

// 	anna.Hello() // 匿名字段，相当于匿名结构直接嵌入，在同一层面

// 	//fmt.Printf( "My name is %s, and my tel number is %s\n", anna.person.name, anna.person.tel)
// }
//
// TEST_7
package main

import "fmt"

type Person struct {
	name string
	age  int
	tel  string
}

type Student struct {
	person *Person // 指针指向一个字段
	school string
}

func main() {
	anna := new(Student)
	anna.person = new(Person) // 必须，否则panic
	anna.person.name = "Anna"
	anna.person.age = 22
	anna.person.tel = "12345"
	anna.school = "mit"

	fmt.Printf("My name is %s, and my tel number is %s\n", anna.person.name, anna.person.tel)

}

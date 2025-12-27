package pgms

import "fmt"

type Person struct {
  name string
  age int
  job string
  salary int
}

func Struct() {
	/*
		type struct_name struct {
			member1 datatype;
			member2 datatype;
			member3 datatype;
			...
		}
	*/
	var person1 Person
	person1.name = "Bob"
	person1.age = 23
	person1.job = "Developer"
	person1.salary = 50000

	fmt.Println("Name: ", person1.name)
	fmt.Println("Age: ", person1.age)
	fmt.Println("Job: ", person1.job)
	fmt.Println("Salary: ", person1.salary)

	// Another way to initialize a struct
	person2 := Person{"Alice", 30, "Designer", 60000}
	fmt.Println("\nName: ", person2.name)
	fmt.Println("Age: ", person2.age)
	fmt.Println("Job: ", person2.job)
	fmt.Println("Salary: ", person2.salary)

	// Initialize with named fields (order doesn't matter)
	person3 := Person{
		name:   "Charlie",
		salary: 70000,
		age:    35,
		job:    "Manager",
	}
	fmt.Println("\nName: ", person3.name)
	fmt.Println("Age: ", person3.age)
	fmt.Println("Job: ", person3.job)
	fmt.Println("Salary: ", person3.salary)

	// Structs as pointers
	personPtr := &person1
	fmt.Println("\nPerson 1 (via pointer):", personPtr.name)
	personPtr.age = 24 // Modify through pointer
	fmt.Println("Person 1's new age:", person1.age)

	// Anonymous structs
	anonPerson := struct {
		name string
		city string
	}{
		name: "David",
		city: "New York",
	}
	fmt.Println("\nAnonymous Person:", anonPerson.name, "from", anonPerson.city)
}
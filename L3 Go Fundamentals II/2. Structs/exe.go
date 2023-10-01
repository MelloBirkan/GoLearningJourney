package main

import "fmt"

type Student struct {
	id   int8
	name string
}

type Classroom struct {
	id          int8
	capacity    int8
	subject     string
	studentList []Student
}

func main() {
	c1 := Classroom{id: 1, capacity: 40, subject: "Math", studentList: []Student{
		{
			id:   20,
			name: "Eric",
		},
		{
			id:   30,
			name: "Sloan",
		},
	},
	}

	c2 := new(Classroom)
	c2.id = 2
	c2.capacity = 80
	c2.subject = "Law"
	c2.studentList = []Student{{id: 32, name: "Carlo"}, {id: 42, name: "Monte"}}

	fmt.Println(c1)
	fmt.Println(c2)
}

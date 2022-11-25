package main

import "fmt"

type Employee struct {
	LastName string
	Salary   int
}

func sumSalary(vector []Employee) int {
	sum := 0
	for _, employee := range vector {
		sum += employee.Salary
	}
	return sum
}

func main() {
	var employees []Employee

	employees = append(employees, Employee{"Emp1", 10000})
	employees = append(employees, Employee{"Emp2", 12000})
	employees = append(employees, Employee{"Emp3", 14000})
	employees = append(employees, Employee{"Emp4", 16000})
	employees = append(employees, Employee{"Emp5", 20000})

	fmt.Printf("sum of salaries of created %d emplyees: %d", len(employees), sumSalary(employees))
}

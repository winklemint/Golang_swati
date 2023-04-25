package main

import "fmt"

func main() {
	employeeSalary := make(map[string]int)
	fmt.Println(employeeSalary)

	employeeSalary = make(map[string]int)
	employeeSalary["steve"] = 12000
	employeeSalary["Raja"] = 17000
	employeeSalary["Rani"] = 15500
	fmt.Println(employeeSalary)

	employeeSalary = map[string]int{
		"Nick":  220,
		"Jonas": 900,
	}
	employeeSalary["Mike"] = 700
	fmt.Println(employeeSalary)

	employeeSalary = map[string]int{
		"Jamie":  3000,
		"Nikita": 5000,
		"Azhar":  6500,
	}

	employee := "Nikita"
	salary := employeeSalary[employee] //calling the salary of nikita
	fmt.Println("Salary of", employee, "is", salary)

	employeeSalary = map[string]int{
		"Swati":   4700,
		"Swap":    6300,
		"Prateek": 5900,
	}
	fmt.Println("Salary of uma is", employeeSalary["Uma"])

	employeeSalary = map[string]int{
		"hiteshi": 3000,
		"Saroj":   980,
	}
	newEmp := "Vaish"
	value, ok := employeeSalary[newEmp]
	if ok == true {
		fmt.Println("Salary of", newEmp, "is", value)
		return
	}
	fmt.Println(newEmp, "not found")

	fmt.Println("Content of the map")
	for key, value := range employeeSalary {
		fmt.Printf("employee [%s] = Salary %d\n", key, value)
	}
	fmt.Println("Map before deletion", employeeSalary)
	delete(employeeSalary, "Saroj")
	fmt.Println("Map after deletion", employeeSalary)

}

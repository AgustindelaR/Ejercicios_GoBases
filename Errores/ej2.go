package main

import (
	"errors"
	"fmt"
)

func validateSalary(salary int) error {
	if salary <= 10000 {
		return MyError{"Error: salary is less than 10000"}
	}
	return nil
}

func ej2() {
	salary := 9999
	err := validateSalary(salary)
	if errors.Is(err, MyError{}) {
		fmt.Printf("%s\n", err)
	} else {
		fmt.Printf("Salary is valid\n")
	}

}

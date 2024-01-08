package main

type MyError struct {
	msg string
}

func (e MyError) Error() string {
	return e.msg
}

func ej1() {
	salary := 150000

	if salary <= 150000 {
		err := MyError{msg: "Error: the salary entered does not reach the taxable minimum"}
		println(err.Error())
		return
	}
	println("Must pay taxes")
}

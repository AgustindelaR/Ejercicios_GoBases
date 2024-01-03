package main

func ej1() {
	productoUno := Product{
		name:        "Termo",
		description: "Guarda agua",
		category:    "Bebidas",
		price:       5.43}

	productoUno.Save()
	productoUno.GetAll()
	productoUno.GetById(0)
}

func ej2() {
	empleadoNuevo := Employee{
		id:       1234,
		position: "Software Developer",
		person: Person{
			id:          123456789,
			name:        "Agustín",
			dateOfBirth: "14/06/2002",
		},
	}

	empleadoNuevo.PrintEmployee()
}

func main() {
	ej1()
	println("-----------------------")
	ej2()
}

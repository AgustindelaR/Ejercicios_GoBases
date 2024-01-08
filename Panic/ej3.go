package main

import (
	"fmt"
	"os"
)

type Client struct {
	File        string
	Name        string
	Home        string
	Id          int
	PhoneNumber int
}

var rutaArchivo = "customers2.txt" // cambiar a que cada cliente tenga la suya

func abrirArchivo() []byte {
	datos, errorRead := os.ReadFile(rutaArchivo)
	if errorRead != nil {
		panic(errorRead)
	}
	return datos
}

func añadirCliente(clientes *[]Client, client *Client) error {
	err := revisarClienteYaExistente(*client, *clientes)
	if err != nil {
		return err
	}
	*clientes = append(*clientes, *client)
	return nil
}

type Myerror struct {
	msg string
}

func (e Myerror) Error() string {
	return e.msg
}

func revisarClienteYaExistente(newClient Client, clientes []Client) error {
	for _, cliente := range clientes {
		if cliente.Name == newClient.Name {
			return Myerror{"Cliente ya existe"}
		}
	}
	return nil
}

func validarCliente(client Client) (bool, error) {
	if client.File == "" {
		return false, Myerror{"Cliente tiene archivo invalido"}
	}
	if client.Name == "" {
		return false, Myerror{"Cliente tiene Nombre invalido"}
	}
	if client.Id == 0 {
		return false, Myerror{"Cliente tiene ID invalido"}
	}
	if client.PhoneNumber == 0 {
		return false, Myerror{"Cliente tiene telefono invalido"}
	}
	if client.Home == "" {
		return false, Myerror{"Cliente tiene home invalido"}
	}

	return true, nil
}

func ej3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("No se encontró el atchivo indicado")
		}
		println("End of Executin")
	}()

	cliente := Client{
		File:        rutaArchivo,
		Name:        "Agustin",
		Id:          1234,
		PhoneNumber: 23497,
		Home:        "Banfield",
	}
	clientes := []Client{cliente}

	clienteDos := Client{
		File:        rutaArchivo,
		Name:        "Diego",
		Id:          1235,
		PhoneNumber: 123432,
		Home:        "Lomas",
	}
	añadirCliente(&clientes, &clienteDos)

	for _, cliente := range clientes {
		os.WriteFile(rutaArchivo, []byte(fmt.Sprintf("%s,", cliente.Name)), 'w')
	}

}

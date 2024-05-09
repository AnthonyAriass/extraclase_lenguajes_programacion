package main

import (
	"fmt"
	"sort"
	"strings"
)

// Definición del tipo infoCliente
type infoCliente struct {
	nombre string
	correo string
	edad   int8
	// Otros campos relevantes
}

// Definición del tipo ListaClientes como un slice de infoCliente
type ListaClientes []infoCliente

// Método agregarCliente para ListaClientes
func (lc *ListaClientes) agregarCliente(nombre, correo string, edad int8) {
	*lc = append(*lc, infoCliente{nombre, correo, edad})
}

// Función listaClientes_ApellidoEnCorreo para filtrar clientes cuyos correos contienen el apellido
func listaClientes_ApellidoEnCorreo(lista *ListaClientes, apellido string) ListaClientes {
	var clientesFiltrados ListaClientes
	apellido = strings.ToLower(apellido)
	for _, cliente := range *lista {
		if strings.Contains(strings.ToLower(cliente.correo), apellido) {
			clientesFiltrados = append(clientesFiltrados, cliente)
		}
	}
	return clientesFiltrados
}

// Función para verificar si el dominio del correo es de Costa Rica
func esCorreoCostaRica(correo string) bool {
	if strings.HasSuffix(correo, ".cr") {
		return true
	}
	return false
}

// Función cantidadCorreosCostaRica para contar la cantidad de clientes cuyos correos son de Costa Rica
func cantidadCorreosCostaRica(lista *ListaClientes) int {
	// Filtrar solo los clientes cuyos correos son de Costa Rica
	clientesFiltrados := filter(*lista, func(cliente interface{}) bool {
		return esCorreoCostaRica(cliente.(infoCliente).correo)
	})

	// Verificar si el resultado de filter es nil
	if clientesFiltrados == nil {
		return 0
	}

	// Contar la cantidad de clientes filtrados
	cantidad := len(clientesFiltrados.([]infoCliente))

	return cantidad
}

// Función filter para filtrar elementos de una lista según un predicado
func filter(slice interface{}, predicate func(interface{}) bool) interface{} {
	switch slice := slice.(type) {
	case ListaClientes: // Corregido el tipo de slice esperado
		result := make([]infoCliente, 0) // Cambio de []infoCliente a []infoCliente
		for _, v := range slice {
			if predicate(v) {
				result = append(result, v)
			}
		}
		return result
	}
	return nil
}

// Función mapFunc para aplicar una función a cada elemento de una lista
func mapFunc(slice interface{}, mapper func(interface{}) interface{}) interface{} {
	switch slice := slice.(type) {
	case []infoCliente:
		result := make([]interface{}, len(slice))
		for i, v := range slice {
			result[i] = mapper(v)
		}
		return result
	}
	return nil
}

// Función reduce para reducir una lista a un solo valor
func reduce(slice interface{}, reducer func(interface{}, interface{}) interface{}, initial interface{}) interface{} {
	switch slice := slice.(type) {
	case []interface{}:
		if slice == nil || len(slice) == 0 {
			return initial
		}
		result := initial
		for _, v := range slice {
			result = reducer(result, v)
		}
		return result
	}
	return nil
}

func clientesSugerenciasCorreos(lista *ListaClientes) ListaClientes {
	var clientesConSugerencias ListaClientes

	for _, cliente := range *lista {
		// Dividir el correo en la parte local y el dominio
		partesCorreo := strings.Split(cliente.correo, "@")
		parteLocal := partesCorreo[0]
		dominio := partesCorreo[1]

		// Verificar si el apellido está contenido en la parte local del correo
		if !strings.Contains(strings.ToLower(parteLocal), strings.ToLower(cliente.nombre)) {
			// Generar una sugerencia de correo utilizando la primera letra del nombre y el apellido completo
			partesNombre := strings.Split(cliente.nombre, " ")
			sugerenciaCorreo := strings.ToLower(string(partesNombre[0][0])) + strings.ToLower(partesNombre[len(partesNombre)-1]) + "@" + dominio
			// Agregar el cliente con la sugerencia de correo a la lista
			clientesConSugerencias = append(clientesConSugerencias, infoCliente{
				nombre: cliente.nombre,
				correo: sugerenciaCorreo,
				edad:   cliente.edad,
			})
		} else {
			// Mantener el cliente original en la lista
			clientesConSugerencias = append(clientesConSugerencias, cliente)
		}
	}

	return clientesConSugerencias
}

// Función correosOrdenadosAlfabeticos para devolver una lista nueva con todos los correos de clientes ordenados alfabéticamente
func correosOrdenadosAlfabeticos(lista *ListaClientes) []string {
	var correos []string

	// Recorrer la lista de clientes y agregar los correos a la lista
	for _, cliente := range *lista {
		correos = append(correos, cliente.correo)
	}

	// Ordenar la lista de correos alfabéticamente
	sort.Strings(correos)

	return correos
}

func main() {
	// Inicialización de una ListaClientes
	listaClientes := ListaClientes{
		{nombre: "Cliente1", edad: 30, correo: "cliente1@example.com"},
		{nombre: "Cliente2", edad: 35, correo: "cliente2@example.com"},
		// Puedes agregar más clientes aquí
	}

	// Agregar nuevos clientes usando el método
	listaClientes.agregarCliente("Oscar Viquez", "oviquez@tec.ac.cr", 44)
	listaClientes.agregarCliente("Pedro Perez", "elsegundo@gmail.com", 30)
	listaClientes.agregarCliente("Maria Lopez", "mlopez@hotmail.com", 18)
	listaClientes.agregarCliente("Juan Rodriguez", "jrodriguez@gmail.com", 25)
	listaClientes.agregarCliente("Luisa Gonzalez", "muyseguro@tec.ac.cr", 67)
	listaClientes.agregarCliente("Marco Rojas", "loquesea@hotmail.cr", 47)
	listaClientes.agregarCliente("Marta Saborio", "msaborio@gmail.com", 33)
	listaClientes.agregarCliente("Camila Segura", "csegura@ice.co.cr", 19)
	listaClientes.agregarCliente("Fernando Rojas", "frojas@estado.gov", 27)
	listaClientes.agregarCliente("Rosa Ramirez", "risuenna@gmail.com", 50)

	// Mostrar la lista de clientes actualizada
	fmt.Println("Lista de clientes actualizada:")
	for _, cliente := range listaClientes {
		fmt.Println(cliente)

	}

	// Filtrar clientes cuyos correos contienen el apellido "Rojas"
	apellido := "Rojas"
	clientesFiltrados := listaClientes_ApellidoEnCorreo(&listaClientes, apellido)

	// Mostrar la lista de clientes filtrada
	fmt.Printf("\nClientes cuyos correos contienen el apellido '%s':\n", apellido)
	for _, cliente := range clientesFiltrados {
		fmt.Println(cliente)
	}

	// Obtener la cantidad de clientes cuyos correos son de Costa Rica
	cantidad := cantidadCorreosCostaRica(&listaClientes)
	fmt.Printf("\nCantidad de clientes cuyos correos son de Costa Rica: %d\n", cantidad)

	// Mostrar la lista de clientes antes de aplicar sugerencias de correos
	fmt.Println("\nLista de clientes antes de aplicar sugerencias de correos:\n-----------------------")
	for _, cliente := range listaClientes {
		fmt.Println("Nombre:", cliente.nombre, "Correo:", cliente.correo)
		fmt.Println()
	}

	// Aplicar sugerencias de correos
	clientesConSugerencias := clientesSugerenciasCorreos(&listaClientes)

	// Mostrar la lista de clientes con sugerencias de correos
	fmt.Println("\nLista de clientes con sugerencias de correos:\n-----------------------")
	for _, cliente := range clientesConSugerencias {
		fmt.Println("Nombre:", cliente.nombre, "Correo sugerido:", cliente.correo)
		fmt.Println()
	}

	// Obtener correos ordenados alfabéticamente
	correosOrdenados := correosOrdenadosAlfabeticos(&listaClientes)

	// Mostrar los correos ordenados
	fmt.Println("Correos ordenados alfabéticamente:")
	for _, correo := range correosOrdenados {
		fmt.Println(correo)
	}
}

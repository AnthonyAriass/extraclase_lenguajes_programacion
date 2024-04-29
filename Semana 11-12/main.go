package main

import (
	"fmt"
	"strconv"
)

type producto struct {
	nombre   string
	cantidad int
	precio   int
}

type listaProductos []producto

var lProductos listaProductos

const existenciaMinima int = 10 //la existencia mínima es el número mínimo debajo de el cual se deben tomar eventuales desiciones

func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {

	// modificar el código para que cuando se agregue un producto, si este ya se encuentra, incrementar la cantidad
	// de elementos del producto y eventualmente el precio si es que es diferente

	//Código Anthony -------------------------------------------------------

	for i := range *l {
		if (*l)[i].nombre == nombre {
			(*l)[i].cantidad += cantidad
			(*l)[i].precio = precio
			return
		}
	}

	//-----------------------------------------------------------------------

	//Si no se encontró un artículo con el mismo nombre, se crea uno y se agrega
	*l = append(*l, producto{nombre: nombre, cantidad: cantidad, precio: precio})
}

// Hacer una función adicional que reciba una cantidad potencialmente infinita de productos previamente creados y los agregue a la lista

//Código Anthony -------------------------------------------------------

func (l *listaProductos) agregarProductosInf(productos ...producto) {
	for _, p := range productos {
		// Iterar sobre la lista de productos
		for i := range *l {
			// Si encontramos un producto con el mismo nombre
			if (*l)[i].nombre == p.nombre {
				// Incrementar la cantidad del producto existente o el precio
				(*l)[i].cantidad += p.cantidad
				(*l)[i].precio = p.precio
				break
			}
		}

		// Si no se encontró un producto con el mismo nombre, agregar uno nuevo
		*l = append(*l, p)
	}
}

//----------------------------------------------------------------------

func (l *listaProductos) buscarProducto(nombre string) (*producto, int) { //el retorno es el índice del producto encontrado y -1 si no existe
	var err = -1
	var p *producto = nil
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			err = 0
			p = &((*l)[i])
		}
	}
	return p, err
}

//modifique la función para que esta vez solo retorne el producto como tal y que retorne otro valor adicional "err" conteniendo un 0 si no hay error y números mayores si existe algún
//tipo de error como por ejemplo que el producto no exista. Una vez implementada la nueva función, cambie los usos de la anterior función en funciones posteriores, realizando los cambios
//que el uso de la nueva función ameriten

//Código Anthony -------------------------------------------------------

func (l *listaProductos) buscarProducto_Err(nombre string) (*producto, int, int) {

	var p *producto
	var err = -1
	var ind = 0

	for i := 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			p = &((*l)[i])
			err = 0
			ind = i
			break
		}
	}

	// Retornamos el producto y el código de error
	return p, err, ind
}

//NOTA: Agregué un índice para que después sea más fácil eliminar elementos de la lista
//----------------------------------------------------------------------

func (l listaProductos) venderProducto(nombre string, cant int) {

	var prod, err, ind = l.buscarProducto_Err(nombre)

	if err == 0 {
		//validar si hay unidades de las que se quieren comprar
		if (*prod).cantidad >= cant {

			(*prod).cantidad -= cant
			//modificar para que cuando no haya existencia de cantidad de productos, el producto se elimine de "la lista" y notificar dicha acción
			//modifique posteriormente para que se ingrese de parámetro la cantidad de productos a vender

			//Código Anthony -------------------------------------------------------

			// Eliminar el producto de la lista si la cantidad llega a cero
			if prod.cantidad == 0 {
				l = append((l)[:ind], (l)[ind+1:]...)
				fmt.Println("Producto ", nombre, " eliminado de la lista")
			}
		} else {
			fmt.Println("ERROR: Solo hay " + strconv.Itoa((*prod).cantidad) + " unidades en venta")
		}

	} else {
		fmt.Println("ERROR: Producto no encontrado")
	}

	//----------------------------------------------------------------------

}

//haga una función para a partir del nombre del producto, se pueda modificar el precio del mismo Pero utilizando la función buscarProducto MODIFICADA PREVIAMENTE

//Código Anthony -------------------------------------------------------

// NOTA: Ya agregarProductosInf() actualiza el precio de un producto en caso de que ya exista previamente.

//----------------------------------------------------------------------

func llenarDatos() {
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("frijoles", 4, 2000)
	lProductos.agregarProducto("leche", 8, 1200)
	lProductos.agregarProducto("café", 12, 4500)

}

func (l *listaProductos) listarProductosMinimos() listaProductos {
	// debe retornar una nueva lista con productos con existencia mínima NOTA: Ya este ejercicio estaba hecho previamente
	newSlice := make(listaProductos, 0)

	for _, p := range *l {
		if p.cantidad <= existenciaMinima {
			newSlice = append(newSlice, p)
		}
	}
	return newSlice
}

func main() {
	llenarDatos()
	fmt.Println(lProductos)
	//venda productos
	lProductos.venderProducto("arroz", 10)
	lProductos.venderProducto("frijoles", 2)
	fmt.Println(lProductos)

	pminimos := lProductos.listarProductosMinimos()
	fmt.Println(pminimos)
	println(pminimos)

}

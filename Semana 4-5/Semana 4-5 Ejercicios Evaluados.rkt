
;;Esta porción de Código es la vista en clase ------------------------------------------

(define listaProductos ;define define funciones, en este caso solo es una declaración
  '(("Arroz" 8 1800) ("Frijoles" 5 1200) ("Azúcar" 6 1100) ("Café" 2 2800) ("Leche" 6 1800)))

(define (agregarProducto Lista nombre cantidad precio)
  (cond ((null? Lista)
         (list (list nombre cantidad precio)))
        ((equal? nombre (caar Lista))
         (cons (list
                (caar Lista)                ;construya una lista con nombre
                (+ (cadar Lista) cantidad)  ;el car del cdr del car ó (list-ref () ())
                precio)                     ;
                (cdr Lista)))               ; 
        (else
         (cons (car Lista) (agregarProducto (cdr Lista) nombre cantidad precio)))))

;;Esta función no será usada para efectos prácticos del ejercicio en cambio se usará la que está más abajo.
(define (venderProductoNOSEUSA Lista nombre cantidad)
  (cond ((null? Lista)
         (display "No existe ese producto para vender")
         '())
         (list (list nombre cantidad))
         ((equal? nombre (caar Lista))
          (cons (list
                 (caar Lista)
                 (- (list-ref (car Lista) 1) cantidad)
                 (list-ref (car Lista) 2))
                 (cdr Lista)))
         (else
          (cons (car Lista) 
                (venderProducto (cdr Lista) nombre cantidad)))))

(define (existenciasMinimas Lista cantidad)
  (filter (lambda (x) (<= (cadr x) cantidad))
          Lista))

;--------------------------------------------------------------------------------

;Ejercicio 1 (EVALUADO)

;Primero vamos a imprimir la lista antes de hacer cualquier compra para asegurarnos de que sí se modifica
(display "-----------------------------------\nLista Original:\n")
(display listaProductos)
(display "\n-----------------------------------\n\n\n")

;; Función auxiliar para calcular el total de la venta
(define (calcularTotal productosVenta)
  (if (null? productosVenta)
      0
      (+ (* (list-ref (car productosVenta) 1) (list-ref (car productosVenta) 2))
         (calcularTotal (cdr productosVenta)))))

;; Función auxiliar para generar el detalle de la venta
(define (generarDetalleVenta productosVenta)
  (if (null? productosVenta)
      '()
      (list (list (caar productosVenta)
                  (list-ref (car productosVenta) 1)
                  (list-ref (car productosVenta) 2))
            (generarDetalleVenta (cdr productosVenta)))))

;; Función auxiliar para calcular el impuesto del 13%
(define (calcularImpuesto total montoLimite)
  (if (> total montoLimite)
      (* total 0.13) ; Aplicar el 13% de impuesto si el total supera el monto límite
      0))           ; No aplicar impuesto si el total no supera el monto límite

;; Función para generar la factura, incluyendo impuesto si es necesario
(define (generarFactura listaProductos listaVenta montoLimiteImpuesto)
  ;; Calcular el total y el detalle de la venta
  (define total (calcularTotal listaVenta))
  (define detalleVenta (generarDetalleVenta listaVenta))
  ;; Calcular el impuesto
  (define impuesto (calcularImpuesto total montoLimiteImpuesto))
  ;; Crear la estructura de la factura con productos, total y, si aplica, impuesto
  (list 'factura
        (list 'productos detalleVenta)
        (list 'total total)
        (if (> impuesto 0)
            (list 'impuesto impuesto)
            '())))

;; Función auxiliar para vender un producto específico de la lista
(define (venderProductoAux lista nombre cantidad)
  (cond ((null? lista)
         (display "No existe ese producto para vender\n\n") ; Mensaje si el producto no está en la lista
         '())
        ((equal? nombre (caar lista))
         (let* ((cantidadExistente (list-ref (car lista) 1))
                (productoVendido (list nombre cantidad (list-ref (car lista) 2))))
           (if (<= cantidadExistente cantidad)
               (begin
                 (display "No hay suficiente cantidad en inventario para vender\n\n")
                 lista) ; Devolver la lista original si no hay suficiente cantidad en inventario
               (let* ((nuevaListaProductos (cons (list (caar lista)
                                                        (- cantidadExistente cantidad)
                                                        (list-ref (car lista) 2))
                                                  (cdr lista)))
                      (factura (generarFactura nuevaListaProductos (list productoVendido) 3000)));; Este valor es el monto límite desde el cuál se empezará a calcular el impuesto
                 (display "Venta realizada exitosamente. Detalle de la factura:\n")
                 (display factura)
                 (display "\n\n")
                 nuevaListaProductos)))) ; Devuelve la lista actualizada después de la venta
        (else
         (cons (car lista)
               (venderProductoAux (cdr lista) nombre cantidad))))) ; Mantener el producto sin modificar en la lista

;; Función principal para vender una lista de productos
(define (venderProductos lista productosVenta)
  (if (null? productosVenta)
      (begin
        (display "La lista de productos a vender está vacía\n\n")
        lista)  ; Devolver la lista original si no hay productos para vender
      ; Iterar sobre la lista de productos a vender
      (let loop ((listaProductos lista)
                 (productosRestantes productosVenta))
        (if (null? productosRestantes)
            ; Devolver la lista de productos actualizada
            listaProductos
            ; Actualizar la lista de productos para el producto actual
            (let ((producto (car productosRestantes)))
              (set! listaProductos (venderProductoAux listaProductos (car producto) (cadr producto)))
              ; Llamar recursivamente con el resto de la lista de productos a vender
              (loop listaProductos (cdr productosRestantes)))))))


;;Acá se pueden enviar los artículos que se deseen comprar como argumento
(define listaProductos (venderProductos listaProductos '(("Arroz" 2) ("Frijoles" 2)("Azúcar" 7))))


;Ahora vamos a imprimir la lista con las compras realizadas
(display "\n-----------------------------------\nLista Modificada:\n")
(display listaProductos)
(display "\n-----------------------------------\n\n\n")

;Ejercicio 2 (EVALUADO)

;; Función que verifica si una subcadena está presente en una cadena principal
(define (string-contains str sub)
  (regexp-match? (regexp sub) str))

;; Función que filtra una lista de cadenas para incluir solo aquellas que contienen una subcadena específica
(define (filtrar-por-substring lista substring)
  (filter (lambda (cadena) (string-contains cadena substring)) lista))

;Pruebas Ejercicio 2
(filtrar-por-substring '("la casa" "el perro" "pintando la cerca" "carga la mesa") "la")

(filtrar-por-substring '("aroma" "prado" "se rompió el aro" "caro") "aro")


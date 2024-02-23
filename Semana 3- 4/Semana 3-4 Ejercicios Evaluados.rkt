
;Ejercicio 1 (EVALUADO)

(define (interes capital porcentaje años)
  (cond
    [(> 0 años) (display "Number of years must be 1 or bigger")]
    [(= años 1) (display capital)]
    [else (interes (+ capital (* capital porcentaje)) porcentaje (- años 1))]
  )
 )

;Pruebas Ejercicio 1
(display "\n")
(interes 2000 0.1 0)
(display "\n")
(interes 2000 0.1 2)
(display "\n")
(interes 2000 0.1 4)

;-----------------------------------------

;Ejercicio 6 (EVALUADO)
(define (merge lista1 lista2)
  (cond
    [(empty? lista1) lista2]
    [(empty? lista2) lista1]
    ;Esta validación es la que se encarga de ordenar la lista final
    ;Comparando las cabezas de ambas listas.
    [(< (car lista1) (car lista2))
     (cons (car lista1) (merge (cdr lista1) lista2))]
    [else
     (cons (car lista2) (merge lista1 (cdr lista2)))]))

;Pruebas Ejercicio 6
(display "\n")
(display "\n")
(display (merge '(2 4 6 8 10) '(1 3 5 7 9 11)))
(display "\n")
(display (merge '(1 2 3 4) '(1 3 5 6)))

;-----------------------------------------

;Ejercicio 8 (EVALUADO)
(define (sub-conjunto? lista1 lista2)
  (cond
    ((null? lista1) #t)  ; Si la primera lista está vacía es un subconjunto
    ((null? lista2) #f)  ; Si la segunda está vacía la primera no puede ser subconjunto
    ((member (car lista1) lista2) ; Verifica si el primer elemento de list1 está en list2
     (sub-conjunto? (cdr lista1) lista2));En caso de que sí sigue revisando con la cola.
    (else #f)))


;Pruebas Ejercicio 8
(display "\n")
(display "\n")
(sub-conjunto? '(1 2 3) '(1 2 3 4 5))
(display "\n")
(sub-conjunto? '() '(1 2))
(display "\n")
(sub-conjunto? '(1 2 3 4) '(1 0 3 5))

;-----------------------------------------

;Ejercicio 9 (EVALUADO)
(define (eliminar-elemento lista elemento)
  (filter (lambda (x) (not (null? x))) ; Filtra las listas vacías
          (map (lambda (x) 
                 (if (equal? x elemento) 
                     '() ;El lambda verifica si el elemento es igual a x y lo reemplaza por una lista vacía
                     (list x)))
               lista)))


;Pruebas Ejercicio 9
(display "\n")
(display "\n")
(eliminar-elemento '(1 2 3) 3)
(display "\n")
(eliminar-elemento '(1 2 3 4 5) 4)






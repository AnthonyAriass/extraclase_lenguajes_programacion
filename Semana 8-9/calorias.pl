% Definición de las calorías de cada alimento
calorias(guacamole, 200).
calorias(ensalada, 150).
calorias(consome, 300).
calorias(tostadas_caprese, 250).
calorias(filete_cerdo, 400).
calorias(pollo_horno, 280).
calorias(carne_salsa, 320).
calorias(tilapia, 160).
calorias(salmon, 300).
calorias(trucha, 225).
calorias(flan, 200).
calorias(nueces_con_miel, 500).
calorias(naranja_confitada, 450).
calorias(flan_coco, 375).

% Definición de las categorías de alimentos: entrada, carne, pescado, postre

entrada(guacamole).
entrada(ensalada).
entrada(consome).
entrada(tostadas_caprese).

carne(filete_cerdo).
carne(pollo_horno).
carne(carne_salsa).

pescado(tilapia).
pescado(salmon).
pescado(trucha).

postre(flan).
postre(nueces_con_miel).
postre(naranja_confitada).
postre(flan_coco).

% Definición de una combinación válida de alimentos con sus calorías totales
combinacion(Entrada, Carne, Postre, Calorias) :-
    entrada(Entrada),
    carne(Carne),
    postre(Postre),
    calorias(Entrada, CaloriasEntrada),
    calorias(Carne, CaloriasCarne),
    calorias(Postre, CaloriasPostre),
    Calorias is CaloriasEntrada + CaloriasCarne + CaloriasPostre.

% Predicado para encontrar las primeras 5 combinaciones que no exceden un límite de calorías dado
combinaciones_hasta(CaloriasMax, Combinaciones) :-
    findall((Entrada, Carne, Postre, Calorias),
            (combinacion(Entrada, Carne, Postre, Calorias),
            Calorias =< CaloriasMax),
            Combinaciones),
    length(Combinaciones, NumCombinaciones),
    length(Combinaciones5, 5),
    append(Combinaciones5, _, Combinaciones),
    NumCombinaciones >= 5.


% Ejemplo para probar: combinaciones_hasta(1000, Combinaciones).

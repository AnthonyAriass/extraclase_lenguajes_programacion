% Predicado para calcular la distancia de Hamming entre dos strings
hamming_distance(String1, String2, Distance) :-
    atom_chars(String1, List1),
    atom_chars(String2, List2),
    min_length(List1, List2, Min),
    hamming_distance_aux(List1, List2, Min, Distance).

% Predicado auxiliar para calcular la distancia de Hamming entre dos listas
hamming_distance_aux(_, _, 0, 0).
hamming_distance_aux([X|Xs], [X|Ys], N, Distance) :-
    N > 0,
    N1 is N - 1,
    hamming_distance_aux(Xs, Ys, N1, Distance). % Llama recursivamente con las colas de las listas
hamming_distance_aux([_|Xs], [_|Ys], N, Distance) :-
    N > 0,
    N1 is N - 1,
    hamming_distance_aux(Xs, Ys, N1, SubDistance),
    Distance is SubDistance + 1.

% Predicado para obtener la longitud mínima entre dos listas
min_length(List1, List2, Min) :-
    length(List1, Length1),
    length(List2, Length2),
    Min is min(Length1, Length2).


% Para probar
% hamming_distance('romano', 'ron', Distance).

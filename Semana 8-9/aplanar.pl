% Aplanar
flatten([], []).

% Si la cabeza de la lista es a su vez una lista, aplana esa lista y luego aplana la cola.
flatten([X|Xs], FlatList) :-
    is_list(X),
    flatten(X, FlatX),
    flatten(Xs, FlatXs),
    append(FlatX, FlatXs, FlatList).

% Si la cabeza no es una lista, se mantiene tal como está y se aplana la cola.
flatten([X|Xs], [X|FlatXs]) :-
    \+ is_list(X),
    flatten(Xs, FlatXs).

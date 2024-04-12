% Subconjunto
is_subset([], _).

is_subset([X|Xs], Ys) :-
    member(X, Ys),
    is_subset(Xs, Ys).

# Parole
Questo repository contiene il progetto d'esame dell'insegnamento di **"Algoritmi e Strutture Dati" presso l'Università degli Studi di Milano all'a.a 2024/2025**.


Obiettivo del progetto è gestire un **dizionario di parole e schemi**.

## Parole e catene di parole
Una parola è una sequenza finita di caratteri appartenenti all’alfabeto inglese minuscolo {a, b, c, . . . , y,
z}.
Definiamo operazioni elementari di editing su una parola x le seguenti:

• Inserzione di un carattere in qualsiasi posizione in x. Ad esempio, “pippo” diventa “pioppo” tramite
inserzione di “o”.

• Cancellazione di un carattere in qualsiasi posizione in x. Ad esempio, “capra” diventa “capa”
tramite cancellazione di “r”.

• Sostituzione di un carattere con un altro in qualsiasi posizione in x. Ad esempio, “cane” diventa
“rane” tramite la sostituzione di “c” con “r”.

• Scambio della posizione di due caratteri adiacenti qualsiasi in x. Ad esempio “trota” diventa “torta”
scambiando di posizione “r” con “o”.

Date due parole x e y la loro distanza di editing è data dal minor numero di operazioni elementari di
editing necessarie per passare da x a y. Ad esempio, “cavolo” e “cavallo” hanno distanza 2, in quanto
si passa dalla prima alla seconda sostituendo la prima “o” con una “a” e inserendo una “l”; la distanza
fra “capra” e “arpa” è 2, in quanto occorre cancellare la “c” e scambiare la posizione di “p” e “r”; la
distanza fra “pesce” e “sedia” è 4 (si noti che ci sono più modi per passare da “pesce” a “sedia” con 4
operazioni elementari).

Due parole che hanno distanza di editing pari a 1 sono dette simili.

Una catena tra due parole x e y è una sequenza di parole che inizia con x, finisce con y e tale che la
distanza di editing sia 1 per ogni coppia di parole consecutive nella sequenza.

Un gruppo è un insieme massimale di parole del dizionario che possono essere trasformate l’una nell’altra
con una catena di parole tutte interne al gruppo.

## Schemi
Uno schema è una sequenza finita di caratteri appartenenti all’alfabeto inglese {a, b, c, . . . , y, z } ∪ {A,
B, C, . . . , Y, Z}, che contiene almeno una lettera maiuscola in {A, B, C, . . . , Y,Z}.
3

Un’assegnazione è una funzione σ da {A, B, C, . . . , Z} a {a, b, c, . . . , z}. In altri termini, a ogni lettera
maiuscola l’assegnazione associa una lettera minuscola.

Dato uno schema S = α1 . . . αn e un’assegnazione σ, denotiamo con σ(S) la parola β1 . . . βn tale che, per
ogni 1 ≤ i ≤ n,: se αi è una lettera maiuscola allora βi = σ(αi); se αi è una lettera minuscola allora βi = αi.
Una parola x è compatibile con uno schema S se esiste un’assegnazione σ tale che x = σ(S).

## Specifiche di progettazione
alcune delle operazioni nel progetto:
- **crea ()**: Crea un nuovo dizionario vuoto (eliminando l’eventuale dizionario gi`a esistente).
- **carica (file)**: Inserisce nel dizionario le parole e/o gli schemi contenuti nel file di nome file; file è di tipo testo e le parole / gli schemi sono separati da uno o più caratteri di spaziatura (compresi tabulatori e newline). Se file non esiste non viene eseguita alcuna operazione.
- **stampa parole ()**: Stampa tutte le parole del dizionario.
- **stampa schemi ()**: Stampa tutti gli schemi del dizionario.
- **inserisci (w)**: Inserisce nel dizionario la parola / lo schema w; se w è già presente non viene eseguita alcuna operazione.
- **elimina (w)**: Elimina dal dizionario la parola / lo schema w; se w non è nel dizionario non viene eseguita alcunaoperazione.
- **ricerca (S)**: Stampa lo schema S e poi l’insieme di tutte le parole nel dizionario che sono compatibili con lo schema S.
- **distanza (x, y)**: Stampa la distanza di editing fra le due parole x e y.
- **catena (x, y)**: Stampa una catena di lunghezza minima tra x e y di parole nel dizionario. Se tale catena non esiste o se x o y non sono nel dizionario, stampa “non esiste”.

package main

/*
Progetto "Parole" appello del 16 giugno 2025 per il corso Algoritmi e Strutture Dati
Matricola: 13375A
Studente: Lonigro Luigi
*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Tipo dizionario: contiene le parole, gli schemi e il grafo delle parole con distanza 1
type dizionario struct {
	parole map[string]bool     // insieme delle parole inserite
	schemi map[string]bool     // insieme degli schemi inseriti
	grafoP map[string][]string // grafo delle parole, con archi tra parole a distanza 1
}

// Funzione main: legge continuamente righe dallo stdin ed esegue comandi finché non arriva un comando 't'
func main() {
	var d dizionario
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		line := scanner.Text()
		if len(line) == 0 {
			os.Exit(0) // Termina se riga vuota
		}

		// Comando 'c' singolo: crea nuovo dizionario
		if line[0] == 'c' && len(line) == 1 {
			d = newDizionario()
		} else {
			esegui(d, line)
		}
	}
}

//----------------------------------------------------- COSTRUTTORE DELLA STRUTTRA ---------------------------------------------------------------------------------------------//

// Crea e restituisce un nuovo dizionario vuoto
func newDizionario() dizionario {
	return dizionario{
		parole: make(map[string]bool),
		schemi: make(map[string]bool),
		grafoP: make(map[string][]string),
	}
}

//----------------------------------------------------- FUNZIONI ELEMENTARI ---------------------------------------------------------------------------------------------//

// Stampa tutte le parole presenti nel dizionario
func (d *dizionario) stampa_parole() {
	fmt.Println("[")
	for k := range d.parole {
		fmt.Println(k)
	}
	fmt.Println("]")
}

// Stampa tutti gli schemi presenti nel dizionario
func (d *dizionario) stampa_schemi() {
	fmt.Println("[")
	for k := range d.schemi {
		fmt.Println(k)
	}
	fmt.Println("]")
}

// Elimina una parola o uno schema dal dizionario
func (d *dizionario) elimina(word string) {
	if isParola(word) {
		delete(d.parole, word)
		d.grafoP = nil //resetta il grafo poiché c'è stata una modifica al dizionario
	} else {
		delete(d.schemi, word)
	}
}

// Inserisce una parola o uno schema nel dizionario
func (d *dizionario) inserisci(word string) {
	if isParola(word) {
		d.parole[word] = false
		d.grafoP = nil //resetta il grafo poiché c'è stata una modifica
	} else {
		d.schemi[word] = false
	}
}

// Carica parole e/o schemi da file di testo
func (d *dizionario) carica(file string) {
	text, err := os.Open(file)
	if err != nil {
		fmt.Println("file non caricato correttamente")
		return
	}
	defer text.Close()

	scanner := bufio.NewScanner(text)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		d.inserisci(line)
	}

}

// Ritorna true se la stringa è una parola (tutta minuscola), altrimenti è uno schema
func isParola(line string) bool {
	lettere := []rune(line)
	for i := 0; i < len(line); i++ {
		if lettere[i] >= 65 && lettere[i] <= 90 {
			return false
		}
	}
	return true
}

//----------------------------------------------------- FUNZIONE DI GESTIONE DEI COMANDI ---------------------------------------------------------------------------------------------//

// Interpreta la stringa in input e chiama la funzione corrispondete
func esegui(d dizionario, s string) {
	data := strings.Split(s, " ")
	if len(data) == 0 || data[0] == "" {
		fmt.Println("Nessun comando inserito.")
		return
	}
	switch data[0] {
	case "c":
		if len(data) == 2 {
			d.carica(data[1])
		} else {
			d.catena(data[1], data[2])
		}
	case "p":
		d.stampa_parole()
	case "s":
		d.stampa_schemi()
	case "i":
		if len(data) < 2 {
			fmt.Println("Errore: comando 'i' richiede una parola o schema.")
			return
		}
		d.inserisci(data[1])
	case "e":
		if len(data) < 2 {
			fmt.Println("Errore: comando 'e' richiede una parola o schema.")
			return
		}
		d.elimina(data[1])
	case "r":
		if len(data) < 2 {
			fmt.Println("Errore: comando 'r' richiede uno schema.")
			return
		}
		d.ricerca(data[1])
	case "d":
		if len(data) < 3 {
			fmt.Println("Errore: comando 'd' richiede due parole.")
			return
		}
		fmt.Println(distanza(data[1], data[2]))
	case "t":
		os.Exit(0)
	default:
		fmt.Println("Comando non riconosciuto. Usa: c, p, s, i, e, r, d, t.")
	}

}

//----------------------------------------------------- FUNZIONI PRINCIPALI ---------------------------------------------------------------------------------------------//

// Stampa le parole compatibili con lo schema dato
func (d *dizionario) ricerca(schema string) {
	result := []string{}
	for parola := range d.parole {
		if compatibileConSchema(schema, parola) {
			result = append(result, parola)
		}
	}

	fmt.Print(schema, ":[", "\n")
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
	fmt.Println("]")
}

// Cerca una catena tra due parole usando BFS sul grafo delle parole simili
func (d *dizionario) catena(x, y string) {
	//controlla che la parola esista nel dizionario
	if _, ok := d.parole[x]; !ok {
		fmt.Println("non esiste")
		return
	}

	//controlla che la parola esista nel dizionario
	if _, ok := d.parole[y]; !ok {
		fmt.Println("non esiste")
		return
	}
	//se il grafo non è stato inizializzato viene creato ora
	if d.grafoP == nil || len(d.grafoP) == 0 {
		d.costruisciGrafo()
	}

	listaDiEsplorazione := []string{x}
	nodiVisitati := map[string]bool{x: true}
	predecessori := map[string]string{}

	//finché abbiamo nodi da esplorare
	for len(listaDiEsplorazione) > 0 {
		nodoCorrente := listaDiEsplorazione[0]
		listaDiEsplorazione = listaDiEsplorazione[1:]

		//se siamo arrivati alla fine del cammino
		if nodoCorrente == y {
			break
		}

		//controlliamo tutti i nodi vicini al nodo corrente, cioè alla testa della lista di esplorazione
		for _, vicini := range d.grafoP[nodoCorrente] {
			if !nodiVisitati[vicini] {
				nodiVisitati[vicini] = true
				predecessori[vicini] = nodoCorrente
				listaDiEsplorazione = append(listaDiEsplorazione, vicini)
			}
		}
	}

	//se non abbiamo trovato una strada per arrivare alla fine del cammino
	if !nodiVisitati[y] {
		fmt.Println("non esiste")
		return
	}

	//ripercorriamo la strada dalla fine all'inizio e salviamo nella slice al contrario così che l'ordine sia dall'inizio alla fine
	catena := []string{}
	for nodo := y; nodo != ""; nodo = predecessori[nodo] {
		catena = append([]string{nodo}, catena...)
	}

	//stampa secondo richiesta del progetto
	fmt.Println("(")
	for i := 0; i < len(catena); i++ {
		fmt.Println(catena[i])
	}
	fmt.Println(")")
}

// Calcola la distanza di Damerau-Levenshtein tra due stringhe
func distanza(str1, str2 string) int {
	// conversione in array di rune
	runeStr1 := []rune(str1)
	runeStr2 := []rune(str2)

	// prende e salva la len delle due parole
	runeStr1len := len(runeStr1)
	runeStr2len := len(runeStr2)
	if runeStr1len == 0 {
		return runeStr2len
	} else if runeStr2len == 0 {
		return runeStr1len
	} else if str1 == str2 {
		return 0
	}

	// Crea alfabeto
	ultimaRiga := make(map[rune]int)
	for i := 0; i < runeStr1len; i++ {
		ultimaRiga[runeStr1[i]] = 0
	}
	for i := 0; i < runeStr2len; i++ {
		ultimaRiga[runeStr2[i]] = 0
	}

	// Crea matrice
	matrice := make([][]int, runeStr1len+2)
	for i := 0; i <= runeStr1len+1; i++ {
		matrice[i] = make([]int, runeStr2len+2)
		for j := 0; j <= runeStr2len+1; j++ {
			matrice[i][j] = 0
		}
	}

	// Distanza max
	maxDist := runeStr1len + runeStr2len

	// Inizializza la matrice
	matrice[0][0] = maxDist
	for i := 0; i <= runeStr1len; i++ {
		matrice[i+1][0] = maxDist
		matrice[i+1][1] = i
	}
	for j := 0; j <= runeStr2len; j++ {
		matrice[0][j+1] = maxDist
		matrice[1][j+1] = j
	}

	var cost int

	for i := 1; i <= runeStr1len; i++ {
		match := 0
		for j := 1; j <= runeStr2len; j++ {
			uO := ultimaRiga[runeStr2[j-1]] //ULTIMA OCCORRENZA
			uM := match                     // ULTIMO MATCH
			if runeStr1[i-1] == runeStr2[j-1] {
				cost = 0
				match = j
			} else {
				cost = 1
			}
			//Sceglie l'azione meno costosa tra: inserimento, eliminazione, sostituzione e trasposizione
			matrice[i+1][j+1] = Min(Min(matrice[i+1][j]+1, matrice[i][j+1]+1), Min(matrice[i][j]+cost, matrice[uO][uM]+(i-uO-1)+1+(j-uM-1)))
		}
		ultimaRiga[runeStr1[i-1]] = i
	}

	return matrice[runeStr1len+1][runeStr2len+1]
}

//----------------------------------------------------- FUNZIONI AUSILIARIE ---------------------------------------------------------------------------------------------//

// Verifica se una parola è compatibile con uno schema
func compatibileConSchema(schema, parola string) bool {
	if len(schema) != len(parola) {
		return false
	}
	mapAssegnazioni := make(map[byte]byte) //Assegna maiuscola a minuscola

	for i := 0; i < len(schema); i++ {
		s := schema[i]
		p := parola[i]

		if s >= 'A' && s <= 'Z' {
			if val, assegnata := mapAssegnazioni[s]; assegnata {
				if val != p {
					return false //violazione mappatura
				}
			} else {
				mapAssegnazioni[s] = p
			}
		} else {
			if s != p {
				return false
			}
		}
	}
	return true
}

// Costruisce il grafo delle parole connesse da distanza di editing 1
func (d *dizionario) costruisciGrafo() {
	d.grafoP = make(map[string][]string)

	parole := make([]string, 0, len(d.parole))
	for p := range d.parole {
		parole = append(parole, p)
	}

	for i := 0; i < len(parole); i++ { //Scorre e trova per ogni parola quali sono le sue parole "simili" cioè con distanza 1
		for j := i + 1; j < len(parole); j++ {
			p1 := parole[i]
			p2 := parole[j]
			if distanza(p1, p2) == 1 {
				d.grafoP[p1] = append(d.grafoP[p1], p2)
				d.grafoP[p2] = append(d.grafoP[p2], p1)
			}
		}
	}
}

// Restituisce il minimo tra due interi
func Min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

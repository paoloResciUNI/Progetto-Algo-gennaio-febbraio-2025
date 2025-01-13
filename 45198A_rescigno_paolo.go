package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

type punto struct {
	coordinataX int
	coordinataY int
	id          string
	richiamo    bool
	successivo  *punto
	precedente  *punto
}

// rappresentazione del grafo che rappresenta il piano
type piano struct {
	inizio *punto
	fine   *punto
}

// Il camapo dove si muovono gli automi
var Campo piano

// Il punto da cui viene il richiamo
var Sorgente *punto

func esegui(p piano, s string) {
	comandi := strings.Split(s, " ")
	switch comandi[0] {
	case "c":
		Campo = newPiano()
	case "S":
		stampa()
	case "s":
		a, _ := strconv.Atoi(comandi[1])
		b, _ := strconv.Atoi(comandi[2])
		stato(a, b)
	case "a":
		a, _ := strconv.Atoi(comandi[1])
		b, _ := strconv.Atoi(comandi[2])
		automa(a, b, comandi[3])
	case "o":
		a, _ := strconv.Atoi(comandi[1])
		b, _ := strconv.Atoi(comandi[2])
		c, _ := strconv.Atoi(comandi[3])
		d, _ := strconv.Atoi(comandi[4])
		ostacolo(a, b, c, d)
	case "p":
		posizioni(comandi[1])
	case "e":
		a, _ := strconv.Atoi(comandi[1])
		b, _ := strconv.Atoi(comandi[2])
		esistePercorso(a, b, comandi[3])
	case "f":
		return
	}
}

func newPiano() piano {
	var nuovPiano piano
	return nuovPiano
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		esegui(Campo, scanner.Text())
	}
}

func stampa() {
	percorrente := new(punto)
	Println("(")
	percorrente = Campo.inizio
	for percorrente != nil && !strings.Contains(percorrente.id, "ostacolo") {
		Printf("%s: %d,%d\n", percorrente.id, percorrente.coordinataX, percorrente.coordinataY)
		percorrente = percorrente.successivo
	}
	Println(")")
	Println("[")
	percorrente = Campo.fine
	for percorrente != nil && strings.Contains(percorrente.id, "ostacolo") {
		x0, y0, x1, y1 := estraiCoordinate(percorrente.id)
		Printf("(%d,%d)(%d,%d)\n", x0, y0, x1, y1)
		percorrente = percorrente.precedente
	}
	Println("]")
}

func stato(x, y int) {
	if dentroAreaOstacolo(x, y) {
		Println("O")
		return
	}
	if Campo.cerca(x, y, "") != nil {
		Println("A")
		return
	} else {
		Println("E")
	}
}

// Aggiunge un automa al campo se le coordinate (x, y) non fanno già parte di un ostacolo, se no non fa nulla
// Se le coordinate (x, y) non fanno parte di un ostacolo allora controlla se l'automa eta esiste già e in caso affermativo
// sposta l'automa altrimenti lo crea nuovo.
func automa(x, y int, eta string) {
	puntoCercato := Campo.cerca(x, y, eta)
	if puntoCercato != nil {
		if strings.Contains(puntoCercato.id, "ostacolo") && !dentroAreaOstacolo(x, y) {
			return
		} else {
			puntoCercato.coordinataX = x
			puntoCercato.coordinataY = y
		}
	} else {
		puntoCercato = new(punto)
		puntoCercato.coordinataX = x
		puntoCercato.coordinataY = y
		puntoCercato.id = eta
		if Campo.inizio == nil {
			Campo.inizio = puntoCercato
			Campo.fine = puntoCercato
			return
		}
		Campo.inizio.precedente = puntoCercato
		puntoCercato.successivo = Campo.inizio
		Campo.inizio = puntoCercato
	}
}

// Un ostacolo è un punto in cui non possono esserci automi
// Ogni ostacolo ha un'adiacenza e ogni punto dell'ostacolo fa parte del campo
// Assunzione: x_1 < x_2 e y_1 < y_2
func ostacolo(x0, y0, x1, y1 int) {
	percorrente := Campo.inizio
	for percorrente != nil && !strings.Contains(percorrente.id, "ostacolo") {
		if (percorrente.coordinataX <= x1 && percorrente.coordinataX >= x0) && (percorrente.coordinataY <= y1 && percorrente.coordinataY >= y0) {
			return
		}
		percorrente = percorrente.successivo
	}
	newOstacolo := new(punto)
	newOstacolo.coordinataX = x0
	newOstacolo.coordinataY = y1
	newOstacolo.id = Sprintf("%d,%d,%d,%d,ostacolo", x0, y0, x1, y1)
	if Campo.fine == nil {
		Campo.fine = newOstacolo
		Campo.inizio = newOstacolo
		return
	}
	newOstacolo.precedente = Campo.fine
	Campo.fine.successivo = newOstacolo
	Campo.fine = newOstacolo
}

func richiamo(x, y int, alpha string) {
	Sorgente = new(punto)
	Sorgente.coordinataX = x
	Sorgente.coordinataY = y
	Sorgente.id = alpha
	percorrente := Campo.inizio
	for percorrente != nil && !strings.Contains(percorrente.id, "ostacolo") {
		if strings.HasPrefix(percorrente.id, alpha) {
			percorrente.richiamo = true
		}
		percorrente = percorrente.successivo
	}
}

func posizioni(alpha string) {
	percorrente := new(punto)
	Println("(")
	percorrente = Campo.inizio
	for percorrente != nil && !strings.Contains(percorrente.id, "ostacolo") {
		if strings.HasPrefix(percorrente.id, alpha) {
			Printf("%s: %d,%d\n", percorrente.id, percorrente.coordinataX, percorrente.coordinataY)
		}
		percorrente = percorrente.successivo
	}
	Println(")")
}

// La funzione prende in input un punto e il nome di un'automa e restituisce "SI" se esiste un percorso di lunghezza D che va dal punto (x, y)
// alla sorgente, "NO" altrimenti. Restituisce "NO" se il punto x, y fa parte di un ostacolo.
func esistePercorso(x, y int, eta string) {
	richiamo(x, y, eta)
	if dentroAreaOstacolo(x, y) {
		Println("NO")
		return
	}
	percorrente := Campo.cerca(x, y, eta)
	if percorrente == nil {
		Println("NO")
		return
	}
	distanza := calcolaDistanza(percorrente.coordinataX, percorrente.coordinataY, x, y)
	percorsoEffettuato, passiMancanti := (avanza(percorrente, distanza))
	if percorsoEffettuato.coordinataX == x && percorsoEffettuato.coordinataY == y && passiMancanti == 0 {
		Println("SI")
		return
	} else {
		Println("NO")
	}
}

func avanza(p *punto, passi int) (*punto, int) {
	var conatoreOstacoliX, contatoreOstacoliY int
	if passi == 0 {
		return p, passi
	}
	possibilePasso := new(punto)
	possibilePasso.coordinataX = p.coordinataX
	possibilePasso.coordinataY = p.coordinataY
	possibilePasso.id = p.id

	sX := passi
	sY := passi
	for possibilePasso.presenzaOstacoloPercorsoX(sY) {
		conatoreOstacoliX++
		sY--
	}
	for possibilePasso.presenzaOstacoloPercorsoY(sX) {
		contatoreOstacoliY++
		sX--
	}
	if p.coordinataX < Sorgente.coordinataX && conatoreOstacoliX >= contatoreOstacoliY {
		possibilePasso.coordinataX = p.coordinataX + 1
		passi--
		return avanza(possibilePasso, passi)
	} else if p.coordinataX > Sorgente.coordinataX && conatoreOstacoliX >= contatoreOstacoliY {
		possibilePasso.coordinataX = p.coordinataX - 1
		passi--
		return avanza(possibilePasso, passi)
	}
	if p.coordinataY < Sorgente.coordinataY {
		possibilePasso.coordinataY = p.coordinataY + 1
		passi--
		return avanza(possibilePasso, passi)
	} else if p.coordinataY > Sorgente.coordinataY {
		possibilePasso.coordinataY = p.coordinataY - 1
		return avanza(possibilePasso, passi)
	}

	passi--
	return avanza(p, passi)
}

func (*piano) cerca(x, y int, id string) *punto {
	percorrente := Campo.inizio
	for percorrente != nil {
		if (percorrente.coordinataX == x && percorrente.coordinataY == y) || percorrente.id == id {
			return percorrente
		}
		percorrente = percorrente.successivo
	}
	return nil
}

func dentroAreaOstacolo(x, y int) bool {
	percorrente := Campo.fine
	for percorrente != nil && strings.Contains(percorrente.id, "ostacolo") {
		x0, y0, x1, y1 := estraiCoordinate(percorrente.id)
		if (x <= x1 && x >= x0) && (y <= y1 && y >= y0) {
			return true
		}
		percorrente = percorrente.precedente
	}
	return false
}

// La funzione controlla se ci sono ostacoli sullo specifico asse delle x di p
func (p *punto) presenzaOstacoloPercorsoX(y int) bool {
	for i := p.coordinataY; i < y; i++ {
		if dentroAreaOstacolo(p.coordinataX, i) {
			return true
		}
	}
	return false
}

// La funzione controlla se ci sono ostacoli sullo specifico asse delle y di p
func (p *punto) presenzaOstacoloPercorsoY(x int) bool {
	for i := p.coordinataX; i < x; i++ {
		if dentroAreaOstacolo(i, p.coordinataY) {
			return true
		}
	}
	return false
}

// Questa funzione estrae le coordinate da un ostacolo
func estraiCoordinate(id string) (x0 int, y0 int, x1 int, y1 int) {
	coordinate, _ := strings.CutSuffix(id, "ostacolo")
	slCoordinate := strings.Split(coordinate, ",")
	x0, _ = strconv.Atoi(slCoordinate[0])
	y0, _ = strconv.Atoi(slCoordinate[1])
	x1, _ = strconv.Atoi(slCoordinate[2])
	y1, _ = strconv.Atoi(slCoordinate[3])
	return
}

func calcolaDistanza(x0, y0, x1, y1 int) int {
	return (x1 - x0) + (y1 - y0)
}

package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type punto struct {
	coordinataX int
	coordinataY int
	id          string
	successivo  *punto
	precedente  *punto
}
type piano struct {
	inizio *punto
	fine   *punto
}
type Richiamati struct {
	testa *nodoPila
}
type nodoPila struct {
	chiamato *punto
	distanza int
	prossimo *nodoPila
}

var Campo piano
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
	case "r":
		a, _ := strconv.Atoi(comandi[1])
		b, _ := strconv.Atoi(comandi[2])
		richiamo(a, b, comandi[3])
	case "e":
		a, _ := strconv.Atoi(comandi[1])
		b, _ := strconv.Atoi(comandi[2])
		esistePercorso(a, b, comandi[3])
	case "f":
		os.Exit(0)
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
	minDistance := math.MaxInt
	pilaChiamata := new(Richiamati)
	Sorgente = new(punto)
	Sorgente.coordinataX = x
	Sorgente.coordinataY = y
	Sorgente.id = alpha
	percorrente := Campo.inizio
	for percorrente != nil && !strings.Contains(percorrente.id, "ostacolo") {
		if strings.HasPrefix(percorrente.id, alpha) {
			distanza := calcolaDistanza(percorrente.coordinataX, percorrente.coordinataY, x, y)
			possibileAvanzamento, _ := avanza(percorrente, distanza)
			if possibileAvanzamento.coordinataX == x && possibileAvanzamento.coordinataY == y {
				if distanza <= minDistance {
					minDistance = distanza
				}
				nodoChiamata := new(nodoPila)
				nodoChiamata.chiamato = percorrente
				nodoChiamata.prossimo = pilaChiamata.testa
				nodoChiamata.distanza = distanza
				pilaChiamata.testa = nodoChiamata
			}
		}
		percorrente = percorrente.successivo
	}
	attraversoPila := pilaChiamata.testa
	for attraversoPila != nil {
		if attraversoPila.distanza == minDistance {
			possibileArrivo, _ := avanza(attraversoPila.chiamato, minDistance)
			if possibileArrivo.coordinataX == x && possibileArrivo.coordinataY == y {
				attraversoPila.chiamato.coordinataX = x
				attraversoPila.chiamato.coordinataY = y
			}
		}
		attraversoPila = attraversoPila.prossimo
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

func esistePercorso(x, y int, eta string) {
	Sorgente = new(punto)
	Sorgente.coordinataX = x
	Sorgente.coordinataY = y
	Sorgente.id = eta
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
	var conatoreOstacoliX, contatoreOstacoliY, latoX, latoY int
	if passi <= 0 {
		return p, 0
	}
	possibilePasso := new(punto)
	possibilePasso.coordinataX = p.coordinataX
	possibilePasso.coordinataY = p.coordinataY
	possibilePasso.id = p.id
	if Sorgente.coordinataX < possibilePasso.coordinataX {
		latoX = possibilePasso.coordinataX - passi
	} else {
		latoX = possibilePasso.coordinataX + passi
	}
	if Sorgente.coordinataY < possibilePasso.coordinataY {
		latoY = possibilePasso.coordinataY - passi
	} else {
		latoY = possibilePasso.coordinataY + passi
	}
	for possibilePasso.presenzaOstacoloPercorsoX(latoY) {
		// Println("ciclo")
		conatoreOstacoliX++
		if latoY < possibilePasso.coordinataY {
			latoY++
		} else {
			latoY--
		}
	}
	for possibilePasso.presenzaOstacoloPercorsoY(latoX) {
		contatoreOstacoliY++
		if latoX < possibilePasso.coordinataX {
			latoX++
		} else {
			latoX--
		}
	}
	if p.coordinataX < Sorgente.coordinataX && conatoreOstacoliX >= contatoreOstacoliY && !dentroAreaOstacolo(possibilePasso.coordinataX+1, possibilePasso.coordinataY) {
		possibilePasso.coordinataX++
		passi--
		return avanza(possibilePasso, passi)
	} else if p.coordinataX > Sorgente.coordinataX && conatoreOstacoliX >= contatoreOstacoliY && !dentroAreaOstacolo(possibilePasso.coordinataX-1, possibilePasso.coordinataY) {
		possibilePasso.coordinataX--
		passi--
		return avanza(possibilePasso, passi)
	}
	if p.coordinataY < Sorgente.coordinataY {
		possibilePasso.coordinataY++
		passi--
		return avanza(possibilePasso, passi)
	} else if p.coordinataY > Sorgente.coordinataY {
		possibilePasso.coordinataY--
		passi--
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

func (p *punto) presenzaOstacoloPercorsoX(y int) bool {
	if p.coordinataY > y {
		for i := y; i < p.coordinataY; i++ {
			if dentroAreaOstacolo(p.coordinataX, i) {
				return true
			}
		}
	} else {
		for i := p.coordinataY; i < y; i++ {
			if dentroAreaOstacolo(p.coordinataX, i) {
				return true
			}
		}
	}
	return false
}

func (p *punto) presenzaOstacoloPercorsoY(x int) bool {
	if p.coordinataX > x {
		for i := x; i < p.coordinataX; i++ {
			if dentroAreaOstacolo(i, p.coordinataY) {
				return true
			}
		}
	} else {
		for i := p.coordinataX; i < x; i++ {
			if dentroAreaOstacolo(i, p.coordinataY) {
				return true
			}
		}
	}
	return false
}

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
	Distanza := math.Abs(float64(x1-x0)) + math.Abs(float64(y1-y0))
	return int(Distanza)
}

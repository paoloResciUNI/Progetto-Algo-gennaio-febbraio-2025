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

type piano *Piano


type Piano struct {
	inizio *punto
	fine   *punto
}
type nodoPila struct {
	chiamato *punto
	distanza int
	prossimo *nodoPila
}

var Sorgente *punto

func esegui(p piano, s string) {
	comandi := strings.Split(s, " ")
	switch comandi[0] {
	case "c":
		P := newPiano()
		p.inizio = P.inizio
		p.fine = P.fine
	case "S":
		(*p).stampa()
	case "s":
		a, _ := strconv.Atoi(comandi[1])
		b, _ := strconv.Atoi(comandi[2])
		(*p).stato(a, b)
	case "a":
		a, _ := strconv.Atoi(comandi[1])
		b, _ := strconv.Atoi(comandi[2])
		(*p).automa(a, b, comandi[3])
	case "o":
		a, _ := strconv.Atoi(comandi[1])
		b, _ := strconv.Atoi(comandi[2])
		c, _ := strconv.Atoi(comandi[3])
		d, _ := strconv.Atoi(comandi[4])
		(*p).ostacolo(a, b, c, d)
	case "p":
		(*p).posizioni(comandi[1])
	case "r":
		a, _ := strconv.Atoi(comandi[1])
		b, _ := strconv.Atoi(comandi[2])
		(*p).richiamo(a, b, comandi[3])
	case "e":
		a, _ := strconv.Atoi(comandi[1])
		b, _ := strconv.Atoi(comandi[2])
		(*p).esistePercorso(a, b, comandi[3])
	case "f":
		os.Exit(0)
	}
}

func newPiano() piano {
	var nuovPiano Piano
	return &nuovPiano
}

func main() {
	var Campo Piano
	var pCampo piano
	pCampo = &Campo
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		esegui(pCampo, scanner.Text())
	}
	Println(&Campo)
}

func (Campo Piano) stampa() {
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

func (Campo Piano) stato(x, y int) {
	if Campo.dentroAreaOstacolo(x, y) {
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

func (Campo *Piano) automa(x, y int, eta string) {
	puntoCercato := Campo.cerca(x, y, eta)
	if puntoCercato != nil {
		if strings.Contains(puntoCercato.id, "ostacolo") && !(Campo.dentroAreaOstacolo(x, y)) {
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

func (Campo *Piano) ostacolo(x0, y0, x1, y1 int) {
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

func (Campo *Piano) richiamo(x, y int, alpha string) {
	minDistance := math.MaxInt
	pilaChiamata := new(nodoPila)
	percorrente := Campo.inizio
	for percorrente != nil && !strings.Contains(percorrente.id, "ostacolo") {
		if strings.HasPrefix(percorrente.id, alpha) {
			distanza := calcolaDistanza(percorrente.coordinataX, percorrente.coordinataY, x, y)
			possibileAvanzamCampo, entro := avanza(Campo, percorrente, distanza)
			if possibileAvanzamCampo.coordinataX == x && possibileAvanzamCampo.coordinataY == y && entro == 0 {
				if distanza <= minDistance {
					minDistance = distanza
				}
				nodoChiamata := new(nodoPila)
				nodoChiamata.chiamato = percorrente
				nodoChiamata.prossimo = pilaChiamata
				nodoChiamata.distanza = distanza
				pilaChiamata = nodoChiamata
			}
		}
		percorrente = percorrente.successivo
	}
	attraversoPila := pilaChiamata
	for attraversoPila != nil {
		if attraversoPila.distanza == minDistance {
			possibileArrivo, _ := avanza(Campo, attraversoPila.chiamato, minDistance)
			if possibileArrivo.coordinataX == x && possibileArrivo.coordinataY == y {
				attraversoPila.chiamato.coordinataX = x
				attraversoPila.chiamato.coordinataY = y
			}
		}
		attraversoPila = attraversoPila.prossimo
	}
}

func (Campo Piano) posizioni(alpha string) {
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

func (Campo *Piano) esistePercorso(x, y int, eta string) {
	Sorgente = new(punto)
	Sorgente.coordinataX = x
	Sorgente.coordinataY = y
	Sorgente.id = eta
	if Campo.dentroAreaOstacolo(x, y) {
		Println("NO")
		return
	}
	percorrente := Campo.cerca(x, y, eta)
	if percorrente == nil {
		Println("NO")
		return
	}
	distanza := calcolaDistanza(percorrente.coordinataX, percorrente.coordinataY, x, y)
	percorsoEffettuato, passiMancanti := (avanza(Campo, percorrente, distanza))
	if percorsoEffettuato.coordinataX == x && percorsoEffettuato.coordinataY == y && passiMancanti == 0 {
		Println("SI")
		return
	} else {
		Println("NO")
	}
}

func avanza(Campo piano, p *punto, passi int) (*punto, int) {
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
	for possibilePasso.presenzaOstacoloPercorsoX(Campo, latoY) {
		conatoreOstacoliX++
		if latoY < possibilePasso.coordinataY {
			latoY++
		} else {
			latoY--
		}
	}
	for possibilePasso.presenzaOstacoloPercorsoY(Campo, latoX) {
		contatoreOstacoliY++
		if latoX < possibilePasso.coordinataX {
			latoX++
		} else {
			latoX--
		}
	}
	if p.coordinataX < Sorgente.coordinataX && conatoreOstacoliX >= contatoreOstacoliY && !((*Campo).dentroAreaOstacolo(possibilePasso.coordinataX+1, possibilePasso.coordinataY)) {
		possibilePasso.coordinataX++
		passi--
		return avanza(Campo, possibilePasso, passi)
	} else if p.coordinataX > Sorgente.coordinataX && conatoreOstacoliX >= contatoreOstacoliY && !(*Campo).dentroAreaOstacolo(possibilePasso.coordinataX-1, possibilePasso.coordinataY) {
		possibilePasso.coordinataX--
		passi--
		return avanza(Campo, possibilePasso, passi)
	}
	if p.coordinataY < Sorgente.coordinataY {
		possibilePasso.coordinataY++
		passi--
		return avanza(Campo, possibilePasso, passi)
	} else if p.coordinataY > Sorgente.coordinataY {
		possibilePasso.coordinataY--
		passi--
		return avanza(Campo, possibilePasso, passi)
	}
	passi--
	return avanza(Campo, p, passi)
}

func (Campo *Piano) cerca(x, y int, id string) *punto {
	percorrente := Campo.inizio
	for percorrente != nil {
		if (percorrente.coordinataX == x && percorrente.coordinataY == y) || percorrente.id == id {
			return percorrente
		}
		percorrente = percorrente.successivo
	}
	return nil
}

func (Campo *Piano) dentroAreaOstacolo(x, y int) bool {
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

func (p *punto) presenzaOstacoloPercorsoX(Campo piano, y int) bool {
	if p.coordinataY > y {
		for i := y; i < p.coordinataY; i++ {
			if (*Campo).dentroAreaOstacolo(p.coordinataX, i) {
				return true
			}
		}
	} else {
		for i := p.coordinataY; i < y; i++ {
			if (*Campo).dentroAreaOstacolo(p.coordinataX, i) {
				return true
			}
		}
	}
	return false
}

func (p *punto) presenzaOstacoloPercorsoY(Campo piano, x int) bool {
	if p.coordinataX > x {
		for i := x; i < p.coordinataX; i++ {
			if (*Campo).dentroAreaOstacolo(i, p.coordinataY) {
				return true
			}
		}
	} else {
		for i := p.coordinataX; i < x; i++ {
			if (*Campo).dentroAreaOstacolo(i, p.coordinataY) {
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

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

func main() {
	var Campo Piano
	pCampo := &Campo
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
	Sorgente := new(punto)
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
	percorsoEffettuato, _ := (avanza(Campo, percorrente, Sorgente, distanza))
	if percorsoEffettuato.coordinataX == x && percorsoEffettuato.coordinataY == y {
		Println("SI")
		return
	} else {
		Println("NO")
	}
}

func newPiano() piano {
	var nuovPiano Piano
	return &nuovPiano
}

func (Campo *Piano) automa(x, y int, eta string) {
	puntoCercato := Campo.cerca(x, y, eta)
	if puntoCercato != nil {
		if strings.Contains(puntoCercato.id, "ostacolo") && !Campo.dentroAreaOstacolo(x, y) {
			return
		} else {
			puntoCercato.coordinataX = x
			puntoCercato.coordinataY = y
		}
	} else if !Campo.dentroAreaOstacolo(x, y) {
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
	Sorgente := new(punto)
	Sorgente.coordinataX = x
	Sorgente.coordinataY = y
	minDistance := math.MaxInt
	pilaChiamata := new(nodoPila)
	percorrente := Campo.inizio
	for percorrente != nil && !strings.Contains(percorrente.id, "ostacolo") {
		if strings.HasPrefix(percorrente.id, alpha) {
			distanza := calcolaDistanza(percorrente.coordinataX, percorrente.coordinataY, x, y)
			possibileAvanzamento, entro := avanza(Campo, percorrente, Sorgente, distanza)
			if possibileAvanzamento.coordinataX == x && possibileAvanzamento.coordinataY == y && entro == 0 {
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
			possibileArrivo, _ := avanza(Campo, attraversoPila.chiamato, Sorgente, minDistance)
			if possibileArrivo.coordinataX == x && possibileArrivo.coordinataY == y {
				attraversoPila.chiamato.coordinataX = x
				attraversoPila.chiamato.coordinataY = y
			}
		}
		attraversoPila = attraversoPila.prossimo
	}
}

func (Campo *Piano) ostacoliPercorso(partenza, arrivo *punto) (distanza_O_Ascisse, distanza_O_Ordinate int) {
	_, ostacoloVicino := partenza.posizioneOstacoloVerticale(Campo, arrivo.coordinataY)
	if ostacoloVicino != nil {
		_, y0, _, y1 := estraiCoordinate(ostacoloVicino.id)
		if arrivo.coordinataY < partenza.coordinataY {
			distanza_O_Ordinate = calcolaDistanza(partenza.coordinataX, partenza.coordinataY, partenza.coordinataX, y1)
		} else if arrivo.coordinataY > partenza.coordinataY {
			distanza_O_Ordinate = calcolaDistanza(partenza.coordinataX, partenza.coordinataY, partenza.coordinataX, y0)
		}
	} else {
		distanza_O_Ordinate = calcolaDistanza(partenza.coordinataX, partenza.coordinataY, partenza.coordinataX, arrivo.coordinataY)
	}
	_, ostacoloVicino = partenza.posizioneOstacoloOrizzontale(Campo, arrivo.coordinataX)
	if ostacoloVicino != nil {
		x0, _, x1, _ := estraiCoordinate(ostacoloVicino.id)
		if arrivo.coordinataX < partenza.coordinataX {
			distanza_O_Ascisse = calcolaDistanza(partenza.coordinataX, partenza.coordinataY, x1, partenza.coordinataY)
		} else {
			distanza_O_Ascisse = calcolaDistanza(partenza.coordinataX, partenza.coordinataY, x0, partenza.coordinataY)
		}
	} else {
		distanza_O_Ascisse = calcolaDistanza(partenza.coordinataX, partenza.coordinataY, arrivo.coordinataX, partenza.coordinataY)
	}
	return
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

// Cerca un ostacolo sull'asse delle y del punto dove si trova un punto p considerato.
func (p *punto) posizioneOstacoloVerticale(Campo piano, y int) (bool, *punto) {
	if p.coordinataY > y {
		for i := p.coordinataY - 1; i >= y; i-- {
			ostacolo := (*Campo).cercaOstacolo(p.coordinataX, i)
			if ostacolo != nil {
				return true, ostacolo
			}
		}
	} else if p.coordinataY < y {
		for i := p.coordinataY + 1; i < y; i++ {
			ostacolo := (*Campo).cercaOstacolo(p.coordinataX, i)
			if ostacolo != nil {
				return true, ostacolo
			}
		}
	}
	return false, nil
}

func (p *punto) posizioneOstacoloOrizzontale(Campo piano, x int) (bool, *punto) {
	if p.coordinataX > x {
		for i := p.coordinataX - 1; i >= x; i-- {
			ostacolo := (*Campo).cercaOstacolo(i, p.coordinataY)
			if ostacolo != nil {
				return true, ostacolo
			}
		}
	} else if p.coordinataX < x {
		for i := p.coordinataX + 1; i < x; i++ {
			ostacolo := (*Campo).cercaOstacolo(i, p.coordinataY)
			if ostacolo != nil {
				return true, ostacolo
			}
		}
	}
	return false, nil
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

func avanza(Campo piano, p *punto, Sorgente *punto, passi int) (*punto, int) {
	var distanzaVerticaleO, distanzaOrizzontaleO int
	if passi <= 0 || p.coordinataX == Sorgente.coordinataX && p.coordinataY == Sorgente.coordinataY {
		return p, 0
	}
	possibilePasso := new(punto)
	possibilePasso.coordinataX = p.coordinataX
	possibilePasso.coordinataY = p.coordinataY
	possibilePasso.id = p.id
	distanzaOrizzontaleO, distanzaVerticaleO = (*Campo).ostacoliPercorso(possibilePasso, Sorgente)
	if distanzaVerticaleO < distanzaOrizzontaleO {
		possibilePasso = (*Campo).forwardingX(possibilePasso, Sorgente)
	} else if distanzaOrizzontaleO == distanzaVerticaleO && p.coordinataX < Sorgente.coordinataX && !(*Campo).dentroAreaOstacolo(p.coordinataX+1, p.coordinataY) {
		possibilePasso.coordinataX++
	} else if distanzaOrizzontaleO == distanzaVerticaleO && p.coordinataX > Sorgente.coordinataX && !(*Campo).dentroAreaOstacolo(p.coordinataX-1, p.coordinataY) {
		possibilePasso.coordinataX--
	} else if distanzaVerticaleO > distanzaOrizzontaleO {
		possibilePasso = (*Campo).forwardingY(possibilePasso, Sorgente)
	}
	if p.coordinataX == possibilePasso.coordinataX && p.coordinataY == possibilePasso.coordinataY {
		return possibilePasso, passi
	}
	passi--
	return avanza(Campo, possibilePasso, Sorgente, passi)
}

func (Campo *Piano) forwardingX(start *punto, destination *punto) *punto {
	var forward punto
	_, ostacoloVicino := start.posizioneOstacoloVerticale(Campo, destination.coordinataY)
	if ostacoloVicino != nil {
		var puntoE, puntoO int
		x0, _, x1, _ := estraiCoordinate(ostacoloVicino.id)
		puntoE = calcolaDistanza(destination.coordinataX, destination.coordinataY, x1, destination.coordinataY)
		puntoO = calcolaDistanza(destination.coordinataX, destination.coordinataY, x0, destination.coordinataY)
		if puntoE < puntoO {
			forward.coordinataX = x1 + 1
			forward.coordinataY = start.coordinataY
		} else {
			forward.coordinataX = x0 - 1
			forward.coordinataY = start.coordinataY
		}
		return &forward
	}
	_, ostacoloVicino = start.posizioneOstacoloOrizzontale(Campo, destination.coordinataX)
	if ostacoloVicino != nil {
		x0, _, x1, _ := estraiCoordinate(ostacoloVicino.id)
		if start.coordinataX < destination.coordinataX {
			forward.coordinataX = x0 - 1
		} else if start.coordinataX > destination.coordinataX {
			forward.coordinataX = x1 + 1
		}
	} else {
		forward.coordinataX = destination.coordinataX
	}
	forward.coordinataY = start.coordinataY
	_, osostacoloVicino := forward.posizioneOstacoloVerticale(Campo, destination.coordinataX)
	if osostacoloVicino != nil {
		var puntoE, puntoO int
		x0, _, x1, _ := estraiCoordinate(osostacoloVicino.id)
		puntoE = calcolaDistanza(start.coordinataX, start.coordinataY, x1, start.coordinataY)
		puntoO = calcolaDistanza(start.coordinataX, start.coordinataY, x0, start.coordinataY)
		if puntoE < puntoO {
			forward.coordinataX = x1 + 1
		} else {
			forward.coordinataX = x0 - 1
		}
	}
	return &forward
}

func (Campo *Piano) forwardingY(start *punto, destination *punto) *punto {
	var forward punto
	_, ostacoloVicino := start.posizioneOstacoloOrizzontale(Campo, destination.coordinataX)
	if ostacoloVicino != nil {
		var puntoN, puntoS int
		_, y0, _, y1 := estraiCoordinate(ostacoloVicino.id)
		puntoN = calcolaDistanza(destination.coordinataX, destination.coordinataY, destination.coordinataX, y1)
		puntoS = calcolaDistanza(destination.coordinataX, destination.coordinataY, destination.coordinataX, y0)
		if puntoN < puntoS {
			forward.coordinataY = y1 + 1
			forward.coordinataX = start.coordinataX
		} else {
			forward.coordinataY = y0 - 1
			forward.coordinataX = start.coordinataX
		}
		return &forward
	}
	forward.coordinataX = start.coordinataX
	_, ostacoloVicino = start.posizioneOstacoloVerticale(Campo, destination.coordinataY)
	if ostacoloVicino != nil {
		_, y0, _, y1 := estraiCoordinate(ostacoloVicino.id)
		if start.coordinataY < destination.coordinataY {
			forward.coordinataY = y0 - 1
		} else if start.coordinataY > destination.coordinataY {
			forward.coordinataY = y1 + 1
		}
	} else {
		forward.coordinataY = destination.coordinataY
	}
	_, ostacoloVicino = forward.posizioneOstacoloOrizzontale(Campo, destination.coordinataX)
	if ostacoloVicino != nil {
		var puntoN, puntoS int
		_, y0, _, y1 := estraiCoordinate(ostacoloVicino.id)
		puntoN = calcolaDistanza(start.coordinataX, start.coordinataY, start.coordinataX, y1)
		puntoS = calcolaDistanza(start.coordinataX, start.coordinataY, start.coordinataX, y0)
		if puntoN < puntoS {
			forward.coordinataY = y1 + 1
		} else {
			forward.coordinataY = y0 - 1
		}
		return &forward
	}
	return &forward
}

func (Campo *Piano) cercaOstacolo(x int, y int) *punto {
	percorrente := Campo.fine
	for percorrente != nil && strings.Contains(percorrente.id, "ostacolo") {
		x0, y0, x1, y1 := estraiCoordinate(percorrente.id)
		if (x <= x1 && x >= x0) && (y <= y1 && y >= y0) {
			return percorrente
		}
		percorrente = percorrente.precedente
	}
	return nil
}

func (Campo *Piano) cerca(x, y int, id string) *punto {
	percorrente := Campo.inizio
	for percorrente != nil {
		if percorrente.coordinataX == x && percorrente.coordinataY == y || percorrente.id == id {
			return percorrente
		}
		percorrente = percorrente.successivo
	}
	return nil
}

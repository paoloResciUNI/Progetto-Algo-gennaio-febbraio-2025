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
	Println(s)
	comandi := strings.Split(s, " ")
	Println(comandi)
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
		Println(comandi[1])
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

// ______________________________________________________________________________________________________________________
// __________________________Sezione di funzioni e metodi per la manipolazione del campo_____________________________________
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

// ______________________________________________________________________________________________________________________
// __________________________Sezione di funzioni e meotdi per il calcolo della presenza di ostacoli_____________________
func (Campo *Piano) ostacoliPercorso(partenza, arrivo *punto) (vicinanza_O_Ascisse, vicinanza_O_Ordinate int) {
	// Cerca l'ostacolo più vicino che si appoggia su questo specifico asse delle X.
	_, ostacoloVicinoAsseX := partenza.presenzaOstacoloVerticale(Campo, arrivo.coordinataY)
	if ostacoloVicinoAsseX != nil {
		if arrivo.coordinataX < partenza.coordinataX {
			vicinanza_O_Ordinate = partenza.coordinataY - ostacoloVicinoAsseX.coordinataY
		} else {
			vicinanza_O_Ordinate = ostacoloVicinoAsseX.coordinataY - partenza.coordinataY
		}

	} else {
		vicinanza_O_Ordinate = 0
	}
	_, ostacoloVicinoAsseY := partenza.presenzaOstacoloOrizzontale(Campo, arrivo.coordinataX)
	if ostacoloVicinoAsseY != nil {
		if arrivo.coordinataY < partenza.coordinataY {
			vicinanza_O_Ascisse = partenza.coordinataX - ostacoloVicinoAsseY.coordinataX
		} else {
			vicinanza_O_Ascisse = ostacoloVicinoAsseY.coordinataX - partenza.coordinataX
		}
	} else {
		vicinanza_O_Ascisse = 0
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
func (p *punto) presenzaOstacoloVerticale(Campo piano, y int) (bool, *punto) {
	if p.coordinataY > y {
		for i := p.coordinataY - 1; i >= y; i-- {
			ostacolo := (*Campo).cercaOstacolo(p.coordinataX, i)
			if ostacolo != nil {
				return true, ostacolo
			}
		}
	} else {
		for i := p.coordinataY + 1; i < y; i++ {
			ostacolo := (*Campo).cercaOstacolo(p.coordinataX, i)
			if ostacolo != nil {
				return true, ostacolo
			}
		}
	}
	return false, nil
}

// cerca l'ostacolo più vicino sull'asse delle x dell'automa che va verso la sorgente del segnale
func (p *punto) presenzaOstacoloOrizzontale(Campo piano, x int) (bool, *punto) {
	if p.coordinataX > x {
		for i := p.coordinataX - 1; i >= x; i-- {
			ostacolo := (*Campo).cercaOstacolo(i, p.coordinataY)
			if ostacolo != nil {
				return true, ostacolo
			}
		}
	} else {
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

// ______________________________________________________________________________________________________________________
// __________________________Sezione di funzioni e metodi di supporto per le operazioni di manipolazione e osservazione sul campo____________________________________
func calcolaDistanza(x0, y0, x1, y1 int) int {
	Distanza := math.Abs(float64(x1-x0)) + math.Abs(float64(y1-y0))
	return int(Distanza)
}

func avanza(Campo piano, p *punto, Sorgente *punto, passi int) (*punto, int) {
	var contatoreOstacoliX, contatoreOstacoliY int
	if passi <= 0 || p.coordinataX == Sorgente.coordinataX && p.coordinataY == Sorgente.coordinataY {
		return p, 0
	}
	possibilePasso := new(punto)
	possibilePasso.coordinataX = p.coordinataX
	possibilePasso.coordinataY = p.coordinataY
	possibilePasso.id = p.id
	contatoreOstacoliY, contatoreOstacoliX = (*Campo).ostacoliPercorso(Sorgente, possibilePasso)
	if p.coordinataX < Sorgente.coordinataX && contatoreOstacoliX >= contatoreOstacoliY && !(*Campo).dentroAreaOstacolo(possibilePasso.coordinataX+1, possibilePasso.coordinataY) {
		possibilePasso = (*Campo).forwardingX(possibilePasso, Sorgente)
	} else if p.coordinataX > Sorgente.coordinataX && contatoreOstacoliX >= contatoreOstacoliY && !(*Campo).dentroAreaOstacolo(possibilePasso.coordinataX-1, possibilePasso.coordinataY) {
		possibilePasso = (*Campo).forwardingX(possibilePasso, Sorgente)
	}
	if p.coordinataY < Sorgente.coordinataY {
		possibilePasso = (*Campo).forwardingY(possibilePasso, Sorgente)
	} else if p.coordinataY > Sorgente.coordinataY {
		possibilePasso = (*Campo).forwardingY(possibilePasso, Sorgente)
	}
	if p.coordinataX == possibilePasso.coordinataX && p.coordinataY == possibilePasso.coordinataY {
		return possibilePasso, passi
	}
	passi--
	return avanza(Campo, possibilePasso, Sorgente, passi)
}

// start è il punto da dove parte il forwarding e destination è il segnale.
// Il metodo deve restituire il punto nel quale si è mosso l'automa dopo il forwarding. Cioè dopo aver fatto i calcoli
// sui peercorsi possibili dell'avanzamento, o meglio: se dopo il forwarding l'automa si trova bloccato da un ostacolo allora l'automa si muovarà
// sul punto subito dopo il vertice che si avvicina al sugnale senza aumentare la distanza di manattan.
// Con forwardX assumo già che dovrò fare il forwarding sull'asse X.
// Mi muovo sull'asse delle X.
func (Campo *Piano) forwardingX(start *punto, destination *punto) *punto {
	var forward punto
	_, ostacoloVicino := start.presenzaOstacoloPercorsoX(Campo, destination.coordinataY)
	if ostacoloVicino != nil {
		var puntoSE, puntoSO int
		// Clcolo la distanza di tutti i verici dall'automa e dal segnale.
		// Questo mi serve per capire quale vertice si avvicina di più al segnale restando comunque vicino all'automa.
		x0, y0, x1, _ := estraiCoordinate(ostacoloVicino.id)
		puntoSE = calcolaDistanza(start.coordinataX, start.coordinataY, x1, y0)
		puntoSO = calcolaDistanza(start.coordinataX, start.coordinataY, x0, y0)
		if puntoSE < puntoSO {
			forward.coordinataX = x1 + 1
			forward.coordinataY = start.coordinataY
		} else {
			forward.coordinataX = x0 - 1
			forward.coordinataY = start.coordinataY
		}
		return &forward
	}
	forward.coordinataX = destination.coordinataX
	forward.coordinataY = start.coordinataY
	_, osostacoloVicino := forward.presenzaOstacoloPercorsoY(Campo, destination.coordinataX)
	if osostacoloVicino != nil {
		var puntoNO, puntoSO int
		x0, y0, _, y1 := estraiCoordinate(osostacoloVicino.id)
		puntoNO = calcolaDistanza(forward.coordinataX, forward.coordinataY, x0, y1)
		puntoSO = calcolaDistanza(forward.coordinataX, forward.coordinataY, x0, y0)
		if puntoNO < puntoSO {
			forward.coordinataY = y1 + 1
			forward.coordinataX = start.coordinataX
		} else {
			forward.coordinataY = y0 - 1
			forward.coordinataX = start.coordinataX
		}
	}
	return &forward
}

// Con forwardY assumo già che dovrò fare il forwarding sull'asse Y.
func (Campo *Piano) forwardingY(start *punto, destination *punto) *punto {
	var forward punto
	_, ostacoloVicino := start.presenzaOstacoloPercorsoY(Campo, destination.coordinataX)
	if ostacoloVicino != nil {
		var puntoNO, puntoSO int
		x0, y0, _, y1 := estraiCoordinate(ostacoloVicino.id)
		puntoNO = calcolaDistanza(start.coordinataX, start.coordinataY, x0, y1)
		puntoSO = calcolaDistanza(start.coordinataX, start.coordinataY, x0, y0)
		if puntoNO < puntoSO {
			forward.coordinataY = y1 + 1
			forward.coordinataX = start.coordinataX
		} else {
			forward.coordinataY = y0 - 1
			forward.coordinataX = start.coordinataX
		}
		return &forward
	}
	forward.coordinataY = destination.coordinataY
	forward.coordinataX = start.coordinataX
	_, osostacoloVicino := forward.presenzaOstacoloPercorsoX(Campo, destination.coordinataY)
	if osostacoloVicino != nil {
		var puntoSE, puntoSO int
		x0, y0, x1, _ := estraiCoordinate(osostacoloVicino.id)
		puntoSE = calcolaDistanza(forward.coordinataX, forward.coordinataY, x1, y0)
		puntoSO = calcolaDistanza(forward.coordinataX, forward.coordinataY, x0, y0)
		if puntoSE < puntoSO {
			forward.coordinataX = x1 + 1
			forward.coordinataY = start.coordinataY
		} else {
			forward.coordinataX = x0 - 1
			forward.coordinataY = start.coordinataY
		}
	}
	return &forward
}

func (Campo *Piano) cercaOstacolo(x int, y int) *punto {
	percorrente := Campo.fine
	for percorrente != nil && strings.Contains(percorrente.id, "ostacolo") {
		if Campo.dentroAreaOstacolo(x, y) {
			return percorrente
		percorrente = percorrente.precedente
	}
	return nil
}

func (Campo *Piano) cerca(x, y int, id string) *punto {
	percorrente := Campo.fine
	for percorrente != nil && strings.Contains(percorrente.id, "ostacolo") {
		x0, y0, x1, y1 := estraiCoordinate(percorrente.id)
		if (x <= x1 && x >= x0) && (y <= y1 && y >= y0) {
			return true
		percorrente = percorrente.precedente
		return percorrente
		}
		percorrente = percorrente.successivo
	}
	return nil
}

// Aggiungere un metodo di forwarding che va, quando ci snono ostacoli da entrambe le parti, verso il verice più vicino alla sorgente del segnale ma verso la direzione dell'automa.
// Aggiungere il controllo: Prima del forwarding controllare se ci sono percorsi liberi che riducono la distanza dalla sorgente del segnale. Altrimenti non fare il forwarding ed esci dalla funzione

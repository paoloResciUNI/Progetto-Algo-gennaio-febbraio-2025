package main

import (
	. "fmt"
	"strconv"
	"strings"
)

// nodo
//
// - Se il punto contiene un automa la stringa id ne conterrà il nome.
// - Se il punto è vuoto allora la stringa srà vuota.
// - Se il punto fa parte di un ostacolo allora id srà uguale alla stringa "ostacolo"
type punto struct {
	coordinataX int
	coordinataY int
	id          string
	richiamo    bool
	adiacenza   []*punto
}

// rappresentazione del grafo che rappresenta il piano
type piano []*punto

// Il camapo dove si muovono gli automi
var Campo piano

// Il punto da cui viene il richiamo
var Sorgente *punto

func esegui(p piano, s string) {

}

func newPiano() piano {
	crea()
	return Campo
}

func main() {
	newPiano()
	automa(2, 1, "1")
	ostacolo(1, 5, 5, 8)
	ostacolo(14, 1, 17, 4)
	ostacolo(7, 3, 10, 5)
	ostacolo(9, 5, 16, 7)
	ostacolo(6, 8, 10, 10)
	ostacolo(3, 0, 4, 2)
	esistePercorso(15, 9, "1")
	stato(2, 1)
	stato(3, 6)
	stato(5, 6)
	Println(Campo[0].presenzaOstacoloPercorsoY(12))
	Println(calcolaDistanza(2, 1, 15, 9))
	stampa()
}

// Crea un piano vuoto, eliminando il piano già esistente
func crea() {
	Campo = *new(piano)
}

// Si assume che il campo sia già ordinato nel formato corretto:
//
// Prima gli automi e poi gli ostacoli
func stampa() {
	i := 0
	Println("(")
	for i < len(Campo) {
		if !strings.Contains(Campo[i].id, "ostacolo") {
			Printf("%s:%d,%d \n", Campo[i].id, Campo[i].coordinataX, Campo[i].coordinataY)
		}
		i++
	}
	Println(")")
	Println("[")
	i = len(Campo) - 1
	for i >= 0 {
		if strings.Contains(Campo[i].id, "ostacolo") {
			x0, y0, x1, y1 := estraiCoordinate(Campo[i].id)
			Printf("(%d,%d)(%d,%d)\n", x0, y0, x1, y1)
		}
		i--
	}
	Println("]")
}

// Questa funzone stampa il contenuto nella posizione (x, y).
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
// Questa funzione deve essere ottimizzata!!
func automa(x, y int, eta string) {
	puntoCercato := Campo.cerca(x, y, eta)
	if puntoCercato != nil {
		if strings.Contains(puntoCercato.id, "ostacolo") {
			return
		} else {
			puntoCercato.coordinataX = x
			puntoCercato.coordinataY = y
			Campo.remove(eta)
			puntoCercato.adiacenze()
		}
	} else {
		puntoCercato = new(punto)
		puntoCercato.coordinataX = x
		puntoCercato.coordinataY = y
		puntoCercato.id = eta
		puntoCercato.adiacenze()
		Campo = append([]*punto{puntoCercato}, Campo...)
	}
}

// Un ostacolo è un punto in cui non possono esserci automi
// Ogni ostacolo ha un'adiacenza e ogni punto dell'ostacolo fa parte del campo
// Assunzione: x_1 < x_2 e y_1 < y_2
func ostacolo(x0, y0, x1, y1 int) {
	for i := 0; i < len(Campo); i++ {
		if (Campo[i].coordinataX <= x1 && Campo[i].coordinataX >= x0) && (Campo[i].coordinataY <= y1 && Campo[i].coordinataY >= y0) {
			return
		}
	}
	newOstacolo := new(punto)
	newOstacolo.coordinataX = x0
	newOstacolo.coordinataY = y1
	newOstacolo.id = Sprintf("%d,%d,%d,%d,ostacolo", x0, y0, x1, y1)
	Campo = append(Campo, newOstacolo)
}

func richiamo(x, y int, alpha string) {
	Sorgente = new(punto)
	Sorgente.coordinataX = x
	Sorgente.coordinataY = y
	Sorgente.id = alpha
	for i := 0; i < len(Campo); i++ {
		if strings.HasPrefix(Campo[i].id, alpha) {
			Campo[i].richiamo = true
		}
	}
}

func posizioni(alpha string) {
	Println("(")
	for i := 0; i < len(Campo); i++ {
		if strings.HasPrefix(Campo[i].id, alpha) {
			Printf("%s :%d,%d \n", Campo[i].id, Campo[i].coordinataX, Campo[i].coordinataY)
		}
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
	percorsoEffettuato, passiMancanti := avanza(percorrente, distanza)
	if percorsoEffettuato.coordinataX == x && percorsoEffettuato.coordinataY == y && passiMancanti == 0 {
		Println("SI")
		return
	} else {
		Println("NO")
	}
}

func avanza(p *punto, passi int) (*punto, int) {
	if passi == 0 {
		return p, passi
	}
	possibilePasso := new(punto)
	possibilePasso.coordinataX = p.coordinataX
	possibilePasso.coordinataY = p.coordinataY
	possibilePasso.id = p.id
	p.adiacenze()
	nord := calcolaDistanza(Sorgente.coordinataX, Sorgente.coordinataY, p.coordinataX, p.coordinataY+1)
	sud := calcolaDistanza(Sorgente.coordinataX, Sorgente.coordinataY, p.coordinataX, p.coordinataY-1)
	west := calcolaDistanza(Sorgente.coordinataX-1, Sorgente.coordinataY, p.coordinataX, p.coordinataY)
	est := calcolaDistanza(Sorgente.coordinataX+1, Sorgente.coordinataY, p.coordinataX, p.coordinataY)
	slCoords := [4]*int{&nord, &sud, &west, &est}
	direzione := scegliDirezione(p, slCoords[:])
	switch {
	case direzione == &nord:
		p.coordinataY++
		return avanza(p, passi-1)
	case direzione == &sud:
		p.coordinataY--
		return avanza(p, passi-1)
	case direzione == &est:
		p.coordinataX++
		return avanza(p, passi-1)
	case direzione == &west:
		p.coordinataX--
		return avanza(p, passi-1)
	}

	return p, passi
}

// La funzionde in questione dice quale direzione è meglio intraprendere per raggiungere
// la sorgente del segnale
func scegliDirezione(p *punto, coord []*int) *int {
	direzionePossibile1 := findMin(coord)
	direzionePossibile2 := findMin(coord[0:])
	if *direzionePossibile1 == *direzionePossibile2 {
		sX := Sorgente.coordinataX
		sY := Sorgente.coordinataY
		for p.presenzaOstacoloPercorsoX(sY) {
			sY--
		}
		for p.presenzaOstacoloPercorsoY(sX) {
			sX--
		}
		if sX < sY {
			return direzionePossibile2
		} else {
			return direzionePossibile1
		}
	} else if *direzionePossibile1 < *direzionePossibile2 {
		return direzionePossibile1
	} else {
		return direzionePossibile2
	}

}

// Cerco il puntatore che contiene il valore minimo tra le coordinate in input.
func findMin(coords []*int) *int {
	min := coords[0]
	for i := 1; i < 4; i++ {
		if *coords[i] < *min {
			min = coords[i]
		}
	}
	return min
}

func (*piano) cerca(x, y int, id string) *punto {
	for i := 0; i < len(Campo); i++ {
		if (Campo[i].coordinataX == x && Campo[i].coordinataY == y) || Campo[i].id == id {
			return Campo[i]
		}
	}
	return nil
}

func (p *punto) adiacenze() {
	var x, y int
	x = p.coordinataX - 1
	y = p.coordinataY - 1
	p.adiacenza = []*punto{}

	for y < p.coordinataY+2 {
		for x < p.coordinataX+2 {
			if dentroAreaOstacolo(x, y) {
				Ostacolo := new(punto)
				Ostacolo.coordinataX = x
				Ostacolo.coordinataY = y
				Ostacolo.id = "ostacolo"
				p.adiacenza = append(p.adiacenza, Ostacolo)
				x++
				continue
			}
			puntoAdiacente := Campo.cerca(x, y, "")
			if (puntoAdiacente != nil) && (x != p.coordinataX || y != p.coordinataY) {
				p.adiacenza = append(p.adiacenza, puntoAdiacente)
				puntoAdiacente.adiacenza = append(puntoAdiacente.adiacenza, p)
			}
			x++
		}
		x = p.coordinataX - 1
		y++
	}
}

// Rimuove un elemento dalla slice Campo.
func (*piano) remove(eta string) {
	for i := 0; i < len(Campo); i++ {
		if eta == Campo[i].id {
			if Campo[i].adiacenza == nil {
				return
			}
			for adiacente := 0; adiacente < len(Campo[i].adiacenza); adiacente++ {
				Campo[i].adiacenza[adiacente].adiacenze()
			}
		}
	}
}

func dentroAreaOstacolo(x, y int) bool {
	for i := len(Campo) - 1; i >= 0; i-- {
		if strings.Contains(Campo[i].id, "ostacolo") {
			x0, y0, x1, y1 := estraiCoordinate(Campo[i].id)
			if (x <= x1 && x >= x0) && (y <= y1 && y >= y0) {
				return true
			}
		}
	}
	return false
}

// La funzione controlla se ci sono ostacoli sull'asse delle x del percorso intrapreso dal punto p
func (p *punto) presenzaOstacoloPercorsoX(y int) bool {
	for i := p.coordinataY; i < y; i++ {
		if dentroAreaOstacolo(p.coordinataX, i) {
			return true
		}
	}
	return false
}

// La funzione controlla se ci sono ostacoli sull'asse delle y del percorso intrapreso dal punto p
func (p *punto) presenzaOstacoloPercorsoY(x int) bool {
	for i := p.coordinataX; i < x; i++ {
		if dentroAreaOstacolo(i, p.coordinataY) {
			return true
		}
	}
	return false
}

// Assumo che la stringa in input sia nel formato corretto
// Questa funzione è da sistemare.
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

// func visitaAdiacenze(possibiliStrade []*punto) *punto {}

# Relazione progetto d'esame di algoritmi e strutture dati

## Approccio al problema

Per rappresentare il piano è stata usata una lista doppiamente concatenata. È stata scelta una lista per rappresentare il piano perché come struttura dati è efficiente per il consumo della memoria ed è semplice da manipolare.

Gli automi e gli ostacoli condividono la stessa lista. I primi occupano soltanto la parte superiore della lista mentre gli altri solo quella inferiore. 
Degli ostacoli viene salvato in memoria solo il punto in alto a destra. Così facendo l'entità ostacolo risulta più leggera da rappresentare in memoria e non devo salvare, ogni volta che viene creato un nuovo ostacolo, tutti i punti facenti parte del perimetro e dell'area ma solamente uno. 

Gli *automi* sono punti all'interno del piano che si possono spostare durante l'esecuzione del programma. Essi si distinguono dagli ostacoli mediante l'id. Infatti se l'id non contiene la segnatura *"ostacolo"*  al suo interno allora il punto considerato è un automa. 

Gli *ostacoli* sono anche essi punti del piano, ma a differenza degli automi non si spostano all'interno di esso. Occupano uno spazio più ampio, all'interno del campo. L'area di un rettangolo viene calcolata tramite il suo punto più in basso a sinistra e il suo punto più in alto a destra e all'interno non possono esserci automi. 

Il *richiamo* viene emesso da un generico punto $(x, y)$ sul piano, esso riguarda soltanto alcuni automi. Degli automi interessati dal segnale si spostano soltanto quelli con distanza minore. 
Per controllare quali automi si sposteranno, vengono inseriti tutti quelli con un percorso libero fino al punto $(x,y)$ del richiamo, all'interno di una pila e viene memorizzata la distanza minima che si è trovata tra uno di questi automi e il richiamo. Scorrendo la pila, si trovano gli automi che corrispondono alla distanza minima trovata. Si cambiano le coordinate di quell'automa con le coordinate $(x, y)$.  

## Strutture dati coinvolte e tipi di dato principali

Il tipo `piano` è usato per rappresentare il campo, mediante una lista doppiamente concatenata, contenete l'elenco degli automi e degli ostacoli. Ogni ostacolo o automa è memorizzato tramite una struttura `punto` che rappresenta i nodi della lista. Il `piano` è composto dai campi `inizio` e `fine` che sono puntatori rispettivamente al `punto` che sta all'inizio della lista e a quello alla fine. Gli automi staranno nella prima metà della lista e gli ostacoli nella seconda metà, questo rende più rapida la ricerca all'interno del campo di un automa o di un ostacolo.

Il tipo `punto` è usato per rappresentare un generico punto nel piano, e un generico nodo della lista. Esso è composto da un campo `id`, due valori interi che rappresentano rispettivamente la coordinata $x$ e la coordinata $y$ e due puntatori ai punti: `successivo` e `precedente`, che sono il nodo successivo e il nodo precedente del punto considerato nella lista.    

Il tipo `nodoPila` è usato esclusivamente nella funzione `richiamo` per gestire i possibili nodi che possono avanzare. Serve esclusivamente per implementare una pila che contenga gli elementi del campo da considerare come possibili punti in movimento.     

## Metodi e funzioni

In questa sezione sono spiegate tutte le implementazioni e il comportamento delle entità introdotte precedentemente.  

#### Funzioni

- `esegui(p piano, s string)` : Questa funzione prende in input un piano `p` e una stringa `s` e, in base al contenuto di `s` vengono eseguite operazioni diverse sul piano.  

- `newPiano() piano` : Questa funzione viene invocata ogni volta che bisogna creare un nuovo piano. È invocata dalla lettera `c` all'interno della stringa `s`. Viene associato una nuova variabile di tipo `piano` alla variabile globale `Campo`. Questa funzione implementa l'operazione ***crea()***.  

- `stampa()` : Questa funzione stampa prima un elenco contenente tutti gli automi che fanno parte del piano e poi un altro contenente tutti gli ostacoli.     

- `stato(x, y int)` : Questa funzione stampa il tipo di entità presente in un determinato punto. Se nelle coordinate `x`, `y` è presente un automa, allora la funzione stamperà `A`, se è un punto facente parte di un ostacolo, stamperà `O` e invece se non c'è niente, stamperà `E`.  Questa funzione fa uso del metodo `cerca` che un punto nella lista a partire dalla testa.  

- `posizioni(alpha string)` : Questa funzione stampa gli automi che hanno come prefisso `alpha`. Questa funzione stampa gli automi nel formato output richiesto. 

- `automa(x, y int, eta string)` :  Questa funzione aggiunge un automa al campo se il punto $(x,y)$ non fanno parte di nessun ostacolo, altrimenti non fa nulla. Se le coordinate $(x, y)$ non fanno parte di un'ostacolo allora controlla se l'automa `eta` esiste già e in caso affermativo sposta l'automa di nome `eta` nella posizione $(x, y)$, altrimenti ne crea uno nuovo.

- `ostacolo(x0, y0, x1, y1 int)` : Questa funzione aggiunge un ostacolo all'interno del `piano`. Se l'area del quadrato compresa tra il vertice in basso a sinistra, `x0` e `y0`, e il vertice in alto a destra, `x1` e `y1`, contiene un'automa la funzione termina. Altrimenti aggiunge un punto con coordinate, `(x0, y1)` con id nel seguente formato: `(x0, y0, x1, y1) ostacolo`.
  
  > ##### Considerazioni sul comportamento dell'entità ostacolo:
  > 
  >  Per il calcolo dell'area del quadrato verranno usate le funzioni `dentroAreaOstacolo` e `estraiCoordinate` all'occorrenza all'interno del programma. Ho scelto di implementare in questo modo il comportamento e la rappresentazione nel progetto perchè mi sembrava il modo più efficiente per verificare se due coordinate $x$, $y$ facessero parte dell'area di un ostacolo. In questo modo ho potuto mantenere un  unico tipo di dato per automi e ostacoli. 
  > 
  > - `dentroAreaOstacolo(x, y int) bool` : Questa funzione serve per calcolare se  un punto è all'interno dell'area di un ostacolo. Restituisce `true` se il punto `(x, y)` fa parte dell'area di un generico ostacolo all'interno del campo. 
  > 
  > - `estraiCoordinate(id string) (x0 int, y0 int, x1 int, y1 int)` : Questa funzione è ustata per estrarre le coordinare del punto in basso a sinistra e in alto a destra dall'id.  Esso infatti è formattato in modo che vi siano le coordina dei due punti e la dicitura `ostacolo `subito dopo, in qusto modo: `Ostacolo.id = "x0,y0,x1,y1,ostacolo"`.

- `richiamo(x, y int, alpha string)` : Questa funzione avvia il richiamo per, e soltanto, tutti gli automi che come prefisso hanno la stringa `alpha`. Setta il campo `richiamo` degli automi interessati a `true`. Inoltre alloca una variabile globale `Sorgente` che sarà il punto che gli automi con prefisso `alpha` dovranno raggiungere. 

- `esistePercorso(x, y int, eta string)` : Questa funzione è usata per far stampare a schermo se esista o meno un percorso di distanza $D(P(\eta), (x,y)$) che in caso affermativo stampa su standard output, `SI` altrimenti `NO`. La funzione stampa `NO` anche se il punto $(x,y)$ fa parte dell'area di un ostacolo.
  
  > ##### Considerazione sulla funzione `esistePercorso`
  > 
  > L'implementazione della funzione delega alcune operazioni su altre due funzioni: `avanza` e `calcolaDistanza`. Questa scelta è stata presa per *appesantire*  il meno possibile il codice della  funzione principale e per poter dividere concettualmente il problema della ricerca del percorso. Di fatto questa scelta non è stata presa per fini legati alla performance del programma ma ai fini di rendere più facile implementare la risoluzione del problema della ricerca del percorso. 
  > 
  > - `avanza(p *punto, passi int) (*punto, int)` : Questa è una funzione ricorsiva che deve cercare il percorso minore per arrivare al punto dove si trova `Sorgente`. Come caso base, la funzione richiede `passi == 0`, `passi` è la distanza minima $D(P(\eta), (x,y))$. Il passo ricorsivo controlla dove il punto `p` si avvicina di più alla sorgente del segnale. Controlla poi su quale asse del piano, relativo alle coordinate del punto, ci sono più ostacoli. Per fare ciò mi sono appoggiato ai metodi `(p *punto) presenzaOstacoloPercorsoX(y int)` e `(p *punto) presenzaOstacoloPercorsoY(x int)` (questi metodi sono discussi nella sezione relativa ai metodi). Se non ci sono percorsi, che riducano la distanza dal punto `p` alla sorgente e che non passino attraverso un ostacolo,  il punto `p` va in *stallo* mentre i `passi` diminuiscono.
  > 
  > - `calcolaDistanza(x0, y0, x1, y1 int) int` : Questa funzione calcola, e restituisce, la distanza tra il punto $(x0, y0)$ e il punto $(x1, y1)$. 

#### Metodi

- `(*piano)cerca(x, y int, id string)*punto` : Questo metodo serve per cercare uno specifico punto, attraverso le coordinate ed eventualmente il nome, all'interno del piano. Se il punto non è presente il metodo restituisce `nil`, altrimenti da il punto all'interno del piano. 
- `(p *punto) presenzaOstacoloPercorsoX(y int) bool` : Questo metodo controlla se sull'asse delle $x$ del punto considerato ci sono ostacoli, in particolare se trova ostacoli sull'asse allora ritorna `false` altrimenti restituisce `true`. 
- `(p *punto) presenzaOstacoloPercorsoY(x int) bool` : Questo metodo controlla se sull'asse delle $y$ del punto considerato ci sono ostacoli, in particolare se trova ostacoli sull'asse allora ritorna `false` altrimenti restituisce `true` 

## Esempi di esecuzione e casi limite
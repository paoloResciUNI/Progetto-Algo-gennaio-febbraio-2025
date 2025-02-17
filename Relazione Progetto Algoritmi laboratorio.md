# Relazione progetto d'esame di algoritmi e strutture dati (revisione)

### Introduzione

Il progetto implementa un sistema per la gestione del movimento di automi puntiformi su un piano, rispettando vincoli di ostacoli e segnali di richiamo. L'obiettivo è studiare il comportamento degli automi in un contesto dove vi sono ostacoli e richiami che definiscono un percorso minimo che l'automa dovrà intraprendere, se questo esiste. I movimenti possibili di un'automa in posizione $A(x_A, y_A)$ ad un richiamo di posizione $R(x_R, y_R)$ devono essere compresi nella distanza $D(A, R) = |x_R - x_A|+|y_R-y_A|$.

### Strutture dati e scelte progettuali

Per la rappresentazione del piano si è utilizzata una **lista doppiamente concatenata**. Questa scelta permette di gestire dinamicamente l'aggiunta e la modifica di automi e l'aggiunta di ostacoli. Inoltre, è una struttura relativamente leggera in termini di consumo di memoria. Gli automi sono salvati nella parte superiore della lista mentre gli ostacoli nella parte inferiore. Questo permette di cercare gli automi e gli ostacoli in maniera più efficiente.

#### Strutture dati principali

- **`punto`**: rappresenta un nodo del piano, contenente le coordinate `(x, y)`, un identificativo `id` e i riferimenti al nodo precedente e successivo.
- **`Piano`**: struttura principale che mantiene riferimenti ai nodi iniziale e finale della lista doppiamente concatenata.
- **`piano`**: alias di un tipo puntatore ad una variabile `Piano`.
- **`nodoPila`**: usato per gestire le operazioni di richiamo degli automi, memorizzando i candidati allo spostamento. Questa struttura è usata solamente nel metodo `richiamo`.

### Implementazione delle operazioni 

Le operazioni implementate nel programma seguono le specifiche fornite. Di seguito una descrizione delle principali operazioni.
 
- **`avanza(Campo piano, p *punto, Sorgente *punto) *punto`**: Questa funzione simula lo spostamento dell'automa verso la sorgente del richiamo. È una funzione ricorsiva che come caso base ha $D = 0$, dove $D$ è la distanza di Manhattan fra `p` e `Sorgente`. La funzione `avanza` restituisce il punto nel quale la simulazione si è fermata, ovvero se l'automa ha raggiunto la sorgente del segnale oppure è andato in stallo (non ci sono percorsi liberi). Questa funzione dispone dei metodi `forwardX` e `forwardY` che spostano l'automa su un rispettivo asse, in modo che, se non raggiunge direttamente la sorgente del richiamo, esso possa spostarsi in futuro. L'`avanza` assume che `forwardX` e `forwardY` spostino sempre l'automa, se ciò non accade l'automa è andato in stallo e la funzione termina. Altri dettagli sulla logiaca di forward sono forniti più avanti nel paragrafo. La funzione `avanza` calcola la distanza dagli ostacoli sull'asse orizzontale e verticale che l'automa ha in un detrminato momento, in caso non vi siano ostacoli su uno dei due assi si utilizza la distanzache fra le rispettive coordinate  dei due punti, se la distanza è maggiore per l'asse verticale, allora verrà usata la `forwardY` altrimenti la `forwardX`. Un caso particolare si ha quando la distanza dagli ostacoli è uguale su entrambi gli assi. In questo caso si esegue un passo unitario sull'asse orizzontale, se questo non è possibile allora l'automa è andato in stallo e la funzione termina.  
    >Nota: Dato che la logica per `forwardX` e `forwardY` è la stessa, di seguito viene fornita esclusivamente la spiegazione di `forwardX`. 
La `forwardX` prende in ingresso due punti, `start` e `destination`, e restituisce il punto che si trova sullo stesso asse orizzontale di `start`, il più vicino possibile all'asse verticale di `destination` ma facendo in modo che dal punto restituito si possa fare uno spostamento varticale, senza ostacoli in mezzo. Questo punto restituito verrà chiamato d'ora in poi $\lambda$. In pasrticolare `forwardX` si divide in tre macro controlli: 
    - Sull'asse verticale di `start`: si controlla se c'è un ostacolo e si procede ad aggirarlo posizionando $\lambda$ sulla coordinata $x$ successiva ad esso. Si restituisce $\lambda$.  
    - Sull'asse orizzontale di `start`: questo controllo viene fatto solo se non ci sono ostacoli sull'asse verticale. Se c'è un'ostacolo sull'asse orizzontale allora $\lambda$ viene posizionato subito prima di esso altrimenti viene posizionato sull'asse $x$ di `destination`. $\lambda$ non viene ancora restituito. 
    - Sull'asse verticale di $\lambda$: 



- **`richiamo(x, y, alpha)`**: emette un segnale che richiama gli automi con prefisso `alpha` verso il punto `(x, y)`. Gli automi si spostano solatnto se hanno distanza $D$ minore rispetto agli altri automi richiamati e se il loro percorso non è bloccatto da ostacoli. Questa operazione è implementata dal metodo `richiamo(x, y int, alpha string)`.
- **`esistePercorso(x, y, eta)`**: verifica se esiste un percorso libero di distanza minima da un'automa alla destinazione. Se il percorso esiste viene stampato a schermo `SI` altrimenti `NO`. L'operazione `esistePercorso(x, y, eta)` è implementata dal metodo `esistePercorso(x, y int, eta string)`.

### Funzioni e metodi principali

- **`esegui(p piano, s string)`**: interpreta ed esegue i comandi ricevuti da standard input. Permette di eseguire le operazioni precedentemente descritte.
- **`newPiano() piano`**: inizializza una nuova struttura `piano`.
- **`(Piano)stampa()`**: stampa automi e ostacoli. Scorre la lista dall'inizio per stampare gli automi mentre per gli ostacoli scorre la lista dalla fine. 
- **`(Piano)stato(x, y int)`**: restituisce informazioni sulla posizione `(x, y)`. Questo metodo scorre la lista dall'inizio e controlla prima che le coordinate `(x, y)` contengano un'automa (in questo caso viene stampato su standard output il carattere `A`), poi controlla se faccia parte dell'area di un ostacolo (in questo caso viene stampato su standard output il carattere `O`) e se nessuno dei casi precedenti si è verificato, allora viene stampato su standard output il carattere `E`.  
- **`(*Piano)posizioni(alpha string)`**: stampa le posizioni degli automi con prefisso `alpha`.
- **`(Campo *Piano)automa(x, y int, eta string)`**: aggiunge un'automa o ne modifica la posizione.
- **`(Campo *Piano)ostacolo(x0, y0, x1, y1 int)`**: aggiunge un ostacolo.
- **`(Campo *Piano)richiamo(x, y int, alpha string)`**: gestisce il richiamo degli automi compatibili.
- **`(Campo *Piano)esistePercorso(x, y int, eta string)`**: verifica se un'automa `eta` ha un percorso libero di distanza $D$ verso il punto `(x, y)`.

### Metodi e funzioni secondarie 


### Analisi delle prestazioni

L'uso di una **lista doppiamente concatenata** consente di mantenere basso l'uso di memoria. La gestione delle operazioni principali avviene con complessità ottimizzata:

- **Inserimento di automi**: tempo medio .
- **Inserimento di ostacoli**: tempo medio .
- **Verifica di percorsi liberi**: tempo peggiore .
- **Gestione dei richiami**: tempo peggiore .


### Esempi di esecuzione

#### **Esempio 1: Inserimento di automi e ostacoli**

```
c
a 2 3 101
a 5 6 110
o 4 4 6 6
S
```

**Output atteso:**

```
(
101: 2,3
110: 5,6
)
[
(4,4)(6,6)
]
```

#### **Esempio 2: Verifica esistenza percorso libero**

```
c
a -2000 -200 101
o 3 2 5 4
e 6 2 101
```

**Output atteso:**

```
NO
```

#### **Caso limite 1: Automa circondato da ostacoli**

```
c
a 7 4 11001
a 10 6 001
a 5 5 101
o 3 2 5 4
o 2 1 10 4
o 8 5 12 10
o 0 7 6 15
r 16 1 1
p 1
```

**Output atteso:**

```
(
101: 5,5
11001: 16,1
)
```

### Conclusione

Il progetto implementato è conforme alle specifiche e fornisce un'efficace gestione degli automi e degli ostacoli nel piano. Ulteriori miglioramenti potrebbero includere **ottimizzazioni sugli algoritmi di percorso** per ridurre ulteriormente il tempo di esecuzione in scenari complessi.
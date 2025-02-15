# Relazione progetto d'esame di algoritmi e strutture dati

### Introduzione

Il progetto implementa un sistema per la gestione del movimento di automi puntiformi su un piano, rispettando vincoli di ostacoli e segnali di richiamo. L'obiettivo è studiare il comportamento degli automi in un contesto dove vi sono ostacoli e richiami che definiscono un percorso minimo che l'automa dovrà intraprendere, se questo esiste. I movimenti possibili di un'automa in posizione $A(x_A, y_A)$ ad un richiamo di posizione $R(x_R, y_R)$ devono essere compresi nella distanza $D(A, R) = |x_R - x_A|+|y_R-y_A|$.

### Strutture dati e scelte progettuali

Per la rappresentazione del piano si è utilizzata una **lista doppiamente concatenata**. Questa scelta permette di gestire dinamicamente l'aggiunta e la modifica di automi e l'aggiunta di ostacoli. Inoltre, è una struttura relativamente leggera in termini di consumo di memoria. Gli automi sono salvati nella parte superiore della lista mentre gli ostacoli nella parte inferiore. Questo permette di cercare gli automi e gli ostacoli in maniera più efficiente.

#### Strutture dati principali

- **`punto`**: rappresenta un nodo del piano, contenente le coordinate `(x, y)`, un identificativo `id` e i riferimenti al nodo precedente e successivo.
- **`Piano`**: struttura principale che mantiene riferimenti ai nodi iniziale e finale della lista doppiamente concatenata.
- **`piano`**: alias di un tipo puntatore ad una variabile `Piano`.
- **`nodoPila`**: usato per gestire le operazioni di richiamo degli automi, memorizzando i candidati allo spostamento. Questa struttura è usata solamente nel metodo `richiamo`.

Questa modellazione consente un accesso rapido agli automi e agli ostacoli, facilitando operazioni come la ricerca, l'inserimento e la gestione dei percorsi.

### Implementazione delle operazioni 

Le operazioni implementate nel programma seguono le specifiche fornite. Di seguito una descrizione delle principali operazioni.
 
>*P.S.: Qui di seguito è presente solamente la descrizione delle principali operazioni implementate. L'implementazione vera e propria è descritta dettagliatamente nella sezione relativa a metodi e funzioni*.

#### Creazione e gestione del piano

- **`crea()`**: crea un nuovo piano vuoto. Se vi era già un piano vengono eliminati i dati esistenti. In particolare vengono assegnati due puntatori vuoti al campo `inizio` e al campo `fine` del piano esistente. L'operazione `crea()` è impleemntata grazie alla funzione `newPiano()`.
- **`stampa()`**: stampa la lista degli automi e degli ostacoli secondo il formato output richiesto. Questa operazione è implementata dal metodo `stampa()`.
- **`stato(x, y)`**: restituisce il contenuto della posizione `(x, y)` (`A` per automi, `O` per ostacoli, `E` per posizioni vuote). Questa operazione è implemenata dal metodo `stato(x, y int)`.
- **`posizioni(alpha)`**: stampa la posizione degli automi che nell'id hanno prefisso `alpha`. L'operazione è implementata dal metodo `posizioni(alpha string)`.

#### Inserimento e gestione di automi e ostacoli

- **`automa(x, y, eta)`**: permette di posizionare un nuovo automa o di riposizionarne uno già presente. Infatti se l'automa `eta` esiste già viene riposizionato in `(x, y)`. Questa operazione viene implementata dal metodo `automa(x, y int, eta string)`.
- **`ostacolo(x0, y0, x1, y1)`**: aggiunge un nuovo ostacolo se non vi sono automi nella sua area. L'ostacolo viene distinto dagli automi per il suo `id`. All'interno dell'`id` dell'ostacolo vengono infatti salvate: le coordinate del punto più in basso a sinistra e del punto più in alto a destra seguite dalla stringa `"ostacolo"`. L'operazione di inserimento degli ostacoli all'interno del piano è implementata dal metodo `ostacolo(x0, y0, x1, y1)`.

#### Movimenti e richiami

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
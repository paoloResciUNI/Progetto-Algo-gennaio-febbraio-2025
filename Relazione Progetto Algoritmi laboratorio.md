# Paolo Rescigno

**Numero di matricola: 45198A**

## Relazione progetto d'esame di algoritmi e strutture dati

### Introduzione

Il progetto implementa un sistema per la gestione del movimento di automi puntiformi su un piano, rispettando vincoli di ostacoli e segnali di richiamo. L'obiettivo è studiare il comportamento degli automi in un contesto dove vi sono ostacoli e richiami che definiscono un percorso minimo che l'automa dovrà intraprendere, se questo esiste. I movimenti possibili di un'automa in posizione ad un richiamo di posizione devono essere compresi nella distanza:

---

### Strutture dati e scelte progettuali

Per la rappresentazione del piano si è utilizzata una **lista doppiamente concatenata**. Questa scelta permette di gestire dinamicamente l'aggiunta e la modifica di automi e l'aggiunta di ostacoli. Inoltre, è una struttura relativamente leggera in termini di consumo di memoria. Gli automi sono salvati nella parte superiore della lista mentre gli ostacoli nella parte inferiore. Questo permette di cercare gli automi e gli ostacoli in maniera più efficiente.

#### Strutture dati principali

- **punto**: rappresenta un nodo del piano, contenente coordinate , un identificativo (**id**) e riferimenti al nodo precedente e successivo.
- **piano**: alias di un tipo puntatore ad una variabile **Piano**.
- **Piano**: struttura principale che mantiene riferimenti ai nodi iniziale e finale della lista doppiamente concatenata.
- **nodoPila**: usato per gestire le operazioni di richiamo degli automi, memorizzando i candidati allo spostamento.

Questa modellazione consente un accesso rapido agli automi e agli ostacoli, facilitando operazioni come la ricerca, l'inserimento e la gestione dei percorsi.

---

### Implementazione delle operazioni

#### Creazione e gestione del piano

- `newPiano()`: crea un nuovo piano vuoto, eliminando i dati esistenti.
- `stampa()`: stampa la lista degli automi e degli ostacoli secondo il formato richiesto.
- `stato(x, y)`: restituisce il contenuto della posizione  
  **(A per automi, O per ostacoli, E per posizioni vuote)**.

#### Inserimento e gestione di automi e ostacoli

- `automa(x, y, eta)`: posiziona un nuovo automa o riposiziona uno esistente.
- `ostacolo(x0, y0, x1, y1)`: aggiunge un nuovo ostacolo se non vi sono automi nella sua area.

#### Movimenti e richiami

- `richiamo(x, y, alpha)`: emette un segnale che richiama gli automi compatibili verso .
- `esistePercorso(x, y, eta)`: verifica se esiste un percorso libero minimo da un automa alla destinazione .

---

### Funzioni principali

- `esegui(p piano, s string)`: interpreta ed esegue i comandi ricevuti.
- `newPiano() piano`: inizializza una nuova struttura **Piano**.
- `stampa()`: stampa automi e ostacoli.
- `stato(x, y)`: restituisce informazioni sulla posizione .
- `posizioni(alpha)`: stampa le posizioni degli automi con prefisso **alpha**.
- `automa(x, y, eta)`: aggiunge o modifica la posizione di un automa.
- `ostacolo(x0, y0, x1, y1)`: aggiunge un ostacolo.
- `richiamo(x, y, alpha)`: gestisce il richiamo degli automi compatibili.
- `esistePercorso(x, y, eta)`: verifica se un automa ha un percorso libero verso .

---

### Analisi delle prestazioni

L'uso di una **lista doppiamente concatenata** consente di mantenere basso l'uso di memoria. La gestione delle operazioni principali avviene con complessità ottimizzata:

- **Inserimento di automi**: tempo medio .
- **Inserimento di ostacoli**: tempo medio .
- **Verifica di percorsi liberi**: tempo peggiore .
- **Gestione dei richiami**: tempo peggiore .

---

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

---

### Conclusione

Il progetto implementato è conforme alle specifiche e fornisce un'efficace gestione degli automi e degli ostacoli nel piano. Ulteriori miglioramenti potrebbero includere **ottimizzazioni sugli algoritmi di percorso** per ridurre ulteriormente il tempo di esecuzione in scenari complessi.

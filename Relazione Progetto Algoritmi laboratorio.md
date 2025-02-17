# Relazione progetto d'esame di algoritmi e strutture dati (revisione)

### Introduzione

Questa relazione presenta le specifiche delle funzioni implementate allo scopo di risolvere il problema dato nella traccia. Nella relzione si farà riferimento alla distanza di Manhattan con $D$, al numero di automi nel campo con $a$ e al numero di ostacoli con $m$.

## Strutture dati e scelte progettuali

Per rappresentare il piano sono state usate due liste concatenate, una per contenere gli automi e un'altra per gli ostacoli. Questa scelta permette di gestire dinamicamente l'aggiunta e la modifica di automi e di ostacoli. Inoltre, la lista è una struttura relativamente leggera in termini di consumo di memoria.

#### Strutture dati principali

- **`punto`**: rappresenta un nodo del piano, contenente le coordinate `x` e `y`, un identificativo `id` e un riferimento al nodo e successivo.
- **`Piano`**: struttura principale che mantiene riferimenti ai nodi iniziali delle due liste.
- **`piano`**: alias di un tipo puntatore ad una variabile `Piano`.
- **`nodoPila`**: usato per gestire le operazioni di richiamo degli automi, memorizzando i candidati allo spostamento. Questa struttura è usata solamente nel metodo `richiamo`.

### Implementazione delle operazioni

**`avanza(Campo piano, p *punto, Sorgente *punto) *punto`**: Questa funzione simula lo spostamento dell'automa verso la sorgente del richiamo. È una funzione ricorsiva che come caso base ha $D = 0$, dove $D$ è la distanza di Manhattan fra `p` e `Sorgente`. La funzione `avanza` restituisce il punto nel quale la simulazione si è fermata, ovvero se l'automa ha raggiunto la sorgente del segnale oppure è andato in stallo (non ci sono percorsi liberi). Questa funzione dispone dei metodi `forwardX` e `forwardY` che spostano l'automa su un rispettivo asse, in modo che, se non raggiunge direttamente il richiamo, esso possa avanzare ulteriormente nella funzione. L'`avanza` assume che `forwardX` e `forwardY` spostino sempre l'automa, se ciò non accade l'automa è andato in stallo e la funzione termina. Altri dettagli sulla logica di forward sono forniti più avanti nel paragrafo. La funzione `avanza` calcola la distanza dagli ostacoli sull'asse orizzontale e verticale che l'automa ha in un detrminato momento, in caso non vi siano ostacoli su uno dei due assi si utilizza la distanza
 fra le rispettive coordinate  dei due punti, se la distanza è maggiore per l'asse verticale, allora verrà usata la `forwardY` altrimenti la `forwardX`. Un caso particolare si ha quando la distanza dagli ostacoli è uguale su entrambi gli assi. In questo caso si esegue un passo unitario sull'asse orizzontale, se questo non è possibile allora l'automa è andato in stallo e la funzione termina.  

> Nota: Dato che la logica per `forwardX` e `forwardY` è la stessa, di seguito viene fornita esclusivamente la spiegazione di `forwardX`.

La `forwardX` prende in ingresso due punti, `start` e `destination`, e restituisce il punto che si trova sullo stesso asse orizzontale di `start`, il più vicino possibile all'asse verticale di `destination` ma facendo in modo che dal punto restituito si possa fare uno spostamento varticale, senza ostacoli in mezzo. Questo punto restituito verrà chiamato d'ora in poi $\lambda$. In particolare `forwardX` si divide in tre macro controlli: 

- Sull'asse verticale di `start`: si controlla se c'è un ostacolo e si procede ad aggirarlo posizionando $\lambda$ sulla coordinata $x$ successiva ad esso. Si restituisce $\lambda$.  
- Sull'asse orizzontale di `start`: questo controllo viene fatto solo se non ci sono ostacoli sull'asse verticale. Se c'è un ostacolo sull'asse orizzontale allora $\lambda$ viene posizionato subito prima di esso, altrimenti viene posizionato sull'asse $x$ di `destination`. $\lambda$ non viene ancora restituito dal metodo. 
- Sull'asse verticale di $\lambda$: dopo aver posizionato $\lambda$ si controlla se sul suo asse verticale si incontrano ostacoli, se ciò non avviene $\lambda$ viene restituito, altrimenti si fa retrocedere (verso la $x$ di `start`) in modo che venga aggirato l'ostacolo. $\lambda$ viene restituito. 

**TODO** parlare delle funzioni: `richiamo` e `esistePercorso`.

#### Breve descrizione dell'implementazione delle operazioni richieste

### Analisi delle prestazioni

**TODO** inserire spiegazione. 

- **Inserimento di automi**: tempo medio $O(1)$.
- **Inserimento di ostacoli**: tempo medio $O(1)$.
- **Verifica di percorsi liberi**: caso peggiore $O(D \cdot m)$ .
- **Gestione dei richiami**: tempo peggiore $O(a \cdot m \cdot D)$.

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
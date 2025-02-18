# Relazione progetto d'esame di algoritmi e strutture dati (revisione)

## Introduzione

Questa relazione presenta le specifiche delle funzioni implementate allo scopo di risolvere il problema dato nella traccia. Nella relzione si farà riferimento alla distanza di Manhattan fra due punti del piano con $D$, al numero di automi nel piano con $a$ e al numero di ostacoli con $m$.

## Strutture dati e scelte progettuali

Per rappresentare il piano è stata usata una struttura con riferimento a due liste concatenate, una per contenere gli automi e un'altra per gli ostacoli. Questa scelta permette di gestire dinamicamente l'aggiunta e la modifica di automi e di ostacoli. Inoltre, la lista è una struttura relativamente leggera in termini di consumo di memoria e facile da manipolare. 

### Strutture dati principali

- **`punto`**: rappresenta un punto del piano, contenente le coordinate `x` e `y`, un identificativo `id` e un puntatore ad un tipo `punto`. Questo tipo di dato viene usato per rappresentare sia automi che ostacoli 
- **`Piano`**: struttura principale che mantiene riferimenti ad una lista di ostacoli e ad una di automi.
- **`piano`**: alias di un tipo puntatore ad una variabile `Piano`.
- **`elementoPila`**: struttura usata per gestire l'operazione di richiamo, memorizzando gli automi candidati allo spostamento. Composta da: un tipo `*punto`, che rappresenta l'automa candidato, un tipo `int` che rappresenta la distaza che il candidato ha dal richiamo e un tipo `*elementoPila` che collega la struttura ad una pila. Questa struttura è usata esclusivamente nel metodo `richiamo`.

### Moviementi degli automi

**`avanza(Campo piano, p *punto, Sorgente *punto) *punto`**: Questa funzione simula lo spostamento dell'automa verso la sorgente del richiamo. È una funzione ricorsiva che come caso base ha $D = 0$, dove $D$ è la distanza di Manhattan fra `p` e `Sorgente`. La funzione `avanza` restituisce il punto nel quale la simulazione si è fermata, ovvero se l'automa ha raggiunto la sorgente del segnale oppure è andato in stallo (non ci sono percorsi liberi). Questa funzione dispone dei metodi `forwardX` e `forwardY` che spostano l'automa su un rispettivo asse, in modo che, se non raggiunge direttamente il richiamo, esso possa avanzare ulteriormente nella funzione. L'`avanza` assume che `forwardX` e `forwardY` spostino sempre l'automa, se ciò non accade l'automa è andato in stallo e la funzione termina. Altri dettagli sulla logica di forward sono forniti più avanti nel paragrafo.
La funzione `avanza` calcola la distanza dagli ostacoli sull'asse orizzontale e verticale che l'automa ha in un determinato momento, in caso non vi siano ostacoli su almeno uno dei due assi si utilizza la distanza fra le rispettive coordinate dei due punti, se la distanza è maggiore per l'asse verticale, allora verrà usata la `forwardY` altrimenti la `forwardX`. Un caso particolare si ha quando la distanza dagli ostacoli è uguale su entrambi gli assi. In questo caso si esegue un passo unitario sull'asse orizzontale, se questo non è possibile allora l'automa è andato in stallo e la funzione termina.

La funzione `avanza` impiega $O(D^2 \cdot m)$ questo perchè nel caso peggiore dovrà fare $D$ passi ricorsivi e quindi eseguire altrettante volte la logica di forwarding.

> Nota: Dato che la logica per `forwardX` e `forwardY` è la stessa, di seguito viene fornita esclusivamente la spiegazione di `forwardX`.

La `forwardX` prende in ingresso due punti, `start` e `destination`, e restituisce il punto che si trova sullo stesso asse orizzontale di `start`, il più vicino possibile all'asse verticale di `destination` ma facendo in modo che dal punto restituito si possa fare uno spostamento varticale, senza ostacoli in mezzo. Questo punto restituito verrà chiamato d'ora in poi $\lambda$. In particolare `forwardX` si divide in due macro controlli:

- Sull'asse verticale di `start`: si controlla se c'è un ostacolo e si procede ad aggirarlo posizionando $\lambda$ sulla coordinata $x$ successiva ad esso. Si restituisce $\lambda$ e `forwardX` termina. In caso non ci siano ostacoli sull'asse verticale di `start` viene fatto un controllo sul suo asse orizzontale.
- Sull'asse orizzontale di `start`: questo controllo viene fatto solo se non ci sono ostacoli sull'asse verticale. Se c'è un ostacolo sull'asse orizzontale allora $\lambda$ viene posizionato subito prima di esso, altrimenti viene posizionato sull'asse $x$ di `destination`. Dopo aver posizionato $\lambda$ si controlla se sul suo asse verticale si incontrano ostacoli, se ciò non avviene $\lambda$ viene restituito, altrimenti si fa retrocedere (verso la $x$ di `start`) in modo che venga aggirato l'ostacolo. A questo punto $\lambda$ viene restituito.

Il mtodo `fowardX`, quindi anche `forwardY`, impiega $O(m\cdot D)$. Questo tempo è dovuto al fatto che per ogni punto dell'asse verticale e dell'asse orizzontale, che il metodo deve controllare, scorre, nel caso peggiore, tutta la lista degli ostacoli del piano.

L'operazione `richiamo` è implementata da un metodo omonimo. Questo metodo controlla i punti più vicini al richiamo e li fa spostare verso di esso. Per richiamare gli automi si controlla che il prefisso dell'id dell'automa corrisponda al richiamo, se ciò avviene si esegue la funzione `avanza` per simulare lo spostamento dell'automa verso il richiamo. Quando l'avanza termina, se il punto restituito ha le stesse coordinate del richiamo, viene inserito all'interno di una struttura `elementoPila`.  Dopo aver inserito tutti gli automi che si possono spostare all'intrno della pila si controllano le distanze minime che essi hanno dal richiamo. Verranno effettivamente spostati solo quelli con distanza minima. Il tempo d'esecusione del metodo `richiamo` è $O(a \cdot D^2 \cdot m)$ nel caso peggiore, ovvero se tutti gli automi del piano si possono muovere, sono tutti di eguale distanza dal richiamo e devono per forza eseguire $D$ passi ricorsivi di `avanza`.

### Implementazione e tempi delle altre operazioni richieste

L'operazione `crea` viene implementata restituendo una nuova variabile di tipo `Piano`, se questa non esiste già, altrimenti sostituisce i puntatori alle liste degli ostacoli e degli automi del piano con puntatori vuoti. Questa operazione impiega tempo costante essere eseguita.

L'operazione `stato` viene implementata scorrendo la lista degli automi e degli ostacoli. L'operazione richiede tempo pari a $O(a+m)$ nel caso peggiore.

L'operazione `stampa` scorre entrambe le liste, degli ostacoli e degli automi, del piano. Impiega tempo pari a $\Theta(a+m)$

L'operazione `automa` viene implementata scorrendo la lista degli ostacoli e la lista degli automi facenti parte della struttura `Piano`. Questa operazione impiega $\Theta(a+m)$.

L'operazione `ostacolo` viene implementata scorrendo esclusivamente la lista degli automi del piano. Questa operazione impiega $\Theta(a)$.

L'operazione `posizioni` è implementata scorrendo la lista degli automi e verificando che l'automa considerato abbia il prefisso giusto. Questa operazione impiega $O(a)$.

L'operazione `esistePercorso` è implementata controllando prima che l'id cercato corrisponda ad un automa esistenete e poi controllando se il punto di arrivo non corrisponda all'area di un ostacolo. Se almeno una di queste condizioni non si verifica viene stampato su standard output `NO`. Altrimenti viene usata la funzione `avanza`. La funzione prende in ingresso l'automa e il punto d'arrivo e restituisce un punto. Se il punto restituito dalla funzione `avanza` ha le stesse coordinata del punto d'arrivo viene stampato su standard output `SI`, altrimenti `NO`.

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
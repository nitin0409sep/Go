STEP-BY-STEP EXECUTION (CHANNEL + GOROUTINE)

## CODE:

func processNum(numChan chan int) { for num := range numChan { fmt.Println("processing num", num) time.Sleep(time.Second \* 2) } }

func main() { numChan := make(chan int) go processNum(numChan)

    for {
        numChan <- rand.Intn(100)
    }

## }

STEP 1: PROGRAM START

- Go runtime starts
- main() goroutine is created automatically

Goroutines:

- main

STEP 2: CHANNEL CREATION

- make(chan int) creates a channel object in memory
- Channel is unbuffered (buffer size = 0)

Important:

- numChan is NOT the channel
- numChan is a reference to the channel object

Memory: numChan ───► [ channel object ]

STEP 3: STARTING A GOROUTINE

- go processNum(numChan) creates a new goroutine
- A COPY of the channel reference is passed

Memory: main.numChan ───► [ channel object ] process.numChan ───► [ channel object ]

Both goroutines share the SAME channel

STEP 4: processNum STARTS EXECUTING

- for num := range numChan tries to RECEIVE
- No value is available yet

Result:

- processNum BLOCKS waiting for data

STEP 5: main ENTERS INFINITE LOOP

- main tries to SEND a random number

numChan <- rand.Intn(100)

STEP 6: CHANNEL SYNCHRONIZATION

- Channel is unbuffered
- Send blocks until a receiver is ready
- Receive blocks until a sender is ready

Runtime sees:

- main is sending
- processNum is receiving

Result:

- Value is transferred
- Both goroutines UNBLOCK

STEP 7: processNum RECEIVES VALUE

- num := <-numChan succeeds
- Value is printed

Output: processing num 42

STEP 8: processNum SLEEPS

- time.Sleep(2s)
- processNum stops receiving during sleep

STEP 9: main TRIES TO SEND AGAIN

- main sends next random number
- No receiver available

Result:

- main BLOCKS

STEP 10: SLEEP ENDS

- processNum loops again
- Tries to receive from channel

STEP 11: SECOND SYNCHRONIZATION

- main is blocked on send
- processNum is ready to receive

Result:

- Value transferred
- Both unblock
- processNum prints and sleeps again

STEP 12: LOOP CONTINUES FOREVER

- One number processed every 2 seconds
- main can NEVER outrun processNum

## WHY NO POINTER IS NEEDED

- Channels are reference types
- Passing a channel copies the reference, NOT the channel
- Both goroutines talk through the same pipe

You are NOT updating numChan You are sending values THROUGH numChan

## WHEN POINTER TO CHANNEL IS NEEDED (RARE)

Only if you want to REPLACE the channel itself:

func reset(ch *chan int) { *ch = make(chan int) }

## FINAL ONE-LINE TAKEAWAY

Channels don’t need pointers because they already reference shared, synchronized communication state between goroutines.

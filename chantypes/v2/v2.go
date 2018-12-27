package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

const (
	timeFormat = "15:04:05.000" // time.StampMilli
	timeUnit   = time.Millisecond

	generatorSleep = 10
	filterSleep    = 20
	adderSleep     = 30
)

var total = 0
var wg = sync.WaitGroup{}

type Number struct {
	no    int
	start time.Time
	end   time.Time
}

func (n *Number) String() string {
	return fmt.Sprintf("no=%d, taken=%v, start=%s, end=%s", n.no, n.end.Sub(n.start), n.start.Format(timeFormat), n.end.Format(timeFormat))
}

func main() {
	start := time.Now()
	log.Printf("main, total=%d, start=%s\n", total, start.Format(timeFormat))

	noCh := make(chan Number)
	evenCh := make(chan Number)

	wg.Add(3)
	go adder(evenCh)
	go filter(noCh, evenCh)
	generator(noCh)

	wg.Wait()

	end := time.Now()
	log.Printf("main, total=%d, taken=%s, end=%s\n", total, end.Sub(start), end.Format(timeFormat))

	if total != 30 {
		log.Fatalf("total want=30, got=%d\n", total)
	}
}

func generator(noCh chan<- Number) {
	for i := 1; i <= 10; i++ {
		generatorWork()
		noCh <- Number{no: i, start: time.Now()}
	}
	close(noCh)
	wg.Done()
}

func filter(noCh <-chan Number, evenCh chan<- Number) {
	for num := range noCh {
		if num.no%2 == 0 {
			filterWork()
			evenCh <- num
		}
	}
	close(evenCh)
	wg.Done()
}

func adder(evenCh <-chan Number) {
	for evenNum := range evenCh {
		adderWork()
		total += evenNum.no
		evenNum.end = time.Now()
		log.Printf("adder, total=%d, got=%s\n", total, evenNum.String())
	}
	wg.Done()
}

func generatorWork() {
	time.Sleep(timeUnit * generatorSleep)
}

func filterWork() {
	time.Sleep(timeUnit * filterSleep)
}

func adderWork() {
	time.Sleep(timeUnit * adderSleep)
}

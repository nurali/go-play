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

	chanBuffSize = 6
)

var total = 0
var wg = sync.WaitGroup{}

type Number struct {
	no    int
	start time.Time
	end   time.Time
}

func (n *Number) String() string {
	return fmt.Sprintf("no=%d, %s", n.no, formatTime(n.end, n.start))
}

func formatTime(end, start time.Time) string {
	return fmt.Sprintf("taken=%v, start=%s, end=%s", end.Sub(start), start.Format(timeFormat), end.Format(timeFormat))
	// return fmt.Sprintf("taken=%v", end.Sub(start))
}

func formatTimeUnit(timeUnit time.Duration) string {
	switch {
	case timeUnit == time.Millisecond:
		return "ms"
	case timeUnit == time.Second:
		return "sec"
	}
	return "unknown"
}

func main() {
	start := time.Now()
	fmt.Printf("main, start, total=%d, chanBuffSize=%d, timeUnit=%s, generator=%d, filter=%d, adder=%d\n",
		total, chanBuffSize, formatTimeUnit(timeUnit), generatorSleep, filterSleep, adderSleep)

	noCh := make(chan Number, chanBuffSize)
	evenCh := make(chan Number, chanBuffSize)

	wg.Add(3)
	go adder(evenCh)
	go filter(noCh, evenCh)
	generator(noCh)

	wg.Wait()

	end := time.Now()
	fmt.Printf("main, end, total=%d, %s\n", total, formatTime(end, start))

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
		fmt.Printf("adder, total=%d, got=%s\n", total, evenNum.String())
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

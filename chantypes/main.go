package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

const (
	timeFormat = ".000" // time.StampMilli
	timeUnit   = time.Millisecond

	generatorSleep = 10
	validatorSleep = 20
	adderSleep     = 30
)

var total = 0
var wg = sync.WaitGroup{}

type work struct {
	id    int
	start time.Time
	end   time.Time
}

func timeTaken(end, start time.Time) string {
	// return fmt.Sprintf("taken=%v ms, start=%s, end=%s", end.Sub(start).Nanoseconds()/(1000*1000), start.Format(timeFormat), end.Format(timeFormat))
	return fmt.Sprintf("taken=%v", end.Sub(start).Nanoseconds()/(1000*1000))
}

func main() {
	fmt.Printf("start, timeUnit=%s, generator=%d, validator=%d, adder=%d\n\n",
		"ms", generatorSleep, validatorSleep, adderSleep)
	for buffSize := 1; buffSize <= 5; buffSize++ {
		total = 0
		run(buffSize)

		time.Sleep(time.Millisecond * 10)
		fmt.Println()
	}
}

func run(chanBuffSize int) {
	start := time.Now()
	fmt.Printf("start, total=%d, buffChanSize=%d\n", total, chanBuffSize)

	ch1 := make(chan work, chanBuffSize)
	ch2 := make(chan work, chanBuffSize)

	wg.Add(3)
	go adder(ch2)
	go validator(ch1, ch2)
	generator(ch1)

	wg.Wait()

	if total != 15 {
		log.Fatalf("total want=15, got=%d\n", total)
	}

	end := time.Now()
	fmt.Printf("end, %s\n", timeTaken(end, start))
}

func generator(ch1 chan<- work) {
	for i := 1; i <= 5; i++ {
		w := work{id: i, start: time.Now()}
		generatorWork()
		ch1 <- w
	}
	close(ch1)
	wg.Done()
}

func validator(ch1 <-chan work, ch2 chan<- work) {
	for w := range ch1 {
		if w.id > 0 {
			validatorWork()
			ch2 <- w
		}
	}
	close(ch2)
	wg.Done()
}

func adder(ch2 <-chan work) {
	for w := range ch2 {
		adderWork()
		total += w.id
		w.end = time.Now()
		fmt.Printf("adder, %s\n", timeTaken(w.end, w.start))
	}
	wg.Done()
}

func generatorWork() {
	time.Sleep(timeUnit * generatorSleep)
}

func validatorWork() {
	time.Sleep(timeUnit * validatorSleep)
}

func adderWork() {
	time.Sleep(timeUnit * adderSleep)
}

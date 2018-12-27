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

	validatorSleep = 30
	adderSleep     = 20
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
	for buffSize := 1; buffSize <= 6; buffSize++ {
		total = 0
		run(buffSize)

		time.Sleep(time.Millisecond * 10)
		fmt.Println()
	}
}

func run(chanBuffSize int) {
	start := time.Now()
	fmt.Printf("start, total=%d, buffChanSize=%d, timeUnit=%s, validator=%d, adder=%d\n",
		total, chanBuffSize, "ms", validatorSleep, adderSleep)

	ch1 := make(chan work, chanBuffSize)
	ch2 := make(chan work, chanBuffSize)

	wg.Add(3)
	go adder(ch2)
	go validator(ch1, ch2)
	generator(ch1)

	wg.Wait()

	end := time.Now()
	fmt.Printf("end, %s\n", timeTaken(end, start))

	if total != 15 {
		log.Fatalf("total want=15, got=%d\n", total)
	}
}

func generator(ch1 chan<- work) {
	for i := 1; i <= 5; i++ {
		ch1 <- work{id: i, start: time.Now()}
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

func validatorWork() {
	time.Sleep(timeUnit * validatorSleep)
}

func adderWork() {
	time.Sleep(timeUnit * adderSleep)
}

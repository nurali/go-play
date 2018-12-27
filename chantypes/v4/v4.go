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
	validatorSleep = 30
	adderSleep     = 20
)

var total = 0
var wg = sync.WaitGroup{}

type Number struct {
	no    int
	start time.Time
	end   time.Time
}

func formatTime(end, start time.Time) string {
	// return fmt.Sprintf("taken=%v ms, start=%s, end=%s", end.Sub(start).Nanoseconds()/(1000*1000), start.Format(timeFormat), end.Format(timeFormat))
	return fmt.Sprintf("taken=%v", end.Sub(start).Nanoseconds()/(1000*1000))
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
	for i := 1; i <= 6; i++ {
		total = 0
		run(i)
		time.Sleep(time.Millisecond * 10)
		fmt.Println()
	}
}

func run(buffChanSize int) {
	start := time.Now()
	fmt.Printf("start, total=%d, buffChanSize=%d, timeUnit=%s, generator=%d, validator=%d, adder=%d\n",
		total, buffChanSize, formatTimeUnit(timeUnit), generatorSleep, validatorSleep, adderSleep)

	noCh := make(chan Number, buffChanSize)
	evenCh := make(chan Number, buffChanSize)

	wg.Add(3)
	go adder(evenCh)
	go validator(noCh, evenCh)
	generator(noCh)

	wg.Wait()

	end := time.Now()
	fmt.Printf("end, %s\n", formatTime(end, start))

	if total != 15 {
		log.Fatalf("total want=15, got=%d\n", total)
	}
}

func generator(noCh chan<- Number) {
	for i := 1; i <= 5; i++ {
		generatorWork()
		noCh <- Number{no: i, start: time.Now()}
	}
	close(noCh)
	wg.Done()
}

func validator(noCh <-chan Number, evenCh chan<- Number) {
	for num := range noCh {
		if num.no > 0 {
			validatorWork()
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
		fmt.Printf("adder, %s\n", formatTime(evenNum.end, evenNum.start))
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

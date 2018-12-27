package main

import (
	"log"
	"time"
)

const (
	timeUnit       = time.Second
	generatorSleep = 1
	filterSleep    = 0
	adderSleep     = 0
)

var total = 0

func main() {
	log.Printf("start, total=%d\n", total)

	noCh := make(chan int)
	evenCh := make(chan int)

	go filter(noCh, evenCh)
	go adder(evenCh)

	generator(noCh)

	close(noCh)
	close(evenCh)

	log.Printf("end, total=%d\n", total)
}

func generator(noCh chan<- int) {
	for i := 1; i <= 10; i++ {
		// fmt.Printf("generator, no=%d\n", i)
		noCh <- i
		time.Sleep(timeUnit * generatorSleep)
	}
}

func filter(noCh <-chan int, evenCh chan<- int) {
	for no := range noCh {
		if no%2 == 0 {
			// fmt.Printf("filter, evenNo=%d\n", no)
			evenCh <- no
			time.Sleep(timeUnit * filterSleep)
		}
	}
}

func adder(evenCh <-chan int) {
	for evenNo := range evenCh {
		total += evenNo
		log.Printf("adder, total=%d\n", total)
		time.Sleep(timeUnit * adderSleep)
	}
}

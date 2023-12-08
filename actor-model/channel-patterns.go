// channle partters runs the Subscribe example with a real RSS fetcher.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	counterA := Counter(1, 10)
	counterB := Counter(15, 50)

	inputA := TurnOutChanType{Channel: counterA, ChanType: InputChannel}
	inputB := TurnOutChanType{Channel: counterB, ChanType: InputChannel}
	outputA := TurnOutChanType{Channel: make(chan int), ChanType: OutputChannel}
	outputB := TurnOutChanType{Channel: make(chan int), ChanType: OutputChannel}
	TurnOut(inputA, inputB, outputA, outputB)

	result := Merge(outputA.Channel, outputB.Channel)
	for ctr := range result {
		println(ctr)
	}
	return
	// create a WaitGroup
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for ctr := range outputA.Channel {
			println(ctr)
		}
		wg.Done()
	}()

	go func() {
		for ctr := range outputB.Channel {
			println(ctr)
		}
		wg.Done()
	}()

	wg.Wait()

	//result := Merge(outputA.Channel, outputB.Channel)
	//for ctr := range result {
	//	println(ctr)
	//}
}

func Counter(start int, end int) chan int {
	c := make(chan int)
	go func() {
		for i := start; i <= end; i++ {
			c <- i
			time.Sleep(time.Millisecond * 100)
		}
		close(c)
	}()
	return c
}

type ChanType string

const (
	InputChannel  ChanType = "INPUT_CHANNEL"
	OutputChannel          = "OUTPUT_CHANNEL"
)

type TurnOutChanType struct {
	Channel  chan int
	ChanType ChanType
}

// TurnOut returns a merged channel of `outputsChan` channels
// this produce fan-in channel
// this is variadic function
func TurnOut(channels ...TurnOutChanType) {
	inputCounts := 0
	for _, c := range channels {
		if c.ChanType == InputChannel {
			inputCounts++
		}
	}
	// create a WaitGroup
	var wg sync.WaitGroup
	wg.Add(inputCounts)

	output := func(sc <-chan int) {
		// run until channel (square numbers sender) closes
		for sqr := range sc {
			for _, mg := range channels {
				if mg.ChanType == OutputChannel {
					mg.Channel <- sqr
				}
			}
		}
		// call `Done` on `WaitGroup` to decrement counter
		wg.Done()
	}

	for _, optChan := range channels {
		if optChan.ChanType == InputChannel {
			go output(optChan.Channel)
		}
	}

	// run goroutine to close merged channel once done
	go func() {
		// wait until WaitGroup finishes
		wg.Wait()
		for _, mg := range channels {
			if mg.ChanType == OutputChannel {
				close(mg.Channel)
			}
		}
	}()

}

// Merge returns a merged channel of `outputsChan` channels
// this produce fan-in channel
// this is variadic function
func Merge(outputsChan ...<-chan int) <-chan int {
	// create a WaitGroup
	var wg sync.WaitGroup

	// make return channel
	merged := make(chan int, 100)

	// increase counter to number of channels `len(outputsChan)`
	// as we will spawn number of goroutines equal to number of channels received to merge
	wg.Add(len(outputsChan))

	// function that accept a channel (which sends square numbers)
	// to push numbers to merged channel
	output := func(sc <-chan int) {
		// run until channel (square numbers sender) closes
		for sqr := range sc {
			merged <- sqr
		}
		// once channel (square numbers sender) closes,
		// call `Done` on `WaitGroup` to decrement counter
		wg.Done()
	}

	// run above `output` function as groutines, `n` number of times
	// where n is equal to number of channels received as argument the function
	// here we are using `for range` loop on `outputsChan` hence no need to manually tell `n`
	for _, optChan := range outputsChan {
		go output(optChan)
	}

	// run goroutine to close merged channel once done
	go func() {
		// wait until WaitGroup finishesh
		wg.Wait()
		close(merged)
	}()

	return merged
}

func SelectAndNilChannel() {
	a, b := make(chan string), make(chan string)
	go func() { a <- "a" }()
	go func() { b <- "b" }()

	if rand.Int()%2 == 0 {
		a = nil
	} else {
		b = nil
	}

	select {
	case s := <-a:
		fmt.Println("Received A", s)
	case s := <-b:
		fmt.Println("Received B", s)
	}
}

func FanIn(inA, inB <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		more := 0
		for {
			select {
			case result, ok := <-inA:
				if !ok {
					more = +1
					continue
				}
				out <- result
			case result, ok := <-inB:
				if !ok {
					more = +1
					continue
				}
				out <- result
			}
			fmt.Println("More: ", more)
			if more >= 2 {
				close(out)
				break
			}

		}
	}()
	return out
}

// FanOut https://www.youtube.com/watch?v=YEKjSzIwAdA&t=10s
func FanOut(in <-chan int, outA, outB chan<- int) {
	for data := range in { // Receive until closed
		select { // Send to first non blocking
		case outA <- data:
		case outB <- data:
		}
	}
}

func Funnel(inA, inB <-chan int, out chan<- int) {
	for {
		data := 0
		more := false
		select {
		case data, more = <-inA:
		case data, more = <-inB:
		}
		if !more {
			return
		}
		select {
		case out <- data:
		}
	}
}

func TurnOut_v2(quit <-chan bool, inA, inB chan int, outA, outB chan<- int) {

	for {
		data := 0
		select {
		case data = <-inA:
		case data = <-inB:
		case <-quit:
			close(inA)
			close(inB)
			FanOut(inA, outA, outB) // Flush the remaining data.
			FanOut(inB, outA, outB)
			return
		}

		select {
		case outA <- data: // Send data to out channels.
		case outB <- data:
		}

	}
}

func TurnOut_v1(inA, inB <-chan int, outA, outB chan<- int) {

	for {
		data := 0
		more := false
		select {
		case data, more = <-inA:
		case data, more = <-inB:
		}

		if !more {
			return
		}

		select {
		case outA <- data:
		case outB <- data:
		}

	}
}

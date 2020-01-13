package worker_pool

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type foreman struct {
	channel chan string
}

// StartWork starts goroutines work
func (f *foreman) StartWork(countOfWorkers int) {
	var wg sync.WaitGroup

	for i := 1; i <= countOfWorkers; i++ {
		wg.Add(1)
		go f.checkWork(i, &wg)
	}

	for i := 1; i <= countOfWorkers; i++ {
		i := <-f.channel
		fmt.Println(i)
	}
	wg.Wait()
}

func (f *foreman) checkWork(i int, wg *sync.WaitGroup) {
	fmt.Println("I am worker " + strconv.Itoa(i) + " and i start my job")
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(3) + 1
	time.Sleep(time.Duration(n) * time.Second)
	f.channel <- "I am worker " + strconv.Itoa(i) + " and i finished my job in " + strconv.Itoa(n) + " seconds"
	wg.Done()
}

// NewForemaner returns new copy of foreman
func NewForemaner(channelIn chan string) Foremaner {
	return &foreman{channel: channelIn}
}

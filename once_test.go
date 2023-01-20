package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce() {
	counter++
}

// Once adalah fitur di Go-Lang yang bisa kita gunakan untuk memastikan bahsa sebuah function di eksekusi hanya sekali
func TestOnce(t *testing.T) {
	var once sync.Once
	var group sync.WaitGroup

	for i := 0; i < 100; i++ {
		go func ()  {
			add := 1
			group.Add(add)
			once.Do(OnlyOnce)
			group.Done()
		
		}()
	}

	group.Wait()
	fmt.Println("Counter Ke- ",counter)
}
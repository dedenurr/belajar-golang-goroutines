package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Mutex adalah solusi dari permasalahan Race Condition (permasalahan manipulasi data variable yang sama oleh beberapa goroutine secara bersamaan)
func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex //kode mutex
	for i := 0; i <= 1000; i++ {
		go func() {
			for j := 0; j <= 100; j++ {
				mutex.Lock() //kode mutex lock
				x = x + 1
				mutex.Unlock() ////kode mutex unlock
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)
}
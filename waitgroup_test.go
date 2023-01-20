package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 1. WaitGroup adalah fitur yang bisa digunakan untuk menunggu sebuah proses selesai dilakukan, dengan fitur ini kita tdak perlu lagi menebak proses goroutines selesai menggunakan timeSleep

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)

	
}

func TestWaitGroup(t *testing.T)  {
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(&group)
	}
	group.Wait()
	fmt.Println("complete")
}
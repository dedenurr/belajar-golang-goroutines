package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

// 1. Create Channel
func TestCreateChannel(t *testing.T) {
	channel := make (chan string)
	defer close(channel)

	go func ()  {
		time.Sleep(2 * time.Second)
		channel <- "Dede Nurrahman"
		fmt.Println("Selesai Mengirim Data Ke Channel")
	}()

	

    data := <- channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
	
	// close(channel)
}

// 2. Create Channel As A Parameter
func GiveMeResponse(channel chan string)  {
	time.Sleep(2 * time.Second)
	channel <- "Dede Nurrahman"
}

func TestChannelAsParameter(t *testing.T)  {
	channel := make(chan string)
	
	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
	close(channel)
}

// 3. channel in (mengirim data)
func OnlyIn(channel chan<- string)  {
	time.Sleep(2 * time.Second)
	channel <- "Dede Nurrahman"
}

// 4. channel out (menerima data)
func OnlyOut(channel <-chan string)  {
	
	data := <-channel
	fmt.Println(data)
	
}

func TestInOutChannel(t *testing.T)  {
	channel := make(chan string)
	defer close(channel)
	
	go OnlyIn(channel)
	go OnlyOut(channel)
	time.Sleep(2 * time.Second)
	
}

// 5 channel buffer -> dilakukan agar data yang dimasukan bisa masuk ke dalam buffer time
func TestBufferedChannel(t *testing.T)  {
	channel := make(chan string, 2)
	
	// nilai channel yang dikirimkan harus sama dengan jumlah buffer di kode yang diatas yaitu 2, jika tidak sesuai akan deadlock
	go func ()  {
		channel <- "dede"
		channel <- "nurrahman"
	}()
	

	// jumlah channel yang diterima pun harus sama yaitu 2
	go func ()  {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("selesai buffer")
	close(channel)
}
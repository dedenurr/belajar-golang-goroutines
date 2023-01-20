package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 1. Mutex adalah solusi dari permasalahan Race Condition (permasalahan manipulasi data variable yang sama oleh beberapa goroutine secara bersamaan)
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

// 2. RWmutex dilakukan melakukan locking tidak hanya pada proses mengubah data, tapi juga membaca data
type BankAccount struct{
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount)AddBalance(ammount int)  {
	account.RWMutex.Lock()
	account.Balance = account.Balance + ammount
	account.RWMutex.Unlock()
}

func (account *BankAccount)GetBalance()int  {
	account.RWMutex.RLock()	
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}


func TestRwMutex(t *testing.T)  {
	account := BankAccount{}

	for i := 0; i < 100; i++ {		
		go  func ()  {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance ",account.GetBalance())
}

// 3. Deadlock keadaan dimana sebuah proses goroutine saling menunggu lock sehingga tidak ada satupun goroutine yang bisa jalan
type UserBalance struct{
	sync.Mutex
	Name string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock()  {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int)  {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int)  {
	user1.Lock()
	fmt.Println("Lock User 1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock User 2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)
	
	user1.Unlock()
	user2.Unlock()
}

func TestDeadLock(t *testing.T)  {
	user1 := UserBalance{
		Name :"Dede",
		Balance: 1000000,
	
	}

	user2 := UserBalance{
		Name :"Zahra",
		Balance: 1000000,
	
	}
	
	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)

	time.Sleep(10 * time.Second)


	fmt.Println("User ",user1.Name,", Balance", user1.Balance)
	fmt.Println("User ",user2.Name,", Balance", user2.Balance)
	
	
}
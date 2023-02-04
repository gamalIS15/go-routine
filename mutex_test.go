package go_routine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	for i := 0; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(4 * time.Second)
	fmt.Println("Counter = ", x)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (b *BankAccount) AddBalance(amount int) {
	b.RWMutex.Lock()
	b.Balance = b.Balance + amount
	b.RWMutex.Unlock()
}

func (b *BankAccount) GetBalance() int {
	b.RWMutex.RLock()
	balance := b.Balance
	b.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	acc := BankAccount{}
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				acc.AddBalance(1)
				fmt.Println(acc.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance", acc.GetBalance())
}

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock User 1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)
	user2.Lock()
	fmt.Println("Lock user 2", user2.Name)
	user2.Change(amount)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Eko",
		Balance: 100000,
	}

	user2 := UserBalance{
		Name:    "Dono",
		Balance: 100000,
	}

	go Transfer(&user1, &user2, 10000)
	go Transfer(&user2, &user1, 20000)

	time.Sleep(10 * time.Second)
	fmt.Println("User ", user1.Name, ", Balance ", user1.Balance)
	fmt.Println("User ", user2.Name, ", Balance ", user2.Balance)

}

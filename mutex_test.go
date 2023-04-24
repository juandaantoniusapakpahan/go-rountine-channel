package gorountinechannel

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Counter: ", x)
}

type EMoneyAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *EMoneyAccount) Add(nominal int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + 1
	account.RWMutex.Unlock()
}

func (account *EMoneyAccount) Get() int {
	account.RWMutex.RLock()
	total := account.Balance
	account.RWMutex.RUnlock()
	return total
}

func TestRWMutex(t *testing.T) {
	emoneyAccount := EMoneyAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				emoneyAccount.Add(j)
				fmt.Println(emoneyAccount.Get())
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Total: ", emoneyAccount.Balance)
}

type UserAccount struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserAccount) Lock() {
	user.Mutex.Lock()
}

func (user *UserAccount) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserAccount) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserAccount, user2 *UserAccount, amount int) {
	user1.Lock()
	fmt.Println("Lock user1", user1.Name)
	user1.Change(-amount)
	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadLock(t *testing.T) {
	data1 := UserAccount{
		Name:    "Richart",
		Balance: 1000000,
	}
	data2 := UserAccount{
		Name:    "Robert",
		Balance: 1000000,
	}

	go Transfer(&data1, &data2, 200000) // 1

	go Transfer(&data2, &data1, 200000) // 2

	time.Sleep(3 * time.Second)

	fmt.Println(data1.Name, data1.Balance)

	fmt.Println(data2.Name, data2.Balance)

}

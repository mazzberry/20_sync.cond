package main

import (
	"fmt"
	"sync"
	"time"
)

var userList = []int{}
var ready = false

func main() {

	// condition := sync.Cond{L: &sync.Mutex{}}
	condition := sync.NewCond(&sync.Mutex{})

	for i := 0; i < 1000; i++ {
		go NewRequest(i, condition)

	}

	time.Sleep(time.Second * 10)

}

func NewRequest(userId int, condition *sync.Cond) {
	Checking(userId, condition)
	condition.L.Lock()
	defer condition.L.Unlock()
	for !ready {
		condition.Wait()
	}
	fmt.Println("User ", userId, "start streaming")

}

func Checking(userId int, condition *sync.Cond) {
	fmt.Println(userId, "waiting to start streaming")
	time.Sleep(time.Millisecond * 150)
	condition.L.Lock()
	defer condition.L.Unlock()
	userList = append(userList, userId)
	if len(userList) == 55 {
		ready = true
		condition.Broadcast()
		fmt.Println("User ", userId, "no more waiting")
	}

}
